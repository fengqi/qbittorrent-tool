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
	StatusTag     StatusTag     `json:"status_tag"`
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
	Resume bool                `json:"resume"` // 已经被暂停的种子，符合条件后是否开启继续做种
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
	Limits       *Limits  `json:"limits"`        // 限速
}

type StatusTag struct {
	Enable    bool              `json:"enable"`
	MapConfig map[string]string `json:"map_Config"`
}

type Limits struct {
	Download            int     `json:"download"`              // 下载限速，bytes/s，大于0生效
	Upload              int     `json:"upload"`                // 上传限速，bytes/s，大于0生效
	Ratio               float64 `json:"ratio"`                 // 分享率，浮点数如：1.2，-2使用全局限制，-1不限制，等于0不生效
	SeedingTime         int     `json:"seeding_time"`          // 做种时间（分钟），-2使用全局限制，-1不限制，等于0不生效
	InactiveSeedingTime int     `json:"inactive_seeding_time"` // 不活跃的做种时间（分钟），-2使用全局限制，-1不限制，等于0不生效
}
