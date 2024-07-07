package api_delete_segment

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type SegmentService interface {
	Delete(ctx context.Context, slug string) error
}

func New(log *slog.Logger, segmentService SegmentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req Request

		// TODO добавить вариативную обработку ошибок
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("Failed to decode req")

			render.JSON(w, r, Error())

			return
		}

		if err := validator.New().Struct(req); err != nil {
			log.Error("Request validation error")

			render.JSON(w, r, Error())

			return
		}

		err = segmentService.Delete(r.Context(), req.Slug)
		if err != nil {
			log.Error("Failed to delete segment")

			render.JSON(w, r, Error())

			return
		}

		render.JSON(w, r, OK())

		return
	}
}
