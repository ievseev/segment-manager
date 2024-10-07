package user

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

var qb = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

const (
	usersTable = "users"

	idField   = "id"
	nameField = "name"
)

func buildInsertQuery(userName string) (string, []interface{}, error) {
	returningClause := fmt.Sprintf("RETURNING %s", idField)

	return qb.Insert(usersTable).Columns(nameField).Values(userName).Suffix(returningClause).ToSql()
}
