package CaptchaRunApi

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"request/WebV"
	"time"
)

//https://captcha.run/sso?inviter=ec5a2891-b02d-44c8-8660-5554c28bb627
// Captcha 判断是 V2 还是 V3
// 请在控制台中执行命令 ___grecaptcha_cfg.clients

type CaptchaPost struct {
	CaptchaType string `json:"captchaType"`
	SiteKey     string `json:"siteKey"`
	SiteReferer string `json:"siteReferer"`
	DeveloperId string `json:"developerId"`
}
type CaptchaRes struct {
	Id          string `json:"_id"`
	Account     string `json:"account"`
	Price       string `json:"price"`
	Deducted    bool   `json:"deducted"`
	CaptchaType string `json:"captchaType"`
	Request     struct {
		IsInvisible bool   `json:"isInvisible"`
		UseCache    bool   `json:"useCache"`
		CaptchaType string `json:"captchaType"`
		SiteKey     string `json:"siteKey"`
		SiteReferer string `json:"siteReferer"`
		DeveloperId string `json:"developerId"`
	} `json:"request"`
	Cache    bool      `json:"cache"`
	Status   string    `json:"status"`
	Ip       string    `json:"ip"`
	Created  time.Time `json:"created"`
	Finished time.Time `json:"finished"`
	Response struct {
		GRecaptchaResponse string `json:"gRecaptchaResponse"`
	} `json:"response"`
}

// ReCaptchaV2
func tasksV2(token string, dataSiteKey string, referer string) (string, error) {
	marshal, _ := json.Marshal(CaptchaPost{
		CaptchaType: "ReCaptchaV2",
		SiteKey:     dataSiteKey,
		SiteReferer: referer,
		DeveloperId: "0bf9e614-e56c-421e-9b82-a3ccbdddb59a",
	})
	req := &WebV.RequestX{
		Url: "https://api.captcha.run/v2/tasks",
		Header: map[string][]string{"Content-Type": []string{"application/json"},
			"Authorization": []string{"Bearer " + token},
		},
		Submit: string(marshal),
	}
	resp, err := req.Do()
	bytes, _ := ioutil.ReadAll(resp.Body)
	var temp interface{}
	_ = json.Unmarshal(bytes, &temp)
	data := temp.(map[string]interface{})
	id := data["taskId"].(string)
	return id, err
}
func getTasksResult(token string, taskId string) (string, error) {
	var gRecaptchaResponse string
	for i := 0; i < 30; i++ {
		//Working
		req := &WebV.RequestX{
			Url: "https://api.captcha.run/v2/tasks/" + taskId,
			Header: map[string][]string{"Content-Type": []string{"application/json"},
				"Authorization": []string{"Bearer " + token},
			},
		}
		resp, _ := req.Do()
		str := CaptchaRes{}
		bytes, _ := ioutil.ReadAll(resp.Body)
		_ = json.Unmarshal([]byte(bytes), &str)
		if len(str.Response.GRecaptchaResponse) > 0 {
			gRecaptchaResponse = str.Response.GRecaptchaResponse
			break
		}
		time.Sleep(5 * time.Second)
	}
	return gRecaptchaResponse, errors.New("not Response")
}

func ResponseV2(token string, dataSiteKey string, referer string) string {
	taskId, _ := tasksV2(token, dataSiteKey, referer)
	result, _ := getTasksResult(token, taskId)
	return result
}

//实例
/*	res := CaptchaRunApi.ResponseV2(
		"0bf9e614-e56c-421e-9b82-a3ccbdddb59a",
		"6LdxWtcgAAAAAIVZ6LcqknOZERq-pU6KE_C8Sc_d",
		"https://test2.gno.land/faucet",
	)
	resp := WebV.RequestX{
		Url:    "https://faucet.test2.gno.land/",
		Submit: "toaddr=g1f86r8j2umd879c69usnmgmsxgz7h88spm2pshx&g-recaptcha-response=" + res,
	}
	resp.SetHeader("content-type: application/x-www-form-urlencoded")
	do, _ := resp.Do()
	all, _ := ioutil.ReadAll(do.Body)
	fmt.Println(string(all))*/
