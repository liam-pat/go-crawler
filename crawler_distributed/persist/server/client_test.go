package main

import (
	"fmt"
	"go-crawler/crawler_distributed/rpcsupport"
	"go-crawler/engine"
	"go-crawler/model"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"
	var result string

	go serverRpc(host, "dating_profile")
	time.Sleep(1 * time.Second)

	client, err := rpcsupport.NewClient(host)

	if err != nil {
		panic(err)
	}

	item := engine.Item{
		Url:  "test.com",
		Type: "true_love",
		Id:   "weishenme",
		PayLoad: model.Profile{
			Name:       "test",
			Gender:     "男",
			Age:        "24",
			Height:     "110",
			Income:     "1.5w",
			Marriage:   "单身",
			Education:  "本科",
			Registered: "广东省",
			ImageUrl:   "",
		}}

	err = client.Call("ItemSaverService.Save", item, &result)
	if err != nil || result != "okay" {
		t.Errorf("result %s,err:%s", result, err)
	} else {
		fmt.Println(result)
	}
}
