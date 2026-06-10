package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var version = flag.Bool("version", false, "测试")
var (
	Version   string
	Branch    string
	Commit    string
	BuildTime string
	GOOS      string
	GOARCH    string
)
var saveFileChanel = make(chan SaveChanMsg, 100)
var config Configuration

func init() {
	conf.MustLoad("config.yml", &config)
	var c logx.LogConf
	conf.MustLoad("config-log.yml", &c) // 加载配置文件
	logx.MustSetup(c)                   // 设置日志配置
	if config.Log.Stdout {
		logx.Info("主函数-日志开启控制台输出")
		logx.AddWriter(logx.NewWriter(os.Stdout)) // 添加控制台输出
	}

	flag.Parse()
	if *version {
		fmt.Println("|------------------------------ VERSION INFO -----------------------------|")
		fmt.Printf("VERSION: \t%s\n", Version)
		fmt.Printf("GOOS: \t\t%s\n", GOOS)
		fmt.Printf("GOARCH: \t%s\n", GOARCH)
		fmt.Printf("BRANCH: \t%s\n", Branch)
		fmt.Printf("COMMIT: \t%s\n", Commit)
		fmt.Printf("BUILDTIME: \t%s\n", BuildTime)
		fmt.Println("|-------------------------------------------------------------------------|")

		os.Exit(0)
	}
}

// Configuration 分层配置
type Configuration struct {
	Gin GinConfig `json:"gin" yaml:"gin"`
	Log struct {
		Stdout bool `json:"stdout"`
	} `json:"log"`
}

// GinConfig Gin 服务配置
type GinConfig struct {
	Port int    `json:"port" yaml:"port,default=8080"` // HTTP 端口
	Mode string `json:"mode" yaml:"mode"`              // 运行模式: debug/release/test
}

type SaveChanMsg struct {
	Key   string
	Value []byte
}
