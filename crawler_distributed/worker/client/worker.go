package client

import (
	"fmt"
	"go-crawler/crawler_distributed/config"
	"go-crawler/crawler_distributed/rpcsupport"
	"go-crawler/crawler_distributed/worker"
	"go-crawler/engine"
)

func CreateProcessor() (engine.Processor, error) {

	client, err := rpcsupport.NewClient(fmt.Sprintf(":%s", config.WorkPort))

	if err != nil {
		return nil, err
	}

	return func(request engine.Request) (result engine.ParseResult, e error) {
		sReq := worker.SerializeRequest(request)

		var sResult worker.ParseResult

		err := client.Call(config.CrawlerServiceRpc, sReq, &sResult)

		if err != nil {
			return engine.ParseResult{}, err
		}

		return worker.DeSerializeResult(sResult), nil
	}, nil
}
