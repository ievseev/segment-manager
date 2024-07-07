package segment

import (
	"context"
)

type segmentRepo interface {
	Save(ctx context.Context, segmentName string) error
	Delete(ctx context.Context, segmentName string) error
}

type Service struct {
	segmentRepo segmentRepo
}

func New(segmentRepo segmentRepo) *Service {
	return &Service{segmentRepo: segmentRepo}
}

func (s *Service) CreateSegment(ctx context.Context, name string) error {
	err := s.segmentRepo.Save(ctx, name)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Delete(ctx context.Context, name string) error {
	err := s.segmentRepo.Delete(ctx, name)
	if err != nil {
		return err
	}

	return nil
}
