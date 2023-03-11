package web

import (
	"fmt"
	"jxnew/mode"
	my "jxnew/mysql"
	"time"
)

var t = "https://wxc.jxnews.com.cn/Weixin/Answertzbershida/"               // 地址
var cip = "222.104.171.128"                                                // ip
var institution = "%E5%8D%97%E6%98%8C%E8%81%8C%E4%B8%9A%E5%A4%A7%E5%AD%A6" // 学校 // 答案
var getArr []mode.Tm

func getAnswerArr() (answer string) {
	getArr = Get(urlGet)
	for i := 0; i < 20; i++ {
		one := my.QueryOne(getArr[i].Question)
		if i != 19 {
			answer = fmt.Sprint(answer, one.Answer, "||")
		} else {
			answer = fmt.Sprint(answer, one.Answer)
		}
	}
	return answer
}

var urlGet = fmt.Sprint(t, "sign_do.html?p=cH_App_zgwxRB_7F&name=%E5%88%98%E7%82%B3%E8%8D%A3&tel=15970645022&cip=", cip, "&area=360101&institution=", institution)
var urlPost = fmt.Sprint(t, "tj_do.html?p=cH_App_zgwxRB_7F&cip=", cip, "&answer=", getAnswerArr())

func WebX() {
	//jxnews := "https://wxc.jxnews.com.cn/Weixin/Answertzbershida"
	//ip := "59.62.79.151"

	// 创建一个一个定时器
	fmt.Println("定时器开始")
	myT := time.NewTimer(17.6554 * 1000 * 1000 * 1000) // 设置7秒后执行一次
	<-myT.C
	fmt.Println(" 18 s 时间到 ", time.Now().Unix())

	answers := Post(urlPost)
	//for i := 0; i < 20; i++ {
	//	fmt.Println(answers[i], getArr[i].Question)
	//}

	for i := 0; i < 20; i++ {
		my.Add(mode.TM{
			Question: getArr[i].Question,
			Answer:   answers[i], // 答案
		})
	}
}
