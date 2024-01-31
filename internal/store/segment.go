package store

import (
	"fmt"
	"segment-manager/internal/storage/postgres"
)

type PG struct {
	db *postgres.Storage
}

func New(db *postgres.Storage) *PG {
	return &PG{db: db}
}

// TODO прокинуть контекст ?
func (p *PG) SaveSegment(segmentName string) error {
	query := fmt.Sprintf("INSERT INTO segments (name) VALUES (%s)", "$1")

	_, err := p.db.Exec(query, segmentName)
	if err != nil {
		return err
	}

	return nil
}
