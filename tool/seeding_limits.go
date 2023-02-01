package tool

import (
	"fengqi/qbittorrent-tool/config"
	"fengqi/qbittorrent-tool/qbittorrent"
)

// SeedingLimits 做种限制加强版
// 相比较于qb自带，增加根据标签、分类、关键字精确限制
func SeedingLimits(c *config.Config, api *qbittorrent.Api) {
	// TODO
}
