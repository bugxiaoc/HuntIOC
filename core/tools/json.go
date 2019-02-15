package tools

import (
	"encoding/json"
)

func JsonStrToMap(str string) map[string]interface{} {
	var Maps map[string]interface{}
	err := json.Unmarshal([]byte(str), &Maps)
	if err != nil {
		panic(err)
	}
	return Maps
}

func JsonToStr(maps interface{}) string {
	resultByte, err := json.Marshal(maps)
	if err != nil {
		panic(err)
	}
	return string(resultByte)
}
