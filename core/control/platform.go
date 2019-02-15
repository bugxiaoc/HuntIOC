package control

import (
	"HuntIOC/core/platform"

	"github.com/gin-gonic/gin"
)

var PlatForms map[string]platform.Base

func init() {
	PlatForms = platform.Init()
}

func Router(c *gin.Context) {
	pf := c.Param("platform")
	if PlatForms[pf] != nil {
		PlatForms[pf].Query(c)
	} else {
		c.JSON(200, gin.H{"status": "00001", "msg": "该平台未实现."})
	}

}
