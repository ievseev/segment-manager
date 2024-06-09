package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Импортируем драйвер для PostgreSQL

	"segment-manager/internal/config"
)

type Storage struct {
	db *sql.DB
}

func New(cfg *config.Config) (*Storage, error) {
	const op = "storage.postgres.New"

	// TODO: Прикрутить SSL
	db, err := sql.Open("postgres", cfg.StoragePath)
	if err != nil {
		return nil, fmt.Errorf("DB open error while %s operation: %v", op, err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("DB ping error while %s operation: %v", op, err)
	}

	return &Storage{db: db}, nil
}

// TODO Сlose() и другие ребята?

func (s *Storage) Exec(query string, args ...interface{}) (sql.Result, error) {
	return s.db.Exec(query, args...)
}
