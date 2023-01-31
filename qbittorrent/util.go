package qbittorrent

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

func (t *Torrent) GetTrackerTag() (string, error) {
	if t.Tracker == "" {
		return "", errors.New("tracker empty")
	}

	parse, err := url.Parse(t.Tracker)
	if err != nil {
		return "", err
	}

	split := strings.Split(parse.Host, ".")
	return fmt.Sprintf("%s.%s", split[len(split)-2], split[len(split)-1]), nil
}
