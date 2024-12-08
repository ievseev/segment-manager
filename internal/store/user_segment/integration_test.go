package user_segment

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"segment-manager/db/migrations"
	"segment-manager/internal/config"
	"segment-manager/internal/storage/postgres"
	"segment-manager/internal/store/segment"
	"segment-manager/internal/store/user"
)

type testSuite struct {
	suite.Suite
	db *postgres.Storage
}

func (s *testSuite) TestUpsertAndGetUserSegments() {
	userSegmentsStore := New(s.db)
	userStore := user.New(s.db)
	segmentStore := segment.New(s.db)

	ctx := context.Background()

	userID, err := userStore.Create(ctx, "test_name")
	require.NoError(s.T(), err)
	require.Equal(s.T(), int64(1), userID)

	firstSegment, err := segmentStore.Create(ctx, "uno")
	require.NoError(s.T(), err)
	require.Equal(s.T(), int64(1), firstSegment)

	err = userSegmentsStore.UpsertUserSegments(ctx, userID, []string{"uno"}, []string{"des"})
	require.NoError(s.T(), err)

	segments, err := userSegmentsStore.SelectUserSegments(ctx, 1)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), 1, len(segments))
	assert.Equal(s.T(), "uno", segments[0].Slug)

}

func (s *testSuite) TestFailedUpsertUserSegmentsUserDoesntExist() {
	store := New(s.db)

	ctx := context.Background()
	err := store.UpsertUserSegments(ctx, 1, []string{"uno"}, nil)
	require.Error(s.T(), err)
}

func (s *testSuite) SetupSuite() {
	cfg := config.MustLoad("../../../.test.env")
	err := migrations.Run(cfg)
	require.NoError(s.T(), err)

	s.db, err = postgres.New(cfg)
	require.NoError(s.T(), err)
}

func TestUserSegmentsSuite(t *testing.T) {
	suite.Run(t, &testSuite{})
}

func (s *testSuite) TearDownTest() {
	ctx := context.Background()

	queries := []string{
		"TRUNCATE TABLE users RESTART IDENTITY CASCADE",
		"TRUNCATE TABLE segments RESTART IDENTITY CASCADE",
		"TRUNCATE TABLE users_segments RESTART IDENTITY CASCADE",
	}

	for _, query := range queries {
		_, err := s.db.ExecContext(ctx, query)
		require.NoError(s.T(), err)
	}
}
