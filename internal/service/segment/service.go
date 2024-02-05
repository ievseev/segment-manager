package segment

import (
	"context"
)

type segmentRepo interface {
	SaveSegment(ctx context.Context, segmentName string) error
	DeleteSegment(ctx context.Context, segmentName string) error
}

type Service struct {
	segmentRepo segmentRepo
}

func New(segmentRepo segmentRepo) *Service {
	return &Service{segmentRepo: segmentRepo}
}

func (s *Service) CreateSegment(ctx context.Context, name string) error {
	err := s.segmentRepo.SaveSegment(ctx, name)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteSegment(ctx context.Context, name string) error {
	err := s.segmentRepo.DeleteSegment(ctx, name)
	if err != nil {
		return err
	}

	return nil
}
