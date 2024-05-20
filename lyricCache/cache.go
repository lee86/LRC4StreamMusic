package lyricCache

import (
	"log"
	"os"
	"path"
)

func CacheSelect(keys string) (lyrics string, ok bool) {
	lyric, err := os.ReadFile(path.Join(config.Base.CacheDir, keys))
	if err != nil {
		return "", false
	}
	return string(lyric), true
}

func CacheSave(keys string, lyric []byte) bool {
	fileName := path.Join(config.Base.CacheDir, keys)
	if _, err := os.Stat(fileName); err == nil {
		err := os.Remove(fileName)
		if err != nil {
			return false
		}
	}
	if file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755); err == nil {
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Println(err)
				return
			}
		}(file)
		// 写入歌词文件
		_, err = file.Write(lyric)
		if err != nil {
			return false
		}
	}
	return true
}
