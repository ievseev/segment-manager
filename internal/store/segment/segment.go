package segment

import (
	"context"
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
func (p *PG) SaveSegment(ctx context.Context, segmentName string) error {
	query := fmt.Sprintf("INSERT INTO segments (slug) VALUES (%s)", "$1")

	_, err := p.db.Exec(query, segmentName)
	if err != nil {
		return err
	}

	return nil
}

// TODO прокинуть контекст ?
func (p *PG) DeleteSegment(ctx context.Context, segmentName string) error {
	query := fmt.Sprintf("DELETE FROM segments WHERE (slug)=(%s)", "$1")

	_, err := p.db.Exec(query, segmentName)
	if err != nil {
		return err
	}

	return nil
}
