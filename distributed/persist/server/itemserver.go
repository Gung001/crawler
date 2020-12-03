package main

import (
	"crawler.com/distributed/config"
	"crawler.com/distributed/persist"
	"crawler.com/distributed/rpcsupport"
	"flag"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Printf("must specify a port")
		return
	}
	log.Fatal(saverRpc(
		fmt.Sprintf(":%d", *port),
		config.ElasticIndex))
}

func saverRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(host, &persist.ItemServerService{
		Client: client,
		Index:  index,
	})
}
