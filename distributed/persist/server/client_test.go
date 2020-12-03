package main

import (
	"crawler.com/concurrent/engine"
	"crawler.com/concurrent/model"
	"crawler.com/distributed/config"
	"crawler.com/distributed/rpcsupport"
	"fmt"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {

	go saverRpc(fmt.Sprintf(":%d", config.ItemServerPort), "test2")
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(fmt.Sprintf(":%d", config.ItemServerPort))
	if err != nil {
		panic(err)
	}

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

	result := ""
	err = client.Call(config.ItemServerRpc, item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result %s , err %v", result, err)
	}
}
