package WebV

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
	Url      string
	Submit   string
	Header   http.Header
	Cookies  []*http.Cookie
	ProxyStr string
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
		req.Header.Add(l[0], l[1])
	}
}
func (req *RequestX) SetCookies(ck []*http.Cookie) {
	req.Cookies = ck
}
func (req *RequestX) SetProxyStr(proxy string) {
	ip := strings.ReplaceAll(proxy, "\n", "")
	ip = strings.ReplaceAll(proxy, " ", "")
	if strings.Index(proxy, "http") == -1 {
		ip = "http://" + ip
	}
	req.ProxyStr = ip
}
func (req *RequestX) Do() (resp *http.Response, err error) {
	req.Url = strings.ReplaceAll(req.Url, " ", "")
	//var netTransport http.RoundTripper = nil
	var netTransport *http.Transport = &http.Transport{}
	var ck http.CookieJar = nil
	if req.ProxyStr != "" {
		//netTransport = &http.Transport{}
		//fmt.Println("proxy", req.ProxyStr)
		proxy, _ := url.Parse(req.ProxyStr)
		//fmt.Println("proxy", proxy)
		netTransport = &http.Transport{
			Proxy:                 http.ProxyURL(proxy),
			ResponseHeaderTimeout: time.Second * time.Duration(15),
			TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		}
	}
	if req.Cookies != nil {
		ck = &cookiejar.Jar{}
		u, _ := url.Parse(req.Url)
		ck.SetCookies(u, req.Cookies)
	}
	https := &http.Client{
		Timeout:   time.Second * 20,
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

	fmt.Println(resp, err)
	if err != nil {
		return "error:" + err.Error()
	}
	ioBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "error:" + err.Error()
	}
	return string(ioBody)
}
