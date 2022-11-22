package mode

type ModePost struct {
	Code   int `json:"code"`
	Result struct {
		Ret         int    `json:"ret"`
		Score       int    `json:"score"`
		Username    string `json:"username"`
		AnswerRight string `json:"answer_right"`
		Pv          int    `json:"pv"`
		Tj          int    `json:"tj"`
		Ctime       string `json:"ctime"`
	} `json:"result"`
	Msg string `json:"msg"`
}
