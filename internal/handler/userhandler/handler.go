package userhandler

import (
	"encoding/json"
	"errors"
	"github.com/ChristopherVennemann/Go-AcademyDay/internal/apperrors"
	"github.com/ChristopherVennemann/Go-AcademyDay/internal/model"
	"github.com/ChristopherVennemann/Go-AcademyDay/internal/service"
	"log"
	"net/http"
)

type Handler struct {
	userService service.UserService
}

func NewHandler(s service.UserService) *Handler {
	return &Handler{userService: s}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &model.User{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.userService.CreateUser(r.Context(), user); err != nil {
		var appErr *apperrors.AppError
		if errors.As(err, &appErr) {
			http.Error(w, appErr.Message, appErr.HttpStatus)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	body, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if _, err = w.Write(body); err != nil {
		log.Printf("error writing response body: %s", err.Error())
	}
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	// todo: implement me!
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
}
