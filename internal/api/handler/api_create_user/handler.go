package api_create_user

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"

	"segment-manager/internal/api/model"
)

type UserService interface {
	CreateUser(ctx context.Context, name string) (int64, error)
}

type CreateUser struct {
	userService UserService
	log         *slog.Logger
}

func New(userService UserService, log *slog.Logger) *CreateUser {
	return &CreateUser{
		userService: userService,
		log:         log,
	}
}

func (c *CreateUser) Handler(w http.ResponseWriter, r *http.Request) {
	var req Request

	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		c.log.Error("Failed to decode req")
		model.SendErrorResponse(w, http.StatusBadRequest, "Failed to decode req")

		return
	}

	if err := validator.New().Struct(req); err != nil {
		c.log.Error("Request validation error")
		model.SendErrorResponse(w, http.StatusBadRequest, "Request validation error")

		return
	}

	userID, err := c.userService.CreateUser(r.Context(), req.Name)
	if err != nil {
		c.log.Error("Failed to create segment")
		model.SendErrorResponse(w, http.StatusInternalServerError, "Failed to create user")

		return
	}

	render.JSON(w, r, Response{
		UserID: userID,
	})

	return
}
