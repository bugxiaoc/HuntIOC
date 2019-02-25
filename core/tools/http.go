package tools

import (
	"github.com/valyala/fasthttp"
)

func SendHttp(method, URI string, arg *fasthttp.Args, cookies map[string]interface{}) ([]byte, int, error) {
	req := &fasthttp.Request{}
	switch method {
	case "GET":
		req.Header.SetMethod(method)
		URI = URI + "?" + arg.String()
	case "POST":
		req.Header.SetMethod(method)
		arg.WriteTo(req.BodyWriter())
	}
	LogInfo.Println("Query [Arg]:", arg.String())
	if cookies != nil {
		for key, v := range cookies {
			req.Header.SetCookie(key, v.(string))
		}
	}
	req.SetRequestURI(URI)
	resp := &fasthttp.Response{}
	var client fasthttp.Client
	if err := client.Do(req, resp); err != nil {
		LogInfo.Printf("Http Do Filed %v", err)
		return []byte{}, 500, err
	} else {
		//LogInfo.Println(req.String())
		return resp.Body(), resp.StatusCode(), err
	}

}
