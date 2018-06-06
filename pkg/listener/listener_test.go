// Package listener_test contains tests for the listener package
package listener_test

import (
	"math/big"
	"runtime"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	cutils "github.com/joincivil/civil-events-crawler/pkg/contractutils"
	"github.com/joincivil/civil-events-crawler/pkg/generated/watcher"
	"github.com/joincivil/civil-events-crawler/pkg/listener"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/utils"
)

func TestCivilListener(t *testing.T) {
	contracts, err := cutils.SetupAllTestContracts()
	if err != nil {
		t.Fatalf("Unable to setup the contracts: %v", err)
	}
	listener := setupListener(t, contracts.Client, contracts.CivilTcrAddr.Hex())
	defer listener.Stop()
}

func TestCivilListenerStop(t *testing.T) {
	contracts, err := cutils.SetupAllTestContracts()
	if err != nil {
		t.Fatalf("Unable to setup the contracts: %v", err)
	}
	listener := setupListener(t, contracts.Client, contracts.CivilTcrAddr.Hex())

	// Simple test is if each watcher loop goroutine is shut down by Stop()
	initialNumRoutines := runtime.NumGoroutine()
	listener.Stop()
	if initialNumRoutines <= runtime.NumGoroutine() {
		t.Errorf("Number of goroutines has not gone down since listener.Stop")
	}
}

// TestCivilListenerEventChan mainly tests the EventRecvChan to ensure it can
// pass along a CivilEvent object
func TestCivilListenerEventChan(t *testing.T) {
	contracts, err := cutils.SetupAllTestContracts()
	if err != nil {
		t.Fatalf("Unable to setup the contracts: err: %v", err)
	}
	listener := setupListener(t, contracts.Client, contracts.CivilTcrAddr.Hex())
	defer listener.Stop()
	quitChan := make(chan interface{})
	eventRecv := make(chan bool)

	go func(quit <-chan interface{}, recv chan<- bool) {
		for {
			select {
			case event := <-listener.EventRecvChan:
				if event.EventType != "_Application" {
					t.Errorf("Eventtype is not correct: %v", event.EventType)
				}
				recv <- true
			case <-quit:
				return
			}
		}
	}(quitChan, eventRecv)

	newEvent := &model.CivilEvent{
		EventType: "_Application",
		Timestamp: utils.CurrentEpochSecsInInt(),
		Payload:   &model.CivilEventPayload{},
	}
	listener.EventRecvChan <- *newEvent

	select {
	case <-eventRecv:
		close(quitChan)
	case <-time.After(5 * time.Second):
		t.Errorf("Event not received, should have been received by goroutine")
		close(quitChan)
	}
}

// TestCivilListenerContractEvents tests event fired from a call to Apply()
// on a simulated TCR on a simulated backend. Tests two events so ensure
// we are handling two different events on the same channel.
func TestCivilListenerContractEvents(t *testing.T) {
	contracts, err := cutils.SetupAllTestContracts()
	if err != nil {
		t.Fatalf("Unable to setup the contracts: err: %v", err)
	}
	listener := setupListener(t, contracts.Client, contracts.CivilTcrAddr.Hex())
	defer listener.Stop()

	quitChan := make(chan interface{})
	eventRecv := make(chan bool)

	go func(quit <-chan interface{}, recv chan<- bool) {
		for {
			select {
			case event := <-listener.EventRecvChan:
				if event.EventType != "_Application" && event.EventType != "_Withdrawal" &&
					event.EventType != "_Deposit" {
					t.Errorf("EventType is not correct: eventType: %v", event.EventType)
				}
				recv <- true
			case <-quit:
				return
			}
		}
	}(quitChan, eventRecv)

	expectedNumEvents := 3

	_, err = contracts.CivilTcrContract.Apply(contracts.Auth, contracts.NewsroomAddr, big.NewInt(400), "")
	if err != nil {
		t.Fatalf("Application failed: err: %v", err)
	}

	contracts.Client.Commit()

	_, err = contracts.CivilTcrContract.Withdraw(contracts.Auth, contracts.NewsroomAddr, big.NewInt(50))
	if err != nil {
		t.Fatalf("Withdrawal failed: err: %v", err)
	}

	contracts.Client.Commit()

	_, err = contracts.CivilTcrContract.Deposit(contracts.Auth, contracts.NewsroomAddr, big.NewInt(50))
	if err != nil {
		t.Fatalf("Deposit failed: err: %v", err)
	}

	contracts.Client.Commit()

	numEvents := 0
Loop:
	for {
		select {
		case <-eventRecv:
			numEvents++
			if numEvents == expectedNumEvents {
				close(quitChan)
				break Loop
			}
		case <-time.After(20 * time.Second):
			t.Errorf("Not all events were received")
			close(quitChan)
			break Loop
		}
	}
}

func setupListener(t *testing.T, client bind.ContractBackend, address string) *listener.CivilEventListener {
	watchers := []model.ContractWatchers{
		&watcher.CivilTCRContractWatchers{},
		&watcher.NewsroomContractWatchers{},
	}
	listener := listener.NewCivilEventListener(client, address, watchers)
	if listener == nil {
		t.Fatal("Listener should not be nil")
	}
	err := listener.Start()
	if err != nil {
		t.Errorf("Listener should have started with no errors: %v", err)
	}
	return listener
}