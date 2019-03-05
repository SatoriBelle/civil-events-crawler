// Code generated by 'gen/watchergen'  DO NOT EDIT.
// IT SHOULD NOT BE EDITED BY HAND AS ANY CHANGES MAY BE OVERWRITTEN
// Please reference 'listener/watchergen' for more details
// File was generated at 2019-03-05 21:05:06.715724 +0000 UTC
package watcher

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	log "github.com/golang/glog"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"

	"github.com/joincivil/civil-events-crawler/pkg/model"

	"github.com/joincivil/go-common/pkg/generated/contract"
	ctime "github.com/joincivil/go-common/pkg/time"
)

func NewCVLTokenContractWatchers(contractAddress common.Address) *CVLTokenContractWatchers {
	return &CVLTokenContractWatchers{
		contractAddress: contractAddress,
	}
}

type CVLTokenContractWatchers struct {
	contractAddress common.Address
	contract        *contract.CVLTokenContract
	activeSubs      []event.Subscription
}

func (w *CVLTokenContractWatchers) ContractAddress() common.Address {
	return w.contractAddress
}

func (w *CVLTokenContractWatchers) ContractName() string {
	return "CVLTokenContract"
}

func (w *CVLTokenContractWatchers) StopWatchers() error {
	for _, sub := range w.activeSubs {
		sub.Unsubscribe()
	}
	w.activeSubs = nil
	return nil
}

func (w *CVLTokenContractWatchers) StartWatchers(client bind.ContractBackend,
	eventRecvChan chan *model.Event) ([]event.Subscription, error) {
	return w.StartCVLTokenContractWatchers(client, eventRecvChan)
}

// StartCVLTokenContractWatchers starts up the event watchers for CVLTokenContract
func (w *CVLTokenContractWatchers) StartCVLTokenContractWatchers(client bind.ContractBackend,
	eventRecvChan chan *model.Event) ([]event.Subscription, error) {
	contract, err := contract.NewCVLTokenContract(w.contractAddress, client)
	if err != nil {
		log.Errorf("Error initializing StartCVLTokenContract: err: %v", err)
		return nil, err
	}
	w.contract = contract

	var sub event.Subscription
	subs := []event.Subscription{}

	sub, err = w.startWatchApproval(eventRecvChan)
	if err != nil {
		return nil, fmt.Errorf("Error starting startApproval: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = w.startWatchOwnershipRenounced(eventRecvChan)
	if err != nil {
		return nil, fmt.Errorf("Error starting startOwnershipRenounced: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = w.startWatchOwnershipTransferred(eventRecvChan)
	if err != nil {
		return nil, fmt.Errorf("Error starting startOwnershipTransferred: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = w.startWatchTransfer(eventRecvChan)
	if err != nil {
		return nil, fmt.Errorf("Error starting startTransfer: err: %v", err)
	}
	subs = append(subs, sub)

	w.activeSubs = subs
	return subs, nil
}

func (w *CVLTokenContractWatchers) startWatchApproval(eventRecvChan chan *model.Event) (event.Subscription, error) {
	return event.NewSubscription(func(quit <-chan struct{}) error {
		maxRetries := 5
		startupFn := func() (event.Subscription, chan *contract.CVLTokenContractApproval, error) {
			retry := 0
			for {
				opts := &bind.WatchOpts{}
				recvChan := make(chan *contract.CVLTokenContractApproval)
				log.Infof("startupFn: Starting WatchApproval")
				sub, err := w.contract.WatchApproval(
					opts,
					recvChan,
					[]common.Address{},
					[]common.Address{},
				)
				if err != nil {
					if sub != nil {
						log.Infof("startupFn: Unsubscribing WatchApproval")
						sub.Unsubscribe()
					}
					if retry >= maxRetries {
						return nil, nil, err
					}
					retry++
					log.Warningf("startupFn: Retrying start WatchApproval: %v", err)
					continue
				}
				return sub, recvChan, nil
			}
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchApproval: %v", err)
			return err
		}
		defer sub.Unsubscribe()
		log.Infof("Starting up WatchApproval for contract %v", w.contractAddress.Hex())
		for {
			select {
			// 15 min premptive resubscribe
			case <-time.After(time.Second * time.Duration(60*15)):
				// log.Infof("Premptive restart of Approval")
				oldSub := sub
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("Error starting Approval: %v", err)
					return err
				}
				oldSub.Unsubscribe()
				log.Infof("Done preemptive restart Approval")
			case event := <-recvChan:
				if log.V(2) {
					log.Infof("Received event on WatchApproval: %v", spew.Sprintf("%#+v", event))
				} else {
					log.Info("Received event on WatchApproval")
				}
				modelEvent, err := model.NewEventFromContractEvent("Approval", w.ContractName(), w.contractAddress, event, ctime.CurrentEpochSecsInInt64(), model.Watcher)
				if err != nil {
					log.Errorf("Error creating new event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- modelEvent:
				case err := <-sub.Err():
					log.Errorf("Error with WatchApproval, fatal (a): %v", err)
					sub.Unsubscribe()
					sub, recvChan, err = startupFn()
					if err != nil {
						log.Errorf("Error restarting WatchApproval, fatal (a): %v", err)
						return err
					}
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				log.Errorf("Error with WatchApproval, fatal (b): %v", err)
				sub.Unsubscribe()
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("WATCHER: Error restarting WatchApproval, fatal (b): %v", err)
					return err
				}
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (w *CVLTokenContractWatchers) startWatchOwnershipRenounced(eventRecvChan chan *model.Event) (event.Subscription, error) {
	return event.NewSubscription(func(quit <-chan struct{}) error {
		maxRetries := 5
		startupFn := func() (event.Subscription, chan *contract.CVLTokenContractOwnershipRenounced, error) {
			retry := 0
			for {
				opts := &bind.WatchOpts{}
				recvChan := make(chan *contract.CVLTokenContractOwnershipRenounced)
				log.Infof("startupFn: Starting WatchOwnershipRenounced")
				sub, err := w.contract.WatchOwnershipRenounced(
					opts,
					recvChan,
					[]common.Address{},
				)
				if err != nil {
					if sub != nil {
						log.Infof("startupFn: Unsubscribing WatchOwnershipRenounced")
						sub.Unsubscribe()
					}
					if retry >= maxRetries {
						return nil, nil, err
					}
					retry++
					log.Warningf("startupFn: Retrying start WatchOwnershipRenounced: %v", err)
					continue
				}
				return sub, recvChan, nil
			}
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchOwnershipRenounced: %v", err)
			return err
		}
		defer sub.Unsubscribe()
		log.Infof("Starting up WatchOwnershipRenounced for contract %v", w.contractAddress.Hex())
		for {
			select {
			// 15 min premptive resubscribe
			case <-time.After(time.Second * time.Duration(60*15)):
				// log.Infof("Premptive restart of OwnershipRenounced")
				oldSub := sub
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("Error starting OwnershipRenounced: %v", err)
					return err
				}
				oldSub.Unsubscribe()
				log.Infof("Done preemptive restart OwnershipRenounced")
			case event := <-recvChan:
				if log.V(2) {
					log.Infof("Received event on WatchOwnershipRenounced: %v", spew.Sprintf("%#+v", event))
				} else {
					log.Info("Received event on WatchOwnershipRenounced")
				}
				modelEvent, err := model.NewEventFromContractEvent("OwnershipRenounced", w.ContractName(), w.contractAddress, event, ctime.CurrentEpochSecsInInt64(), model.Watcher)
				if err != nil {
					log.Errorf("Error creating new event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- modelEvent:
				case err := <-sub.Err():
					log.Errorf("Error with WatchOwnershipRenounced, fatal (a): %v", err)
					sub.Unsubscribe()
					sub, recvChan, err = startupFn()
					if err != nil {
						log.Errorf("Error restarting WatchOwnershipRenounced, fatal (a): %v", err)
						return err
					}
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				log.Errorf("Error with WatchOwnershipRenounced, fatal (b): %v", err)
				sub.Unsubscribe()
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("WATCHER: Error restarting WatchOwnershipRenounced, fatal (b): %v", err)
					return err
				}
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (w *CVLTokenContractWatchers) startWatchOwnershipTransferred(eventRecvChan chan *model.Event) (event.Subscription, error) {
	return event.NewSubscription(func(quit <-chan struct{}) error {
		maxRetries := 5
		startupFn := func() (event.Subscription, chan *contract.CVLTokenContractOwnershipTransferred, error) {
			retry := 0
			for {
				opts := &bind.WatchOpts{}
				recvChan := make(chan *contract.CVLTokenContractOwnershipTransferred)
				log.Infof("startupFn: Starting WatchOwnershipTransferred")
				sub, err := w.contract.WatchOwnershipTransferred(
					opts,
					recvChan,
					[]common.Address{},
					[]common.Address{},
				)
				if err != nil {
					if sub != nil {
						log.Infof("startupFn: Unsubscribing WatchOwnershipTransferred")
						sub.Unsubscribe()
					}
					if retry >= maxRetries {
						return nil, nil, err
					}
					retry++
					log.Warningf("startupFn: Retrying start WatchOwnershipTransferred: %v", err)
					continue
				}
				return sub, recvChan, nil
			}
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchOwnershipTransferred: %v", err)
			return err
		}
		defer sub.Unsubscribe()
		log.Infof("Starting up WatchOwnershipTransferred for contract %v", w.contractAddress.Hex())
		for {
			select {
			// 15 min premptive resubscribe
			case <-time.After(time.Second * time.Duration(60*15)):
				// log.Infof("Premptive restart of OwnershipTransferred")
				oldSub := sub
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("Error starting OwnershipTransferred: %v", err)
					return err
				}
				oldSub.Unsubscribe()
				log.Infof("Done preemptive restart OwnershipTransferred")
			case event := <-recvChan:
				if log.V(2) {
					log.Infof("Received event on WatchOwnershipTransferred: %v", spew.Sprintf("%#+v", event))
				} else {
					log.Info("Received event on WatchOwnershipTransferred")
				}
				modelEvent, err := model.NewEventFromContractEvent("OwnershipTransferred", w.ContractName(), w.contractAddress, event, ctime.CurrentEpochSecsInInt64(), model.Watcher)
				if err != nil {
					log.Errorf("Error creating new event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- modelEvent:
				case err := <-sub.Err():
					log.Errorf("Error with WatchOwnershipTransferred, fatal (a): %v", err)
					sub.Unsubscribe()
					sub, recvChan, err = startupFn()
					if err != nil {
						log.Errorf("Error restarting WatchOwnershipTransferred, fatal (a): %v", err)
						return err
					}
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				log.Errorf("Error with WatchOwnershipTransferred, fatal (b): %v", err)
				sub.Unsubscribe()
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("WATCHER: Error restarting WatchOwnershipTransferred, fatal (b): %v", err)
					return err
				}
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (w *CVLTokenContractWatchers) startWatchTransfer(eventRecvChan chan *model.Event) (event.Subscription, error) {
	return event.NewSubscription(func(quit <-chan struct{}) error {
		maxRetries := 5
		startupFn := func() (event.Subscription, chan *contract.CVLTokenContractTransfer, error) {
			retry := 0
			for {
				opts := &bind.WatchOpts{}
				recvChan := make(chan *contract.CVLTokenContractTransfer)
				log.Infof("startupFn: Starting WatchTransfer")
				sub, err := w.contract.WatchTransfer(
					opts,
					recvChan,
					[]common.Address{},
					[]common.Address{},
				)
				if err != nil {
					if sub != nil {
						log.Infof("startupFn: Unsubscribing WatchTransfer")
						sub.Unsubscribe()
					}
					if retry >= maxRetries {
						return nil, nil, err
					}
					retry++
					log.Warningf("startupFn: Retrying start WatchTransfer: %v", err)
					continue
				}
				return sub, recvChan, nil
			}
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchTransfer: %v", err)
			return err
		}
		defer sub.Unsubscribe()
		log.Infof("Starting up WatchTransfer for contract %v", w.contractAddress.Hex())
		for {
			select {
			// 15 min premptive resubscribe
			case <-time.After(time.Second * time.Duration(60*15)):
				// log.Infof("Premptive restart of Transfer")
				oldSub := sub
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("Error starting Transfer: %v", err)
					return err
				}
				oldSub.Unsubscribe()
				log.Infof("Done preemptive restart Transfer")
			case event := <-recvChan:
				if log.V(2) {
					log.Infof("Received event on WatchTransfer: %v", spew.Sprintf("%#+v", event))
				} else {
					log.Info("Received event on WatchTransfer")
				}
				modelEvent, err := model.NewEventFromContractEvent("Transfer", w.ContractName(), w.contractAddress, event, ctime.CurrentEpochSecsInInt64(), model.Watcher)
				if err != nil {
					log.Errorf("Error creating new event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- modelEvent:
				case err := <-sub.Err():
					log.Errorf("Error with WatchTransfer, fatal (a): %v", err)
					sub.Unsubscribe()
					sub, recvChan, err = startupFn()
					if err != nil {
						log.Errorf("Error restarting WatchTransfer, fatal (a): %v", err)
						return err
					}
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				log.Errorf("Error with WatchTransfer, fatal (b): %v", err)
				sub.Unsubscribe()
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("WATCHER: Error restarting WatchTransfer, fatal (b): %v", err)
					return err
				}
			case <-quit:
				return nil
			}
		}
	}), nil
}
