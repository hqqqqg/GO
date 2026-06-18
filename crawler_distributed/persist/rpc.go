package persist

import (
	"google/crawler/engine"
	"google/crawler/persist"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

type ItemSaverService struct { //供RPC注册的Item保存服务
	Client *elasticsearch.Client
	Index  string
}

func (s *ItemSaverService) Save(
	item engine.Item, result *string) error {
	err := persist.Save(s.Client, s.Index, &item) //写入 ES
	log.Printf("Item %v saved.", item)
	if err == nil {
		*result = "ok"
	} else {
		log.Printf("Error saving item %v:%v",
			item, err)
	}
	return err
}
