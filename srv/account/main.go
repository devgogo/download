package main

import (
	"github.com/micro/go-micro"
	"github.com/wenmingtang/download/srv/account/db"
	"github.com/wenmingtang/download/srv/account/handler"
	"github.com/wenmingtang/download/srv/account/proto/account"
	"log"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.account"),
	)
	service.Init()

	session, err := db.CreateConnection()
	if err != nil {
		log.Fatal(err)
	}

	repo := &db.UserRepository{DB: session}
	tokenService := &db.TokenService{Repo: repo}
	accountService := &handler.Account{Repo: repo, TokenService: tokenService}

	if err := account.RegisterAccountHandler(service.Server(), accountService); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
