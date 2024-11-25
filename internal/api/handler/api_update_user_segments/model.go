package api_update_user_segments

import "segment-manager/internal/api/model"

type Request struct {
	UserID        int64    `json:"user_id" validate:"required"`
	SlugsToAdd    []string `json:"slugs_to_add,omitempty"`
	SlugsToDelete []string `json:"slugs_to_delete,omitempty"`
}

type Response struct {
	// TODO возможно стоит возвращать список сегментов после обновления
	Error *model.ResultError `json:"error,omitempty"`
}
