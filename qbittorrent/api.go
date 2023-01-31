package qbittorrent

import (
	"fengqi/qbittorrent-tool/config"
	"fmt"
	"net/http"
	"strings"
)

func Login(c *config.Config) (*Api, error) {
	url := fmt.Sprintf("%s/api/v2/auth/login", c.Host)
	data := fmt.Sprintf("username=%s&password=%s", c.Username, c.Password)
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &Api{
		Host:     c.Host,
		Username: c.Username,
		Password: c.Password,
		Cookie:   resp.Cookies(),
	}, nil
}
