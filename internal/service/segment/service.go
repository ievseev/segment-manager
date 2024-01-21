package segment

import "context"

type segmentCreateRepo interface {
	Create(name string) error
}

type Service struct {
	// TODO создание сегмента в БД
}

func New() *Service {
	return &Service{}
}

func (s *Service) CreateSegment(ctx context.Context, name string) (int64, error) {
	return 1, nil
}
