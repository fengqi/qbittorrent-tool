package tool

import (
	"fengqi/qbittorrent-tool/config"
	"fengqi/qbittorrent-tool/qbittorrent"
	"fmt"
	"log"
)

// DomainTag 根据域名设置tag, 主要是给webui用
// 等webui可以和桌面端一样自动合并tracker的时候, 可以放弃使用
func DomainTag(c *config.Config) {
	if !c.DomainTag.Enable || c.DomainTag.MapConfig == nil {
		return
	}

	params := map[string]string{
		"filter": "all",
		"tag":    "",
		"limit":  "1000",
	}
	torrentList, err := qbittorrent.Api.GetTorrentList(params)
	if err != nil {
		log.Printf("[ERR] get torrent list without tag err %v\n", err)
		return
	}

	log.Printf("[INFO] get torrent list without tag count: %d\n", len(torrentList))
	for _, i := range torrentList {
		if i.Tracker == "" {
			trackerList, err := qbittorrent.Api.GetTorrentTrackers(i.Hash)
			if err == nil && len(trackerList) > 0 {
				i.Tracker = trackerList[0].Url // 暂时默认用第一个
			}
		}

		tag, err := i.GetTrackerHost()
		if err != nil {
			fmt.Printf("[ERR] get %s tag err: %v\n", i.Name, err)
			continue
		}

		if custom, ok := c.DomainTag.MapConfig[tag]; ok {
			tag = custom
		}

		err = qbittorrent.Api.AddTags(i.Hash, tag)
		if err != nil {
			fmt.Printf("[ERR] add tag %s to %s err: %v\n", tag, i.Name, err)
			continue
		}

		fmt.Printf("[INFO] add tag %s to %s\n", tag, i.Name)
	}
}
