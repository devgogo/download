package db

import (
	pb "github.com/wenmingtang/download/srv/account/proto/account"
	"math/rand"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var tokenLength = 128

func randomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

type TokenService struct {
	Repo Repository
}

func (t *TokenService) Generate(user *pb.User) (string, error) {
	token := randomString(tokenLength)
	if err := t.Repo.UpdateToken(user, token); err != nil {
		return "", err
	}
	return token, nil
}
