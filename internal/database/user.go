package database

import (
	"context"
	"errors"
	"github.com/ChristopherVennemann/Go-AcademyDay/internal/apperrors"
	"github.com/ChristopherVennemann/Go-AcademyDay/internal/model"
	"github.com/lib/pq"
)

func (db *Database) SaveUser(ctx context.Context, user *model.User) error {
	query := `
		INSERT INTO users (username, email)
		VALUES ($1, $2) RETURNING id, created_at
	`
	err := db.Connection.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Email,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)
	if err != nil {
		var pgErr *pq.Error
		if errors.As(err, &pgErr) && pgErr.Code == apperrors.PgUniqueViolation {
			return apperrors.UserAlreadyExists
		}
		return err
	}
	return nil
}

func (db *Database) GetUsers(ctx context.Context) []*model.User {
	return []*model.User{
		{ID: 1, Username: "1", Email: "1", CreatedAt: "timestamp"},
		{ID: 1, Username: "1", Email: "1", CreatedAt: "timestamp"},
	}
}
