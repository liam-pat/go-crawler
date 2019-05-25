package persist

import (
	"go-crawler/engine"
	"go-crawler/persist"
	"gopkg.in/olivere/elastic.v7"
	"log"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	res, err := persist.Save(s.Client, s.Index, item)

	log.Printf("res: %+v", res)

	if err == nil {
		log.Printf("Item %+v saved.\n", item)
		*result = "okay"
	} else {
		log.Printf("error saving error :  %v -> %v\n", item, err)
	}
	return err
}
