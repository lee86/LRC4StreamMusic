package main

import (
	"qqlrc/openapi"
)

func main() {
	openapi.RouteLoadHandle("/lyric", openapi.LyricInfoHandler)
	openapi.RouteLoadHandle("/lyricInfo", openapi.LyricInfoHandler)
	openapi.StartServer()
}
