package main

import (
	"go-crawler/crawler_distributed/persist/client"
	"go-crawler/engine"
	"go-crawler/scheduler"
	"go-crawler/sites/com.zhenai/parser"
)

func main() {
	itemChan, err := client.ItemSaver(":1234")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkCount: 10,
		ItemChan:  itemChan,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: parser.ParseCity,
	//})
}
