package main

import (
	"github.com/micro/go-micro"
	"github.com/wenmingtang/download/api/user/handler"
	"github.com/wenmingtang/download/api/user/proto/user"
	"log"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.user"),
	)
	service.Init()

	_ = user.RegisterUsersHandler(service.Server(), &handler.User{Client: service.Client()})
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
