package user_segment

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"segment-manager/internal/service/user_segment/mocks"
	usDB "segment-manager/internal/store/user_segment"
)

func TestService_UpdateUserSegments(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	ctx := context.Background()

	testCases := []struct {
		name                string
		userID              int64
		slugsToAdd          []string
		slugsToDelete       []string
		userSegmentRepoMock *mocks.MockUserSegmentRepo
		expResult           []Segment
		wantErr             bool
	}{
		{
			name:          "success",
			userID:        1,
			slugsToAdd:    []string{"seg1"},
			slugsToDelete: []string{"seg2"},
			userSegmentRepoMock: func() *mocks.MockUserSegmentRepo {
				mock := mocks.NewMockUserSegmentRepo(mc)
				mock.EXPECT().
					UpsertUserSegments(ctx, int64(1), []string{"seg1"}, []string{"seg2"}).
					Return(nil)

				return mock
			}(),
			wantErr: false,
		},
		{
			name:          "failed",
			userID:        1,
			slugsToAdd:    []string{"seg1"},
			slugsToDelete: []string{"seg2"},
			userSegmentRepoMock: func() *mocks.MockUserSegmentRepo {
				mock := mocks.NewMockUserSegmentRepo(mc)
				mock.EXPECT().
					UpsertUserSegments(ctx, int64(1), []string{"seg1"}, []string{"seg2"}).
					Return(errors.New("error"))

				return mock
			}(),
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := New(tc.userSegmentRepoMock)
			err := service.Update(ctx, tc.userID, tc.slugsToAdd, tc.slugsToDelete)

			if tc.wantErr {
				assert.Error(t, err)
			}
		})
	}
}

func TestService_GetUserSegments(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	ctx := context.Background()

	testCases := []struct {
		name                string
		userID              int64
		userSegmentRepoMock *mocks.MockUserSegmentRepo
		expResult           []Segment
		wantErr             bool
	}{
		{
			name:   "success",
			userID: 1,
			userSegmentRepoMock: func() *mocks.MockUserSegmentRepo {
				mock := mocks.NewMockUserSegmentRepo(mc)
				mock.EXPECT().
					SelectUserSegments(ctx, int64(1)).
					Return(
						[]usDB.UserSegmentDB{
							{
								ID:   1,
								Slug: "seg1",
							},
						},
						nil)

				return mock
			}(),
			expResult: []Segment{
				{
					ID:   1,
					Slug: "seg1",
				},
			},
			wantErr: false,
		},
		{
			name:   "failed",
			userID: 1,
			userSegmentRepoMock: func() *mocks.MockUserSegmentRepo {
				mock := mocks.NewMockUserSegmentRepo(mc)
				mock.EXPECT().
					SelectUserSegments(ctx, int64(1)).
					Return(nil, errors.New("error"))

				return mock
			}(),
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := New(tc.userSegmentRepoMock)
			slugs, err := service.Get(ctx, tc.userID)

			if tc.wantErr {
				assert.Error(t, err)
			}

			assert.Equal(t, tc.expResult, slugs)
		})
	}
}
