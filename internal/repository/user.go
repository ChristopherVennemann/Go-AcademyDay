package repository

import (
	"context"
	"github.com/ChristopherVennemann/Go-AcademyDay/internal/model"
)

type UserRepository interface {
	SaveUser(ctx context.Context, user *model.User) error
	GetUsers(ctx context.Context) []*model.User
}
