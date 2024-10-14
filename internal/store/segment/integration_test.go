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

// TODO
// добавить очистку БД

func (s *testSuite) TestSaveAndDeleteSegment() {
	store := New(s.db)
	defer s.tearDown()

	ctx := context.Background()
	firstSegment, err := store.Create(ctx, "test_name")
	require.NoError(s.T(), err)
	require.True(s.T(), firstSegment == 1)

	err = store.Delete(ctx, "test_name")
	require.NoError(s.T(), err)

}

func TestSegmentSuite(t *testing.T) {
	ts := testSuite{}

	cfg := config.MustLoad("../../../.test.env")
	err := migrations.Run(cfg)
	require.NoError(t, err)

	ts.db, _ = postgres.New(cfg)

	suite.Run(t, &ts)
}

func (s *testSuite) tearDown() {
	ctx := context.Background()
	_, err := s.db.ExecContext(ctx, "TRUNCATE TABLE segments RESTART IDENTITY")
	if err != nil {
		return
	}
}
