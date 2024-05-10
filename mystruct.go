package main

type (
	Lrcjson struct {
		Retcode int    `json:"retcode"`
		Code    int    `json:"code"`
		Subcode int    `json:"subcode"`
		Songt   int    `json:"songt"`
		Lyric   string `json:"lyric"`
	}
)

type (
	Musicjson struct {
		Code int `json:"code"`
		Data struct {
			Keyword string `json:"keyword"`
			Song    struct {
				Curnum int `json:"curnum"`
				List   []List
			}
		}
	}
)
type List struct {
	Albumname string `json:"albumname"`
	Lyric     string `json:"lyric"`
	Songname  string `json:"songname"`
	Songmid   string `json:"songmid"`
	Singer    []Singer
}
type Singer struct {
	Name string `json:"name"`
}

type Conf struct {
	Net struct {
		Port string `yaml:"port"`
	}
}
