package client

import (
	"go-crawler/crawler_distributed/config"
	"go-crawler/crawler_distributed/worker"
	"go-crawler/engine"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {

	return func(request engine.Request) (result engine.ParseResult, e error) {
		sReq := worker.SerializeRequest(request)

		var sResult worker.ParseResult

		c := <-clientChan
		err := c.Call(config.CrawlerServiceRpc, sReq, &sResult)

		if err != nil {
			return engine.ParseResult{}, err
		}

		return worker.DeSerializeResult(sResult), nil
	}
}
