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
	// Создаем подзапрос для извлечения segment_ids
	subQuery := qb.Select("unnest(segment_ids)").
		From(usersSegmentsTable).
		Where(sq.Eq{userIDField: userID})

	// Получаем SQL и аргументы для подзапроса
	subSQL, args, err := subQuery.ToSql()
	if err != nil {
		return "", nil, fmt.Errorf("ошибка построения подзапроса: %w", err)
	}

	// Создаем основной запрос, используя подзапрос в секции WHERE
	mainQuery := qb.Select("id", "slug").
		From("segments").
		Where(fmt.Sprintf("id IN (%s)", subSQL))

	// Получаем SQL и аргументы для основного запроса
	mainSQL, _, err := mainQuery.ToSql()
	if err != nil {
		return "", nil, fmt.Errorf("ошибка построения основного запроса: %w", err)
	}

	return mainSQL, args, nil
}
