package openapi

import (
	"encoding/base64"
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

func getLrc(songmid string) string {
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

func getSongId(word, num string) string {
	urlword :=
		"https://u.y.qq.com/cgi-bin/musicu.fcg?pcachetime=" + strconv.FormatInt(time.Now().UnixMilli(), 10)
	data := "{\n   \"comm\" : {\n      \"_channelid\" : \"0\",\n      \"_os_version\" : \"6.1.7601-2%2C+Service+Pack+1\",\n      \"authst\" : \"\",\n      \"ct\" : \"19\",\n      \"cv\" : \"1873\",\n      \"guid\" : \"\",\n      \"patch\" : \"118\",\n      \"psrf_access_token_expiresAt\" : 0,\n      \"psrf_qqaccess_token\" : \"\",\n      \"psrf_qqopenid\" : \"\",\n      \"psrf_qqunionid\" : \"\",\n      \"tmeAppID\" : \"qqmusic\",\n      \"tmeLoginType\" : 2,\n      \"uin\" : \"0\",\n      \"wid\" : \"0\"\n   },\n   \"music.search.SearchCgiService\" : {\n      \"method\" : \"DoSearchForQQMusicDesktop\",\n      \"module\" : \"music.search.SearchCgiService\",\n      \"param\" : {\n         \"grp\" : 1,\n         \"num_per_page\" : " + num + ",\n         \"page_num\" : 1,\n         \"query\" : \"" + word + "\",\n         \"remoteplace\" : \"txt.newclient.top\",\n         \"search_type\" : 0,\n         \"searchid\" : \"\"\n      }\n   }\n}"
	urls := string(POST(urlword, strings.NewReader(data)))
	//fmt.Println(urlword)
	//fmt.Println(data)
	urls = strings.Replace(urls, "<em>", "", -1)
	urls = strings.Replace(urls, "</em>", "", -1)
	urls = strings.Replace(urls, "callback({", "{", -1)
	urls = strings.Replace(urls, "})", "}", -1)
	urls = strings.Replace(urls, "music.search.SearchCgiService", "music", -1)

	//sbyte := []byte(urls)
	////fmt.Println(urls)
	//// 将字节切片映射到指定结构上
	//var mj Musicjson
	//json.Unmarshal(sbyte, &mj)
	//fmt.Println(mj.Data.Song.List)
	//fmt.Println(mj.Data.Song.List)
	return urls
	//for key, value := range mj.Data.Song.List {
	//	fmt.Println(key+1, value)
	//}
	//return mj.Data.Song.List[0].Songmid
}
