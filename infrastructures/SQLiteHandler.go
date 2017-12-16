package infrastructures

import (
	"database/sql"
	"fmt"
	"github.com/irahardianto/service-pattern-go/interfaces"
)

type SQLiteHandler struct {
	Conn *sql.DB
}

func (handler *SQLiteHandler) Execute(statement string) {
	handler.Conn.Exec(statement)
}

func (handler *SQLiteHandler) Query(statement string) (interfaces.IRow, error) {
	//fmt.Println(statement)
	rows, err := handler.Conn.Query(statement)

	if err != nil {
		fmt.Println(err)
		return new(SqliteRow),err
	}
	row := new(SqliteRow)
	row.Rows = rows

	return row, nil
}

type SqliteRow struct {
	Rows *sql.Rows
}

func (r SqliteRow) Scan(dest ...interface{}) error {
	err := r.Rows.Scan(dest...)
	if err != nil {
		return err
	}

	return  nil
}

func (r SqliteRow) Next() bool {
	return r.Rows.Next()
}