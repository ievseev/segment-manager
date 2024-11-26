package user

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

var qb = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

const (
	UsersTable = "users"

	IdField   = "id"
	NameField = "name"
)

func buildInsertQuery(userName string) (string, []interface{}, error) {
	returningClause := fmt.Sprintf("RETURNING %s", IdField)

	return qb.Insert(UsersTable).Columns(NameField).Values(userName).Suffix(returningClause).ToSql()
}
