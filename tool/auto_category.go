package tool

import (
	"fengqi/qbittorrent-tool/config"
	"fengqi/qbittorrent-tool/qbittorrent"
	"fmt"
	"log"
)

// AutoCategory 根据保存目录设置分类
func AutoCategory(c *config.Config, api *qbittorrent.Api) {
	if !c.AutoCategory.Enable || c.AutoCategory.MapConfig == nil {
		return
	}

	params := map[string]string{
		"filter":   "all",
		"category": "",
		"limit":    "1000",
	}
	torrentList, err := api.GetTorrentList(params)
	if err != nil {
		log.Printf("[ERR] get torrent list without category err %v\n", err)
		return
	}

	log.Printf("[INFO] get torrent list without category count: %d\n", len(torrentList))
	for _, i := range torrentList {
		category, ok := c.AutoCategory.MapConfig[i.SavePath]
		if !ok {
			log.Printf("[WARN] get path %s categroy empty\n", i.SavePath)
			continue
		}

		err = api.SetCategory(i.Hash, category)
		if err != nil {
			log.Printf("[ERR] set category: %s \tto: %s err: %v\n", category, i.Name, err)
			continue
		}

		fmt.Printf("[INFO] set category: %s \tto: %s\n", category, i.Name)
	}
}
