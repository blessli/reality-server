package util

import (
	"reality-server/pkg/setting"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// GetPage 查询queryContext page值
func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}

	return result
}

// GetPageSize 查询单页数值
func GetPageSize(c *gin.Context) int {
	size, _ := com.StrTo(c.Query("size")).Int()
	if size == 0 {
		size = setting.PageSize
	}
	return size
}
