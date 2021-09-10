package main

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type dbConn interface {
	Begin(context.Context) (pgx.Tx, error)
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	Ping(context.Context) error
	Prepare(context.Context, string, string) (*pgconn.StatementDescription, error)
	Close(context.Context) error
}

func ReadFromDatabase(conn dbConn) ([]int, error) {
	var result []int

	rows, err := conn.Query(context.TODO(), "SELECT nullable_int_column FROM applications LIMIT 2")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var columnValue *int
		err = rows.Scan(&columnValue)
		if err != nil {
			return nil, err
		}

		if columnValue != nil {
			result = append(result, *columnValue)
		}
	}
	return result, nil
}
