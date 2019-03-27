package handler

import (
	"context"
	"github.com/micro/go-micro/errors"
	"github.com/wenmingtang/download/srv/account/db"
	pb "github.com/wenmingtang/download/srv/account/proto/account"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	Repo         db.Repository
	TokenService *db.TokenService
}

func (s *Account) Create(ctx context.Context, req *pb.User, rsp *pb.Response) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.InternalServerError("go.micro.srv.account", err.Error())
	}
	req.Password = string(hashedPass)
	if err := s.Repo.Create(req); err != nil {
		return errors.BadRequest("go.micro.srv.account", err.Error())
	}
	rsp.User = req
	return nil
}

func (s *Account) Get(ctx context.Context, req *pb.User, rsp *pb.Response) error {
	// TODO
	return nil
}

func (s *Account) Update(ctx context.Context, req *pb.User, rsp *pb.Response) error {
	// TODO
	return nil
}

func (s *Account) Auth(ctx context.Context, req *pb.User, rsp *pb.Token) error {
	user, err := s.Repo.GetByEmail(req.Email)
	if err != nil {
		return errors.BadRequest("go.micro.srv.account", "user not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return errors.BadRequest("go.micro.srv.account", "email or password invalid")
	}
	token, err := s.TokenService.Generate(user)
	rsp.Token = token
	return nil
}

func (s *Account) ValidateToken(ctx context.Context, req *pb.Token, rsp *pb.Token) error {
	_, err := s.Repo.GetByToken(req.Token)
	if err != nil {
		return errors.NotFound("go.micro.srv.account", "invalid token")
	}
	rsp.Token = req.Token
	rsp.Valid = true
	return nil
}
