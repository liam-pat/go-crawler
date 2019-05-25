package main

import (
	"go-crawler/crawler_distributed/config"
	"go-crawler/crawler_distributed/persist/client"
	client2 "go-crawler/crawler_distributed/worker/client"
	"go-crawler/engine"
	"go-crawler/scheduler"
	"go-crawler/sites/com.zhenai/parser"
)

func main() {
	itemChan, err := client.ItemSaver(":1234")
	if err != nil {
		panic(err)
	}

	processor, err := client2.CreateProcessor()

	if err != nil {
		panic(err)
	}

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
