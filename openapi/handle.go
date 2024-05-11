package openapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func LyricInfoHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.RequestURI)
	query := r.URL.Query()
	title := query.Get("title")
	artist := query.Get("artist")
	num := query.Get("limit")

	var mj MusicjsonN
	err := json.Unmarshal([]byte(getsongid(fmt.Sprintf("%v %v", artist, title), num)), &mj)
	fmt.Println(err, mj)
	if err != nil {
		fmt.Println(err)
	}
	if len(mj.MusicSearchSearchCgiService.Data.Body.Song.List) >= 1 {
		lrc := getlrc(mj.MusicSearchSearchCgiService.Data.Body.Song.List[0].Mid)
		lrc = strings.ReplaceAll(lrc, "[by:]", "[by: Jiangwe Leo QQLrc]")
		fmt.Fprint(w, lrc)
		return
	}
	fmt.Fprint(w, "have no msg")
}
