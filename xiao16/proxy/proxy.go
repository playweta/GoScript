package proxy

import (
	"request/webV"
	"strings"
	"sync"
	"time"
)

var lock = sync.Mutex{}

func GetIp(api string) string {
	lock.Lock()
	res := webV.RequestX{
		Url: api,
	}
	ip := res.DoBodyToString()
	ip = strings.ReplaceAll(ip, "\n", "")
	ip = strings.ReplaceAll(ip, " ", "")
	time.Sleep(1 * time.Second)
	lock.Unlock()
	return ip
}
