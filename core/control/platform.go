package control

import (
	"HuntIOC/core/platform"
	"HuntIOC/core/tools"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

var PlatForms map[string]platform.Base

func init() {
	PlatForms = platform.Init()
}

func Router(c *gin.Context) {
	session := sessions.Default(c)
	pf := c.Param("platform")
	tools.LogInfo.Printf("User:%s, Query : %s, Type %s, Data: %s", session.Get("uname"), c.Param("platform"), c.Param("type"), c.PostForm("data"))
	if PlatForms[pf] != nil {
		PlatForms[pf].Query(c)
	} else {
		c.JSON(200, gin.H{"status": "00001", "msg": "该平台未实现."})
	}
}

func SetFileds(c *gin.Context){
	filedName := c.Param("type")
	data := c.Query("data")

	if tools.SetFiled(filedName, data) {
		c.String(200, data)
	} else {
		c.String(500, "")
	}
}