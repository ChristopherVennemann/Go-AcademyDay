package userhandler

import (
	"github.com/go-chi/chi/v5"
)

const (
	CreateUserPath = "/user"
	GetUsersPath   = "/users"
)

func NewUserRouter(handler *Handler) *chi.Mux {
	r := chi.NewRouter()
	r.Post(CreateUserPath, handler.CreateUser)
	r.Get(GetUsersPath, handler.GetUsers)
	return r
}
