package main

import (
	"qqlrc/openapi"
)

func main() {
	openapi.RouteLoadHandle("/lyric/v1", openapi.LyricHandler)
	openapi.RouteLoadHandle("/lyricInfo/v1", openapi.LyricInfoHandler)
	openapi.StartServer()
}
