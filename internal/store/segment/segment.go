package segment

import (
	"context"
)

type Executer interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (interface{}, error)
}

type Segment struct {
	executer Executer
}

func New(executer Executer) *Segment {
	return &Segment{
		executer: executer,
	}
}

func (s *Segment) Save(ctx context.Context, segmentName string) error {
	query, args, err := buildInsertQuery(segmentName)
	if err != nil {
		return err
	}

	_, err = s.executer.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (s *Segment) Delete(ctx context.Context, segmentName string) error {
	query, args, err := buildDeleteQuery(segmentName)
	if err != nil {
		return err
	}

	_, err = s.executer.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
