package db

import (
	"database/sql"
	pb "github.com/wenmingtang/download/srv/account/proto/account"
)

type Repository interface {
	Create(*pb.User) error
	Read(*pb.User) error
	Update(*pb.User) error
	GetByEmail(string) (*pb.User, error)
	GetByToken(string) (*pb.User, error)
	UpdateToken(*pb.User, string) error
}

type UserRepository struct {
	DB *sql.DB
}

func (repo *UserRepository) Create(user *pb.User) error {
	stmt, err := repo.DB.Prepare("INSERT INTO users (name, email, password, created_at) VALUES (?, ?, ?, now())")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.Name, user.Email, user.Password)
	return err
}

func (repo *UserRepository) Read(user *pb.User) error {
	return nil
}

func (repo *UserRepository) Update(user *pb.User) error {
	return nil
}

func (repo *UserRepository) GetByEmail(email string) (*pb.User, error) {
	stmt, err := repo.DB.Prepare("SELECT id, name, email, password, created_at from users WHERE email = ?")
	if err != nil {
		return nil, err
	}
	user := &pb.User{}
	r := stmt.QueryRow(email)
	if err := r.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Created); err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) UpdateToken(user *pb.User, token string) error {
	stmt, err := repo.DB.Prepare("UPDATE users set api_token = ? where id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(token, user.Id)
	return err
}

func (repo *UserRepository) GetByToken(token string) (*pb.User, error) {
	stmt, err := repo.DB.Prepare("SELECT id, name, email, password, created_at from users WHERE api_token = ?")
	if err != nil {
		return nil, err
	}
	user := &pb.User{}
	r := stmt.QueryRow(token)
	if err := r.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Created); err != nil {
		return nil, err
	}
	return user, nil
}
