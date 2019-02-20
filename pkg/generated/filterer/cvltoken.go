// Code generated by 'gen/eventhandlergen.go'  DO NOT EDIT.
// IT SHOULD NOT BE EDITED BY HAND AS ANY CHANGES MAY BE OVERWRITTEN
// Please reference 'gen/filterergen_template.go' for more details
// File was generated at 2019-02-19 22:10:18.059862 +0000 UTC
package filterer

import (
	log "github.com/golang/glog"
	"runtime"
	"sync"

	"github.com/Jeffail/tunny"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	commongen "github.com/joincivil/civil-events-crawler/pkg/generated/common"
	"github.com/joincivil/civil-events-crawler/pkg/model"

	"github.com/joincivil/go-common/pkg/generated/contract"
	ctime "github.com/joincivil/go-common/pkg/time"
)

func NewCVLTokenContractFilterers(contractAddress common.Address) *CVLTokenContractFilterers {
	contractFilterers := &CVLTokenContractFilterers{
		contractAddress:   contractAddress,
		eventTypes:        commongen.EventTypesCVLTokenContract(),
		eventToStartBlock: make(map[string]uint64),
		lastEvents:        make([]*model.Event, 0),
	}
	for _, eventType := range contractFilterers.eventTypes {
		contractFilterers.eventToStartBlock[eventType] = 0
	}
	return contractFilterers
}

type CVLTokenContractFilterers struct {
	contractAddress   common.Address
	contract          *contract.CVLTokenContract
	eventTypes        []string
	eventToStartBlock map[string]uint64
	lastEvents        []*model.Event
	lastEventsMutex   sync.Mutex
	pastEventsMutex   sync.Mutex
}

func (f *CVLTokenContractFilterers) ContractName() string {
	return "CVLTokenContract"
}

func (f *CVLTokenContractFilterers) ContractAddress() common.Address {
	return f.contractAddress
}

func (f *CVLTokenContractFilterers) StartFilterers(client bind.ContractBackend, pastEvents []*model.Event) (error, []*model.Event) {
	return f.StartCVLTokenContractFilterers(client, pastEvents)
}

func (f *CVLTokenContractFilterers) EventTypes() []string {
	return f.eventTypes
}

func (f *CVLTokenContractFilterers) UpdateStartBlock(eventType string, startBlock uint64) {
	f.eventToStartBlock[eventType] = startBlock
}

func (f *CVLTokenContractFilterers) LastEvents() []*model.Event {
	return f.lastEvents
}

// StartCVLTokenContractFilterers retrieves events for CVLTokenContract
func (f *CVLTokenContractFilterers) StartCVLTokenContractFilterers(client bind.ContractBackend, pastEvents []*model.Event) (error, []*model.Event) {
	contract, err := contract.NewCVLTokenContract(f.contractAddress, client)
	if err != nil {
		log.Errorf("Error initializing StartCVLTokenContract: err: %v", err)
		return err, pastEvents
	}
	f.contract = contract

	workerMultiplier := 1
	numWorkers := runtime.NumCPU() * workerMultiplier
	log.Infof("Num of workers: %v", numWorkers)
	pool := tunny.NewFunc(numWorkers, func(payload interface{}) interface{} {
		f := payload.(func())
		f()
		return nil
	})
	defer pool.Close()

	wg := sync.WaitGroup{}
	resultsChan := make(chan []*model.Event)
	done := make(chan bool)

	wg.Add(1)
	go func() {
		filterFunc := func() {
			startBlock := f.eventToStartBlock["Approval"]
			e, pevents := f.startFilterApproval(startBlock, []*model.Event{})
			if e != nil {
				log.Errorf("Error retrieving Approval: err: %v", e)
				return
			}
			if len(pevents) > 0 {
				f.lastEventsMutex.Lock()
				f.lastEvents = append(f.lastEvents, pevents[len(pevents)-1])
				f.lastEventsMutex.Unlock()
				resultsChan <- pevents
			}
		}
		pool.Process(filterFunc)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		filterFunc := func() {
			startBlock := f.eventToStartBlock["OwnershipRenounced"]
			e, pevents := f.startFilterOwnershipRenounced(startBlock, []*model.Event{})
			if e != nil {
				log.Errorf("Error retrieving OwnershipRenounced: err: %v", e)
				return
			}
			if len(pevents) > 0 {
				f.lastEventsMutex.Lock()
				f.lastEvents = append(f.lastEvents, pevents[len(pevents)-1])
				f.lastEventsMutex.Unlock()
				resultsChan <- pevents
			}
		}
		pool.Process(filterFunc)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		filterFunc := func() {
			startBlock := f.eventToStartBlock["OwnershipTransferred"]
			e, pevents := f.startFilterOwnershipTransferred(startBlock, []*model.Event{})
			if e != nil {
				log.Errorf("Error retrieving OwnershipTransferred: err: %v", e)
				return
			}
			if len(pevents) > 0 {
				f.lastEventsMutex.Lock()
				f.lastEvents = append(f.lastEvents, pevents[len(pevents)-1])
				f.lastEventsMutex.Unlock()
				resultsChan <- pevents
			}
		}
		pool.Process(filterFunc)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		filterFunc := func() {
			startBlock := f.eventToStartBlock["Transfer"]
			e, pevents := f.startFilterTransfer(startBlock, []*model.Event{})
			if e != nil {
				log.Errorf("Error retrieving Transfer: err: %v", e)
				return
			}
			if len(pevents) > 0 {
				f.lastEventsMutex.Lock()
				f.lastEvents = append(f.lastEvents, pevents[len(pevents)-1])
				f.lastEventsMutex.Unlock()
				resultsChan <- pevents
			}
		}
		pool.Process(filterFunc)
		wg.Done()
	}()

	go func() {
		wg.Wait()
		done <- true
		log.Info("Filtering routines complete")
	}()

Loop:
	for {
		select {
		case <-done:
			break Loop
		case pevents := <-resultsChan:
			f.pastEventsMutex.Lock()
			pastEvents = append(pastEvents, pevents...)
			f.pastEventsMutex.Unlock()
		}
	}
	log.Infof("Total events found: %v", len(pastEvents))
	return nil, pastEvents
}

func (f *CVLTokenContractFilterers) startFilterApproval(startBlock uint64, pastEvents []*model.Event) (error, []*model.Event) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering events for Approval for contract %v starting at block %v", f.contractAddress.Hex(), startBlock)
	itr, err := f.contract.FilterApproval(
		opts,
		[]common.Address{},
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error getting event Approval: %v", err)
		return err, pastEvents
	}
	beforeCount := len(pastEvents)
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("Approval", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	numEventsAdded := len(pastEvents) - beforeCount
	log.Infof("Approval events added: %v", numEventsAdded)
	return nil, pastEvents
}

func (f *CVLTokenContractFilterers) startFilterOwnershipRenounced(startBlock uint64, pastEvents []*model.Event) (error, []*model.Event) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering events for OwnershipRenounced for contract %v starting at block %v", f.contractAddress.Hex(), startBlock)
	itr, err := f.contract.FilterOwnershipRenounced(
		opts,
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error getting event OwnershipRenounced: %v", err)
		return err, pastEvents
	}
	beforeCount := len(pastEvents)
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("OwnershipRenounced", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	numEventsAdded := len(pastEvents) - beforeCount
	log.Infof("OwnershipRenounced events added: %v", numEventsAdded)
	return nil, pastEvents
}

func (f *CVLTokenContractFilterers) startFilterOwnershipTransferred(startBlock uint64, pastEvents []*model.Event) (error, []*model.Event) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering events for OwnershipTransferred for contract %v starting at block %v", f.contractAddress.Hex(), startBlock)
	itr, err := f.contract.FilterOwnershipTransferred(
		opts,
		[]common.Address{},
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error getting event OwnershipTransferred: %v", err)
		return err, pastEvents
	}
	beforeCount := len(pastEvents)
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("OwnershipTransferred", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	numEventsAdded := len(pastEvents) - beforeCount
	log.Infof("OwnershipTransferred events added: %v", numEventsAdded)
	return nil, pastEvents
}

func (f *CVLTokenContractFilterers) startFilterTransfer(startBlock uint64, pastEvents []*model.Event) (error, []*model.Event) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	log.Infof("Filtering events for Transfer for contract %v starting at block %v", f.contractAddress.Hex(), startBlock)
	itr, err := f.contract.FilterTransfer(
		opts,
		[]common.Address{},
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error getting event Transfer: %v", err)
		return err, pastEvents
	}
	beforeCount := len(pastEvents)
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("Transfer", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	numEventsAdded := len(pastEvents) - beforeCount
	log.Infof("Transfer events added: %v", numEventsAdded)
	return nil, pastEvents
}
