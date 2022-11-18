package proxy

import (
	"request/WebV"
	"strings"
	"sync"
	"time"
)

var lock = sync.Mutex{}

func GetIp(api string) string {
	lock.Lock()
	res := WebV.RequestX{
		Url: api,
	}
	ip := res.DoBodyToString()
	ip = strings.ReplaceAll(ip, "\n", "")
	ip = strings.ReplaceAll(ip, " ", "")
	time.Sleep(1 * time.Second)
	lock.Unlock()
	return ip
}
