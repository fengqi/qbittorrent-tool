package config

import "regexp"

type Config struct {
	Host         string       `json:"host"`
	Username     string       `json:"username"`
	Password     string       `json:"password"`
	AutoCategory AutoCategory `json:"auto_category"`
	DomainTag    DomainTag    `json:"domain_tag"`
	//UploadLimit  UploadLimit  `json:"upload_limit"`
}

type AutoCategory struct {
	Enable    bool              `json:"enable"`
	MapConfig map[string]string `json:"map_Config"`
}

type DomainTag struct {
	Enable    bool              `json:"enable"`
	MapConfig map[string]string `json:"map_Config"`
}

type UploadLimit struct {
	Enable  bool    `json:"enable"`
	Limiter Limiter `json:"limiter"`
}

type Limiter struct {
	LimiterName LimiterName `json:"limiter_name"`
	LimiterTag  LimiterTag  `json:"limiter_tag"`
}

type LimiterName struct {
	Name    string        `json:"name"`
	Limit   int           `json:"limit"`
	Keyword string        `json:"keyword"`
	Regexp  regexp.Regexp `json:"regexp"`
}

type LimiterTag struct {
	Tags  []string `json:"tags"`
	Limit int      `json:"limit"`
}
