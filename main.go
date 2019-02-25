package main

import (
	"HuntIOC/core/control"
	"HuntIOC/core/tools"
	"flag"
	"path/filepath"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func main() {
	var h bool
	configFilePath := flag.String("C", "./conf.json", "config file path")
	portConfig := flag.String("P", "80", "Port case: 80")
	flag.BoolVar(&h, "h", false, "Helper Menu")
	flag.Parse()
	if h {
		flag.PrintDefaults()
		return
	}
	//Load Config To Global
	if err := tools.LoadConfiguration(*configFilePath); err != nil {
		tools.LogInfo.Println("Config Init Err ", err.Error())
		return
	}
	router := gin.Default()
	store := sessions.NewCookieStore([]byte("secret"))
	store.Options(sessions.Options{
		MaxAge: int(30 * 60), //120min = 2H
		Path:   "/",
	})

	router.Use(sessions.Sessions("session", store))

	router.GET("/", control.Index)
	router.POST("/login", control.Login)

	router.POST("/q/:platform/:type", control.Router)

	router.GET("/set/:type", control.SetFileds)

	router.Static("/static", filepath.Join(tools.GetCurrentDirectory(), "/views/static/"))
	router.Static("/src", filepath.Join(tools.GetCurrentDirectory(), "/src/"))
	router.LoadHTMLGlob(filepath.Join(tools.GetCurrentDirectory(), "/views/template/**/*"))

	router.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{"host": c.Request.Host})
	})

	router.Run(":" + *portConfig)
}
