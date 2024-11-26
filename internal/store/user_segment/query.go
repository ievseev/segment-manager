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

func buildSelectUserSegmentsQuery(userID int64) (string, []interface{}, error) {
	// Создаем подзапрос
	subQuery := sq.Select("unnest(segment_ids)").
		From("users_segments").
		Where(sq.Eq{"user_id": userID})

	// Создаем основной запрос с использованием подзапроса
	queryBuilder := sq.Select("slug").
		From("segments").
		Where(sq.Expr(`id IN (?)`, subQuery))

	return queryBuilder.ToSql()
}
