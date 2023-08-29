package tool

import (
	"fengqi/qbittorrent-tool/config"
	"fengqi/qbittorrent-tool/qbittorrent"
	"fmt"
	"strings"
)

// DomainTag 根据域名设置tag, 主要是给webui用
// 等webui可以和桌面端一样自动合并tracker的时候, 可以放弃使用
func DomainTag(c *config.Config, torrent *qbittorrent.Torrent) {
	if !c.DomainTag.Enable || c.DomainTag.MapConfig == nil {
		return
	}

	if torrent.Tracker == "" {
		trackerList, err := qbittorrent.Api.GetTorrentTrackers(torrent.Hash)
		if err == nil && len(trackerList) > 0 {
			torrent.Tracker = trackerList[0].Url // 暂时默认用第一个
		}
	}

	tag, err := torrent.GetTrackerHost()
	if err != nil {
		fmt.Printf("[ERR] get %s tag err: %v\n", torrent.Name, err)
		return
	}

	if custom, ok := c.DomainTag.MapConfig[tag]; ok {
		tag = custom
	}

	if strings.Contains(torrent.Tags, tag) {
		return
	}

	err = qbittorrent.Api.AddTags(torrent.Hash, tag)
	if err != nil {
		fmt.Printf("[ERR] add tag %s to %s err: %v\n", tag, torrent.Name, err)
		return
	}

	fmt.Printf("[INFO] add tag %s to %s\n", tag, torrent.Name)
}
