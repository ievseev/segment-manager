package user

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

func buildUpsertQuery(userID string, segmentIDs []string) (string, []interface{}, error) {
	return qb.Insert(usersSegmentsTable).
		Columns(userIDField, segmentIDsField).
		Values(userID, pq.Array(segmentIDs)).
		Suffix(fmt.Sprintf("ON CONFLICT (%s) DO UPDATE SET %s = EXCLUDED.%s", userIDField, segmentIDs, segmentIDs)).
		ToSql()
}
