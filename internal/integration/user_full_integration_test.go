package integration

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/ChristopherVennemann/Go-AcademyDay/internal/apperrors"
	"github.com/ChristopherVennemann/Go-AcademyDay/internal/database"
	"github.com/ChristopherVennemann/Go-AcademyDay/internal/handler"
	"github.com/ChristopherVennemann/Go-AcademyDay/internal/handler/userhandler"
	"github.com/ChristopherVennemann/Go-AcademyDay/internal/model"
	"github.com/ChristopherVennemann/Go-AcademyDay/internal/repository"
	"github.com/ChristopherVennemann/Go-AcademyDay/internal/service"
	"github.com/ChristopherVennemann/Go-AcademyDay/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func setUpApp(t *testing.T) (http.Handler, func(), error) {
	conn, cancel, err := testutils.SetupTestPostgres(context.Background())
	require.NoError(t, err)
	useRepo := repository.NewRepository(&database.Database{Connection: conn})
	useService := service.NewService(useRepo)
	router := handler.NewRouter(useService)

	return router, cancel, nil
}

func TestCreateUser_ShouldCreateAndReturnNewUser(t *testing.T) {
	testApi, cancel, err := setUpApp(t)
	require.NoError(t, err)
	defer cancel()
	inputUser := testutils.NewUser("testuser", "test@email.com")
	inputBytes, err := json.Marshal(inputUser)
	require.NoError(t, err)
	expectedId := 1

	req := httptest.NewRequest(http.MethodPost, userhandler.CreateUserPath, bytes.NewReader(inputBytes)).
		WithContext(context.Background())
	rec := httptest.NewRecorder()
	testApi.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))
	actualUser := &model.User{}
	err = json.Unmarshal(rec.Body.Bytes(), actualUser)
	assert.NoError(t, err)
	assert.Equal(t, expectedId, actualUser.ID)
	assert.Equal(t, inputUser.Username, actualUser.Username)
	assert.Equal(t, inputUser.Email, actualUser.Email)
}
func TestCreateUser_ShouldReturnStatus409AndErrorMessageWhenDuplicateIsCreated(t *testing.T) {
	testApi, cancel, err := setUpApp(t)
	require.NoError(t, err)
	defer cancel()

	user1 := testutils.NewUser("testuser", "test@email.com")

	inputBytes, err := json.Marshal(user1)
	require.NoError(t, err)

	req1 := httptest.NewRequest(http.MethodPost, userhandler.CreateUserPath, bytes.NewReader(inputBytes)).
		WithContext(context.Background())
	req2 := httptest.NewRequest(http.MethodPost, userhandler.CreateUserPath, bytes.NewReader(inputBytes)).
		WithContext(context.Background())
	rec1 := httptest.NewRecorder()
	rec2 := httptest.NewRecorder()
	testApi.ServeHTTP(rec1, req1)
	testApi.ServeHTTP(rec2, req2)

	assert.Equal(t, apperrors.UserAlreadyExists.HttpStatus, rec2.Code)
	assert.Equal(t, apperrors.UserAlreadyExists.Message, strings.TrimSpace(rec2.Body.String()))
}
