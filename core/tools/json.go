package tools

import (
	"encoding/json"
	"net/url"
	"strconv"
)

func JsonToStr(maps interface{}) string {
	resultByte, err := json.Marshal(maps)
	if err != nil {
		panic(err)
	}
	return string(resultByte)
}

func JsonWriteCsv(name string, data []byte) string {
	csvPath := "./src/csv/" + name + ".csv"
	jsonPath := "./src/json/" + name + ".json"
	WriteFile(jsonPath, data)
	_ = OpenFile(csvPath, true)
	var Maps map[string]interface{}
	err := json.Unmarshal(data, &Maps)
	if err != nil {
		LogErr.Println("Json Unmarshal Error , Err data:", string(data))
		return ""
	}
	if Maps["data"] == "" {
		return ""
	}
	rList := Maps["data"].([]interface{})
	vMap := map[string][]string{}
	for index := range rList {
		line := rList[index].(map[string]interface{})
		for k, v := range line {
			switch v.(type) {
			case string:
				vMap[k] = append(vMap[k], v.(string))
			case int:
				vMap[k] = append(vMap[k], string(v.(int)))
			case float32:
				vMap[k] = append(vMap[k], strconv.FormatFloat(float64(v.(float32)), 'E', -1, 32))
			case float64:
				vMap[k] = append(vMap[k], strconv.FormatFloat(v.(float64), 'E', -1, 64))
			default:
				vMap[k] = append(vMap[k], "Nil Type")
			}
		}
	}
	var csvHeader []string
	var csvHeaderCN []string
	for k, _ := range vMap {
		csvHeader = append(csvHeader, k)
		csvHeaderCN = append(csvHeaderCN, GetLangage(k))
	}
	WriteCsv(csvPath, csvHeaderCN)
	if len(csvHeader) <= 0 {
		return ""
	}
	for index := 0; index < len(vMap[csvHeader[0]]); index ++ {
		line := []string{}
		for head := range csvHeader {
			lsTmp := vMap[csvHeader[head]]
			if index >= len(lsTmp) {
				line = append(line, "无该字段.")
			} else {
				decodeurl, err := url.QueryUnescape(lsTmp[index])
				if err != nil {
					LogErr.Println("Unescape Error:", err.Error(),"---",lsTmp[index])
				}
				line = append(line, decodeurl)
			}
		}
		WriteCsv(csvPath, line)
	}
	return csvPath
}

func GetLangage(key string) string {
	switch key {
	case "url":
		return "URL地址"
	case "rp":
		return "远程地址端口"
	case "filepath":
		return "样本路径"
	case "esid":
		return "主机唯一标识"
	case "ip":
		return "目标IP地址"
	case "_ts":
		return "URL地址"
	case "domain":
		return "域名"
	case "filelevel":
		return "文件等级"
	case "ndtype":
		return "URL地址"
	case "_row":
		return "URL地址"
	case "clientip":
		return "客户端IP"
	case "time":
		return "时间"
	case "HiScl":
		return "当前进程的启动参数"
	case "HiItn":
		return "窗口标题"
	case "HiUzp":
		return "URL地址"
	case "hisrc":
		return "父进程绝对路径"
	case "HiOrn":
		return "当前进程名"
	case "HiDst":
		return "当前进程全路径"
	case "HiCle":
		return "当前进程子进程参数"
	case "HiMD5PAR":
		return "父进程md5"
	case "HiDlt":
		return "HiDlt未知"
	case "count":
		return "访问次数"
	case "first_seen":
		return "第一次访问时间"
	case "last_seen":
		return "最后一次访问时间"
	case "date":
		return "访问时间"
	case "host":
		return "Url的host"
	default:
		return "未知字段" + key
	}
}
