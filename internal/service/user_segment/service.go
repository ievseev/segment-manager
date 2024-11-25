package user_segment

import "context"

type UserSegmentRepo interface {
	Upsert(ctx context.Context, userID int64, slugsToAdd, slugsToDelete []string) error
}

type UserSegment struct {
	UserSegmentRepo UserSegmentRepo
}

func New(userSegmentRepo UserSegmentRepo) *UserSegment {
	return &UserSegment{UserSegmentRepo: userSegmentRepo}
}

func (us *UserSegment) Update(ctx context.Context, userID int64, slugsToAdd, slugsToDelete []string) error {
	err := us.UserSegmentRepo.Upsert(ctx, userID, slugsToAdd, slugsToDelete)

	if err != nil {
		return err
	}

	return nil
}
