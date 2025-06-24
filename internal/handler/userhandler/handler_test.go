package userhandler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/ChristopherVennemann/Go-AcademyDay/internal/apperrors"
	"github.com/ChristopherVennemann/Go-AcademyDay/internal/model"
	"github.com/ChristopherVennemann/Go-AcademyDay/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	ContextMatcher = mock.MatchedBy(func(ctx context.Context) bool { return true })
)

func TestCreateUser_ShouldSucceed(t *testing.T) {
	inputUser := testutils.NewUser("Username", "test@email.com")
	inputBytes, err := json.Marshal(inputUser)
	require.NoError(t, err)

	expectedId := 1

	mockService := &testutils.MockUserService{}
	mockService.On("CreateUser", ContextMatcher, inputUser).
		Run(func(args mock.Arguments) {
			u := args.Get(1).(*model.User)
			u.ID = expectedId
			u.CreatedAt = "2001-01-01T01:01:01.000001Z"
		}).
		Return(nil)
	handler := &Handler{mockService}

	req := httptest.NewRequest(http.MethodPost, CreateUserPath, bytes.NewReader(inputBytes)).
		WithContext(context.Background())
	rec := httptest.NewRecorder()

	handler.CreateUser(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))
	actualUser := &model.User{}
	err = json.Unmarshal(rec.Body.Bytes(), actualUser)
	assert.NoError(t, err)
	assert.Equal(t, expectedId, actualUser.ID)
	assert.Equal(t, inputUser.Username, actualUser.Username)
	assert.Equal(t, inputUser.Email, actualUser.Email)
	mockService.AssertExpectations(t)
}
func TestCreateUser_ShouldHaveStatusAndMessageFromAppError(t *testing.T) {
	inputUser := testutils.NewUser("Username", "test@email.com")
	inputBytes, err := json.Marshal(inputUser)
	require.NoError(t, err)

	mockService := &testutils.MockUserService{}
	mockService.On("CreateUser", ContextMatcher, inputUser).Return(apperrors.UserAlreadyExists)
	handler := &Handler{mockService}

	req := httptest.NewRequest(http.MethodPost, CreateUserPath, bytes.NewReader(inputBytes)).
		WithContext(context.Background())
	rec := httptest.NewRecorder()

	handler.CreateUser(rec, req)

	assert.Equal(t, apperrors.UserAlreadyExists.HttpStatus, rec.Code)
	assert.Equal(t, apperrors.UserAlreadyExists.Message, strings.TrimSpace(rec.Body.String()))
	mockService.AssertExpectations(t)
}

func TestCreateUser_ShouldHaveStatusInternalServerErrorForUnknownError(t *testing.T) {
	inputUser := testutils.NewUser("Username", "test@email.com")
	inputBytes, err := json.Marshal(inputUser)
	require.NoError(t, err)

	mockService := &testutils.MockUserService{}
	mockService.On("CreateUser", ContextMatcher, inputUser).Return(errors.New(""))
	handler := &Handler{mockService}

	req := httptest.NewRequest(http.MethodPost, CreateUserPath, bytes.NewReader(inputBytes)).
		WithContext(context.Background())
	rec := httptest.NewRecorder()

	handler.CreateUser(rec, req)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Empty(t, rec.Body.String())
	mockService.AssertExpectations(t)
}

func TestGetUsers_ShouldReturnStatusOkAndEmptyListIfNoUsersExist(t *testing.T) {
	var noUsers []*model.User
	mockService := &testutils.MockUserService{}
	mockService.On("GetUsers", ContextMatcher).Return(noUsers, nil)
	handler := &Handler{mockService}

	req := httptest.NewRequest(http.MethodGet, GetUsersPath, nil).
		WithContext(context.Background())
	rec := httptest.NewRecorder()

	handler.GetUsers(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	respBody := rec.Body.String()
	assert.JSONEq(t, "[]", respBody)
	mockService.AssertExpectations(t)
}

func TestGetUsers_ShouldReturnStatusOkAndListOfUsers(t *testing.T) {
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
	mockService := &testutils.MockUserService{}
	mockService.On("GetUsers", ContextMatcher).Return(expectedUsers, nil)
	handler := &Handler{mockService}

	req := httptest.NewRequest(http.MethodGet, GetUsersPath, nil).
		WithContext(context.Background())
	rec := httptest.NewRecorder()

	handler.GetUsers(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))
	var actualUsers []*model.User
	err := json.Unmarshal(rec.Body.Bytes(), &actualUsers)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(actualUsers))
	assert.ElementsMatchf(t, expectedUsers, actualUsers, "expectedUsers should match")
	mockService.AssertExpectations(t)
}
