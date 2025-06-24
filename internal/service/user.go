package service

import (
	"context"
	"github.com/ChristopherVennemann/Go-AcademyDay/internal/model"
	"github.com/ChristopherVennemann/Go-AcademyDay/internal/repository"
)

type userService struct {
	repo repository.UserRepository
}

type UserService interface {
	CreateUser(context.Context, *model.User) error
	GetUsers(ctx context.Context) ([]*model.User, error)
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) CreateUser(ctx context.Context, user *model.User) error {
	return s.repo.SaveUser(ctx, user)
}

func (s *userService) GetUsers(ctx context.Context) ([]*model.User, error) {
	// todo: implement me!
	return nil, nil
}
