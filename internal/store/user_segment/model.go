package user_segment

type UserSegmentDB struct {
	ID   int64  `db:"id"`
	Slug string `db:"slug"`
}
