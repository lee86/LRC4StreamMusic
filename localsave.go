package main

import (
	"qqlrc/lyricCache"

	"github.com/zeromicro/go-zero/core/logx"
)

func save() {
	for {
		kv, ok := <-saveFileChanel
		if ok {
			if err, ok := lyricCache.CacheSave(kv.Key, kv.Value); ok {
				logx.Info(kv.Key, " - 本地缓存完成")
			} else {
				logx.Error(kv.Key, " - 本地缓存失败 ", err)
			}
		}
	}
}
