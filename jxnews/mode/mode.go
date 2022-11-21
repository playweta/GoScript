package mode

type Mode struct {
	Code   int    `json:"code"`
	Result Result `json:"result"`
	Msg    string `json:"msg"`
}

type Result struct {
	Ret         int    `json:"ret"`
	StepId      int    `json:"step_id"`
	Tel         string `json:"tel"`
	Tm          []Tm   `json:"tm"`
	AnswerRight string `json:"answer_right"`
	Jxtoken     string `json:"jxtoken"`
}

type Tm struct {
	Kind         int      `json:"kind"`
	Question     string   `json:"question"`
	AnswerOption []string `json:"answer_option"`
}
