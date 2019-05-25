package main

import (
	"fmt"
	"go-crawler/crawler_distributed/config"
	"go-crawler/crawler_distributed/rpcsupport"
	"go-crawler/crawler_distributed/worker"
)

func main() {
	err := rpcsupport.ServeRpc(
		fmt.Sprintf(":%s", config.WorkPort),
		worker.CrawlerService{},
	)

	if err != nil {
		panic(err)
	}
}
