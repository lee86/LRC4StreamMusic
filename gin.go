package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/logx"
)

var router *gin.Engine
var sever *http.Server

func start() {
	setMode()
	router = gin.Default()

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/lyricInfo", LyricInfoHandler)
			v1.POST("/lyricCheck", LyricHandler)
		}
	}
	fmt.Println("|-------------------------------------------------------------------------|")
	fmt.Println("▗▖   ▗▄▄▖  ▗▄▄▖▄  ▗▖▗▖  ▗▖▗▖ ▗▖ ▗▄▄▖▗▄▄▄▖▗▖  ▗▖▗▄▄▄▖▗▄▄▖ \n▐▌   ▐▌ ▐▌▐▌   █  ▐▌▐▛▚▞▜▌▐▌ ▐▌▐▌     █  ▐▌  ▐▌▐▌   ▐▌ ▐▌\n▐▌   ▐▛▀▚▖▐▌   ▀▀▀▜▌▐▌  ▐▌▐▌ ▐▌ ▝▀▚▖  █  ▐▌  ▐▌▐▛▀▀▘▐▛▀▚▖\n▐▙▄▄▖▐▌ ▐▌▝▚▄▄▖   ▐▌▐▌  ▐▌▝▚▄▞▘▗▄▄▞▘▗▄█▄▖ ▝▚▞▘ ▐▙▄▄▖▐▌ ▐▌")
	fmt.Println("|-------------------------------------------------------------------------|")
	fmt.Printf("| be ready to bind PORT: %v    Have fun! \n", config.Gin.Port)
	fmt.Println("|-------------------------------------------------------------------------|")
	startServer()
}

func startServer() {
	sever = &http.Server{
		Addr:           fmt.Sprintf(":%v", config.Gin.Port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf("| bind PORT: %v over ,  Have fun! \n", config.Gin.Port)
	fmt.Println("|-------------------------------------------------------------------------|")
	err := sever.ListenAndServe()
	if err != nil {
		logx.Error(err)
		//os.Exit(-5)
	}
}

// 重载gin server
func reloadServer(ctx context.Context) {
	sever.Shutdown(ctx)
	startServer()
}

func setMode() {
	gin.SetMode(config.Gin.Mode)
}
