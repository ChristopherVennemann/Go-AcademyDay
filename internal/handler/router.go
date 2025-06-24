package handler

import (
	"github.com/ChristopherVennemann/Go-AcademyDay/internal/handler/userhandler"
	"github.com/ChristopherVennemann/Go-AcademyDay/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"time"
)

const (
	apiRoot = "/"
)

func baseRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	return r
}

func NewRouter(useService *service.Service) *chi.Mux {
	router := baseRouter()
	router.Mount(apiRoot, userhandler.NewUserRouter(userhandler.NewHandler(useService.UserService)))
	return router
}
