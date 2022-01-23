package main

import (
	"github.com/asim/go-micro/v3/web"
	"github.com/gin-gonic/gin"

)

func main() {

	ginRouter := gin.Default()

	ginRouter.Handle("GET", "/", func(context *gin.Context) {
		data := make([]interface{}, 0)
		context.JSON(200, gin.H{
			"data": data,
		})
	})

	server := web.NewService(

		web.Address(":8000"),
		web.Handler(ginRouter),
	)
	server.Run()
}
