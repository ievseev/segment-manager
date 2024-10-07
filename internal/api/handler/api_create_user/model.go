package api_create_user

import (
	"segment-manager/internal/api/model"
)

type Request struct {
	Name string `json:"name" validate:"required"`
}

type Response struct {
	UserID int64              `json:"user_id"`
	Error  *model.ResultError `json:"error,omitempty"`
}
