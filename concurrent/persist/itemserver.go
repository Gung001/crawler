package persist

import (
	"context"
	"crawler.com/concurrent/engine"
	"errors"
	"github.com/olivere/elastic/v7"
	"log"
)

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
			log.Printf("ItemSaver got item #%d  %v", itemCount, item)
			itemCount++
			err := Save(client, item, index)
			if err != nil {
				log.Printf("ItemSaver saving error %v %v", err, item)
			}
		}
	}()
	return out, nil
}

func Save(client *elastic.Client, item engine.Item, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil
	}

	if item.Type == "" {
		return errors.New("must supply Type")
	}

	if item.Id == "" {
		return errors.New("must supply Id")
	}

	_, err = client.Index().
		Index(index).
		Type(item.Type).
		Id(item.Id).
		BodyJson(item).Do(context.Background())
	if err != nil {
		return nil
	}

	return nil
}
