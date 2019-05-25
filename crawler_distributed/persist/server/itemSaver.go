package main

import (
	"flag"
	"fmt"
	"go-crawler/crawler_distributed/config"
	"go-crawler/crawler_distributed/persist"
	"go-crawler/crawler_distributed/rpcsupport"
	"gopkg.in/olivere/elastic.v7"
)

var port = flag.Int("port", 0, "the port for me to listen")

/**
 the bash to start listen Port:
`go run itemSaver.go --port=1234`
*/
func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("Must specify a port")
		return
	}

	err := serverRpc(fmt.Sprintf(":%d", *port), config.ElasticSearchIndex)

	if err != nil {
		panic(err)
	}
}
func serverRpc(host string, index string) error {

	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
