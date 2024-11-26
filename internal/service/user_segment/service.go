package user_segment

import "context"

type UserSegmentRepo interface {
	SelectUserSegments(ctx context.Context, userID int64) ([]string, error)
	UpsertUserSegments(ctx context.Context, userID int64, slugsToAdd, slugsToDelete []string) error
}

type UserSegment struct {
	UserSegmentRepo UserSegmentRepo
}

func New(userSegmentRepo UserSegmentRepo) *UserSegment {
	return &UserSegment{UserSegmentRepo: userSegmentRepo}
}

func (us *UserSegment) Get(ctx context.Context, userID int64) ([]string, error) {
	slugs, err := us.UserSegmentRepo.SelectUserSegments(ctx, userID)

	if err != nil {
		return nil, err
	}

	return slugs, nil
}

func (us *UserSegment) Update(ctx context.Context, userID int64, slugsToAdd, slugsToDelete []string) error {
	err := us.UserSegmentRepo.UpsertUserSegments(ctx, userID, slugsToAdd, slugsToDelete)

	if err != nil {
		return err
	}

	return nil
}
