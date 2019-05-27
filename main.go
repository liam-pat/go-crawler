package main

import (
	"fmt"
	"go-crawler/config"
	"go-crawler/engine"
	"go-crawler/persist"
	"go-crawler/scheduler"
	"go-crawler/sites/com.zhenai/parser"
)

func main() {
	// new a  goroutine to waiting the itemChan to input item to save
	itemChan, err := persist.ItemSaver(fmt.Sprintf("%s", basic_config.ElasticsearchIndex))

	if err != nil {
		panic(err)
	}

	concurrentEngine := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkCount:        10,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}
	concurrentEngine.Run(
		engine.Request{
			Url:    "http://www.zhenai.com/zhenghun",
			Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
		})

	//concurrentEngine.Run(engine.Request{
	//	Url:    "http://www.zhenai.com/zhenghun/shanghai",
	//	Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	//})
}
