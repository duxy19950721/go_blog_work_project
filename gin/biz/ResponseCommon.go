package biz

import (
	"log"

	"github.com/gin-gonic/gin"
)

// 返回成功响应
func responseSuccess(c *gin.Context, msg string) {
	c.JSON(200, gin.H{
		"msg": msg,
	})
}
func responseSuccessWithData(c *gin.Context, msg string, data any) {
	c.JSON(200, gin.H{
		"msg":  msg,
		"data": data,
	})
}

// 返回错误响应
func responseFail(c *gin.Context, msg string, err error) {
	c.JSON(500, gin.H{
		"msg": msg,
	})
	log.Panicln(msg, err)
}
