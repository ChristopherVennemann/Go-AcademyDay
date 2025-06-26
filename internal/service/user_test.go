package service

import (
	"context"
	"github.com/ChristopherVennemann/Go-AcademyDay/internal/apperrors"
	"github.com/ChristopherVennemann/Go-AcademyDay/internal/model"
	"github.com/ChristopherVennemann/Go-AcademyDay/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	ContextMatcher = mock.MatchedBy(func(ctx context.Context) bool { return true })
)

func TestUserService_CreatUser_ShouldSucceed(t *testing.T) {
	mockRepo := new(testutils.MockUserRepo)
	userService := NewUserService(mockRepo)
	newUser := testutils.NewUser("Username", "test@email.com")
	expectedUser := *newUser
	expectedUser.ID = 1
	expectedUser.CreatedAt = "2001-01-01T01:01:01.000001Z"
	mockRepo.On("SaveUser", ContextMatcher, newUser).
		Run(func(args mock.Arguments) {
			u := args.Get(1).(*model.User)
			u.ID = expectedUser.ID
			u.CreatedAt = expectedUser.CreatedAt
		}).
		Return(nil)

	err := userService.CreateUser(context.Background(), newUser)
	require.NoError(t, err)

	assert.Equal(t, &expectedUser, newUser)
	mockRepo.AssertExpectations(t)
}

func TestUserService_CreatUser_ShouldReturnErrorAndNotUpdateUser(t *testing.T) {
	mockRepo := new(testutils.MockUserRepo)
	userService := NewUserService(mockRepo)
	newUser := testutils.NewUser("Username", "test@email.com")
	expectedUser := newUser
	mockRepo.On("SaveUser", ContextMatcher, newUser).Return(apperrors.UserAlreadyExists)

	err := userService.CreateUser(context.Background(), newUser)

	assert.Equal(t, apperrors.UserAlreadyExists, err)
	assert.Equal(t, expectedUser, newUser)
	mockRepo.AssertExpectations(t)
}
func TestUserService_GetUsers_ShouldReturnEmptyListIfNoUsersExist(t *testing.T) {
	mockRepo := new(testutils.MockUserRepo)
	userService := NewUserService(mockRepo)
	var noUsers []*model.User
	mockRepo.On("GetUsers", ContextMatcher).Return(noUsers, nil)

	actualUsers := userService.GetUsers(context.Background())

	assert.Empty(t, actualUsers)
	mockRepo.AssertExpectations(t)
}

func TestUserService_GetUsers_ShouldReturnListOfUsers(t *testing.T) {
	mockRepo := new(testutils.MockUserRepo)
	userService := NewUserService(mockRepo)
	expectedUsers := []*model.User{
		{
			ID:        1,
			Username:  "User1",
			Email:     "user1@email.com",
			CreatedAt: "2001-01-01T01:01:01.000001Z",
		},
		{
			ID:        2,
			Username:  "User2",
			Email:     "user2@email.com",
			CreatedAt: "2001-01-01T01:01:01.000001Z",
		},
	}
	mockRepo.On("GetUsers", ContextMatcher).Return(expectedUsers, nil)

	actualUsers := userService.GetUsers(context.Background())

	assert.ElementsMatchf(t, expectedUsers, actualUsers, "users should match")
	mockRepo.AssertExpectations(t)
}
