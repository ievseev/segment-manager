//go:generate mockgen -source=service.go -destination=mocks/mock_userRepo.go -package=mocks

package user

import (
	"context"
)

type userRepo interface {
	Create(ctx context.Context, userName string) (int64, error)
}

type Service struct {
	userRepo userRepo
}

func New(userRepo userRepo) *Service {
	return &Service{userRepo: userRepo}
}

func (s *Service) CreateUser(ctx context.Context, name string) (int64, error) {
	var userID int64

	userID, err := s.userRepo.Create(ctx, name)
	if err != nil {
		return userID, err
	}

	return userID, nil
}
