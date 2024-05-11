package openapi

import (
	"fmt"
	"net/http"
	"time"
)

// RouteLoadHandle 路由加载
func RouteLoadHandle(uri string, handle func(w http.ResponseWriter, r *http.Request)) {
	route.PathPrefix(fmt.Sprintf("%v", uri)).HandlerFunc(handle)
}

func StartServer() {
	srv := &http.Server{
		Handler:      route,
		Addr:         fmt.Sprintf(":%v", config.Api.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("|-------------------------------------------------------------------------|")
	fmt.Println("| ######## ##     ## ##       #### ########  ########   #######  ######## |\n|      ##  ##     ## ##        ##  ##     ## ##     ## ##     ##    ##    |\n|     ##   ##     ## ##        ##  ##     ## ##     ## ##     ##    ##    |\n|    ##    ##     ## ##        ##  ########  ########  ##     ##    ##    |\n|   ##     ##     ## ##        ##  ##        ##     ## ##     ##    ##    |\n|  ##      ##     ## ##        ##  ##        ##     ## ##     ##    ##    |\n| ########  #######  ######## #### ##        ########   #######     ##    |")
	fmt.Println("|-------------------------------------------------------------------------|")
	fmt.Printf("| Listen to PORT: %v    Have fun! \n", config.Api.Port)
	fmt.Println("|-------------------------------------------------------------------------|")
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
