package user

import (
	sq "github.com/Masterminds/squirrel"
)

var qb = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

const (
	UsersTable = "users"

	NameField = "name"
)

func buildInsertQuery(userName string) (string, []interface{}, error) {
	return qb.Insert(UsersTable).Columns(NameField).Values(userName).Suffix("RETURNING id").ToSql()
}
