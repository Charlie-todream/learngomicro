package main

import (
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/selector"
	"github.com/asim/go-micro/v3/util/log"

	"io/ioutil"
	"net/http"
	"time"
)

// 基本方式调用Api
func callAPI(addr string, path string, method string) (string, error) {
	req, _ := http.NewRequest(method, "http://"+addr+path, nil)
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	buf, _ := ioutil.ReadAll(res.Body)
	return string(buf), nil
}

func main() {
	consulReg := consul.NewRegistry(
		registry.Addrs("localhost:8500"),
	)
	for {
		getService, err := consulReg.GetService("prodservice")
		if err != nil {
		log.Fatal()
		}

		next := selector.RoundRobin(getService) // 轮询的方式 Random
		node, err := next()
		if err != nil {
			log.Fatal()
		}
		callRes, err := callAPI(node.Address, "/v1/prods", "POST")
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second * 1)
		fmt.Println(callRes)
	}
}
