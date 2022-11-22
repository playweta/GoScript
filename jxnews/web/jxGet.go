package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"jxnew/mode"
	"net/http"
)

var token string

func Get(url string) []mode.Tm {
	resp, err := http.Get(url)
	if err != nil {
		println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		println(err)
	}
	var m mode.Mode
	if err := json.Unmarshal(body, &m); err == nil {
		fmt.Println(m.Code)
	} else {
		fmt.Println(err)
	}
	if m.Result.Jxtoken != "" {
		token = m.Result.Jxtoken
	}
	return m.Result.Tm
}
