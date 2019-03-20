package crawlermain

import (
	"context"

	log "github.com/golang/glog"
	"github.com/joincivil/civil-events-crawler/pkg/eventcollector"
	"github.com/joincivil/civil-events-crawler/pkg/utils"
)

// StartUpCrawler starts up the crawler process
func StartUpCrawler(config *utils.CrawlerConfig) error {
	killChan := make(chan bool)

	httpClient, err := setupHTTPEthClient(config)
	if err != nil {
		return err
	}

	wsClient, err := setupWebsocketEthClient(config, killChan)
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

	eventCol := eventcollector.NewEventCollector(
		&eventcollector.Config{
			Chain:              httpClient,
			HTTPClient:         httpClient,
			WsClient:           wsClient,
			Filterers:          contractFilterers(config),
			Watchers:           contractWatchers(config),
			RetrieverPersister: retrieverMetaDataPersister(config),
			ListenerPersister:  listenerMetaDataPersister(config),
			EventDataPersister: eventDataPersister(config),
			Triggers:           eventTriggers(config),
			StartBlock:         config.EthStartBlock,
			CrawlerPubSub:      crawlerPubSub(config),
		},
	)

	// Setup shutdown/cleanup hooks
	setupKillNotify(eventCol, killChan)
	defer func() {
		cleanup(eventCol, killChan)
	}()

	log.Info("Crawler starting...")

	// Will block here while handling collection
	return eventCol.StartCollection()
}
