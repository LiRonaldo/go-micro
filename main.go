package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"go-micro/Services"
	"go-micro/ServicesImpl"
)

func main() {
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
	prodService := micro.NewService(
		micro.Name("prodservice"),
		micro.Address("127.0.0.1:8011"),
		micro.Registry(consulReg),
	)
	prodService.Init()
	Services.RegisterProdServiceHandler(prodService.Server(), new(ServicesImpl.ProdService))
	prodService.Run()
}
