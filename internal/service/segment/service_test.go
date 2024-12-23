package segment

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"segment-manager/internal/service/segment/mocks"
)

func TestService_CreateSegment(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	ctx := context.Background()

	testCases := []struct {
		name            string
		slug            string
		segmentRepoMock *mocks.MocksegmentRepo
		expResponse     int64
		wantErr         bool
	}{
		{
			name: "success",
			slug: "test_segment",
			segmentRepoMock: func() *mocks.MocksegmentRepo {
				mock := mocks.NewMocksegmentRepo(mc)
				mock.EXPECT().
					Create(ctx, "test_segment").
					Return(int64(1), nil)

				return mock
			}(),
			expResponse: 1,
			wantErr:     false,
		},
		{
			name: "failed",
			slug: "test_segment",
			segmentRepoMock: func() *mocks.MocksegmentRepo {
				mock := mocks.NewMocksegmentRepo(mc)
				mock.EXPECT().
					Create(ctx, "test_segment").
					Return(int64(0), errors.New("err"))

				return mock
			}(),
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := New(tc.segmentRepoMock)
			segmentID, err := service.CreateSegment(ctx, tc.slug)

			if tc.wantErr {
				assert.Error(t, err)
			}

			assert.Equal(t, segmentID, tc.expResponse)
		})
	}
}

func TestService_DeleteSegment(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	ctx := context.Background()

	testCases := []struct {
		name            string
		slug            string
		segmentRepoMock *mocks.MocksegmentRepo
		wantErr         bool
	}{
		{
			name: "success",
			slug: "test_segment",
			segmentRepoMock: func() *mocks.MocksegmentRepo {
				mock := mocks.NewMocksegmentRepo(mc)
				mock.EXPECT().
					Delete(ctx, "test_segment").
					Return(nil)

				return mock
			}(),
			wantErr: false,
		},
		{
			name: "failed",
			slug: "test_segment",
			segmentRepoMock: func() *mocks.MocksegmentRepo {
				mock := mocks.NewMocksegmentRepo(mc)
				mock.EXPECT().
					Delete(ctx, "test_segment").
					Return(errors.New("err"))

				return mock
			}(),
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := New(tc.segmentRepoMock)
			err := service.Delete(ctx, tc.slug)

			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
