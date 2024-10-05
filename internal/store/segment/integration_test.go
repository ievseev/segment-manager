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
	ctx context.Context
	db  *postgres.Storage
}

// TODO
// добавить очистку БД

func (s *testSuite) TestSaveAndDeleteSegment() {
	store := New(s.db)
	defer s.truncate()

	firstSegment, err := store.Save(s.ctx, "test_name")
	require.NoError(s.T(), err)
	require.True(s.T(), firstSegment == 1)

	err = store.Delete(s.ctx, "test_name")
	require.NoError(s.T(), err)

}

func TestStore_Flow(t *testing.T) {
	ts := testSuite{}
	var cancel context.CancelFunc

	ts.ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	cfg := config.MustLoad("../../../.test.env")
	err := migrations.Run(cfg)
	require.NoError(t, err)

	ts.db, _ = postgres.New(cfg)

	suite.Run(t, &ts)
}

func (s *testSuite) truncate() {

}
