package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	// 新建一个consul注册地址
	consulReg := consul.NewRegistry(
		registry.Addrs("localhost:8500"),
	)

	ginRouter := gin.Default()

	ginRouter.Handle("GET", "/user", func(context *gin.Context) {
		context.String(200, "user api")
	})

	ginRouter.Handle("GET", "/news", func(context *gin.Context) {
		context.String(200, "new api")
	})

	server := web.NewService(
		web.Name("prodservice"), // 注册到consul服务中的service name
		web.Address(":8001"),
		web.Handler(ginRouter),
		web.Registry(consulReg),  // 注册到哪个服务器上的consul中
	)
	server.Run()
}
