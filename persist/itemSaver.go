package persist

import (
	"context"
	"github.com/pkg/errors"
	"go-crawler/engine"
	"log"
)
import "gopkg.in/olivere/elastic.v7"

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("item saver : got Item #%d: %v", itemCount, item)
			itemCount++

			err := save(client,index, item)

			if err != nil {
				log.Printf("item saver error: item %s ;error: %s", item, err)
			}
		}
	}()
	return out, nil
}

func save(client *elastic.Client, index string, item engine.Item) (err error) {

	if item.Type == "" {
		return errors.New("Type is nil")
	}

	idService := client.Index().Index("dating_profile").Type(item.Type).BodyJson(item)

	if item.Id != "" {
		idService.Id(item.Id)
	}

	_, err = idService.Do(context.Background())

	if err != nil {
		return err
	}

	return nil
}
