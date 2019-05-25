package main

import (
	"fmt"
	"go-crawler/crawler_distributed/config"
	"go-crawler/crawler_distributed/persist"
	"go-crawler/crawler_distributed/rpcsupport"
	"gopkg.in/olivere/elastic.v7"
)

func main() {
	err := serverRpc(fmt.Sprintf(":%s", config.ItemSaverPort), config.ElasticSearchIndex)

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
