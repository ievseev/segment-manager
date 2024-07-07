package segment

import (
	"context"
	"fmt"
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

// TODO прокинуть контекст ?
func (s *Segment) Save(ctx context.Context, segmentName string) error {
	query := fmt.Sprintf("INSERT INTO segments (slug) VALUES (%s)", "$1")

	_, err := s.executer.ExecContext(ctx, query, segmentName)
	if err != nil {
		return err
	}

	return nil
}

// TODO прокинуть контекст ?
func (s *Segment) Delete(ctx context.Context, segmentName string) error {
	query := fmt.Sprintf("DELETE FROM segments WHERE (slug)=(%s)", "$1")

	_, err := s.executer.ExecContext(ctx, query, segmentName)
	if err != nil {
		return err
	}

	return nil
}
