package main

import (
	"fmt"
	WebV "request/webV"
)

func main() {
	//resp, _ := http.Get("https://daohang.qq.com/?fr=hmpage")
	//fmt.Println(resp)
	x := WebV.RequestX{
		Url:    "https://discord.com/api/v9/channels/973198791157104710/messages",
		Submit: `{"content":"领不到水","nonce":"1043099472650829824","tts":false}`,
	}
	x.SetHeader("authorization: MTA0MjY3OTgwNTU1MzkzNDM1Ng.GZoi1G.VefG-jieOSqiWrYX0xMtV9Bg8p_4AOyx2XwR5g\ncontent-length: 68\ncontent-type: application/json\ncookie: __dcfduid=5afb1f70544f11edab1feb9f66ee5a80; __sdcfduid=5afb1f71544f11edab1feb9f66ee5a809627b4daf8829b4cd2af53c75b434e0b31d9a1da8b4a4deaca84c1ab63511786; __stripe_mid=afea82c9-dd31-40ff-aab7-d5f1a0efd5f582d4c6; __cfruid=46d080805a5fdd90b042d3e1d3bff1808f100f90-1668764309; __cf_bm=gUK2JnEVE5sVTZSd5kkj1JA34ZOGpqOuZ6gvoFksADI-1668764315-0-AQkEBfsgJG/7WeGZHLXazbUqbCeXHA6HGi02pRZBiyhFUYrSIv/GD/nqNlY5fKnUmDNpYyiot1l+AbFKdkbr38cEnxBKdT/SWaFilmzGuNHc8wC1F267lbsZqKZWIbbWjaNrDfzpDXWRfCBr8TH5H/Y=\norigin: https://discord.com\nreferer: https://discord.com/channels/973057323805311026/973198791157104710\nsec-ch-ua: \"Google Chrome\";v=\"107\", \"Chromium\";v=\"107\", \"Not=A?Brand\";v=\"24\"\nsec-ch-ua-mobile: ?0\nsec-ch-ua-platform: \"Windows\"\nsec-fetch-dest: empty\nsec-fetch-mode: cors\nsec-fetch-site: same-origin\nuser-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36\nx-debug-options: bugReporterEnabled\nx-discord-locale: zh-CN\nx-super-properties: eyJvcyI6IldpbmRvd3MiLCJicm93c2VyIjoiQ2hyb21lIiwiZGV2aWNlIjoiIiwic3lzdGVtX2xvY2FsZSI6InpoLUNOIiwiYnJvd3Nlcl91c2VyX2FnZW50IjoiTW96aWxsYS81LjAgKFdpbmRvd3MgTlQgMTAuMDsgV2luNjQ7IHg2NCkgQXBwbGVXZWJLaXQvNTM3LjM2IChLSFRNTCwgbGlrZSBHZWNrbykgQ2hyb21lLzEwNy4wLjAuMCBTYWZhcmkvNTM3LjM2IiwiYnJvd3Nlcl92ZXJzaW9uIjoiMTA3LjAuMC4wIiwib3NfdmVyc2lvbiI6IjEwIiwicmVmZXJyZXIiOiJodHRwczovL3N1aS5pby8iLCJyZWZlcnJpbmdfZG9tYWluIjoic3VpLmlvIiwicmVmZXJyZXJfY3VycmVudCI6IiIsInJlZmVycmluZ19kb21haW5fY3VycmVudCI6IiIsInJlbGVhc2VfY2hhbm5lbCI6InN0YWJsZSIsImNsaWVudF9idWlsZF9udW1iZXIiOjE1ODE4MywiY2xpZW50X2V2ZW50X3NvdXJjZSI6bnVsbH0=")
	//x.SetProxyAuth(proxy.GetAuth(orderno, secret))
	x.SetProxyStr("127.0.0.1:7890")
	fmt.Println(x.DoBodyToString())
}
