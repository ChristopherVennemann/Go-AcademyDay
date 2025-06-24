package service

import (
	"github.com/ChristopherVennemann/Go-AcademyDay/internal/repository"
)

type Service struct {
	UserService UserService
}

func NewService(r *repository.Repository) *Service {
	return &Service{NewUserService(r.UserRepo)}
}
