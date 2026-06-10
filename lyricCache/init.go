package lyricCache

import (
	"github.com/zeromicro/go-zero/core/conf"
)

var config Configuration

func init() {
	conf.MustLoad("config.yml", &config)
}

type Configuration struct {
	Base Base `json:"base"`
}

type Base struct {
	CacheDir string `json:"cache"`
}
