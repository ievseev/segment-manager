package api_create_segment

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"segment-manager/internal/api/handler/api_create_segment/mocks"
	"segment-manager/internal/api/model" // Импортируйте пакет с моделью ResultError
)

func TestCreateSegment_Handler(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	ctx := context.Background()

	testCases := []struct {
		name           string
		requestBody    interface{}
		setupMocks     func(*mocks.MockSegmentService)
		wantStatusCode int
		wantResponse   Response
	}{
		{
			name:        "success",
			requestBody: Request{Slug: "test-segment"},
			setupMocks: func(mockService *mocks.MockSegmentService) {
				mockService.
					EXPECT().
					CreateSegment(ctx, "test-segment").
					Return(int64(1), nil)
			},
			wantStatusCode: http.StatusOK,
			wantResponse: Response{
				SegmentID: 1,
			},
		},
		{
			name:           "decode error",
			requestBody:    "wtf",
			setupMocks:     func(mockService *mocks.MockSegmentService) {},
			wantStatusCode: http.StatusBadRequest,
			wantResponse: Response{
				Error: &model.ResultError{
					Code:    http.StatusBadRequest,
					Message: "Failed to decode req",
				},
			},
		},
		{
			name:           "invalid payload",
			requestBody:    Request{},
			setupMocks:     func(mockService *mocks.MockSegmentService) {},
			wantStatusCode: http.StatusBadRequest,
			wantResponse: Response{
				Error: &model.ResultError{
					Code:    http.StatusBadRequest,
					Message: "Request validation error",
				},
			},
		},
		{
			name:        "service error",
			requestBody: Request{Slug: "test-segment"},
			setupMocks: func(mockService *mocks.MockSegmentService) {
				mockService.
					EXPECT().
					CreateSegment(ctx, "test-segment").
					Return(int64(0), errors.New("service error"))
			},
			wantStatusCode: http.StatusInternalServerError,
			wantResponse: Response{
				Error: &model.ResultError{
					Code:    http.StatusInternalServerError,
					Message: "Failed to create segment",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockService := mocks.NewMockSegmentService(mc)
			tc.setupMocks(mockService)

			log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelInfo,
			}))
			handler := New(mockService, log)

			requestBytes, err := json.Marshal(tc.requestBody)
			require.NoError(t, err)

			req := httptest.NewRequest(http.MethodPost, "/segments", bytes.NewBuffer(requestBytes))
			w := httptest.NewRecorder()

			handler.Handler(w, req)

			require.Equal(t, tc.wantStatusCode, w.Code, "статус код не совпадает")

			if w.Code == http.StatusOK {
				var gotResponse Response
				err = json.NewDecoder(w.Body).Decode(&gotResponse)
				require.NoError(t, err)
				require.Equal(t, tc.wantResponse.SegmentID, gotResponse.SegmentID, "SegmentID не совпадает")
			} else {
				var gotErrorResponse model.ResultError
				err = json.NewDecoder(w.Body).Decode(&gotErrorResponse)
				require.NoError(t, err)
				require.Equal(t, tc.wantResponse.Error.Code, gotErrorResponse.Code, "код ошибки не совпадает")
				require.Equal(t, tc.wantResponse.Error.Message, gotErrorResponse.Message, "сообщение об ошибке не совпадает")
			}
		})
	}
}
