// Code generated by 'gen/watchergen'  DO NOT EDIT.
// IT SHOULD NOT BE EDITED BY HAND AS ANY CHANGES MAY BE OVERWRITTEN
// Please reference 'listener/watchergen' for more details
// File was generated at 2019-03-14 22:47:21.034091 +0000 UTC
package watcher

import (
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
	log "github.com/golang/glog"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"

	"github.com/joincivil/civil-events-crawler/pkg/model"

	"github.com/joincivil/go-common/pkg/generated/contract"
	ctime "github.com/joincivil/go-common/pkg/time"

	"math/big"
)

func NewCivilPLCRVotingContractWatchers(contractAddress common.Address) *CivilPLCRVotingContractWatchers {
	return &CivilPLCRVotingContractWatchers{
		contractAddress: contractAddress,
	}
}

type CivilPLCRVotingContractWatchers struct {
	errors          chan error
	contractAddress common.Address
	contract        *contract.CivilPLCRVotingContract
	activeSubs      []event.Subscription
}

func (w *CivilPLCRVotingContractWatchers) ContractAddress() common.Address {
	return w.contractAddress
}

func (w *CivilPLCRVotingContractWatchers) ContractName() string {
	return "CivilPLCRVotingContract"
}

func (w *CivilPLCRVotingContractWatchers) StopWatchers(unsub bool) error {
	if unsub {
		for _, sub := range w.activeSubs {
			sub.Unsubscribe()
		}
	}
	w.activeSubs = nil
	return nil
}

func (w *CivilPLCRVotingContractWatchers) StartWatchers(client bind.ContractBackend,
	eventRecvChan chan *model.Event, errs chan error) ([]event.Subscription, error) {
	return w.StartCivilPLCRVotingContractWatchers(client, eventRecvChan, errs)
}

// StartCivilPLCRVotingContractWatchers starts up the event watchers for CivilPLCRVotingContract
func (w *CivilPLCRVotingContractWatchers) StartCivilPLCRVotingContractWatchers(client bind.ContractBackend,
	eventRecvChan chan *model.Event, errs chan error) ([]event.Subscription, error) {
	w.errors = errs
	contract, err := contract.NewCivilPLCRVotingContract(w.contractAddress, client)
	if err != nil {
		log.Errorf("Error initializing StartCivilPLCRVotingContract: err: %v", err)
		return nil, err
	}
	w.contract = contract

	var sub event.Subscription
	subs := []event.Subscription{}

	sub, err = w.startWatchPollCreated(eventRecvChan)
	if err != nil {
		return nil, fmt.Errorf("Error starting startPollCreated: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = w.startWatchTokensRescued(eventRecvChan)
	if err != nil {
		return nil, fmt.Errorf("Error starting startTokensRescued: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = w.startWatchVoteCommitted(eventRecvChan)
	if err != nil {
		return nil, fmt.Errorf("Error starting startVoteCommitted: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = w.startWatchVoteRevealed(eventRecvChan)
	if err != nil {
		return nil, fmt.Errorf("Error starting startVoteRevealed: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = w.startWatchVotingRightsGranted(eventRecvChan)
	if err != nil {
		return nil, fmt.Errorf("Error starting startVotingRightsGranted: err: %v", err)
	}
	subs = append(subs, sub)

	sub, err = w.startWatchVotingRightsWithdrawn(eventRecvChan)
	if err != nil {
		return nil, fmt.Errorf("Error starting startVotingRightsWithdrawn: err: %v", err)
	}
	subs = append(subs, sub)

	w.activeSubs = subs
	return subs, nil
}

func (w *CivilPLCRVotingContractWatchers) startWatchPollCreated(eventRecvChan chan *model.Event) (event.Subscription, error) {
	return event.NewSubscription(func(quit <-chan struct{}) error {
		startupFn := func() (event.Subscription, chan *contract.CivilPLCRVotingContractPollCreated, error) {
			for {
				opts := &bind.WatchOpts{}
				recvChan := make(chan *contract.CivilPLCRVotingContractPollCreated)
				log.Infof("startupFn: Starting WatchPollCreated")
				sub, err := w.contract.WatchPollCreated(
					opts,
					recvChan,
					[]*big.Int{},
					[]common.Address{},
				)
				if err != nil {
					if sub != nil {
						log.Infof("startupFn: Unsubscribing WatchPollCreated")
						sub.Unsubscribe()
					}
					return nil, nil, err
				}
				log.Infof("startupFn: WatchPollCreated started")
				return sub, recvChan, nil
			}
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchPollCreated: %v", err)
			if sub != nil {
				sub.Unsubscribe()
			}
			w.errors <- err
			return err
		}
		defer sub.Unsubscribe()
		log.Infof("Starting up WatchPollCreated for contract %v", w.contractAddress.Hex())
		for {
			select {
			// 15 min premptive resubscribe
			case <-time.After(time.Second * time.Duration(60*15)):
				log.Infof("Premptive restart of PollCreated")
				oldSub := sub
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("Error starting PollCreated: %v", err)
					w.errors <- err
					return err
				}
				log.Infof("Attempting to unsub old PollCreated")
				oldSub.Unsubscribe()
				log.Infof("Done preemptive restart PollCreated")
			case event := <-recvChan:
				if log.V(2) {
					log.Infof("Received event on WatchPollCreated: %v", spew.Sprintf("%#+v", event))
				} else {
					log.Info("Received event on WatchPollCreated")
				}
				modelEvent, err := model.NewEventFromContractEvent("PollCreated", w.ContractName(), w.contractAddress, event, ctime.CurrentEpochSecsInInt64(), model.Watcher)
				if err != nil {
					log.Errorf("Error creating new event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- modelEvent:
					if log.V(2) {
						log.Infof("Sent event to eventRecvChan on WatchPollCreated: %v", spew.Sprintf("%#+v", event))
					} else {
						log.Info("Sent event to eventRecvChan on WatchPollCreated")
					}
				case err := <-sub.Err():
					log.Errorf("Error with WatchPollCreated, fatal (a): %v", err)
					w.errors <- err
					return err
				case <-quit:
					log.Infof("Quit WatchPollCreated (a): %v", err)
					return nil
				}
			case err := <-sub.Err():
				log.Errorf("Error with WatchPollCreated, fatal (b): %v", err)
				w.errors <- err
				return err
			case <-quit:
				log.Infof("Quitting loop for WatchPollCreated")
				return nil
			}
		}
	}), nil
}

func (w *CivilPLCRVotingContractWatchers) startWatchTokensRescued(eventRecvChan chan *model.Event) (event.Subscription, error) {
	return event.NewSubscription(func(quit <-chan struct{}) error {
		startupFn := func() (event.Subscription, chan *contract.CivilPLCRVotingContractTokensRescued, error) {
			for {
				opts := &bind.WatchOpts{}
				recvChan := make(chan *contract.CivilPLCRVotingContractTokensRescued)
				log.Infof("startupFn: Starting WatchTokensRescued")
				sub, err := w.contract.WatchTokensRescued(
					opts,
					recvChan,
					[]*big.Int{},
					[]common.Address{},
				)
				if err != nil {
					if sub != nil {
						log.Infof("startupFn: Unsubscribing WatchTokensRescued")
						sub.Unsubscribe()
					}
					return nil, nil, err
				}
				log.Infof("startupFn: WatchTokensRescued started")
				return sub, recvChan, nil
			}
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchTokensRescued: %v", err)
			if sub != nil {
				sub.Unsubscribe()
			}
			w.errors <- err
			return err
		}
		defer sub.Unsubscribe()
		log.Infof("Starting up WatchTokensRescued for contract %v", w.contractAddress.Hex())
		for {
			select {
			// 15 min premptive resubscribe
			case <-time.After(time.Second * time.Duration(60*15)):
				log.Infof("Premptive restart of TokensRescued")
				oldSub := sub
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("Error starting TokensRescued: %v", err)
					w.errors <- err
					return err
				}
				log.Infof("Attempting to unsub old TokensRescued")
				oldSub.Unsubscribe()
				log.Infof("Done preemptive restart TokensRescued")
			case event := <-recvChan:
				if log.V(2) {
					log.Infof("Received event on WatchTokensRescued: %v", spew.Sprintf("%#+v", event))
				} else {
					log.Info("Received event on WatchTokensRescued")
				}
				modelEvent, err := model.NewEventFromContractEvent("TokensRescued", w.ContractName(), w.contractAddress, event, ctime.CurrentEpochSecsInInt64(), model.Watcher)
				if err != nil {
					log.Errorf("Error creating new event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- modelEvent:
					if log.V(2) {
						log.Infof("Sent event to eventRecvChan on WatchTokensRescued: %v", spew.Sprintf("%#+v", event))
					} else {
						log.Info("Sent event to eventRecvChan on WatchTokensRescued")
					}
				case err := <-sub.Err():
					log.Errorf("Error with WatchTokensRescued, fatal (a): %v", err)
					w.errors <- err
					return err
				case <-quit:
					log.Infof("Quit WatchTokensRescued (a): %v", err)
					return nil
				}
			case err := <-sub.Err():
				log.Errorf("Error with WatchTokensRescued, fatal (b): %v", err)
				w.errors <- err
				return err
			case <-quit:
				log.Infof("Quitting loop for WatchTokensRescued")
				return nil
			}
		}
	}), nil
}

func (w *CivilPLCRVotingContractWatchers) startWatchVoteCommitted(eventRecvChan chan *model.Event) (event.Subscription, error) {
	return event.NewSubscription(func(quit <-chan struct{}) error {
		startupFn := func() (event.Subscription, chan *contract.CivilPLCRVotingContractVoteCommitted, error) {
			for {
				opts := &bind.WatchOpts{}
				recvChan := make(chan *contract.CivilPLCRVotingContractVoteCommitted)
				log.Infof("startupFn: Starting WatchVoteCommitted")
				sub, err := w.contract.WatchVoteCommitted(
					opts,
					recvChan,
					[]*big.Int{},
					[]common.Address{},
				)
				if err != nil {
					if sub != nil {
						log.Infof("startupFn: Unsubscribing WatchVoteCommitted")
						sub.Unsubscribe()
					}
					return nil, nil, err
				}
				log.Infof("startupFn: WatchVoteCommitted started")
				return sub, recvChan, nil
			}
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchVoteCommitted: %v", err)
			if sub != nil {
				sub.Unsubscribe()
			}
			w.errors <- err
			return err
		}
		defer sub.Unsubscribe()
		log.Infof("Starting up WatchVoteCommitted for contract %v", w.contractAddress.Hex())
		for {
			select {
			// 15 min premptive resubscribe
			case <-time.After(time.Second * time.Duration(60*15)):
				log.Infof("Premptive restart of VoteCommitted")
				oldSub := sub
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("Error starting VoteCommitted: %v", err)
					w.errors <- err
					return err
				}
				log.Infof("Attempting to unsub old VoteCommitted")
				oldSub.Unsubscribe()
				log.Infof("Done preemptive restart VoteCommitted")
			case event := <-recvChan:
				if log.V(2) {
					log.Infof("Received event on WatchVoteCommitted: %v", spew.Sprintf("%#+v", event))
				} else {
					log.Info("Received event on WatchVoteCommitted")
				}
				modelEvent, err := model.NewEventFromContractEvent("VoteCommitted", w.ContractName(), w.contractAddress, event, ctime.CurrentEpochSecsInInt64(), model.Watcher)
				if err != nil {
					log.Errorf("Error creating new event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- modelEvent:
					if log.V(2) {
						log.Infof("Sent event to eventRecvChan on WatchVoteCommitted: %v", spew.Sprintf("%#+v", event))
					} else {
						log.Info("Sent event to eventRecvChan on WatchVoteCommitted")
					}
				case err := <-sub.Err():
					log.Errorf("Error with WatchVoteCommitted, fatal (a): %v", err)
					w.errors <- err
					return err
				case <-quit:
					log.Infof("Quit WatchVoteCommitted (a): %v", err)
					return nil
				}
			case err := <-sub.Err():
				log.Errorf("Error with WatchVoteCommitted, fatal (b): %v", err)
				w.errors <- err
				return err
			case <-quit:
				log.Infof("Quitting loop for WatchVoteCommitted")
				return nil
			}
		}
	}), nil
}

func (w *CivilPLCRVotingContractWatchers) startWatchVoteRevealed(eventRecvChan chan *model.Event) (event.Subscription, error) {
	return event.NewSubscription(func(quit <-chan struct{}) error {
		startupFn := func() (event.Subscription, chan *contract.CivilPLCRVotingContractVoteRevealed, error) {
			for {
				opts := &bind.WatchOpts{}
				recvChan := make(chan *contract.CivilPLCRVotingContractVoteRevealed)
				log.Infof("startupFn: Starting WatchVoteRevealed")
				sub, err := w.contract.WatchVoteRevealed(
					opts,
					recvChan,
					[]*big.Int{},
					[]*big.Int{},
					[]common.Address{},
				)
				if err != nil {
					if sub != nil {
						log.Infof("startupFn: Unsubscribing WatchVoteRevealed")
						sub.Unsubscribe()
					}
					return nil, nil, err
				}
				log.Infof("startupFn: WatchVoteRevealed started")
				return sub, recvChan, nil
			}
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchVoteRevealed: %v", err)
			if sub != nil {
				sub.Unsubscribe()
			}
			w.errors <- err
			return err
		}
		defer sub.Unsubscribe()
		log.Infof("Starting up WatchVoteRevealed for contract %v", w.contractAddress.Hex())
		for {
			select {
			// 15 min premptive resubscribe
			case <-time.After(time.Second * time.Duration(60*15)):
				log.Infof("Premptive restart of VoteRevealed")
				oldSub := sub
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("Error starting VoteRevealed: %v", err)
					w.errors <- err
					return err
				}
				log.Infof("Attempting to unsub old VoteRevealed")
				oldSub.Unsubscribe()
				log.Infof("Done preemptive restart VoteRevealed")
			case event := <-recvChan:
				if log.V(2) {
					log.Infof("Received event on WatchVoteRevealed: %v", spew.Sprintf("%#+v", event))
				} else {
					log.Info("Received event on WatchVoteRevealed")
				}
				modelEvent, err := model.NewEventFromContractEvent("VoteRevealed", w.ContractName(), w.contractAddress, event, ctime.CurrentEpochSecsInInt64(), model.Watcher)
				if err != nil {
					log.Errorf("Error creating new event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- modelEvent:
					if log.V(2) {
						log.Infof("Sent event to eventRecvChan on WatchVoteRevealed: %v", spew.Sprintf("%#+v", event))
					} else {
						log.Info("Sent event to eventRecvChan on WatchVoteRevealed")
					}
				case err := <-sub.Err():
					log.Errorf("Error with WatchVoteRevealed, fatal (a): %v", err)
					w.errors <- err
					return err
				case <-quit:
					log.Infof("Quit WatchVoteRevealed (a): %v", err)
					return nil
				}
			case err := <-sub.Err():
				log.Errorf("Error with WatchVoteRevealed, fatal (b): %v", err)
				w.errors <- err
				return err
			case <-quit:
				log.Infof("Quitting loop for WatchVoteRevealed")
				return nil
			}
		}
	}), nil
}

func (w *CivilPLCRVotingContractWatchers) startWatchVotingRightsGranted(eventRecvChan chan *model.Event) (event.Subscription, error) {
	return event.NewSubscription(func(quit <-chan struct{}) error {
		startupFn := func() (event.Subscription, chan *contract.CivilPLCRVotingContractVotingRightsGranted, error) {
			for {
				opts := &bind.WatchOpts{}
				recvChan := make(chan *contract.CivilPLCRVotingContractVotingRightsGranted)
				log.Infof("startupFn: Starting WatchVotingRightsGranted")
				sub, err := w.contract.WatchVotingRightsGranted(
					opts,
					recvChan,
					[]common.Address{},
				)
				if err != nil {
					if sub != nil {
						log.Infof("startupFn: Unsubscribing WatchVotingRightsGranted")
						sub.Unsubscribe()
					}
					return nil, nil, err
				}
				log.Infof("startupFn: WatchVotingRightsGranted started")
				return sub, recvChan, nil
			}
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchVotingRightsGranted: %v", err)
			if sub != nil {
				sub.Unsubscribe()
			}
			w.errors <- err
			return err
		}
		defer sub.Unsubscribe()
		log.Infof("Starting up WatchVotingRightsGranted for contract %v", w.contractAddress.Hex())
		for {
			select {
			// 15 min premptive resubscribe
			case <-time.After(time.Second * time.Duration(60*15)):
				log.Infof("Premptive restart of VotingRightsGranted")
				oldSub := sub
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("Error starting VotingRightsGranted: %v", err)
					w.errors <- err
					return err
				}
				log.Infof("Attempting to unsub old VotingRightsGranted")
				oldSub.Unsubscribe()
				log.Infof("Done preemptive restart VotingRightsGranted")
			case event := <-recvChan:
				if log.V(2) {
					log.Infof("Received event on WatchVotingRightsGranted: %v", spew.Sprintf("%#+v", event))
				} else {
					log.Info("Received event on WatchVotingRightsGranted")
				}
				modelEvent, err := model.NewEventFromContractEvent("VotingRightsGranted", w.ContractName(), w.contractAddress, event, ctime.CurrentEpochSecsInInt64(), model.Watcher)
				if err != nil {
					log.Errorf("Error creating new event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- modelEvent:
					if log.V(2) {
						log.Infof("Sent event to eventRecvChan on WatchVotingRightsGranted: %v", spew.Sprintf("%#+v", event))
					} else {
						log.Info("Sent event to eventRecvChan on WatchVotingRightsGranted")
					}
				case err := <-sub.Err():
					log.Errorf("Error with WatchVotingRightsGranted, fatal (a): %v", err)
					w.errors <- err
					return err
				case <-quit:
					log.Infof("Quit WatchVotingRightsGranted (a): %v", err)
					return nil
				}
			case err := <-sub.Err():
				log.Errorf("Error with WatchVotingRightsGranted, fatal (b): %v", err)
				w.errors <- err
				return err
			case <-quit:
				log.Infof("Quitting loop for WatchVotingRightsGranted")
				return nil
			}
		}
	}), nil
}

func (w *CivilPLCRVotingContractWatchers) startWatchVotingRightsWithdrawn(eventRecvChan chan *model.Event) (event.Subscription, error) {
	return event.NewSubscription(func(quit <-chan struct{}) error {
		startupFn := func() (event.Subscription, chan *contract.CivilPLCRVotingContractVotingRightsWithdrawn, error) {
			for {
				opts := &bind.WatchOpts{}
				recvChan := make(chan *contract.CivilPLCRVotingContractVotingRightsWithdrawn)
				log.Infof("startupFn: Starting WatchVotingRightsWithdrawn")
				sub, err := w.contract.WatchVotingRightsWithdrawn(
					opts,
					recvChan,
					[]common.Address{},
				)
				if err != nil {
					if sub != nil {
						log.Infof("startupFn: Unsubscribing WatchVotingRightsWithdrawn")
						sub.Unsubscribe()
					}
					return nil, nil, err
				}
				log.Infof("startupFn: WatchVotingRightsWithdrawn started")
				return sub, recvChan, nil
			}
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchVotingRightsWithdrawn: %v", err)
			if sub != nil {
				sub.Unsubscribe()
			}
			w.errors <- err
			return err
		}
		defer sub.Unsubscribe()
		log.Infof("Starting up WatchVotingRightsWithdrawn for contract %v", w.contractAddress.Hex())
		for {
			select {
			// 15 min premptive resubscribe
			case <-time.After(time.Second * time.Duration(60*15)):
				log.Infof("Premptive restart of VotingRightsWithdrawn")
				oldSub := sub
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("Error starting VotingRightsWithdrawn: %v", err)
					w.errors <- err
					return err
				}
				log.Infof("Attempting to unsub old VotingRightsWithdrawn")
				oldSub.Unsubscribe()
				log.Infof("Done preemptive restart VotingRightsWithdrawn")
			case event := <-recvChan:
				if log.V(2) {
					log.Infof("Received event on WatchVotingRightsWithdrawn: %v", spew.Sprintf("%#+v", event))
				} else {
					log.Info("Received event on WatchVotingRightsWithdrawn")
				}
				modelEvent, err := model.NewEventFromContractEvent("VotingRightsWithdrawn", w.ContractName(), w.contractAddress, event, ctime.CurrentEpochSecsInInt64(), model.Watcher)
				if err != nil {
					log.Errorf("Error creating new event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- modelEvent:
					if log.V(2) {
						log.Infof("Sent event to eventRecvChan on WatchVotingRightsWithdrawn: %v", spew.Sprintf("%#+v", event))
					} else {
						log.Info("Sent event to eventRecvChan on WatchVotingRightsWithdrawn")
					}
				case err := <-sub.Err():
					log.Errorf("Error with WatchVotingRightsWithdrawn, fatal (a): %v", err)
					w.errors <- err
					return err
				case <-quit:
					log.Infof("Quit WatchVotingRightsWithdrawn (a): %v", err)
					return nil
				}
			case err := <-sub.Err():
				log.Errorf("Error with WatchVotingRightsWithdrawn, fatal (b): %v", err)
				w.errors <- err
				return err
			case <-quit:
				log.Infof("Quitting loop for WatchVotingRightsWithdrawn")
				return nil
			}
		}
	}), nil
}
