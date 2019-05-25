package main

import (
	"flag"
	"go-crawler/crawler_distributed/config"
	"go-crawler/crawler_distributed/persist/client"
	"go-crawler/crawler_distributed/rpcsupport"
	client2 "go-crawler/crawler_distributed/worker/client"
	"go-crawler/engine"
	"go-crawler/scheduler"
	"go-crawler/sites/com.zhenai/parser"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String("itemSaver_host", "", "itemSaver host")
	workerHosts   = flag.String("worker_hosts", "", "work hosts (comma separated)")
)

/**
you can use the bash
eg. `go run main.go -itemSaver_host=":1234" -worker_hosts=":9000,:9001"`
you should start the itemSaver.go and worker.go(one or more)
*/
func main() {
	flag.Parse()
	itemChan, err := client.ItemSaver(*itemSaverHost)

	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))

	processor := client2.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkCount:        100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})
	//e.Run(engine.Request{
	//	Url:    "http://www.zhenai.com/zhenghun/shanghai",
	//	Parser: engine.NewFuncParser(parser.ParseCity, config.ParseCity),
	//})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client

	for _, h := range hosts {
		rpcClient, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, rpcClient)
			log.Printf("connected to %s", h)
		} else {
			log.Printf("Error connection to %s: %v", h, err)
		}
	}
	out := make(chan *rpc.Client)

	go func() {
		for {
			for _, clientSample := range clients {
				out <- clientSample
			}
		}
	}()

	return out
}
