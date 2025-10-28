package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PanicCatch 捕获 panic 的中间件
func PanicCatch() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 使用 defer + recover 捕获 panic
		defer func() {
			if err := recover(); err != nil {
				// 记录 panic 错误
				log.Printf("Panic recovered: %v\n", err)

				// 返回错误响应
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    500,
					"message": err.(string),
					"error":   err,
				})

				// 终止请求处理链
				c.Abort()
			}
		}()

		// 继续处理请求
		c.Next()
	}
}
