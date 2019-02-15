package main

import (
	"HuntIOC/core/control"
	"HuntIOC/core/tools"
	"flag"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	var h bool
	//configFilePath := flag.String("C", "conf.json", "config file path")
	portConfig := flag.String("P", "80", "Port case: 80")
	flag.BoolVar(&h, "h", false, "Helper Menu")
	flag.Parse()
	if h {
		flag.PrintDefaults()
		return
	}

	router := gin.Default()
	router.GET("/q/:platform/:type", control.Router)
	router.POST("/q/:platform/:type", control.Router)

	router.Static("/static", filepath.Join(tools.GetCurrentDirectory(), "/views/static/"))
	router.LoadHTMLGlob(filepath.Join(tools.GetCurrentDirectory(), "/views/template/**/*"))

	router.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{"host": c.Request.Host})
	})

	router.Run(":" + *portConfig)
}
