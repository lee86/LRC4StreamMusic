package openapi

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func POST(s string, data io.Reader) []byte {
	client := &http.Client{}
	reqest, err := http.NewRequest("POST", s, data) //建立一个请求
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(0)
	}
	//Add 头协议
	reqest.Header.Add("Host", "u.y.qq.com")
	reqest.Header.Add("Origin", "https://y.qq.com")
	reqest.Header.Add("Referer", "https://y.qq.com/n/ryqq/player")
	reqest.Header.Add("Content-Type", "application/json")
	response, err := client.Do(reqest) //提交
	fmt.Println(response)
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))
	return body
}
func GET(s string) []byte {
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", s, nil) //建立一个请求
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(0)
	}
	//Add 头协议
	reqest.Header.Add("Host", "")
	reqest.Header.Add("Referer", "https://y.qq.com/")
	response, err := client.Do(reqest) //提交

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))
	return body

}
