/* Go database/sql test driver
 */
package sqltest

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"io"
	"time"
)

// --------------------------------------------------
// testDriver - driver.Driver
// - testConn
// --------------------------------------------------
type testDriver struct {
}

func init() {
	sql.Register("sqltest", &testDriver{})
}

func (d *testDriver) Open(dsn string) (driver.Conn, error) {
	if dsn == "open_error" {
		return nil, errors.New("open error")
	}
	return &testConn{dns: dsn}, nil
}

// --------------------------------------------------
// testConn - driver.Conn
// - testStmt
// - testTx
// --------------------------------------------------
type testConn struct {
	dns string
}

func (c *testConn) Prepare(query string) (driver.Stmt, error) {
	return &testStmt{dns: c.dns, query: query}, nil
}

func (c *testConn) Begin() (driver.Tx, error) {
	return &testTx{}, nil
}

func (c *testConn) Close() error {
	if c.dns == "close_error" {
		return errors.New("close error")
	}
	return nil
}

// --------------------------------------------------
// testTx - driver.Tx
// --------------------------------------------------
type testTx struct{}

func (t *testTx) Commit() error {
	return nil
}

func (t *testTx) Rollback() error {
	return nil
}

// --------------------------------------------------
// testRows - driver.Rows
// --------------------------------------------------
type testRows struct {
	dns    string
	pos    int
	values [][]driver.Value
	cols   []string
	dbType []string
}

var resultRows = testRows{
	dns: "",
	pos: 0,
	values: [][]driver.Value{
		{"value", 3, true, 1.23, time.Now(), time.Now(), `{}`, 4},
		{"value2", nil, false, 2.34, time.Now(), time.Now(), `{}`, 5},
	},
	cols:   []string{"value", "integer", "boolean", "numeric", "date", "datetime", "json", "count"},
	dbType: []string{"TEXT", "INTEGER", "BOOLEAN", "NUMERIC", "DATE", "DATETIME", "JSON", ""},
}

func (r *testRows) Next(dest []driver.Value) error {
	r.pos++
	if r.pos > len(r.values) || r.dns == "not_found" {
		return io.EOF
	}

	copy(dest[:], r.values[r.pos-1])

	return nil
}

func (r *testRows) Scan(value interface{}) error {
	return nil
}

func (r *testRows) Close() error {
	return nil
}

func (r *testRows) Columns() []string {
	return r.cols
}

func (rs *testRows) ColumnTypeDatabaseTypeName(index int) string {
	return resultRows.dbType[index]
}

// --------------------------------------------------
// testStmt - driver.Stmt
// - testRows
// - testResult
// --------------------------------------------------
type testStmt struct {
	dns   string
	query string
}

func (stmt *testStmt) Close() error {
	return nil
}

func (stmt *testStmt) NumInput() int {
	return -1
}

func (stmt *testStmt) Exec(args []driver.Value) (driver.Result, error) {
	return &testResult{dns: stmt.dns}, nil
}

func (stmt *testStmt) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	if stmt.dns == "exec_error" {
		return &testResult{}, errors.New(stmt.dns)
	}
	return &testResult{dns: stmt.dns}, nil
}

func (stmt *testStmt) Query(args []driver.Value) (driver.Rows, error) {
	return &resultRows, nil
}

func (stmt *testStmt) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	if stmt.dns == "query_error" {
		return &testRows{}, errors.New(stmt.dns)
	}
	return &resultRows, nil
}

// --------------------------------------------------
// testResult - driver.Result
// --------------------------------------------------
type testResult struct {
	dns string
}

func (r *testResult) LastInsertId() (int64, error) {
	return 0, nil
}

func (r *testResult) RowsAffected() (int64, error) {
	if r.dns == "rows_affected" {
		return 1, nil
	}
	return 0, nil
}
