package segment

import (
	"context"
	"database/sql"
)

type Executer interface {
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type Segment struct {
	Executer Executer
}

func New(executer Executer) *Segment {
	return &Segment{
		Executer: executer,
	}
}

func (s *Segment) Create(ctx context.Context, segmentName string) (int64, error) {
	var segmentID int64

	query, args, err := buildInsertQuery(segmentName)
	if err != nil {
		return segmentID, err
	}

	err = s.Executer.QueryRowContext(ctx, query, args...).Scan(&segmentID)
	if err != nil {
		return segmentID, err
	}

	return segmentID, nil
}

func (s *Segment) Delete(ctx context.Context, segmentName string) error {
	query, args, err := buildDeleteQuery(segmentName)
	if err != nil {
		return err
	}

	_, err = s.Executer.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
