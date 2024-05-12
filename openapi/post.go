package openapi

import (
	"fmt"
	"io"
	"net/http"
)

func POST(s string, data io.Reader) []byte {
	client := &http.Client{}
	request, err := http.NewRequest("POST", s, data) //建立一个请求
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return []byte("")
		//os.Exit(0)
	}
	//Add 头协议
	request.Header.Add("Host", "u.y.qq.com")
	request.Header.Add("Origin", "https://y.qq.com")
	request.Header.Add("Referer", "https://y.qq.com/n/ryqq/player")
	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request) //提交
	if err != nil {
		// 打印 error
		fmt.Println(err)
		return []byte("")
	}
	//fmt.Println(response)
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		// 打印 error 返回的信息
		fmt.Println(string(body))
	}
	return body
}
func GET(s string) []byte {
	client := &http.Client{}
	request, err := http.NewRequest("GET", s, nil) //建立一个请求
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return []byte("")
	}
	//Add 头协议
	request.Header.Add("Host", "")
	request.Header.Add("Referer", "https://y.qq.com/")
	response, err := client.Do(request) //提交
	if err != nil {
		// 打印 error
		fmt.Println(err)
		return []byte("")
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		// 打印 error 返回的信息
		fmt.Println(string(body))
	}
	return body
}
