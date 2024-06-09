package migrations

import (
	"errors"

	"github.com/golang-migrate/migrate/v4"
)

func Run(migrationsPath, storagePath string) error {
	m, err := migrate.New(
		migrationsPath,
		storagePath,
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
