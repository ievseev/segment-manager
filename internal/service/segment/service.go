package segment

import (
	"context"
	"segment-manager/internal/store/segment"
)

type segmentCreateRepo interface {
	SaveSegment(ctx context.Context, segmentName string) error
}

type Service struct {
	segmentRepo *segment.PG
}

func New(segmentRepo *segment.PG) *Service {
	return &Service{segmentRepo: segmentRepo}
}

func (s *Service) CreateSegment(ctx context.Context, name string) error {
	err := s.segmentRepo.SaveSegment(ctx, name)
	if err != nil {
		return err
	}

	return nil
}
