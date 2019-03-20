package crawlermain

import (
	"context"
	"fmt"

	log "github.com/golang/glog"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/retriever"
	"github.com/joincivil/civil-events-crawler/pkg/utils"
	"github.com/joincivil/go-common/pkg/eth"
)

const (
	dryRun = false
)

// StartUpRecrawler runs the recrawler based on the config
// The recrawler pulls events from Civil epoch, checks the DB for existing events,
// and adds them if they are missing. Can be configured to update existing events
// if their log meta data has changed.
func StartUpRecrawler(config *utils.CrawlerConfig) error {
	httpClient, err := setupHTTPEthClient(config)
	if err != nil {
		return err
	}

	log.Infof("Starting to filter at block number %v", config.EthStartBlock)
	if log.V(2) {
		header, logErr := httpClient.HeaderByNumber(context.TODO(), nil)
		if logErr == nil {
			log.Infof("Latest block number is: %v", header.Number)
		}
		// If v info level logging, include the ethereum lib logging
		enableGoEthereumLogging()
	}

	// Update start block from the configured start
	filterers := contractFilterers(config)
	for _, filterer := range filterers {
		for _, etype := range filterer.EventTypes() {
			filterer.UpdateStartBlock(etype, config.EthStartBlock+1)
		}
	}

	epersister := eventDataPersister(config)

	eventRetriever := retriever.NewEventRetriever(httpClient, filterers)
	err = eventRetriever.Retrieve()
	if err != nil {
		log.Errorf("Error retrieving events: err: %v", err)
	}

	evs := eventRetriever.PastEvents

	eventsToAppend := []*model.Event{}
	var headerCache *eth.BlockHeaderCache
	retryChain := eth.RetryChainReader{ChainReader: httpClient}

	for _, ev := range evs {
		criteria := &model.RetrieveEventsCriteria{
			TxHash: ev.TxHash().Hex(),
		}

		events, _ := epersister.RetrieveEvents(criteria)
		if len(events) > 0 {
			for _, event := range events {
				if event.EventType() != ev.EventType() {
					// log.Infof("FOUND: %v, %v, %v, %v", event.Hash(), event.EventType(), event.Timestamp(), event.LogPayloadToString())
					continue
				}

				if ev.TxIndex() != event.TxIndex() ||
					ev.LogIndex() != event.LogIndex() ||
					ev.BlockNumber() != event.BlockNumber() ||
					ev.BlockHash() != event.BlockHash() {
					log.Infof("FOUND VIA TXHASH: check: %v, %v, %v, %v", ev.Hash(), ev.EventType(), ev.Timestamp(), ev.LogPayloadToString())
					log.Infof("FOUND VIA TXHASH: found: %v, %v, %v, %v", event.Hash(), event.EventType(), event.Timestamp(), event.LogPayloadToString())
					log.Infof("Log data is not the same, should be updated")
					break
				}
			}
			continue
		}

		err := updateEventTimeFromBlockHeader(retryChain, headerCache, ev)
		if err != nil {
			log.Errorf("err = %v", err)
			continue
		}

		eventsToAppend = append(eventsToAppend, ev)
	}

	for _, event := range eventsToAppend {
		log.Infof("Event to add: %v, %v, %v, %v", event.EventType(), event.Timestamp(), event.LogPayloadToString(), event.Timestamp())
	}

	if !dryRun && len(eventsToAppend) > 0 {
		err = epersister.SaveEvents(eventsToAppend)
		if err != nil {
			return fmt.Errorf("Error saving events: err: %v", err)
		}
	} else {
		log.Infof("DRY RUN or no events, did not save")
	}

	return nil
}
