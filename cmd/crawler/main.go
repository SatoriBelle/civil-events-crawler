// Package main contains commands to run
package main

import (
	"flag"
	"os"

	log "github.com/golang/glog"

	"github.com/joincivil/civil-events-crawler/pkg/crawlermain"
	"github.com/joincivil/civil-events-crawler/pkg/utils"
)

func main() {
	config := &utils.CrawlerConfig{}
	flag.Usage = func() {
		config.OutputUsage()
		os.Exit(0)
	}
	flag.Parse()

	err := config.PopulateFromEnv()
	if err != nil {
		config.OutputUsage()
		log.Errorf("Invalid crawler config: err: %v\n", err)
		os.Exit(2)
	}

	err = crawlermain.StartUpCrawler(config)
	if err != nil {
		log.Errorf("Crawler error: err: %v\n", err)
		os.Exit(2)
	}
	log.Info("Crawler stopped")
}
