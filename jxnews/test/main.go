package main

import (
	"fmt"
	"jxnew/mode"
	sql "jxnew/mysql"
)

var answers string

func main() {

}

func queryQues(tms []mode.Tm) {
	for _, tm := range tms {
		one := sql.QueryOne(tm.Question)
		answers = fmt.Sprintf(answers, "|", one.Answer)
	}
}
