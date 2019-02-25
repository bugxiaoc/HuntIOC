package platform

import (
	"HuntIOC/core/tools"

	"github.com/gin-gonic/gin"
	"github.com/valyala/fasthttp"
)

type Threatbook struct {
	Name    string
	ApiPath string
	ApiKey  string
}

func NewTT() *Threatbook {
	return &Threatbook{
		Name:    "threatbook",
		ApiPath: "https://x.threatbook.cn/api/v1",
		ApiKey:  tools.GetPlatform("threatbook"),
	}
}

func (ad Threatbook) Query(c *gin.Context) {
	Method := "POST"
	types := c.Param("type")
	data := c.PostForm("data")
	var arg fasthttp.Args
	arg.Add("apikey", ad.ApiKey)

	switch types {
	case "hash":
		ad.ApiPath = ad.ApiPath + "/file/report"
		arg.Add("resource", data)
		//arg.Add("field", "")
	case "uri":
		ad.ApiPath = ad.ApiPath + "/dns"
		arg.Add("q", data)
	case "domain":
		ad.ApiPath = ad.ApiPath + "/domain/report"
		arg.Add("domain", data)
	case "ip":
		ad.ApiPath = ad.ApiPath + "/v1/ip/query"
		arg.Add("ip", data)
	case "ip-reputation":
		ad.ApiPath = ad.ApiPath + "/ip_reputation"
		arg.Add("ip", data)
	case "mail":
		ad.ApiPath = ad.ApiPath + "/ip-address/report"
		arg.Add("email", data)
	default:
		c.String(200, "{\"status\":\"00001\",\"msg\":\"该平台无法查询此参数.\"}")
	}

	body, status, err := tools.SendHttp(Method, ad.ApiPath, &arg, nil)
	if status == 200 {
		tmPath := tools.WriteFile(data+"-"+ad.Name, body)
		c.String(200, tmPath)
	} else if status == 500 {
		c.String(200, "{\"status\":\"00001\",\"msg\":\""+err.Error()+"\"}")
	} else {
		c.String(200, string(body))
	}
}
