package platform

import "github.com/gin-gonic/gin"

/*
Query 查询类型，查询内容 返回API返回内容
*/
type Base interface {
	Query(c *gin.Context)
}

//初始化插件内容，每次新加插件都需要处理
func Init() map[string]Base {
	return map[string]Base{
		"threatbook": NewTT(),
		"virustotal": NewVirus(),
	}
}
