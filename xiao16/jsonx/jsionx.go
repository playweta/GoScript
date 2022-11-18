package jsonx

import "encoding/json"

func Parse(arg string) map[string]interface{} {
	var temp interface{}
	_ = json.Unmarshal([]byte(arg), &temp)
	return temp.(map[string]interface{})
}

func String(v any) string {
	marshal, _ := json.Marshal(v)
	return string(marshal)
}
