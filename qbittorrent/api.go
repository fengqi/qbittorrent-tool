package qbittorrent

import (
	"fengqi/qbittorrent-tool/config"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

var Api *api

func Init(c *config.Config) error {
	url := fmt.Sprintf("%s/api/v2/auth/login", c.Host)
	data := fmt.Sprintf("username=%s&password=%s", c.Username, c.Password)
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data))
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("[ERR] login err %v\n", err)
		}
	}(resp.Body)

	Api = &api{
		Host:     c.Host,
		Username: c.Username,
		Password: c.Password,
		Cookie:   resp.Cookies(),
	}

	return nil
}
