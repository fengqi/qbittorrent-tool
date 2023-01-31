package qbittorrent

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

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

func (a *Api) GetTorrentTrackers(hash string) ([]*TorrentTracker, error) {
	url := fmt.Sprintf("%s/api/v2/torrents/trackers?hash=%s", a.Host, hash)
	req, err := http.NewRequest("GET", url, nil)
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

func (a *Api) SetCategory(hashes, category string) error {
	url := fmt.Sprintf("%s/api/v2/torrents/setCategory", a.Host)
	data := fmt.Sprintf("hashes=%s&category=%s", hashes, category)

	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	for _, item := range a.Cookie {
		req.AddCookie(item)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
