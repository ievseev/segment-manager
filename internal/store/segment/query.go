package segment

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

var qb = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

const (
	segmentsTable = "segments"

	idField   = "id"
	slugField = "slug"
)

func buildInsertQuery(segmentName string) (string, []interface{}, error) {
	returningClause := fmt.Sprintf("RETURNING %s", idField)

	return qb.Insert(segmentsTable).Columns(slugField).Values(segmentName).Suffix(returningClause).ToSql()
}

func buildDeleteQuery(segmentName string) (string, []interface{}, error) {
	return qb.Delete(segmentsTable).Where(sq.Eq{slugField: segmentName}).ToSql()
}
