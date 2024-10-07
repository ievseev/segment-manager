package user

import (
	"context"
	"database/sql"
)

type Executer interface {
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type UserSegment struct {
	Executer Executer
}

func New(executer Executer) *UserSegment {
	return &UserSegment{
		Executer: executer,
	}
}

func (s *UserSegment) Save(ctx context.Context, userID string, segmentIDs []string) error {
	query, args, err := buildUpsertQuery(userID, segmentIDs)
	if err != nil {
		return err
	}

	_, err = s.Executer.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
