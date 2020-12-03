package main

import (
	"crawler.com/distributed/rpcsupport"
	"crawler.com/distributed/woker"
	"flag"
	"fmt"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Printf("must specify a port")
		return
	}
	log.Fatal(
		rpcsupport.ServeRpc(
			fmt.Sprintf(":%d", *port),
			woker.CrawlService{}))
}
