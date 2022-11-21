package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	mode "jxnew/mode"
	"net/http"
)

var token string

func main() {
	//jxnews := "https://wxc.jxnews.com.cn/Weixin/Answertzbershida"
	//ip := "59.62.79.151"

	get("https://wxc.jxnews.com.cn/Weixin/Answertzbershida/sign_do.html?p=cH_App_zgwxRB_7F&name=%E9%BE%99%E6%96%8C&tel=13576792042&cip=59.62.72.108&area=360102&institution=%E6%B1%9F%E8%A5%BF%E8%BD%AF%E4%BB%B6%E8%81%8C%E4%B8%9A%E6%8A%80%E6%9C%AF%E5%A4%A7%E5%AD%A6")
	//post("https://wxc.jxnews.com.cn/Weixin/Answertzbershida/tj_do.html?p=cH_App_zgwxRB_7F&cip=59.62.79.151&answer=3||2||3||3||3||3||2||4||4||1||4||4||1||3||3||3||2||4||1||4")
}

func post(url string) {
	var body io.Reader
	resp, err := http.Post(url, "", body)
	if err != nil {
		println(err)
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		println(err)
	}
	fmt.Printf("%v\n", string(all))
}

func get(url string) {
	resp, err := http.Get(url)
	if err != nil {
		println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		println(err)
	}
	fmt.Println(resp.Body)
	fmt.Println(string(body))
	var m mode.Mode
	if err := json.Unmarshal(body, &m); err == nil {
		fmt.Println(m.Code)
	} else {
		fmt.Println(err)
	}
	if m.Result.Jxtoken != "" {
		token = m.Result.Jxtoken
	}
}
