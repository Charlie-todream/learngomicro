# learngomicro

go get github.com/micro/micro/v2

go get -u github.com/gin-gonic/gin

go mod download golang.org/x/net

docker run -d --name=cs -p 8500:8500 consul agent -server -bootstrap -ui -client 0.0.0.0

go micro v2 不在默认支持 consul 官方推荐 etcd 

通过 github.com/micro/go-plugins/registry/consul/v2 引入