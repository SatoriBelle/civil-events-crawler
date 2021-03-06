// Package gen contains all the components for code generation.
package gen

const watcherTmpl = `
// Code generated by 'gen/watchergen'  DO NOT EDIT.
// IT SHOULD NOT BE EDITED BY HAND AS ANY CHANGES MAY BE OVERWRITTEN
// Please reference 'listener/watchergen' for more details
// File was generated at {{.GenTime}}
package {{.PackageName}}

import (
	// "fmt"
	"time"
	"context"

	log "github.com/golang/glog"
	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"


	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/utils"

	ctime "github.com/joincivil/go-common/pkg/time"
{{if .ContractImportPath -}}
	"{{.ContractImportPath}}"
{{- end}}
{{if .AdditionalImports -}}
{{- range .AdditionalImports}}
	"{{.}}"
{{- end}}
{{- end}}
)

func New{{.ContractTypeName}}Watchers(contractAddress common.Address) *{{.ContractTypeName}}Watchers {
	return &{{.ContractTypeName}}Watchers{
		contractAddress: contractAddress,
	}
}

type {{.ContractTypeName}}Watchers struct {
	errors chan error
	contractAddress common.Address
	contract *{{.ContractTypePackage}}.{{.ContractTypeName}}
	activeSubs []utils.WatcherSubscription
}

func (w *{{.ContractTypeName}}Watchers) ContractAddress() common.Address {
	return w.contractAddress
}

func (w *{{.ContractTypeName}}Watchers) ContractName() string {
	return "{{.ContractTypeName}}"
}

func (w *{{.ContractTypeName}}Watchers) cancelFunc(cancelFn context.CancelFunc, killCancel <-chan bool) {
}

func (w *{{.ContractTypeName}}Watchers) StopWatchers(unsub bool) error {
	if unsub {
		for _, sub := range w.activeSubs {
			sub.Unsubscribe()
		}
	}
	w.activeSubs = nil
	return nil
}

func (w *{{.ContractTypeName}}Watchers) StartWatchers(client bind.ContractBackend,
	eventRecvChan chan *model.Event, errs chan error) ([]utils.WatcherSubscription, error) {
	return w.Start{{.ContractTypeName}}Watchers(client, eventRecvChan, errs)
}

// Start{{.ContractTypeName}}Watchers starts up the event watchers for {{.ContractTypeName}}
func (w *{{.ContractTypeName}}Watchers) Start{{.ContractTypeName}}Watchers(client bind.ContractBackend,
	eventRecvChan chan *model.Event, errs chan error) ([]utils.WatcherSubscription, error) {
	w.errors = errs
    contract, err := {{.ContractTypePackage}}.New{{.ContractTypeName}}(w.contractAddress, client)
	if err != nil {
        log.Errorf("Error initializing Start{{.ContractTypeName}}: err: %v", err)
		return nil, errors.Wrap(err, "error initializing Start{{.ContractTypeName}}")
	}
	w.contract = contract

    var sub utils.WatcherSubscription
	subs := []utils.WatcherSubscription{}
{{if .EventHandlers -}}
{{- range .EventHandlers}}

    sub, err = w.startWatch{{.EventMethod}}(eventRecvChan)
	if err != nil {
        return nil, errors.WithMessage(err, "error starting start{{.EventMethod}}")
	}
	subs = append(subs, sub)

{{- end}}
{{- end}}

	w.activeSubs = subs
    return subs, nil
}

{{if .EventHandlers -}}
{{- range .EventHandlers}}

func (w *{{$.ContractTypeName}}Watchers) startWatch{{.EventMethod}}(eventRecvChan chan *model.Event) (utils.WatcherSubscription, error) {
	killCancelTimeoutSecs := 10
	preemptiveTimeoutSecs := 60 * 30
	return utils.NewWatcherSubscription("Watch{{.EventMethod}}", func(quit <-chan struct{}) error {
		startupFn := func() (utils.WatcherSubscription, chan *{{$.ContractTypePackage}}.{{.EventType}}, error) {
			ctx := context.Background()
			ctx, cancelFn := context.WithCancel(ctx)
			opts := &bind.WatchOpts{Context: ctx}
			killCancel := make(chan bool)
			// 10 sec timeout mechanism for starting up watcher
			go func(cancelFn context.CancelFunc, killCancel <-chan bool) {
				select {
				case <-time.After(time.Duration(killCancelTimeoutSecs) * time.Second):
					log.Errorf("Watch{{.EventMethod}} start timeout, cancelling...")
					cancelFn()
				case <-killCancel:
				}
			}(cancelFn, killCancel)
			recvChan := make(chan *{{$.ContractTypePackage}}.{{.EventType}})
			log.Infof("startupFn: Starting Watch{{.EventMethod}}")
			sub, err := w.contract.Watch{{.EventMethod}}(
				opts,
				recvChan,
				{{- if .ParamValues -}}
				{{range .ParamValues}}
					[]{{.Type}}{},
				{{- end}}
				{{end}}
			)
			close(killCancel)
			if err != nil {
				if sub != nil {
					log.Infof("startupFn: Unsubscribing Watch{{.EventMethod}}")
					sub.Unsubscribe()
				}
				return nil, nil, errors.Wrap(err, "startupFn: error starting Watch{{.EventMethod}}")
			}
			log.Infof("startupFn: Watch{{.EventMethod}} started")
			return sub, recvChan, nil
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting Watch{{.EventMethod}}: %v", err)
			if sub != nil {
				sub.Unsubscribe()
			}
			w.errors <- err
			return err
		}
		defer sub.Unsubscribe()
		log.Infof("Starting up Watch{{.EventMethod}} for contract %v", w.contractAddress.Hex())
		for {
			select {
			// 30 min premptive resubscribe
			case <-time.After(time.Second * time.Duration(preemptiveTimeoutSecs)):
				log.Infof("Premptive restart of {{.EventMethod}}")
				oldSub := sub
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("Error starting {{.EventMethod}}: %v", err)
					w.errors <- err
					return err
				}
				log.Infof("Attempting to unsub old {{.EventMethod}}")
				oldSub.Unsubscribe()
				log.Infof("Done preemptive restart {{.EventMethod}}")
			case event := <-recvChan:
				if log.V(2) {
					log.Infof("Received event on Watch{{.EventMethod}}: %v", spew.Sprintf("%#+v", event))
				} else {
					log.Info("Received event on Watch{{.EventMethod}}")
				}
				modelEvent, err := model.NewEventFromContractEvent("{{.EventMethod}}", w.ContractName(), w.contractAddress, event, ctime.CurrentEpochSecsInInt64(), model.Watcher)
				if err != nil {
					log.Errorf("Error creating new event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- modelEvent:
					if log.V(2) {
						log.Infof("Sent event to eventRecvChan on Watch{{.EventMethod}}: %v", spew.Sprintf("%#+v", event))
					} else {
						log.Info("Sent event to eventRecvChan on Watch{{.EventMethod}}")
					}
				case err := <-sub.Err():
					log.Errorf("Error with Watch{{.EventMethod}}, fatal (a): %v", err)
					err = errors.Wrap(err, "error with Watch{{.EventMethod}}")
					w.errors <- err
					return err
				case <-quit:
					log.Infof("Quit Watch{{.EventMethod}} (a): %v", err)
					return nil
				}
			case err := <-sub.Err():
				log.Errorf("Error with Watch{{.EventMethod}}, fatal (b): %v", err)
				err = errors.Wrap(err, "error with Watch{{.EventMethod}}")
				w.errors <- err
				return err
			case <-quit:
				log.Infof("Quitting loop for Watch{{.EventMethod}}")
				return nil
			}
		}
	}), nil
}

{{- end}}
{{- end}}
`
