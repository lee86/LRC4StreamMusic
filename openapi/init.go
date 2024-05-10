package openapi

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/zeromicro/go-zero/core/conf"
)

var route *mux.Router

var config Configuration

func init() {
	route = mux.NewRouter()
	conf.MustLoad("./conf.yml", &config)
	fmt.Println("conf: +-+-+-+-+->>>>> ", config)
}

type Configuration struct {
	Api Api `json:"net"`
}
type Api struct {
	Port int `json:"port"`
}
