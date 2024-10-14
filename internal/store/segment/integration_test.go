package segment

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"segment-manager/db/migrations"
	"segment-manager/internal/config"
	"segment-manager/internal/storage/postgres"
)

type testSuite struct {
	suite.Suite
	db *postgres.Storage
}

func (s *testSuite) TestSaveAndDeleteSegment() {
	store := New(s.db)

	ctx := context.Background()
	firstSegment, err := store.Create(ctx, "test_name")
	require.NoError(s.T(), err)
	require.Equal(s.T(), int64(1), firstSegment)

	err = store.Delete(ctx, "test_name")
	require.NoError(s.T(), err)

}

func (s *testSuite) SetupSuite() {
	cfg := config.MustLoad("../../../.test.env")
	err := migrations.Run(cfg)
	require.NoError(s.T(), err)

	s.db, err = postgres.New(cfg)
	require.NoError(s.T(), err)
}

func TestSegmentSuite(t *testing.T) {
	suite.Run(t, &testSuite{})
}

func (s *testSuite) TearDownTest() {
	ctx := context.Background()
	_, err := s.db.ExecContext(ctx, "TRUNCATE TABLE segments RESTART IDENTITY")
	require.NoError(s.T(), err)
}
