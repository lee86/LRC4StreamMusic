package main

import (
	"encoding/base64"
	"encoding/json"
	"strings"
)

func getlrc(songmid string) string {
	var data Lrcjson
	s := "https://c.y.qq.com/lyric/fcgi-bin/fcg_query_lyric.fcg?songmid=" + songmid
	sss := string(GET(s))

	sss = strings.Replace(sss, "MusicJsonCallback(", "", -1)
	sss = strings.Replace(sss, ")", "", -1)
	sbyte := []byte(sss)
	// 将字节切片映射到指定结构上
	json.Unmarshal(sbyte, &data)
	//fmt.Println(data.Lyric)
	//fmt.Println("------------------------------------------------")
	ss, err := base64.StdEncoding.DecodeString(data.Lyric)
	if err != nil {

	}
	return string(ss)
}
