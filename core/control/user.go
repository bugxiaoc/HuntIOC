package control

import (
	"HuntIOC/core/tools"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

//GET Index
func Index(c *gin.Context) {
	session := sessions.Default(c)
	md5 := session.Get("md5")
	uname := session.Get("uname")
	if md5 != nil && md5 != nil {
		err := tools.CheckSession(uname, md5)
		if err {
			c.HTML(200, "index.html", gin.H{"pl": PlatForms})
		} else {
			c.HTML(http.StatusOK, "login.html", nil)
		}
	} else {
		c.HTML(http.StatusOK, "login.html", nil)
	}
}

//Login
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	md5 := tools.GetMd5ByName(username, password)
	if md5 == "" {
		c.JSON(http.StatusOK, gin.H{"status": 1000, "msg": "Login Error"})
	} else {
		session := sessions.Default(c)
		session.Set("md5", md5)
		session.Set("uname", username)
		session.Save()
		c.JSON(http.StatusOK, gin.H{"status": 0000, "msg": "Login Success"})
	}
}
