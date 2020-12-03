package main

import (
	"crawler.com/concurrent/engine"
	"crawler.com/concurrent/scheduler"
	"crawler.com/concurrent/zhenai/parser"
	"crawler.com/distributed/config"
	"crawler.com/distributed/persist/client"
	"crawler.com/distributed/rpcsupport"
	"flag"
	"fmt"
	"log"
	"net/rpc"
	"strings"
)

const url = "http://localhost:8080/mock/www.zhenai.com/zhenghun"
const urlShanghai = "http://localhost:8080/mock/www.zhenai.com/zhenghun/suining2"

var (
	itemServerHost = flag.Int(
		"item_server_host", 0, "the port for me to listen on")
	workerHosts = flag.String(
		"worker_hosts", "", "comma separated")
)

func main() {
	flag.Parse()
	itemChan, err := client.ItemSaver(
		fmt.Sprintf(":%d", *itemServerHost))
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))
	processor := client.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}

	e.Run(engine.Request{
		Url:    url,
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})

	//e.Run(engine.Request{
	//	Url:       urlShanghai,
	//	ParserFunc: parser.ParseCity,
	//})
}

func createClientPool(host []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range host {
		client, err := rpcsupport.NewClient(h)
		if err != nil {
			log.Printf("Error connecting to %s : %v", h, err)
		} else {
			clients = append(clients, client)
			log.Printf("connected to %s", h)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, c := range clients {
				out <- c
			}
		}
	}()
	return out
}
