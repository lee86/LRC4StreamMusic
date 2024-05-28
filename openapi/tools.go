package openapi

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func getLrc(songMid string) string {
	var data Lrcjson
	s := "https://c.y.qq.com/lyric/fcgi-bin/fcg_query_lyric.fcg?songmid=" + songMid
	sss := string(GET(s))
	if len(sss) == 0 {
		return sss
	}
	sss = strings.Replace(sss, "MusicJsonCallback(", "", -1)
	sss = strings.Replace(sss, ")", "", -1)
	sData := []byte(sss)
	// 将字节切片映射到指定结构上
	err := json.Unmarshal(sData, &data)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(data.Lyric)
	// base64解码歌词，返回字符串
	lyric, err := base64.StdEncoding.DecodeString(data.Lyric)
	if err != nil {
		fmt.Println(err)
	}
	return string(lyric)
}

func getSongInfo(word, num string) string {
	urlword :=
		"https://u.y.qq.com/cgi-bin/musicu.fcg?pcachetime=" + strconv.FormatInt(time.Now().UnixMilli(), 10)
	data := "{\n   \"comm\" : {\n      \"_channelid\" : \"0\",\n      \"_os_version\" : \"6.1.7601-2%2C+Service+Pack+1\",\n      \"authst\" : \"\",\n      \"ct\" : \"19\",\n      \"cv\" : \"1873\",\n      \"guid\" : \"\",\n      \"patch\" : \"118\",\n      \"psrf_access_token_expiresAt\" : 0,\n      \"psrf_qqaccess_token\" : \"\",\n      \"psrf_qqopenid\" : \"\",\n      \"psrf_qqunionid\" : \"\",\n      \"tmeAppID\" : \"qqmusic\",\n      \"tmeLoginType\" : 2,\n      \"uin\" : \"0\",\n      \"wid\" : \"0\"\n   },\n   \"music.search.SearchCgiService\" : {\n      \"method\" : \"DoSearchForQQMusicDesktop\",\n      \"module\" : \"music.search.SearchCgiService\",\n      \"param\" : {\n         \"grp\" : 1,\n         \"num_per_page\" : " + num + ",\n         \"page_num\" : 1,\n         \"query\" : \"" + word + "\",\n         \"remoteplace\" : \"txt.newclient.top\",\n         \"search_type\" : 0,\n         \"searchid\" : \"\"\n      }\n   }\n}"
	urls := string(POST(urlword, strings.NewReader(data)))
	if len(urls) == 0 {
		return urls
	}
	urls = strings.Replace(urls, "<em>", "", -1)
	urls = strings.Replace(urls, "</em>", "", -1)
	urls = strings.Replace(urls, "callback({", "{", -1)
	urls = strings.Replace(urls, "})", "}", -1)
	urls = strings.Replace(urls, "music.search.SearchCgiService", "music", -1)

	return urls
}
