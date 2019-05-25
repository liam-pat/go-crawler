package main

import (
	"flag"
	"fmt"
	"go-crawler/crawler_distributed/rpcsupport"
	"go-crawler/crawler_distributed/worker"
)

var port = flag.Int("port", 0, "the port for me to listen")

/**
 the bash to start listen Port:
`go run worker.go --port=9000`
*/
func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("Must specify a port")
		return
	}
	err := rpcsupport.ServeRpc(
		fmt.Sprintf(":%d", *port),
		worker.CrawlerService{},
	)

	if err != nil {
		panic(err)
	}
}
