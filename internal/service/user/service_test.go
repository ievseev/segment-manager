package user

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"segment-manager/internal/service/user/mocks"
)

func TestService_CreateUser(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	ctx := context.Background()

	testCases := []struct {
		name         string
		userName     string
		userRepoMock *mocks.MockuserRepo
		expResponse  int64
		wantErr      bool
	}{
		{
			name:     "success",
			userName: "test_user",
			userRepoMock: func() *mocks.MockuserRepo {
				mock := mocks.NewMockuserRepo(mc)
				mock.EXPECT().
					Create(ctx, "test_user").
					Return(int64(1), nil)

				return mock
			}(),
			expResponse: 1,
			wantErr:     false,
		},
		{
			name:     "failed",
			userName: "test_user",
			userRepoMock: func() *mocks.MockuserRepo {
				mock := mocks.NewMockuserRepo(mc)
				mock.EXPECT().
					Create(ctx, "test_user").
					Return(int64(0), errors.New("err"))

				return mock
			}(),
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userID, err := tc.userRepoMock.Create(ctx, tc.userName)

			if tc.wantErr {
				assert.Error(t, err)
			}

			assert.Equal(t, userID, tc.expResponse)
		})
	}
}
