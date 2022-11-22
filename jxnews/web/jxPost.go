package web

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

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
