package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/web"
)

func main() {

	ginRouter := gin.Default()

	ginRouter.Handle("GET", "/user", func(context *gin.Context) {
		context.String(200, "user api")
	})

	ginRouter.Handle("GET", "/news", func(context *gin.Context) {
		context.String(200, "new api")
	})

	server := web.NewService(
		web.Address(":8001"),
		web.Handler(ginRouter),
	)

	server.Run()
}
