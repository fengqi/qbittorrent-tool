package config

import "regexp"

type Config struct {
	Host         string       `json:"host"`
	Username     string       `json:"username"`
	Password     string       `json:"password"`
	AutoCategory AutoCategory `json:"auto_category"`
	DomainTag    DomainTag    `json:"domain_tag"`
	//UploadLimit  UploadLimit  `json:"upload_limit"`
	SeedingLimits SeedingLimits `json:"seeding_limits"`
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

type SeedingLimits struct {
	Enable bool                `json:"enable"`
	Rules  []SeedingLimitsRule `json:"rules"`
}

type SeedingLimitsRule struct {
	Ratio        float64  `json:"ratio"`         // 当分享率达到，大于0生效
	SeedingTime  int      `json:"seeding_time"`  // 当做种时间达到（分钟），大于0生效
	ActivityTime int      `json:"activity_time"` // 当最后活动时间达到（分钟），大于0生效
	Tag          []string `json:"tag"`           // 当包括这些标签，不为空生效
	Category     []string `json:"category"`      // 当包括这些分类，不为空生效
	Tracker      []string `json:"tracker"`       // 当包括这些tracker，不为空生效
	SeedsGt      int      `json:"seeds_gt"`      // 当做种数大于，大于0生效
	SeedsLt      int      `json:"seeds_lt"`      // 当做种数小于，大于0生效
	Keyword      []string `json:"keyword"`       // 当种子标题包括这些关键字，不为空生效
	Action       int      `json:"action"`        // 动作：0继续做种、1暂停做种、2删除种子、3删除种子及所属文件、4启动超级做种
}
