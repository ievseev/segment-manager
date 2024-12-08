package segment

import (
	sq "github.com/Masterminds/squirrel"
)

var qb = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

const (
	segmentsTable = "segments"

	idField   = "id"
	slugField = "slug"
)

func buildInsertQuery(slug string) (string, []interface{}, error) {
	return qb.Insert(segmentsTable).Columns(slugField).Values(slug).Suffix("RETURNING id").ToSql()
}

func buildDeleteQuery(slug string) (string, []interface{}, error) {
	return qb.Delete(segmentsTable).Where(sq.Eq{slugField: slug}).ToSql()
}

func BuildSelectQuery(slug []string) (string, []interface{}, error) {
	return qb.Select(idField).From(segmentsTable).Where(sq.Eq{slugField: slug}).ToSql()
}
