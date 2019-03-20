package crawlermain

import (
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	elog "github.com/ethereum/go-ethereum/log"
	log "github.com/golang/glog"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joincivil/civil-events-crawler/pkg/eventcollector"
	"github.com/joincivil/civil-events-crawler/pkg/generated/handlerlist"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/utils"
	"github.com/joincivil/go-common/pkg/eth"
)

const (
	websocketPingDelaySecs = 10 // 10 secs
)

// contractFilterers returns the filterers based on the config
func contractFilterers(config *utils.CrawlerConfig) []model.ContractFilterers {
	return handlerlist.ContractFilterers(config.ContractAddressObjs)
}

// contractWatchers returns the watchers based on the config
func contractWatchers(config *utils.CrawlerConfig) []model.ContractWatchers {
	return handlerlist.ContractWatchers(config.ContractAddressObjs)
}

// eventTriggers returns the triggers based on the config
func eventTriggers(config *utils.CrawlerConfig) []eventcollector.Trigger {
	return []eventcollector.Trigger{
		&eventcollector.AddNewsroomWatchersTrigger{},
		&eventcollector.RemoveNewsroomWatchersTrigger{},
	}
}

// isWebsocketURL returns true if the URL indicates it is for websockets
func isWebsocketURL(rawurl string) bool {
	u, err := url.Parse(rawurl)
	if err != nil {
		log.Infof("Unable to parse URL: err: %v", err)
		return false
	}
	if u.Scheme == "ws" || u.Scheme == "wss" {
		return true
	}
	return false
}

// setupHTTPEthClient sets up the HTTP ethclient based on the config
func setupHTTPEthClient(config *utils.CrawlerConfig) (*ethclient.Client, error) {
	if isWebsocketURL(config.EthAPIURL) {
		return nil, fmt.Errorf(
			"Fatal: Valid HTTP eth client URL required: configured url: %v",
			config.EthAPIURL,
		)
	}

	client, err := ethclient.Dial(config.EthAPIURL)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// setupWebsocketEthClient sets up the WS ethclient based on the config
func setupWebsocketEthClient(config *utils.CrawlerConfig, killChan <-chan bool) (*ethclient.Client, error) {
	if config.EthWsAPIURL == "" {
		return nil, nil
	}

	if !isWebsocketURL(config.EthWsAPIURL) {
		return nil, nil
	}

	client, err := ethclient.Dial(config.EthWsAPIURL)
	if err != nil {
		return nil, err
	}

	go eth.WebsocketPing(client, killChan, websocketPingDelaySecs)

	return client, nil
}

// enableGoEthereumLogging enables logging on the go-etheruem library
func enableGoEthereumLogging() {
	glog := elog.NewGlogHandler(elog.StreamHandler(os.Stderr, elog.TerminalFormat(false)))
	glog.Verbosity(elog.Lvl(elog.LvlDebug)) // nolint: unconvert
	glog.Vmodule("")                        // nolint: errcheck, gosec
	elog.Root().SetHandler(glog)
}

func cleanup(eventCol *eventcollector.EventCollector, killChan chan<- bool) {
	log.Info("Stopping crawler...")
	err := eventCol.StopCollection()
	if err != nil {
		log.Errorf("Error stopping collection: err: %v", err)
	}
	close(killChan)
	log.Info("Crawler stopped")
}

func setupKillNotify(eventCol *eventcollector.EventCollector, killChan chan<- bool) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup(eventCol, killChan)
		os.Exit(1)
	}()
}
