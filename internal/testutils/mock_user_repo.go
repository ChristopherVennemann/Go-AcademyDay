package testutils

import (
	"context"
	"github.com/ChristopherVennemann/Go-AcademyDay/internal/model"
	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) SaveUser(ctx context.Context, user *model.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}
