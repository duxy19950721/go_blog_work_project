package main

import (
	"blog/gin/router"
)

func main() {
	router := router.ConfigRouter()

	router.Run(":8080") // 启动服务器，监听 8080 端口
}
