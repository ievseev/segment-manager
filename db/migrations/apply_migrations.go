package migrations

import (
	"errors"

	"github.com/golang-migrate/migrate/v4"

	"segment-manager/internal/config"
)

func Run(cnf *config.Config) error {
	m, err := migrate.New(
		cnf.MigrationsPath,
		cnf.StoragePath,
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
