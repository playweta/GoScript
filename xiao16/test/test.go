package main

import (
	"fmt"
	"request/Path"
	"request/file"
	"request/proxy"
	"request/webV"

	"strings"
	"sync"
	"time"
)

var suitest = "https://faucet.testnet.sui.io/gas"

var wg sync.WaitGroup
var XMApi = "http://route.xiongmaodaili.com/xiongmao-web/api/glip?secret=146aed44a89ebe9bccae2c6f88e96100&orderNo=GL20221122120741Kiuw1U8x&count=1&isTxt=1&proxyType=1"

func main() {
	arr := getAddresses()
	//arr := []string{"0xae9a613c93fbc6866f225a67f0a1d080f05c6782",
	//	"0x8bf89ff885aa31f86896e43c819b1967d3a38cfe",
	//	"0x00c1ee9e1f896dd4aa2bd9bd01711f48bce4d6ca",
	//	"0x8103204b388fd077a0f6d27726bed3364daf8365",
	//	"0x0b7eda03c083ae416adeb65c6f71e0b6bcdf94a8",
	//}
	araLen := len(arr)
	for i := 0; i < araLen; i++ {
		wg.Add(1)
		go reqestFauct(arr[i])
		if i%6 == 5 {
			time.Sleep(time.Second * 5)
		}
	}
	wg.Wait()
}
func reqestFauct(account string) {
	resp := webV.RequestX{
		Url:    suitest,
		Submit: `{"FixedAmountRequest":{"recipient":"` + account + `"}}`,
	}
	//resp.SetProxy("127.0.0.1:7890")
	resp.SetHeader(`content-type: application/json
origin: chrome-extension://opcgpfmipidbgpenhmajoajpbobppdil
sec-ch-ua: "Google Chrome";v="107", "Chromium";v="107", "Not=A?Brand";v="24"
sec-ch-ua-mobile: ?0
sec-ch-ua-platform: "Windows"
sec-fetch-dest: empty
sec-fetch-mode: cors
sec-fetch-site: none
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36`)
	ip := proxy.GetIp(XMApi)
	fmt.Println(account, ip)
	resp.SetProxy(ip)
	//resp.SetProxy(proxyServer)
	//resp.SetProxyAuth(proxy.GetAuth(orderno, secret))
	toString1 := resp.DoBodyToString()
	resp.Do()

	if strings.Index(toString1, "!DOCTYPE") != -1 {
		fmt.Println("DOCTYPE")
	} else {
		fmt.Println(toString1)
	}

	wg.Done()
}
func getAddresses() []string {
	path := Path.GetCurrentAbPathByCaller(1) + "/keys.txt"
	readFile, _ := file.ReadFile(path)
	addresses := strings.Split(string(readFile), "\n")
	arr := make([]string, len(addresses), len(addresses))
	for i := range addresses {
		if len(addresses[i]) < 41 {
			continue
		}
		arr[i] = addresses[i][:42]
	}
	return arr
}
