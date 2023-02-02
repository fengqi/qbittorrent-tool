package tool

import (
	"fengqi/qbittorrent-tool/config"
	"fengqi/qbittorrent-tool/qbittorrent"
	"fengqi/qbittorrent-tool/util"
	"fmt"
	"log"
	"strings"
	"time"
)

// SeedingLimits 做种限制加强版
// 相比较于qb自带，增加根据标签、分类、关键字精确限制
func SeedingLimits(c *config.Config) {
	if !c.SeedingLimits.Enable || len(c.SeedingLimits.Rules) == 0 {
		return
	}

	params := map[string]string{
		"filter": "seeding",
		"sort":   "added_on",
		"limit":  "1000",
	}
	torrentList, err := qbittorrent.Api.GetTorrentList(params)
	if err != nil {
		log.Printf("[ERR] get torrent list prepare set lmits err %v\n", err)
		return
	}

	log.Printf("[INFO] get torrent list prepare set lmits count: %d\n", len(torrentList))
	for _, item := range torrentList {
		action := matchRule(item, c.SeedingLimits.Rules)
		fmt.Printf("action:%d Tag:%s %s\n", action, item.Tags, item.Name)
		executeAction(item.Hash, action)
	}
}

// 规则至少有一个生效，且生效的全部命中，action才有效，后面的规则会覆盖前面的
func matchRule(torrent *qbittorrent.Torrent, rules []config.SeedingLimitsRule) int {
	action := 0
	loc, _ := time.LoadLocation("Asia/Shanghai")

	for _, rule := range rules {
		score := 0

		// 分享率
		if rule.Ratio > 0 {
			if torrent.Ratio < rule.Ratio {
				continue
			}
			score += 1
		}

		// 做种时间，从下载完成算起
		if rule.SeedingTime > 0 {
			completionOn := time.Unix(int64(torrent.CompletionOn), 0).In(loc)
			deadOn := completionOn.Add(time.Minute * time.Duration(rule.SeedingTime))
			if time.Now().In(loc).Before(deadOn) {
				continue
			}
			score += 1
		}

		// 标签
		if len(rule.Tag) != 0 && torrent.Tags != "" {
			tags := strings.Split(torrent.Tags, ",")
			hit := false
		jump:
			for _, item := range rule.Tag {
				for _, item2 := range tags {
					if item == item2 {
						hit = true
						break jump
					}
				}
			}
			if !hit {
				continue
			}
			score += 1
		}

		// 分类
		if len(rule.Category) != 0 && torrent.Category != "" {
			if !util.InArray(torrent.Category, rule.Category) {
				continue
			}
			score += 1
		}

		// tracker  TODO 可能有多个tracker的情况要处理
		tracker, _ := torrent.GetTrackerHost()
		if len(rule.Tracker) != 0 && tracker != "" {
			if !util.InArray(tracker, rule.Tracker) {
				continue
			}
			score += 1
		}

		// 做种数大于
		if rule.SeedsGt > 0 {
			if torrent.NumComplete < rule.SeedsGt {
				continue
			}
			score += 1
		}

		// 做种数小于
		if rule.SeedsLt > 0 {
			if torrent.NumComplete > rule.SeedsLt {
				continue
			}
			score += 1
		}

		// 关键字
		if len(rule.Keyword) != 0 {
			if !util.ContainsArray(torrent.Name, rule.Keyword) {
				continue
			}
			score += 1
		}

		if score > 0 {
			action = rule.Action
		}
	}
	return action
}

func executeAction(hash string, action int) {
	if action == 0 {
		return
	}

	switch action {
	case 1:
		_ = qbittorrent.Api.PauseTorrents(hash)
		break

	case 2:
		_ = qbittorrent.Api.DeleteTorrents(hash, false)
		break

	case 3:
		_ = qbittorrent.Api.DeleteTorrents(hash, true)
		break

	case 4:
		_ = qbittorrent.Api.SetSuperSeeding(hash, true)
		break

	}
}
