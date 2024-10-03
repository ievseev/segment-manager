package segment

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"segment-manager/internal/config"
	"segment-manager/internal/storage/postgres"
)

type testSuite struct {
	suite.Suite
	ctx context.Context
	db  *postgres.Storage
}

// TODO
// добавить запуск теста на изолированной БД в докере
// добавить очистку БД
// добавить тест на удаление сегмента

func (s *testSuite) TestInsertSegment() {
	store := New(s.db)
	//defer s.truncate()

	firstSegment, err := store.Save(s.ctx, "test_name")
	require.NoError(s.T(), err)
	require.True(s.T(), firstSegment > 0)

}

func TestStore_Flow(t *testing.T) {
	ts := testSuite{}
	var cancel context.CancelFunc

	ts.ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	cfg := config.MustLoad("../../../.env")
	ts.db, _ = postgres.New(cfg)

	suite.Run(t, &ts)
}
