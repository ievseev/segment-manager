package user_segment

import (
	"context"

	usDB "segment-manager/internal/store/user_segment"
)

type UserSegmentRepo interface {
	SelectUserSegments(ctx context.Context, userID int64) ([]usDB.UserSegmentDB, error)
	UpsertUserSegments(ctx context.Context, userID int64, slugsToAdd, slugsToDelete []string) error
}

type UserSegment struct {
	UserSegmentRepo UserSegmentRepo
}

func New(userSegmentRepo UserSegmentRepo) *UserSegment {
	return &UserSegment{UserSegmentRepo: userSegmentRepo}
}

func (us *UserSegment) Get(ctx context.Context, userID int64) ([]Segment, error) {
	userSegments, err := us.UserSegmentRepo.SelectUserSegments(ctx, userID)

	if err != nil {
		return nil, err
	}

	return convUserSegmentFromDB(userSegments), nil
}

func convUserSegmentFromDB(userSegments []usDB.UserSegmentDB) []Segment {
	segments := make([]Segment, 0, len(userSegments))

	for _, segment := range userSegments {
		segments = append(segments, Segment{
			ID:   segment.ID,
			Slug: segment.Slug,
		})
	}

	return segments
}

func (us *UserSegment) Update(ctx context.Context, userID int64, slugsToAdd, slugsToDelete []string) error {
	err := us.UserSegmentRepo.UpsertUserSegments(ctx, userID, slugsToAdd, slugsToDelete)

	if err != nil {
		return err
	}

	return nil
}
