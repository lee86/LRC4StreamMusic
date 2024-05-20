package openapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"qqlrc/lyricCache"
	"strconv"
	"strings"
)

// LyricInfoHandler 歌词信息获取，可正常返回1条+10条
func LyricInfoHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	title := query.Get("title")
	artist := query.Get("artist")
	num, err := strconv.Atoi(query.Get("limit"))
	if title == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte("HAVE NO TITLE"))
		if err != nil {
			return
		}
		return
	}
	var mj MusicjsonN
	keys := fmt.Sprintf("%v-%v", artist, title)
	if lyric, ok := lyricCache.CacheSelect(keys); ok {
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte(lyric))
		if err != nil {
			return
		}
		return
	}
	err = json.Unmarshal([]byte(getSongInfo(keys, fmt.Sprint(num))), &mj)
	//fmt.Println(err, mj)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		_, err = w.Write([]byte(""))
		if err != nil {
			return
		}
		return
	}
	if len(mj.MusicSearchSearchCgiService.Data.Body.Song.List) > 0 {
		if num == 1 {
			// 返回1条
			lrc, ok := lyricCache.CacheSelect(keys)
			if !ok {
				lrc = getLrc(mj.MusicSearchSearchCgiService.Data.Body.Song.List[0].Mid)
				lrc = strings.ReplaceAll(lrc, "[by:]", "[by: Jiangwe Leo QQLrc]")
				if !lyricCache.CacheSave(keys, []byte(lrc)) {
					fmt.Println("save error")
				}
			}
			// 写入歌词
			w.WriteHeader(http.StatusOK)
			_, err = w.Write([]byte(lrc))
			if err != nil {
				return
			}
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
			_, err = w.Write(returnMsg)
			if err != nil {
				return
			}
			fmt.Println(string(returnMsg))
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	_, err = w.Write([]byte(""))
	if err != nil {
		return
	}
	return
}

// LyricHandler 歌词信息确认
// gitbobobo说是通知接口而已
func LyricHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := io.ReadAll(r.Body)
	var lyricCheck LyricCheck
	err := json.Unmarshal(data, &lyricCheck)
	if err != nil {
		return
	}
	keys := fmt.Sprintf("%v-%v", lyricCheck.Artist, lyricCheck.Title)
	lrc := getLrc(lyricCheck.LyricsId)
	lrc = strings.ReplaceAll(lrc, "[by:]", "[by: Jiangwe Leo QQLrc]")
	if lyricCache.CacheSave(keys, []byte(lrc)) {
		// 返回200
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte("ok"))
		if err != nil {
			return
		}
		return
	}
	// 返回416
	w.WriteHeader(http.StatusRequestedRangeNotSatisfiable)
	_, err = w.Write([]byte("check error"))
	if err != nil {
		return
	}
}
