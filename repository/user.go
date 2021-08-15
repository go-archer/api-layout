package repository

import "github.com/go-zelus/api-layout/model"

type UserRepo interface {
	Hello(user model.User) string
}

func NewUserRepo() UserRepo {
	return &userRepo{}
}

type userRepo struct{}

func (u userRepo) Hello(user model.User) string {
	return user.Name
}
