package web

import (
	"fmt"
	"jxnew/mode"
	my "jxnew/mysql"
	"time"
)

/*
*
https://wxc.jxnews.com.cn/Weixin/Answerjifa/sign_do.html
?p=cH_App_zgwxRB_7F
&name=%E9%AD%8F%E6%98%A5%E4%B8%BD&tel=18279973274&area=360302&institution=%E8%90%8D%E4%B9%A1%E5%AD%A6%E9%99%A2
*/
var t = "https://wxc.jxnews.com.cn/Weixin/Answerjifa/" // 地址
// var cip = "222.104.171.128"                                                // ip
// var institution = "%E5%8D%97%E6%98%8C%E8%81%8C%E4%B8%9A%E5%A4%A7%E5%AD%A6" // 学校 // 答案
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

// sign_do.html?p=cH_App_zgwxRB_7F&name=秦尘&tel=13576792042&area=360112&institution=江西软件职业技术大学
var urlGet = fmt.Sprint(t, "sign_do.html?p=cH_App_zgwxRB_7F&name=秦尘&tel=13576792042&area=360112&institution=江西软件职业技术大学")

//var urlGet = fmt.Sprint(t, "sign_do.html?p=cH_App_zgwxRB_7F&name=%E9%AD%8F%E6%98%A5%E4%B8%BD&tel=18279973274&area=360302&institution=%E8%90%8D%E4%B9%A1%E5%AD%A6%E9%99%A2")

//var urlGet = fmt.Sprint(t, "sign_do.html?p=cH_App_zgwxRB_7F&name=%E5%88%98%E7%82%B3%E8%8D%A3&tel=15970645022&area=360112&institution=%E5%8D%97%E6%98%8C%E8%81%8C%E4%B8%9A%E5%A4%A7%E5%AD%A6")

// tj_do.html?p=cH_App_zgwxRB_7F&answer=1%7C%7C1%7C%7C2%7C%7C2%7C%7C3%7C%7C4%7C%7C3%7C%7C4%7C%7C1%7C%7C4%7C%7C3%7C%7C4%7C%7C4%7C%7C4%7C%7C3%7C%7C3,4%7C%7C3,4%7C%7C3,4%7C%7C3,4%7C%7C3,4
// var urlPost = fmt.Sprint(t, "tj_do.html?p=cH_App_zgwxRB_7F&cip=", cip, "&answer=", getAnswerArr())
var urlPost = fmt.Sprint(t, "tj_do.html?p=cH_App_zgwxRB_7F&answer=", getAnswerArr())

func WebX() {
	//jxnews := "https://wxc.jxnews.com.cn/Weixin/Answertzbershida"
	//ip := "59.62.79.151"

	// 创建一个一个定时器
	fmt.Println("定时器开始")
	myT := time.NewTimer(18 * 1000 * 1000 * 1000) // 设置7秒后执行一次
	<-myT.C
	fmt.Println(" 30 s 时间到 ", time.Now().Unix())

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
