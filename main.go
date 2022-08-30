package main

import (
	"github.com/mwqnice/gin-demo/internal/router"
)


func main()  {
	router := router.NewRouter()
	router.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}