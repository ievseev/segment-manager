package user

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

func (s *testSuite) TestSaveUser() {
	store := New(s.db)

	ctx := context.Background()
	firstUser, err := store.Create(ctx, "test_name")
	require.NoError(s.T(), err)
	require.Equal(s.T(), int64(1), firstUser)

}

func (s *testSuite) SetupSuite() {
	cfg := config.MustLoad("../../../.test.env")
	err := migrations.Run(cfg)
	require.NoError(s.T(), err)

	s.db, err = postgres.New(cfg)
	require.NoError(s.T(), err)
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, &testSuite{})
}

func (s *testSuite) TearDownTest() {
	ctx := context.Background()
	_, err := s.db.ExecContext(ctx, "TRUNCATE TABLE users RESTART IDENTITY CASCADE")
	require.NoError(s.T(), err)
}
