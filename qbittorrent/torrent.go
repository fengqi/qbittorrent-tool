package qbittorrent

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// GetTorrentList 获取种子列表
// https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#get-torrent-list
func (a *Api) GetTorrentList(params map[string]string) ([]*Torrent, error) {
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

	return torrentList, nil
}

// GetTorrentListWithoutCategory 获取没有分类的种子
func (a *Api) GetTorrentListWithoutCategory() ([]*Torrent, error) {
	infoApi := fmt.Sprintf("%s/api/v2/torrents/info?filter=all&category=&limit=%d", a.Host, 1000)
	req, err := http.NewRequest("GET", infoApi, nil)
	if err != nil {
		return nil, err
	}

	for _, item := range a.Cookie {
		req.AddCookie(item)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var torrentList []*Torrent
	err = json.Unmarshal(bytes, &torrentList)

	return torrentList, nil
}

// GetTorrentTrackers 获取种子的有效tracker列表
func (a *Api) GetTorrentTrackers(hash string) ([]*TorrentTracker, error) {
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
func (a *Api) SetCategory(hashes, category string) error {
	api := fmt.Sprintf("%s/api/v2/torrents/setCategory", a.Host)
	data := fmt.Sprintf("hashes=%s&category=%s", hashes, category)

	_, err := a.request(http.MethodPost, api, strings.NewReader(data))

	return err
}

// AddTags 给种子添加标签
func (a *Api) AddTags(hashes, tag string) error {
	api := fmt.Sprintf("%s/api/v2/torrents/addTags", a.Host)
	data := fmt.Sprintf("hashes=%s&tags=%s", hashes, tag)

	_, err := a.request("POST", api, strings.NewReader(data))

	return err
}

// 发起请求
func (a *Api) request(method, url string, body io.Reader) ([]byte, error) {
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

	return ioutil.ReadAll(resp.Body)
}
