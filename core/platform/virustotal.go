package platform

import (
	"HuntIOC/core/tools"

	"github.com/gin-gonic/gin"
	"github.com/valyala/fasthttp"
)

type Virustotal struct {
	Name    string
	ApiPath string
	ApiKey  string
}

func NewVirus() *Virustotal {
	return &Virustotal{
		Name:    "Virustotal",
		ApiPath: "www.virustotal.com/vtapi/v2/",
		ApiKey:  "",
	}
}

func (ad Virustotal) Query(c *gin.Context) {
	Method := "GET"
	types := c.Param("type")
	data := c.PostForm("data")
	var arg fasthttp.Args
	arg.Add("apikey", ad.ApiKey)
	switch types {
	case "hash":
		ad.ApiPath = ad.ApiPath + "/file/report"
		arg.Add("resource", data)
		arg.Add("allinfo", "true")
	case "uri":
		ad.ApiPath = ad.ApiPath + "/url/report"
		arg.Add("resource", data)
		arg.Add("allinfo", "true")
		arg.Add("scan", "1")
	case "domain":
		ad.ApiPath = ad.ApiPath + "/domain/report"
		arg.Add("domain", data)
	case "ip":
		ad.ApiPath = ad.ApiPath + "/ip-address/report"
		arg.Add("IP", data)
	default:
		c.String(200, "{\"status\":\"00001\",\"msg\":\"该平台无法查询此参数.\"}")
	}

	body, status, err := tools.SendHttp(Method, ad.ApiPath, &arg, nil)
	if status == 204 {
		c.String(200, "{\"status\":\"00001\",\"msg\":\"超出请求率限制.\"}")
	} else if status == 400 {
		c.String(200, "{\"status\":\"00001\",\"msg\":\"请求参数错误.\"}")
	} else if status == 43 {
		c.String(200, "{\"status\":\"00001\",\"msg\":\"权限不足.\"}")
	} else {
		c.String(200, string(body))
	}
}
