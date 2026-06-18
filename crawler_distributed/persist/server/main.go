package main

import (
	"google/crawler_distributed/persist"
	"google/crawler_distributed/rpcsupport"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

func main() {
	log.Fatal(serveRpc(":1234", "dating_profile"))
}

func serveRpc(host, index string) error {
	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		panic(err)
	}
	return rpcsupport.ServeRpc(host,
		&persist.ItemSaverService{
			Client: client,
			Index:  index,
		})
}
