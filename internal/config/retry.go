package config

import (
	"time"
)

type Retry struct {
	// 是否开启
	Enabled *bool `default:"true"`
	// 次数
	Counts int `default:"${5}"`
	// 间隔
	Backoff time.Duration `default:"${5s}"`
	// 时间
	Timeout time.Duration `default:"${60m}"`
}
