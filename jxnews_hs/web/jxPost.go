package web

import (
	"encoding/json"
	"fmt"
	"jxnew/mode"
	"jxnew/webV"
)

func Post(url string) (answers [20]string) {
	x := webV.RequestX{Url: url}
	x.SetHeader(fmt.Sprint("jxtoken: ", Token))

	body := x.DoBodyToString()
	var m mode.ModePost
	if err := json.Unmarshal([]byte(body), &m); err == nil {
		fmt.Println("姓名：", m.Result.Username, "分数：", m.Result.Score, "时间：", m.Result.Ctime, "请求：", m.Msg)
	} else {
		fmt.Println(err)
	}

	answerRight := m.Result.AnswerRight
	a := 0
	arg := ""
	for i := 0; i < len(answerRight); i++ {
		if answerRight[i] == '|' && i+1 < len(answerRight) {
			i++
		} else if (i+1 < len(answerRight)) && (answerRight[i] == ',' || answerRight[i+1] == ',') {
			arg = fmt.Sprint(arg, string(answerRight[i]))
		} else {
			arg = fmt.Sprint(arg, string(answerRight[i]))
			answers[a] = arg
			a++
			arg = ""
		}

	}
	return answers
}
