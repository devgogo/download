package handler

import (
	"context"
	"github.com/micro/go-micro/client"
	pb "github.com/wenmingtang/download/srv/account/proto/account"
)

type User struct {
	Client client.Client
}

func (s *User) Create(ctx context.Context, req *pb.User, rsp *pb.User) error {
	accountClient := pb.NewAccountService("go.micro.srv.account", s.Client)

	accountRsp, err := accountClient.Create(ctx, req)
	if err != nil {
		return err
	}
	rsp = accountRsp.User
	return nil
}
