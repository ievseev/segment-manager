package user

import (
	"context"
	"database/sql"
)

type Executer interface {
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type User struct {
	Executer Executer
}

func New(executer Executer) *User {
	return &User{
		Executer: executer,
	}
}

func (s *User) Create(ctx context.Context, userName string) (int64, error) {
	var userID int64

	query, args, err := buildInsertQuery(userName)
	if err != nil {
		return userID, err
	}

	err = s.Executer.QueryRowContext(ctx, query, args...).Scan(&userID)
	if err != nil {
		return userID, err
	}

	return userID, nil
}
