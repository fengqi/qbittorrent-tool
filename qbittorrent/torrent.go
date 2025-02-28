package qbittorrent

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

// GetTorrentList 获取种子列表
// https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#get-torrent-list
func (a *api) GetTorrentList(params map[string]string) ([]*Torrent, error) {
	query := url.Values{}
	for k, v := range params {
		query.Add(k, v)
	}

	api := a.Host + "/api/v2/torrents/info?" + query.Encode()
	bytes, err := a.request(http.MethodGet, api, nil)
	if err != nil {
		return nil, err
	}

	var torrentList []*Torrent
	err = json.Unmarshal(bytes, &torrentList)

	return torrentList, err
}

// GetTorrentTrackers 获取种子的有效tracker列表
func (a *api) GetTorrentTrackers(hash string) ([]*TorrentTracker, error) {
	api := fmt.Sprintf("%s/api/v2/torrents/trackers?hash=%s", a.Host, hash)

	bytes, err := a.request(http.MethodGet, api, nil)
	if err != nil {
		return nil, err
	}

	var trackerList []*TorrentTracker
	err = json.Unmarshal(bytes, &trackerList)
	if err != nil {
		return nil, err
	}

	sort.SliceStable(trackerList, func(i, j int) bool {
		return trackerList[i].Status < trackerList[j].Status
	})

	i := 0
	for _, item := range trackerList {
		// tier >=0 的才有效
		if item.Tier >= 0 {
			trackerList[i] = item
			i++
		}
	}

	return trackerList[:i], nil
}

// SetCategory 给种子设置分类
func (a *api) SetCategory(hashes, category string) error {
	api := fmt.Sprintf("%s/api/v2/torrents/setCategory", a.Host)
	data := fmt.Sprintf("hashes=%s&category=%s", hashes, category)

	_, err := a.request(http.MethodPost, api, strings.NewReader(data))

	return err
}

// AddTags 给种子添加标签
func (a *api) AddTags(hashes, tag string) error {
	api := fmt.Sprintf("%s/api/v2/torrents/addTags", a.Host)
	data := fmt.Sprintf("hashes=%s&tags=%s", hashes, tag)

	_, err := a.request("POST", api, strings.NewReader(data))

	return err
}

// ResumeTorrents 继续种子
// TODO  根据版本使用新旧接口：/api/v2/app/version
func (a *api) ResumeTorrents(hashes string) error {
	api := fmt.Sprintf("%s/api/v2/torrents/start", a.Host)
	data := fmt.Sprintf("hashes=%s", hashes)

	_, err := a.request(http.MethodPost, api, strings.NewReader(data))

	return err
}

// PauseTorrents 暂停种子
func (a *api) PauseTorrents(hashes string) error {
	api := fmt.Sprintf("%s/api/v2/torrents/stop", a.Host)
	data := fmt.Sprintf("hashes=%s", hashes)

	_, err := a.request(http.MethodPost, api, strings.NewReader(data))

	return err
}

// DeleteTorrents 删除种子
func (a *api) DeleteTorrents(hashes string, deleteFiles bool) error {
	api := fmt.Sprintf("%s/api/v2/torrents/delete", a.Host)
	data := fmt.Sprintf("hashes=%s&deleteFiles=%t", hashes, deleteFiles)

	_, err := a.request(http.MethodPost, api, strings.NewReader(data))

	return err
}

// SetSuperSeeding 设置超级做种
func (a *api) SetSuperSeeding(hashes string, value bool) error {
	api := fmt.Sprintf("%s/api/v2/torrents/setSuperSeeding", a.Host)
	data := fmt.Sprintf("hashes=%s&value=%t", hashes, value)

	_, err := a.request("POST", api, strings.NewReader(data))

	return err
}

// SetDownloadLimit 下载限速
func (a *api) SetDownloadLimit(hashes string, speed int) error {
	api := fmt.Sprintf("%s/api/v2/torrents/setDownloadLimit", a.Host)
	data := fmt.Sprintf("hashes=%s&limit=%d", hashes, speed)

	_, err := a.request("POST", api, strings.NewReader(data))

	return err
}

// SetUploadLimit 上传限速
func (a *api) SetUploadLimit(hashes string, speed int) error {
	api := fmt.Sprintf("%s/api/v2/torrents/setUploadLimit", a.Host)
	data := fmt.Sprintf("hashes=%s&limit=%d", hashes, speed)

	_, err := a.request("POST", api, strings.NewReader(data))

	return err
}

// SetShareLimit 设置分享率、时间限制
func (a *api) SetShareLimit(hashes string, radio float64, seedingTime, inactiveSeedingTime int) error {
	api := fmt.Sprintf("%s/api/v2/torrents/setShareLimits", a.Host)
	data := fmt.Sprintf("hashes=%s&ratioLimit=%f&seedingTimeLimit=%d&inactiveSeedingTimeLimit=%d",
		hashes,
		radio,
		seedingTime,
		inactiveSeedingTime)

	_, err := a.request("POST", api, strings.NewReader(data))

	return err
}

// 发起请求
func (a *api) request(method, url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	for _, item := range a.Cookie {
		req.AddCookie(item)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("[ERR] close http.Response.body err %v\n", err)
		}
	}(resp.Body)

	return io.ReadAll(resp.Body)
}
