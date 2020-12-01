package persist

import (
	"crawler.com/concurrent/engine"
	"crawler.com/concurrent/model"
	"github.com/olivere/elastic/v7"
	"testing"
)

func TestSave(t *testing.T) {
	profile := model.Profile{
		Name:       "心事痕迹迁就",
		Age:        39,
		Gender:     "男",
		Height:     185,
		Weight:     200,
		Income:     "8001-10000元",
		Marriage:   "离异",
		Education:  "大学",
		Occupation: "程序员",
		Hokou:      "西安市",
		Xinzuo:     "双子座",
		House:      "租房",
		Car:        "无车",
	}

	item := engine.Item{
		Id:      "180900910",
		Url:     "http://localhost:8080/mock/album.zhenai.com/u/180900910",
		Type:    "zhenai",
		Payload: profile,
	}

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return
	}
	const index = "dating_profile1"
	save(client, item, index)
}
