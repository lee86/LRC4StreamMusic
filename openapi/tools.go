package openapi

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"
)

func GetLrc(songMid string) (string, error) {
	var data Lrcjson
	s := "https://c.y.qq.com/lyric/fcgi-bin/fcg_query_lyric.fcg?songmid=" + songMid
	sss := string(GET(s))
	if len(sss) == 0 {
		return "", errors.New("结果长度为0")
	}
	sss = strings.Replace(sss, "MusicJsonCallback(", "", -1)
	sss = strings.Replace(sss, ")", "", -1)
	sData := []byte(sss)
	// 将字节切片映射到指定结构上
	err := json.Unmarshal(sData, &data)
	if err != nil {
		return "", err
	}
	//fmt.Println(data.Lyric)
	// base64解码歌词，返回字符串
	lyric, err := base64.StdEncoding.DecodeString(data.Lyric)
	if err != nil {
		return "", err
	}
	return string(lyric), nil
}

func GetSongInfo(word string, num int) string {
	urlword := "https://u.y.qq.com/cgi-bin/musicu.fcg?pcachetime=" + strconv.FormatInt(time.Now().UnixMilli(), 10)
	//data := "{\n   \"comm\" : {\n      \"_channelid\" : \"0\",\n      \"_os_version\" : \"6.1.7601-2%2C+Service+Pack+1\",\n      \"authst\" : \"\",\n      \"ct\" : \"19\",\n      \"cv\" : \"1873\",\n      \"guid\" : \"\",\n      \"patch\" : \"118\",\n      \"psrf_access_token_expiresAt\" : 0,\n      \"psrf_qqaccess_token\" : \"\",\n      \"psrf_qqopenid\" : \"\",\n      \"psrf_qqunionid\" : \"\",\n      \"tmeAppID\" : \"qqmusic\",\n      \"tmeLoginType\" : 2,\n      \"uin\" : \"0\",\n      \"wid\" : \"0\"\n   },\n   \"music.search.SearchCgiService\" : {\n      \"method\" : \"DoSearchForQQMusicDesktop\",\n      \"module\" : \"music.search.SearchCgiService\",\n      \"param\" : {\n         \"grp\" : 1,\n         \"num_per_page\" : " + num + ",\n         \"page_num\" : 1,\n         \"query\" : \"" + word + "\",\n         \"remoteplace\" : \"txt.newclient.top\",\n         \"search_type\" : 0,\n         \"searchid\" : \"\"\n      }\n   }\n}"
	//fmt.Println(data)
	urls := string(POST(urlword, bytes.NewReader(setQMusicQuery(word, num))))
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

func setQMusicQuery(word string, num int) []byte {
	var qMusicQuery QMusicQuery
	qMusicQuery.Comm.Channelid = "0"
	qMusicQuery.Comm.OsVersion = "6.1.7601-2%2C+Service+Pack+1"
	qMusicQuery.Comm.Authst = ""
	qMusicQuery.Comm.Cv = "1873"
	qMusicQuery.Comm.Guid = ""
	qMusicQuery.Comm.Patch = "118"
	qMusicQuery.Comm.PsrfAccessTokenExpiresAt = 0
	qMusicQuery.Comm.PsrfQqaccessToken = ""
	qMusicQuery.Comm.PsrfQqopenid = ""
	qMusicQuery.Comm.PsrfQqunionid = ""
	qMusicQuery.Comm.TmeAppID = "qqmusic"
	qMusicQuery.Comm.TmeLoginType = 2
	qMusicQuery.Comm.Uin = "0"
	qMusicQuery.Comm.Wid = "0"

	qMusicQuery.MusicSearchSearchCgiService.Method = "DoSearchForQQMusicDesktop"
	qMusicQuery.MusicSearchSearchCgiService.Module = "music.search.SearchCgiService"
	qMusicQuery.MusicSearchSearchCgiService.Param.Grp = 1
	qMusicQuery.MusicSearchSearchCgiService.Param.NumPerPage = num
	qMusicQuery.MusicSearchSearchCgiService.Param.PageNum = 1
	qMusicQuery.MusicSearchSearchCgiService.Param.Query = word
	qMusicQuery.MusicSearchSearchCgiService.Param.Remoteplace = "txt.newclient.top"
	qMusicQuery.MusicSearchSearchCgiService.Param.SearchType = 0
	qMusicQuery.MusicSearchSearchCgiService.Param.Searchid = ""
	marshal, _ := json.Marshal(qMusicQuery)
	return marshal
}

type QMusicQuery struct {
	Comm struct {
		Channelid                string `json:"_channelid"`
		OsVersion                string `json:"_os_version"`
		Authst                   string `json:"authst"`
		Ct                       string `json:"ct"`
		Cv                       string `json:"cv"`
		Guid                     string `json:"guid"`
		Patch                    string `json:"patch"`
		PsrfAccessTokenExpiresAt int    `json:"psrf_access_token_expiresAt"`
		PsrfQqaccessToken        string `json:"psrf_qqaccess_token"`
		PsrfQqopenid             string `json:"psrf_qqopenid"`
		PsrfQqunionid            string `json:"psrf_qqunionid"`
		TmeAppID                 string `json:"tmeAppID"`
		TmeLoginType             int    `json:"tmeLoginType"`
		Uin                      string `json:"uin"`
		Wid                      string `json:"wid"`
	} `json:"comm"`
	MusicSearchSearchCgiService struct {
		Method string `json:"method"`
		Module string `json:"module"`
		Param  struct {
			Grp         int    `json:"grp"`
			NumPerPage  int    `json:"num_per_page"`
			PageNum     int    `json:"page_num"`
			Query       string `json:"query"`
			Remoteplace string `json:"remoteplace"`
			SearchType  int    `json:"search_type"`
			Searchid    string `json:"searchid"`
		} `json:"param"`
	} `json:"music.search.SearchCgiService"`
}
