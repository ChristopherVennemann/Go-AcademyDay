package integration

import (
	"context"
	"errors"
	"github.com/ChristopherVennemann/Go-AcademyDay/internal/apperrors"
	"github.com/ChristopherVennemann/Go-AcademyDay/internal/database"
	"github.com/ChristopherVennemann/Go-AcademyDay/internal/model"
	"github.com/ChristopherVennemann/Go-AcademyDay/internal/repository"
	"github.com/ChristopherVennemann/Go-AcademyDay/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestUserRepository_SaveUser_ShouldSaveNewUsers(t *testing.T) {
	conn, cleanup, err := testutils.SetupTestPostgres(context.Background())
	require.NoError(t, err)
	defer cleanup()
	db := &database.Database{Connection: conn}
	userRepo := repository.NewRepository(db).UserRepo
	newUser1 := testutils.NewUser("Username 1", "test1@email.com")
	expectedId1 := 1
	newUser2 := testutils.NewUser("Username 2", "test2@email.com")
	expectedId2 := 2

	savedUser1 := *newUser1
	err = userRepo.SaveUser(context.Background(), &savedUser1)
	require.NoError(t, err)

	assert.Equal(t, newUser1.Username, savedUser1.Username)
	assert.Equal(t, newUser1.Email, savedUser1.Email)
	assert.Equal(t, expectedId1, savedUser1.ID)
	_, err = time.Parse(time.RFC3339Nano, savedUser1.CreatedAt)
	assert.NoError(t, err)

	savedUser2 := *newUser2
	err = userRepo.SaveUser(context.Background(), &savedUser2)
	require.NoError(t, err)

	assert.Equal(t, newUser2.Username, savedUser2.Username)
	assert.Equal(t, newUser2.Email, savedUser2.Email)
	assert.Equal(t, expectedId2, savedUser2.ID)
	_, err = time.Parse(time.RFC3339Nano, savedUser2.CreatedAt)
	assert.NoError(t, err)
}

func TestUserRepository_SaveUser_ShouldReturnAppErrAndNotUpdateUser_IfUserOrEmailExists(t *testing.T) {
	conn, cleanup, err := testutils.SetupTestPostgres(context.Background())
	require.NoError(t, err)
	defer cleanup()
	db := &database.Database{Connection: conn}
	userRepo := repository.NewRepository(db).UserRepo
	newUser1 := testutils.NewUser("Username 1", "test1@email.com")
	againUser1 := *newUser1

	err = userRepo.SaveUser(context.Background(), newUser1)
	require.NoError(t, err)
	err = userRepo.SaveUser(context.Background(), &againUser1)

	var appErr *apperrors.AppError
	assert.True(t, errors.As(err, &appErr), "error should be AppError")
	if appErr != nil {
		assert.Equal(t, apperrors.UserAlreadyExists.Message, appErr.Message)
		assert.Equal(t, http.StatusConflict, appErr.HttpStatus)
	}
}

func TestUserRepository_GetUsers_ShouldReturnEmptyList_IfNoUsersExist(t *testing.T) {
	conn, cleanup, err := testutils.SetupTestPostgres(context.Background())
	require.NoError(t, err)
	defer cleanup()
	db := &database.Database{Connection: conn}
	userRepo := repository.NewRepository(db).UserRepo

	retrievedUsers := userRepo.GetUsers(context.Background())

	assert.Empty(t, retrievedUsers)
}

func TestUserRepository_GetUsers_ShouldReturnListOfUsers(t *testing.T) {
	conn, cleanup, err := testutils.SetupTestPostgres(context.Background())
	require.NoError(t, err)
	defer cleanup()
	db := &database.Database{Connection: conn}
	userRepo := repository.NewRepository(db).UserRepo
	savedUsers := []*model.User{
		testutils.NewUser("user1", "user1@email.com"),
		testutils.NewUser("user2", "user2@email.com"),
	}
	err = userRepo.SaveUser(context.Background(), savedUsers[0])
	require.NoError(t, err)
	err = userRepo.SaveUser(context.Background(), savedUsers[1])
	require.NoError(t, err)

	retrievedUsers := userRepo.GetUsers(context.Background())

	assert.Len(t, retrievedUsers, len(savedUsers))
	assert.Equal(t, testutils.ToUserComparable(retrievedUsers), testutils.ToUserComparable(savedUsers))
}
