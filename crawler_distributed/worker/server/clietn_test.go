package main

import (
	"fmt"
	"go-crawler/crawler_distributed/config"
	"go-crawler/crawler_distributed/rpcsupport"
	"go-crawler/crawler_distributed/worker"
	"testing"
	"time"
)

func TestCrawlerService(t *testing.T) {
	const host = ":2345"

	go rpcsupport.ServeRpc(host, worker.CrawlerService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url: "http://album.zhenai.com/u/1893685027",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: "朵儿",
		},
	}

	var result worker.ParseResult
	err = client.Call(config.CrawlerServiceRpc, req, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}

}
