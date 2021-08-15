package service

import (
	"context"

	"github.com/go-zelus/api-layout/entity"
	"github.com/go-zelus/api-layout/model"
	"github.com/go-zelus/api-layout/repository"
)

type UserService interface {
	Hello(ctx context.Context, in *entity.HelloRequest) (out *entity.HelloResponse, err error)
}

func NewUserService() UserService {
	return &userService{}
}

type userService struct{}

func (u userService) Hello(ctx context.Context, in *entity.HelloRequest) (out *entity.HelloResponse, err error) {
	userRepo := repository.NewUserRepo()
	user := model.User{Name: in.Name}
	res := userRepo.Hello(user)
	return &entity.HelloResponse{
		Value: "Hello " + res,
	}, nil
}
