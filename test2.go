package main

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/plugins/client/http/v3"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/selector"
	"log"
)

func callServiceAPI(s selector.Selector) {

	myCli := http.NewClient(
		client.Selector(s), // 返回当前service的一个节点
		client.ContentType("application/json"),
	)
	// 下面封装了一个请求
	req := myCli.NewRequest("prodservice","/v1/prods",map[string]string{})

	var resp map[string]interface{}
	err := myCli.Call(context.Background(),req,&resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}

func main()  {
	consulReg := consul.NewRegistry(
		registry.Addrs("localhost:8500"),
		)
	mySelector := selector.NewSelector(
		selector.Registry(consulReg),
		selector.SetStrategy(selector.RoundRobin),  // 轮训查询策略
		)

	callServiceAPI(mySelector)
}
