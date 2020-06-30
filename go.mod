module go-micro

go 1.13

//里边使用了一个json 插件，启动项目会报错，所以添加这一行
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/golang/protobuf v1.4.1
	github.com/hashicorp/consul/api v1.5.0 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins v1.5.1
	google.golang.org/protobuf v1.25.0
)
