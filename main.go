package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"qqlrc/openapi"
	"strings"
)

func main() {
	openapi.RouteLoadHandle("/lyric", lyricInfoHandler)
	openapi.RouteLoadHandle("/lyricInfo", lyricInfoHandler)
	openapi.StartServer()
}

func lyricInfoHandler(w http.ResponseWriter, r *http.Request) {
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

type MusicjsonN struct {
	Code                        int    `json:"code"`
	Ts                          int64  `json:"ts"`
	StartTs                     int64  `json:"start_ts"`
	Traceid                     string `json:"traceid"`
	MusicSearchSearchCgiService struct {
		Code int `json:"code"`
		Data struct {
			Body struct {
				Album struct {
					List []interface{} `json:"list"`
				} `json:"album"`
				Gedantip struct {
					Tab int    `json:"tab"`
					Tip string `json:"tip"`
				} `json:"gedantip"`
				Mv struct {
					List []interface{} `json:"list"`
				} `json:"mv"`
				Qc     []interface{} `json:"qc"`
				Singer struct {
					List []interface{} `json:"list"`
				} `json:"singer"`
				Song struct {
					List []struct {
						Act    int `json:"act"`
						Action struct {
							Alert    int `json:"alert"`
							Icon2    int `json:"icon2"`
							Icons    int `json:"icons"`
							Msgdown  int `json:"msgdown"`
							Msgfav   int `json:"msgfav"`
							Msgid    int `json:"msgid"`
							Msgpay   int `json:"msgpay"`
							Msgshare int `json:"msgshare"`
							Switch   int `json:"switch"`
							Switch2  int `json:"switch2"`
						} `json:"action"`
						Album struct {
							Id         int    `json:"id"`
							Mid        string `json:"mid"`
							Name       string `json:"name"`
							Pmid       string `json:"pmid"`
							Subtitle   string `json:"subtitle"`
							TimePublic string `json:"time_public"`
							Title      string `json:"title"`
						} `json:"album"`
						Bpm         int    `json:"bpm"`
						Content     string `json:"content"`
						Desc        string `json:"desc"`
						DescHilight string `json:"desc_hilight"`
						Docid       string `json:"docid"`
						Eq          int    `json:"eq"`
						Es          string `json:"es"`
						File        struct {
							B30S          int           `json:"b_30s"`
							E30S          int           `json:"e_30s"`
							HiresBitdepth int           `json:"hires_bitdepth"`
							HiresSample   int           `json:"hires_sample"`
							MediaMid      string        `json:"media_mid"`
							Size128Mp3    int           `json:"size_128mp3"`
							Size192Aac    int           `json:"size_192aac,omitempty"`
							Size192Ogg    int           `json:"size_192ogg"`
							Size24Aac     int           `json:"size_24aac"`
							Size320Mp3    int           `json:"size_320mp3"`
							Size360Ra     []interface{} `json:"size_360ra"`
							Size48Aac     int           `json:"size_48aac"`
							Size96Aac     int           `json:"size_96aac"`
							Size96Ogg     int           `json:"size_96ogg"`
							SizeApe       int           `json:"size_ape"`
							SizeDolby     int           `json:"size_dolby"`
							SizeDts       int           `json:"size_dts"`
							SizeFlac      int           `json:"size_flac"`
							SizeHires     int           `json:"size_hires"`
							SizeNew       []int         `json:"size_new"`
							SizeTry       int           `json:"size_try"`
							TryBegin      int           `json:"try_begin"`
							TryEnd        int           `json:"try_end"`
							Url           string        `json:"url"`
							Size192Aacc   int           `json:"size_192aacc,omitempty"`
						} `json:"file"`
						Fnote int `json:"fnote"`
						Genre int `json:"genre"`
						Grp   []struct {
							Act    int `json:"act"`
							Action struct {
								Alert    int `json:"alert"`
								Icon2    int `json:"icon2"`
								Icons    int `json:"icons"`
								Msgdown  int `json:"msgdown"`
								Msgfav   int `json:"msgfav"`
								Msgid    int `json:"msgid"`
								Msgpay   int `json:"msgpay"`
								Msgshare int `json:"msgshare"`
								Switch   int `json:"switch"`
								Switch2  int `json:"switch2"`
							} `json:"action"`
							Album struct {
								Id         int    `json:"id"`
								Mid        string `json:"mid"`
								Name       string `json:"name"`
								Pmid       string `json:"pmid"`
								Subtitle   string `json:"subtitle"`
								TimePublic string `json:"time_public"`
								Title      string `json:"title"`
							} `json:"album"`
							Bpm         int    `json:"bpm"`
							Content     string `json:"content"`
							Desc        string `json:"desc"`
							DescHilight string `json:"desc_hilight"`
							Docid       string `json:"docid"`
							Eq          int    `json:"eq"`
							Es          string `json:"es"`
							File        struct {
								B30S           int           `json:"b_30s"`
								E30S           int           `json:"e_30s"`
								HiresBitdepth  int           `json:"hires_bitdepth,omitempty"`
								HiresSample    int           `json:"hires_sample"`
								MediaMid       string        `json:"media_mid"`
								Size128Mp3     int           `json:"size_128mp3"`
								Size192Aac     int           `json:"size_192aac"`
								Size192Ogg     int           `json:"size_192ogg"`
								Size24Aac      int           `json:"size_24aac"`
								Size320Mp3     int           `json:"size_320mp3"`
								Size360Ra      []interface{} `json:"size_360ra"`
								Size48Aac      int           `json:"size_48aac"`
								Size96Aac      int           `json:"size_96aac"`
								Size96Ogg      int           `json:"size_96ogg"`
								SizeApe        int           `json:"size_ape"`
								SizeDolby      int           `json:"size_dolby"`
								SizeDts        int           `json:"size_dts"`
								SizeFlac       int           `json:"size_flac"`
								SizeHires      int           `json:"size_hires"`
								SizeNew        []int         `json:"size_new"`
								SizeTry        int           `json:"size_try"`
								TryBegin       int           `json:"try_begin"`
								TryEnd         int           `json:"try_end"`
								Url            string        `json:"url"`
								HiresBitdepth1 int           `json:"hires__bitdepth,omitempty"`
								HiresSBitdepth int           `json:"hires s_bitdepth,omitempty"`
							} `json:"file"`
							Fnote   int           `json:"fnote"`
							Genre   int           `json:"genre"`
							Grp     []interface{} `json:"grp"`
							Hotness struct {
								Desc     string `json:"desc"`
								IconUrl  string `json:"icon_url"`
								JumpType int    `json:"jump_type"`
								JumpUrl  string `json:"jump_url"`
							} `json:"hotness"`
							Href3      string `json:"href3"`
							Id         int    `json:"id"`
							IndexAlbum int    `json:"index_album"`
							IndexCd    int    `json:"index_cd"`
							Interval   int    `json:"interval"`
							Isonly     int    `json:"isonly"`
							Ksong      struct {
								Id  int    `json:"id"`
								Mid string `json:"mid"`
							} `json:"ksong"`
							Label        string `json:"label"`
							Language     int    `json:"language"`
							Lyric        string `json:"lyric"`
							LyricHilight string `json:"lyric_hilight"`
							Mid          string `json:"mid"`
							Mv           struct {
								Id    int    `json:"id"`
								Name  string `json:"name"`
								Title string `json:"title"`
								Vid   string `json:"vid"`
								Vt    int    `json:"vt"`
							} `json:"mv"`
							Name      string `json:"name"`
							NewStatus int    `json:"newStatus"`
							Ov        int    `json:"ov"`
							Pay       struct {
								PayDown    int `json:"pay_down"`
								PayMonth   int `json:"pay_month"`
								PayPlay    int `json:"pay_play"`
								PayStatus  int `json:"pay_status"`
								PriceAlbum int `json:"price_album"`
								PriceTrack int `json:"price_track"`
								TimeFree   int `json:"time_free"`
							} `json:"pay"`
							Protect int `json:"protect"`
							Sa      int `json:"sa"`
							Singer  []struct {
								Id    int    `json:"id"`
								Mid   string `json:"mid"`
								Name  string `json:"name"`
								Pmid  string `json:"pmid"`
								Title string `json:"title"`
								Type  int    `json:"type"`
								Uin   int    `json:"uin"`
							} `json:"singer"`
							Status       int       `json:"status"`
							Subtitle     string    `json:"subtitle"`
							Tag          int       `json:"tag"`
							Tid          int       `json:"tid"`
							TimePublic   string    `json:"time_public"`
							Title        string    `json:"title"`
							TitleHilight string    `json:"title_hilight"`
							Type         int       `json:"type"`
							Url          string    `json:"url"`
							Version      int       `json:"version"`
							Vf           []float64 `json:"vf"`
							Vi           []int     `json:"vi"`
							Volume       struct {
								Gain float64 `json:"gain"`
								Lra  float64 `json:"lra"`
								Peak float64 `json:"peak"`
							} `json:"volume"`
							Vs []string `json:"vs"`
						} `json:"grp"`
						Hotness struct {
							Desc     string `json:"desc"`
							IconUrl  string `json:"icon_url"`
							JumpType int    `json:"jump_type"`
							JumpUrl  string `json:"jump_url"`
						} `json:"hotness"`
						Href3      string `json:"href3"`
						Id         int    `json:"id"`
						IndexAlbum int    `json:"index_album"`
						IndexCd    int    `json:"index_cd"`
						Interval   int    `json:"interval"`
						Isonly     int    `json:"isonly"`
						Ksong      struct {
							Id  int    `json:"id"`
							Mid string `json:"mid"`
						} `json:"ksong"`
						Label        string `json:"label"`
						Language     int    `json:"language"`
						Lyric        string `json:"lyric"`
						LyricHilight string `json:"lyric_hilight"`
						Mid          string `json:"mid"`
						Mv           struct {
							Id    int    `json:"id"`
							Name  string `json:"name"`
							Title string `json:"title"`
							Vid   string `json:"vid"`
							Vt    int    `json:"vt"`
						} `json:"mv"`
						Name      string `json:"name"`
						NewStatus int    `json:"newStatus"`
						Ov        int    `json:"ov"`
						Pay       struct {
							PayDown    int `json:"pay_down"`
							PayMonth   int `json:"pay_month"`
							PayPlay    int `json:"pay_play"`
							PayStatus  int `json:"pay_status"`
							PriceAlbum int `json:"price_album"`
							PriceTrack int `json:"price_track"`
							TimeFree   int `json:"time_free"`
						} `json:"pay"`
						Protect int `json:"protect"`
						Sa      int `json:"sa"`
						Singer  []struct {
							Id    int    `json:"id"`
							Mid   string `json:"mid"`
							Name  string `json:"name"`
							Pmid  string `json:"pmid"`
							Title string `json:"title"`
							Type  int    `json:"type"`
							Uin   int64  `json:"uin"`
						} `json:"singer"`
						Status       int       `json:"status"`
						Subtitle     string    `json:"subtitle"`
						Tag          int       `json:"tag"`
						Tid          int       `json:"tid"`
						TimePublic   string    `json:"time_public"`
						Title        string    `json:"title"`
						TitleHilight string    `json:"title_hilight"`
						Type         int       `json:"type"`
						Url          string    `json:"url"`
						Version      int       `json:"version"`
						Vf           []float64 `json:"vf"`
						Vi           []int     `json:"vi"`
						Volume       struct {
							Gain float64 `json:"gain"`
							Lra  float64 `json:"lra"`
							Peak float64 `json:"peak"`
						} `json:"volume"`
						Vs []string `json:"vs"`
					} `json:"list"`
				} `json:"song"`
				Songlist struct {
					List []interface{} `json:"list"`
				} `json:"songlist"`
				User struct {
					List []interface{} `json:"list"`
				} `json:"user"`
				Zhida struct {
					List []interface{} `json:"list"`
				} `json:"zhida"`
			} `json:"body"`
			Code        int    `json:"code"`
			FeedbackURL string `json:"feedbackURL"`
			Meta        struct {
				Cid           string        `json:"cid"`
				Curpage       int           `json:"curpage"`
				Dir           string        `json:"dir"`
				DisplayOrder  []interface{} `json:"display_order"`
				Ein           int           `json:"ein"`
				EstimateSum   int           `json:"estimate_sum"`
				Expid         string        `json:"expid"`
				IsFilter      int           `json:"is_filter"`
				NextPageStart struct {
				} `json:"next_page_start"`
				Nextpage   int    `json:"nextpage"`
				Perpage    int    `json:"perpage"`
				Query      string `json:"query"`
				ReportInfo struct {
					Items struct {
					} `json:"items"`
				} `json:"report_info"`
				ResultTrustworthy  int    `json:"result_trustworthy"`
				Ret                int    `json:"ret"`
				SafetyType         int    `json:"safetyType"`
				SafetyUrl          string `json:"safetyUrl"`
				Searchid           string `json:"searchid"`
				Sid                string `json:"sid"`
				Sin                int    `json:"sin"`
				StepRelaSyntaxTree struct {
				} `json:"step_rela_syntax_tree"`
				Sum     int           `json:"sum"`
				TabList []interface{} `json:"tab_list"`
				Uid     string        `json:"uid"`
				V       int           `json:"v"`
			} `json:"meta"`
			Ver int `json:"ver"`
		} `json:"data"`
	} `json:"music"`
}
