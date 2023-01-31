package config

import (
	"encoding/json"
	"io/ioutil"
)

func LoadConfig(path string) (*Config, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var c *Config
	err = json.Unmarshal(bytes, &c)
	return c, err
}
