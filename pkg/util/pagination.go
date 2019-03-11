package util

import (
	"ginblog/config"

	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
)

// GetPage 获取页面
func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * config.PageSize
	}

	return result
}
