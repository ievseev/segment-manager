//go:generate mockgen -source=service.go -destination=mocks/mock_segmentRepo.go -package=mocks

package segment

import (
	"context"
)

type segmentRepo interface {
	Save(ctx context.Context, segmentName string) (int64, error)
	Delete(ctx context.Context, segmentName string) error
}

type Service struct {
	segmentRepo segmentRepo
}

func New(segmentRepo segmentRepo) *Service {
	return &Service{segmentRepo: segmentRepo}
}

func (s *Service) CreateSegment(ctx context.Context, name string) (int64, error) {
	var segmentID int64

	segmentID, err := s.segmentRepo.Save(ctx, name)
	if err != nil {
		return segmentID, err
	}

	return segmentID, nil
}

func (s *Service) Delete(ctx context.Context, name string) error {
	err := s.segmentRepo.Delete(ctx, name)
	if err != nil {
		return err
	}

	return nil
}
