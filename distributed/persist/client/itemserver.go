package client

import (
	"crawler.com/concurrent/engine"
	"crawler.com/distributed/config"
	"crawler.com/distributed/rpcsupport"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error) {

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("ItemSaver got item #%d  %v", itemCount, item)
			itemCount++

			result := ""
			err := client.Call(config.ItemServerRpc, item, &result)
			if err != nil {
				log.Printf("ItemSaver saving error %v %v", err, item)
			}
		}
	}()
	return out, nil
}
