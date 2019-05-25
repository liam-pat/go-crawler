package persist

import (
	"context"
	"encoding/json"
	"go-crawler/engine"
	"go-crawler/model"
	"testing"
)
import "gopkg.in/olivere/elastic.v7"

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url:  "test.com",
		Type: "true_love",
		Id:   "22222",
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

	const index = "dating_profile"

	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	_, err = Save(client, index, expected)

	if err != nil {
		panic(err)
	}

	response, err := client.Get().Index(index).Type(expected.Type).Id(expected.Id).Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("%s", response.Source)

	var actual engine.Item
	err = json.Unmarshal(response.Source, &actual)

	if err != nil {
		panic(err)
	}

	actualProfile, _ := model.FromJsonObj(actual.PayLoad)

	actual.PayLoad = actualProfile

	if actual != expected {
		t.Errorf("got %v ;expect %v", actual, expected)
	}
}
