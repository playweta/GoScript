package email

import (
	"io/ioutil"
	"request/WebV"
	"time"
)

func GetCode(em string) string {
	for i := 0; i < 15; i++ {
		emailReq := WebV.RequestX{
			Url: em,
		}
		resp, _ := emailReq.Do()
		all, _ := ioutil.ReadAll(resp.Body)

		if len(string(all)) >= 1 {
			return string(all)
		}
		time.Sleep(5 * time.Second)
	}
	return string("")
}
