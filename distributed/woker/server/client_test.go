package main

import (
	"crawler.com/distributed/config"
	"crawler.com/distributed/rpcsupport"
	"crawler.com/distributed/woker"
	"fmt"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"

	go rpcsupport.ServeRpc(host, woker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := woker.Request{
		Url: "http://localhost:8080/mock/album.zhenai.com/u/8796617497327185714",
		Parser: woker.SerializedParser{
			Name: config.ParseProfile,
			Args: "独久厌闹草莓裙摆",
		},
	}

	result := woker.ParseResult{}
	err = client.Call(config.ItemCrawlServiceRpc, req, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
