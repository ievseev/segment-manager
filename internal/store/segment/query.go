package segment

import sq "github.com/Masterminds/squirrel"

var qb = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

const (
	segmentsTable = "segments"

	slugField = "slug"
)

func buildInsertQuery(segmentName string) (string, []interface{}, error) {
	return qb.Insert(segmentsTable).Columns(slugField).Values(segmentName).ToSql()
}

func buildDeleteQuery(segmentName string) (string, []interface{}, error) {
	return qb.Delete(segmentsTable).Where(sq.Eq{slugField: segmentName}).ToSql()
}
