package main

import (
	"fleet_backend_final/truck_service/implement"
	"fleet_backend_final/truck_service/proto"
	"fmt"
	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("truckservice"),

	)

	repo, err := implement.NewTruckServiceRepository()
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

	if err := proto.RegisterTruckServiceHandler(service.Server(), handler) ; err != nil{
		panic("couldn't run this1")
	}

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

