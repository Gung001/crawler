package persist

import (
	"crawler.com/concurrent/engine"
	"crawler.com/concurrent/persist"
	"github.com/olivere/elastic/v7"
)

type ItemServerService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemServerService) Save(item engine.Item, result *string) error {
	save := persist.Save(s.Client, item, s.Index)
	if save == nil {
		*result = "ok"
	}
	return nil
}
