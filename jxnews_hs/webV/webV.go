package webV

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

type RequestX struct {
	Url       string
	Submit    string
	Header    http.Header
	Cookies   []*http.Cookie
	ProxyStr  string
	ProxyAuth string
	Timeout   time.Duration
}

func (req *RequestX) SetUrl(url string) {
	req.Url = url
}
func (req *RequestX) SetSubmit(sub string) {
	req.Submit = sub
}
func (req *RequestX) SetHeader(heard string) {
	if req.Header == nil {
		req.Header = map[string][]string{}
	}
	types := strings.Split(heard, "\n")

	for i := range types {
		l := strings.Split(types[i], ":")
		//fmt.Println(l[0], l[1])
		req.Header.Add(l[0], l[1])
	}
}
func (req *RequestX) SetCookies(ck []*http.Cookie) {
	req.Cookies = ck
}

func (req *RequestX) SetProxyAuth(auth string) {

	req.ProxyAuth = auth
	if req.Header == nil {
		req.Header = map[string][]string{}
	}
	req.Header.Add("Proxy-Authorization", auth)
}
func (req *RequestX) SetProxy(proxy string) {
	ip := strings.ReplaceAll(proxy, "\n", "")
	ip = strings.ReplaceAll(proxy, " ", "")
	ip = strings.ReplaceAll(proxy, "\r", "")
	if strings.Index(proxy, "http") == -1 {
		ip = "http://" + ip
	}
	req.ProxyStr = ip
}
func (req *RequestX) Do() (resp *http.Response, err error) {
	req.Url = strings.ReplaceAll(req.Url, " ", "")
	var netTransport = &http.Transport{}
	var ck http.CookieJar = nil
	if req.Timeout == time.Duration(0) {
		req.Timeout = time.Second * 60
	}
	if req.ProxyStr != "" {
		//netTransport = &http.Transport{}
		//fmt.Println("proxy", req.ProxyStr)
		proxy, proxyError := url.Parse(req.ProxyStr)
		if proxyError != nil {
			fmt.Println(req.ProxyStr)
			fmt.Println(proxy)

			panic("proxy error Nil")
		}
		netTransport = &http.Transport{
			Proxy: http.ProxyURL(proxy),
			//ResponseHeaderTimeout: req.Timeout,
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
	if req.Cookies != nil {
		ck = &cookiejar.Jar{}
		u, _ := url.Parse(req.Url)
		ck.SetCookies(u, req.Cookies)
	}
	https := &http.Client{
		//Timeout:   time.Second * 20,
		Transport: netTransport,
		Jar:       ck,
	}
	if req.Submit == "" || len(req.Submit) == 0 {

		//fmt.Println("GET")
		re, err := http.NewRequest("GET", req.Url, nil)
		re.Header = req.Header
		if err != nil {
			return nil, err
		}
		return https.Do(re)
		//return https.Get(req.Url)
	} else {
		if req.Header == nil {
			req.Header = map[string][]string{}
		}
		//fmt.Println("POST")
		if strings.Index(req.Submit, "{") != -1 {
			req.Header["Content-Type"] = []string{"application/json"}
		}
		re, err := http.NewRequest("POST", req.Url, strings.NewReader(req.Submit))

		re.Referer()
		re.Header = req.Header
		if err != nil {
			return nil, err
		}
		/*	js, _ := json.Marshal(re.Header)
			fmt.Println(string(js))*/
		return https.Do(re)
		//return https.Post(req.Url, req.ContentType, strings.NewReader(req.Submit))
	}
}
func (req *RequestX) DoBodyToString() string {
	resp, err := req.Do()

	if err != nil {
		return "error req.Do:" + err.Error()
	}
	ioBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "error ioutil.ReadAll:" + err.Error()
	}

	return string(ioBody)
}
