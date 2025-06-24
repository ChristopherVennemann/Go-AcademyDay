package repository

import (
	"github.com/ChristopherVennemann/Go-AcademyDay/internal/database"
)

type Repository struct {
	UserRepo UserRepository
}

func NewRepository(db *database.Database) *Repository {
	return &Repository{
		db,
	}
}
