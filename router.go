package main

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	musicmid := query.Get("musicmid")
	fmt.Fprint(w, getlrc(musicmid))
}

func GetMidHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	musicname := query.Get("musicname")
	num := query.Get("num")
	//id := query.Get("id")
	//fmt.Fprint(w, musicname)
	if num == "" {
		num = "20"
	}
	fmt.Fprint(w, getsongid(musicname, num))
}
