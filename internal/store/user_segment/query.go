package user_segment

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"
)

var qb = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

const (
	usersSegmentsTable = "users_segments"

	userIDField     = "user_id"
	segmentIDsField = "segment_ids"
)

func buildUpsertQuery(userID int64, updateIDs []int64) (string, []interface{}, error) {
	return qb.Insert(usersSegmentsTable).
		Columns(userIDField, segmentIDsField).
		Values(userID, pq.Array(updateIDs)).
		Suffix(fmt.Sprintf(`ON CONFLICT (%s) DO UPDATE SET %s = EXCLUDED.segment_ids`, userIDField, segmentIDsField)).ToSql()
}

func buildSelectQuery(userID int64) (string, []interface{}, error) {
	return qb.Select(userIDField).From(usersSegmentsTable).Where(sq.Eq{userIDField: userID}).ToSql()
}
