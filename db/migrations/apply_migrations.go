package migrations

import (
	"errors"

	"github.com/golang-migrate/migrate/v4"

	"segment-manager/internal/config"
)

func Run(cfg *config.Config) error {
	m, err := migrate.New(
		cfg.MigrationsPath,
		cfg.StoragePath,
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
