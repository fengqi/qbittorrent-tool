package main

import (
	"fengqi/qbittorrent-tool/config"
	"fengqi/qbittorrent-tool/qbittorrent"
	"fengqi/qbittorrent-tool/tool"
	"flag"
	"fmt"
)

func main() {
	configFile := flag.String("c", "./config.json", "config file path")
	flag.Parse()
	c, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Printf("[ERR] load config err: %v\n", err)
		return
	}

	api, err := qbittorrent.Login(c)
	if err != nil {
		fmt.Printf("[ERR] login to qbittorrent err %v\n", err)
		return
	}

	tool.AutoCategory(c, api)
}
