package openapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// LyricInfoHandler 歌词信息获取，可正常返回1条+10条
func LyricInfoHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	title := query.Get("title")
	artist := query.Get("artist")
	num, err := strconv.Atoi(query.Get("limit"))

	var mj MusicjsonN
	err = json.Unmarshal([]byte(getSongInfo(fmt.Sprintf("%v-%v", artist, title), fmt.Sprint(num))), &mj)
	//fmt.Println(err, mj)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(""))
		return
	}
	if len(mj.MusicSearchSearchCgiService.Data.Body.Song.List) > 0 {
		if num == 1 {
			lrc := getLrc(mj.MusicSearchSearchCgiService.Data.Body.Song.List[0].Mid)
			lrc = strings.ReplaceAll(lrc, "[by:]", "[by: Jiangwe Leo QQLrc]")
			// 写入歌词
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(lrc))
			return
		} else {
			// 返回所有歌词，提供选择
			var lrcLists []LrcList
			for _, v := range mj.MusicSearchSearchCgiService.Data.Body.Song.List {
				lrc := getLrc(v.Mid)
				lrc = strings.ReplaceAll(lrc, "[by:]", "[by: Jiangwe Leo QQLrc]")
				lrcLists = append(lrcLists, LrcList{
					Id:     v.Mid,
					Title:  v.Title,
					Artist: artist,
					Lyrics: lrc,
				})
			}
			returnMsg, _ := json.Marshal(lrcLists)
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(returnMsg)
			fmt.Println(string(returnMsg))
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(""))
	return
}

// LyricHandler 歌词信息确认
// gitbobobo说是通知接口而已
func LyricHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := io.ReadAll(r.Body)
	var lyricCheck LyricCheck
	json.Unmarshal(data, &lyricCheck)

	lrc := getLrc(lyricCheck.LyricsId)
	lrc = strings.ReplaceAll(lrc, "[by:]", "[by: Jiangwe Leo QQLrc]")
	// 写入歌词
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
