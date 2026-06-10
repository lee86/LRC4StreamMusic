package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/proc"
)

func main() {
	// 注册要捕获的信号：SIGTERM（终止信号）和os.Interrupt（中断信号）
	stop := make(chan os.Signal, 1)
	// 启动一个goroutine来处理信号
	signal.Notify(stop, syscall.SIGTERM, os.Interrupt)
	// 阻塞等待接收信号
	go func() {
		// 收到信号后打印日志信息
		<-stop
		// 关闭日志系统
		logx.Info("收到关闭信号，正在关闭...")
		// 退出程序
		logx.Close()
		os.Exit(0)
	}()
	// 初始化voce模块，传入配置文件路径

	//voce.Init(*configFile)
	go start()
	//go start()
	for {
		select {
		case <-proc.Done(): // 检查程序是否需要退出
			os.Exit(0)
		default:
			time.Sleep(time.Second * 60)
			logx.Info(time.Now().Format("2006-01-02 15:04:05"), " health") // 打印当前时间到日志
		}
	}
}
