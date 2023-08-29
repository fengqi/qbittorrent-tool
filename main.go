package main

import (
	"fengqi/qbittorrent-tool/config"
	"fengqi/qbittorrent-tool/qbittorrent"
	"fengqi/qbittorrent-tool/tool"
	"flag"
	"fmt"
	"log"
	"strconv"
)

func main() {
	configFile := flag.String("c", "./config.json", "config file path")
	flag.Parse()
	c, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Printf("[ERR] load config err: %v\n", err)
		return
	}

	if err = qbittorrent.Init(c); err != nil {
		fmt.Printf("[ERR] login to qbittorrent err %v\n", err)
		return
	}

	offset := 0
	limit := 1000
	for {
		params := map[string]string{
			"filter": "all",
			"sort":   "added_on",
			"limit":  strconv.Itoa(limit),
			"offset": strconv.Itoa(offset),
		}
		torrentList, err := qbittorrent.Api.GetTorrentList(params)
		if err != nil {
			log.Printf("[ERR] get torrent list err %v\n", err)
			return
		}

		log.Printf("[INFO] get torrent list count: %d\n", len(torrentList))
		for _, torrent := range torrentList {
			tool.AutoCategory(c, torrent)
			tool.DomainTag(c, torrent)
			tool.SeedingLimits(c, torrent)
			tool.StatusTag(c, torrent)
		}

		if len(torrentList) < limit {
			break
		}
		offset += limit
	}
}
