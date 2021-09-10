package main

import (
	"context"
	"testing"

	pgxmock "github.com/pashagolub/pgxmock"
)

var columns = []string{"null_column"}

func TestReadFromDatabase(t *testing.T) {
	// open database stub
	ctx := context.Background()
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mock.Close(ctx)

	mock.ExpectQuery("SELECT nullable_int_column FROM applications").
		WillReturnRows(
			mock.NewRows(columns).AddRow(nil).AddRow(1),
		)

	_, err = ReadFromDatabase(mock)

	// test fails here with:
	// error: Destination kind 'ptr' not supported for value kind 'int' of column 'null_column'
	if err != nil {
		t.Errorf("error: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
