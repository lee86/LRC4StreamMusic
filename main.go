package main

import (
	"qqlrc/openapi"
)

func main() {
	openapi.RouteLoadHandle("/lyric", openapi.LyricHandler)
	openapi.RouteLoadHandle("/lyricInfo", openapi.LyricInfoHandler)
	openapi.StartServer()
}
