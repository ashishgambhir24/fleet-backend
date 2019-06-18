package main

import (
	"fleet_backend_final/customer_service/implement"
	"fleet_backend_final/customer_service/proto"
	"fmt"
	"github.com/micro/go-micro"
)




func main() {
	service := micro.NewService(
		micro.Name("customerservice"),

	)

	repo, err := implement.NewCustomerServiceRepository()
	if err != nil {
		panic(err)
	}

	defer repo.Close()

	srvs := implement.Service{
		Repository:repo,
	}

	handler := implement.Handler{
		Service: srvs,
	}

	service.Init()

	if err := proto.RegisterCustomerServiceHandler(service.Server(), handler) ; err != nil{
		panic("couldn't run this1")
	}

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
