package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"qqlrc/lyricCache"
	"qqlrc/openapi"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/logx"
)

// LyricInfoHandler 歌词信息获取，可正常返回1条+10条
// RequestMethod = GET
func LyricInfoHandler(ctx *gin.Context) {
	logx.Info("解析请求开始")
	// 参数映射
	title := ctx.Query("title")
	artist := ctx.Query("artist")
	num, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		num = 1
	}
	logx.Infof("解析请求结束 , 歌手 %v , 歌曲 %v , 请求限制 %v", artist, title, num)
	// 判断title是否非空
	if title == "" {
		logx.Info("歌曲名为空，不接受请求，返回")
		ctx.JSONP(http.StatusNotFound, gin.H{"success": false, "errMsg": "request have no title", "errCode": 40401})
		return
	}
	// 简单拼装歌手名+歌曲名
	keys := fmt.Sprintf("%v-%v", artist, title)

	var mj openapi.MusicjsonN
	if num == 1 {
		logx.Info("开始请求缓存")
		// 先从缓存中查找，可命中则返回本地缓存数据
		if lyric, ok := lyricCache.CacheSelect(keys); ok {
			logx.Info("命中缓存，结束请求")
			ctx.String(http.StatusOK, lyric)
			return
		}
		logx.Info("未命中缓存")
	}
	// 缓存中没有，则从QQ音乐中重新请求获取
	logx.Info("开始请求QQ音乐api")
	err = json.Unmarshal([]byte(openapi.GetSongInfo(keys, num)), &mj)
	if err != nil {
		logx.Error("json序列化失败 ", err)
		ctx.JSONP(http.StatusNotFound, gin.H{"success": true, "errMsg": "从QQ音乐获取歌曲信息失败", "errCode": 40402})
		return
	}
	// 获取到歌词
	if len(mj.MusicSearchSearchCgiService.Data.Body.Song.List) > 0 {
		// 直接返回第一组命中歌词
		if num == 1 {
			lyric, err := openapi.GetLrc(mj.MusicSearchSearchCgiService.Data.Body.Song.List[0].Mid)
			if err != nil {
				logx.Error(err)
				ctx.String(http.StatusOK, err.Error())
				return
			}
			lyric = strings.ReplaceAll(lyric, "[by:]", "[by: Jiangwe Leo QQLrc]")

			// 写入歌词
			saveFileChanel <- SaveChanMsg{keys, []byte(lyric)}
			// 返回请求
			ctx.String(http.StatusOK, lyric)
			return
		}
		// 返回所有歌词，提供选择
		var lrcLists []openapi.LrcList
		for _, v := range mj.MusicSearchSearchCgiService.Data.Body.Song.List {
			lyric, err := openapi.GetLrc(v.Mid)
			if err != nil {
				logx.Error(err)
				ctx.String(http.StatusOK, err.Error())
				return
			}
			lyric = strings.ReplaceAll(lyric, "[by:]", "[by: Jiangwe Leo QQLrc]")
			lrcLists = append(lrcLists, openapi.LrcList{
				Id:     v.Mid,
				Title:  v.Title,
				Artist: artist,
				Lyrics: lyric,
			})
		}
		//returnMsg, _ := json.Marshal(lrcLists)
		ctx.Header("content-type", "application/json")
		ctx.JSONP(http.StatusOK, lrcLists)
		return
	}
	ctx.JSONP(http.StatusNotFound, gin.H{"success": false, "errMsg": "从QQ音乐获取歌词失败", "errCode": 40403})
	return
}

// LyricHandler 歌词信息确认
// gitbobobo 说是通知接口仅用于确认而已
// RequestMethod = POST
func LyricHandler(ctx *gin.Context) {
	var lyricCheck openapi.LyricCheck
	err := ctx.ShouldBindJSON(&lyricCheck)
	if err != nil {
		logx.Info("解析请求异常 ", err)
		return
	}
	keys := fmt.Sprintf("%v-%v", lyricCheck.Artist, lyricCheck.Title)
	lyric, err := openapi.GetLrc(lyricCheck.LyricsId)
	if err != nil {
		logx.Error(err)
		ctx.String(http.StatusOK, err.Error())
		return
	}
	lyric = strings.ReplaceAll(lyric, "[by:]", "[by: Jiangwe Leo QQLrc]")

	// 写入歌词
	saveFileChanel <- SaveChanMsg{keys, []byte(lyric)}
	// 无论是否成功缓存，均返回200
	ctx.JSONP(http.StatusOK, "ok")
}
