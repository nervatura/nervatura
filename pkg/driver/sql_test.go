package database

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"reflect"
	"testing"

	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	_ "github.com/nervatura/nervatura/v6/test/sqltest"
)

func Test_registerDriver(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "sql",
			args: args{
				name: "sql",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			registerDriver(tt.args.name)
		})
	}
}

func TestSQLDriver_Properties(t *testing.T) {
	type fields struct {
		alias   string
		connStr string
		engine  string
		Db      *sql.DB
		closed  bool
		Config  cu.IM
	}
	tests := []struct {
		name   string
		fields fields
		want   struct{ SQL, Transaction bool }
	}{
		{
			name:   "call",
			fields: fields{},
			want:   struct{ SQL, Transaction bool }{SQL: true, Transaction: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &SQLDriver{
				alias:   tt.fields.alias,
				connStr: tt.fields.connStr,
				engine:  tt.fields.engine,
				Db:      tt.fields.Db,
				closed:  tt.fields.closed,
				Config:  tt.fields.Config,
			}
			if got := ds.Properties(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SQLDriver.Properties() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSQLDriver_Connection(t *testing.T) {
	type fields struct {
		alias   string
		connStr string
		engine  string
		Db      *sql.DB
		closed  bool
		Config  cu.IM
	}
	tests := []struct {
		name   string
		fields fields
		want   struct {
			Alias     string
			Connected bool
			Engine    string
		}
	}{
		{
			name: "call",
			fields: fields{
				engine: "postgres",
				Db:     nil,
				alias:  "test",
			},
			want: struct {
				Alias     string
				Connected bool
				Engine    string
			}{
				Alias:     "test",
				Connected: false,
				Engine:    "postgres",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &SQLDriver{
				alias:   tt.fields.alias,
				connStr: tt.fields.connStr,
				engine:  tt.fields.engine,
				Db:      tt.fields.Db,
				closed:  tt.fields.closed,
				Config:  tt.fields.Config,
			}
			if got := ds.Connection(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SQLDriver.Connection() = %v, want %v", got, tt.want)
			}
		})
	}
}

type errorDriver struct{}

func (d errorDriver) Open(name string) (driver.Conn, error) {
	return &errorConn{}, nil
}

type errorConn struct{}

func (c *errorConn) Prepare(query string) (driver.Stmt, error) { return nil, nil }
func (c *errorConn) Close() error {
	return errors.New("forced close error")
}
func (c *errorConn) Begin() (driver.Tx, error) { return nil, nil }

func init() {
	sql.Register("error", errorDriver{})
}

func TestSQLDriver_CreateConnection(t *testing.T) {
	type fields struct {
		alias   string
		connStr string
		engine  string
		Db      *sql.DB
		closed  bool
		Config  cu.IM
	}
	type args struct {
		alias   string
		connStr string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "db_close_error",
			fields: fields{
				Db: func() *sql.DB {
					db, _ := sql.Open("sqltest", "close_error")
					db.Ping()
					return db
				}(),
			},
			args: args{
				alias:   "test",
				connStr: "sqltest://:memory:",
			},
			wantErr: true,
		},
		{
			name: "sqlite_pragma",
			fields: fields{
				Db: func() *sql.DB {
					db, _ := sql.Open("sqltest", "test")
					return db
				}(),
			},
			args: args{
				alias:   "test",
				connStr: "sqltest://test",
			},
			wantErr: false,
		},
		{
			name:   "mysql",
			fields: fields{},
			args: args{
				alias:   "test",
				connStr: "mysql://test:test@tcp(localhost:1234)/nervatura?value=1",
			},
			wantErr: true,
		},
		{
			name:   "mysql",
			fields: fields{},
			args: args{
				alias:   "test",
				connStr: "mysql://",
			},
			wantErr: true,
		},
		{
			name:   "invalid_dbs_types",
			fields: fields{},
			args: args{
				alias:   "test",
				connStr: "invalid://test:test@tcp(localhost:1234)/nervatura",
			},
			wantErr: true,
		},
		{
			name:   "mssql",
			fields: fields{},
			args: args{
				alias:   "test",
				connStr: "mssql://sa:Password1234_1@localhost:1433?database=nervatura",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &SQLDriver{
				alias:   tt.fields.alias,
				connStr: tt.fields.connStr,
				engine:  tt.fields.engine,
				Db:      tt.fields.Db,
				closed:  tt.fields.closed,
				Config:  tt.fields.Config,
			}
			if err := ds.CreateConnection(tt.args.alias, tt.args.connStr); (err != nil) != tt.wantErr {
				t.Errorf("SQLDriver.CreateConnection() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSQLDriver_CloseConnection(t *testing.T) {
	type fields struct {
		alias   string
		connStr string
		engine  string
		Db      *sql.DB
		closed  bool
		Config  cu.IM
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "close",
			fields: fields{
				Db: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &SQLDriver{
				alias:   tt.fields.alias,
				connStr: tt.fields.connStr,
				engine:  tt.fields.engine,
				Db:      tt.fields.Db,
				closed:  tt.fields.closed,
				Config:  tt.fields.Config,
			}
			if err := ds.CloseConnection(); (err != nil) != tt.wantErr {
				t.Errorf("SQLDriver.CloseConnection() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSQLDriver_checkConnection(t *testing.T) {
	type fields struct {
		alias   string
		connStr string
		engine  string
		Db      *sql.DB
		closed  bool
		Config  cu.IM
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "reconnect",
			fields: fields{
				closed:  true,
				alias:   "test",
				connStr: "test://test:test@tcp(localhost:1234)/nervatura",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &SQLDriver{
				alias:   tt.fields.alias,
				connStr: tt.fields.connStr,
				engine:  tt.fields.engine,
				Db:      tt.fields.Db,
				closed:  tt.fields.closed,
				Config:  tt.fields.Config,
			}
			ds.checkConnection()
		})
	}
}

func TestSQLDriver_getPrmString(t *testing.T) {
	type fields struct {
		alias   string
		connStr string
		engine  string
		Db      *sql.DB
		closed  bool
		Config  cu.IM
	}
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "postgres",
			fields: fields{
				engine: "postgres",
			},
			args: args{
				index: 1,
			},
			want: "$1",
		},
		{
			name: "sqlite",
			fields: fields{
				engine: "sqlite",
			},
			args: args{
				index: 1,
			},
			want: "?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &SQLDriver{
				alias:   tt.fields.alias,
				connStr: tt.fields.connStr,
				engine:  tt.fields.engine,
				Db:      tt.fields.Db,
				closed:  tt.fields.closed,
				Config:  tt.fields.Config,
			}
			if got := ds.getPrmString(tt.args.index); got != tt.want {
				t.Errorf("SQLDriver.getPrmString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSQLDriver_getFilterString(t *testing.T) {
	type fields struct {
		alias   string
		connStr string
		engine  string
		Db      *sql.DB
		closed  bool
		Config  cu.IM
	}
	type args struct {
		filter    md.Filter
		start     bool
		sqlString string
		params    []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  []any
	}{
		{
			name:   "start_eq",
			fields: fields{},
			args: args{
				start:     true,
				sqlString: "",
				params:    []any{},
				filter: md.Filter{
					Field: "fieldname",
					Comp:  "==",
					Value: "value",
				},
			},
			want:  "(fieldname=?)",
			want1: []any{"value"},
		},
		{
			name:   "and_like",
			fields: fields{},
			args: args{
				sqlString: "",
				params:    []any{},
				filter: md.Filter{
					Or:    false,
					Field: "fieldname",
					Comp:  "like",
					Value: "value",
				},
			},
			want:  " and (fieldname like ?)",
			want1: []any{"value"},
		},
		{
			name:   "or_is",
			fields: fields{},
			args: args{
				sqlString: "",
				params:    []any{},
				filter: md.Filter{
					Or:    true,
					Field: "fieldname",
					Comp:  "is",
					Value: "null",
				},
			},
			want:  " or fieldname is null",
			want1: []any{},
		},
		{
			name:   "and_in",
			fields: fields{},
			args: args{
				sqlString: "",
				params:    []any{},
				filter: md.Filter{
					Or:    false,
					Field: "fieldname",
					Comp:  "in",
					Value: "value1,value2,value3",
				},
			},
			want:  " and (fieldname in(?,?,?))",
			want1: []any{"value1", "value2", "value3"},
		},
		{
			name:   "and_in_query",
			fields: fields{},
			args: args{
				sqlString: "",
				params:    []any{},
				filter: md.Filter{
					Field: "fieldname",
					Comp:  "in",
					Value: md.Query{
						Fields: []string{"field1", "field2"},
						From:   "table",
						Filters: []md.Filter{
							{
								Field: "field1",
								Comp:  "==",
								Value: "value",
							},
						},
					},
				},
			},
			want:  " and (fieldname in(select field1,field2 from table where (field1=?)))",
			want1: []any{"value"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &SQLDriver{
				alias:   tt.fields.alias,
				connStr: tt.fields.connStr,
				engine:  tt.fields.engine,
				Db:      tt.fields.Db,
				closed:  tt.fields.closed,
				Config:  tt.fields.Config,
			}
			got, got1 := ds.getFilterString(tt.args.filter, tt.args.start, tt.args.sqlString, tt.args.params)
			if got != tt.want {
				t.Errorf("SQLDriver.getFilterString() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SQLDriver.getFilterString() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSQLDriver_decodeSQL(t *testing.T) {
	type fields struct {
		alias   string
		connStr string
		engine  string
		Db      *sql.DB
		closed  bool
		Config  cu.IM
	}
	type args struct {
		queries []md.Query
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  []any
		engine string
	}{
		{
			name:   "queries",
			fields: fields{},
			engine: "sqlite",
			args: args{
				queries: []md.Query{
					{
						Fields: []string{"field1", "field2"},
						From:   "table",
						Filters: []md.Filter{
							{
								Field: "field1",
								Comp:  "==",
								Value: "value",
							},
							{
								Or:    false,
								Field: "field2",
								Comp:  "in",
								Value: "1,2,3",
							},
						},
						OrderBy: []string{"field1"},
					},
					{
						Fields: []string{"field1", "field2"},
						From:   "table",
						Filters: []md.Filter{
							{
								Field: "field1",
								Comp:  "==",
								Value: "value",
							},
							{
								Or:    false,
								Field: "field2",
								Comp:  "in",
								Value: "1,2,3",
							},
						},
						OrderBy: []string{"field1"},
						Limit:   10,
						Offset:  5,
					},
				},
			},
			want:  "select field1,field2 from table where (field1=?) and (field2 in(?,?,?)) order by field1 union select field1,field2 from table where (field1=?) and (field2 in(?,?,?)) order by field1 limit 10 offset 5",
			want1: []any{"value", "1", "2", "3", "value", "1", "2", "3"},
		},
		{
			name:   "queries_mssql",
			fields: fields{},
			engine: "mssql",
			args: args{
				queries: []md.Query{
					{
						Fields: []string{"field1", "field2"},
						From:   "table",
						Filters: []md.Filter{
							{
								Field: "field1",
								Comp:  "==",
								Value: "value",
							},
							{
								Or:    false,
								Field: "field2",
								Comp:  "in",
								Value: "1,2,3",
							},
						},
						OrderBy: []string{"field1"},
					},
					{
						Fields: []string{"field1", "field2"},
						From:   "table",
						Filters: []md.Filter{
							{
								Field: "field1",
								Comp:  "==",
								Value: "value",
							},
							{
								Or:    false,
								Field: "field2",
								Comp:  "in",
								Value: "1,2,3",
							},
						},
						OrderBy: []string{"field1"},
						Limit:   10,
						Offset:  5,
					},
				},
			},
			want:  "select top(10) field1,field2 from table where (field1=?) and (field2 in(?,?,?)) order by field1 union select top(10) field1,field2 from table where (field1=?) and (field2 in(?,?,?)) order by field1 offset 5",
			want1: []any{"value", "1", "2", "3", "value", "1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &SQLDriver{
				alias:   tt.fields.alias,
				connStr: tt.fields.connStr,
				engine:  tt.engine,
				Db:      tt.fields.Db,
				closed:  tt.fields.closed,
				Config:  tt.fields.Config,
			}
			got, got1 := ds.decodeSQL(tt.args.queries)
			if got != tt.want {
				t.Errorf("SQLDriver.decodeSQL() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SQLDriver.decodeSQL() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSQLDriver_Query(t *testing.T) {
	type fields struct {
		alias   string
		connStr string
		engine  string
		Db      *sql.DB
		closed  bool
		Config  cu.IM
	}
	type args struct {
		queries []md.Query
		trans   any
	}
	db, _ := sql.Open("sqltest", "query_error")
	trans, _ := db.Begin()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "query",
			fields: fields{
				engine: "sqltest",
				Db:     db,
			},
			args: args{
				queries: []md.Query{
					{
						Fields: []string{"field1", "field2"},
						From:   "table",
						Filters: []md.Filter{
							{
								Field: "field1",
								Comp:  "==",
								Value: "value",
							},
							{
								Or:    false,
								Field: "field2",
								Comp:  "in",
								Value: "1,2,3",
							},
						},
						OrderBy: []string{"field1"},
					},
				},
				trans: trans,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &SQLDriver{
				alias:   tt.fields.alias,
				connStr: tt.fields.connStr,
				engine:  tt.fields.engine,
				Db:      tt.fields.Db,
				closed:  tt.fields.closed,
				Config:  tt.fields.Config,
			}
			_, err := ds.Query(tt.args.queries, tt.args.trans)
			if (err != nil) != tt.wantErr {
				t.Errorf("SQLDriver.Query() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSQLDriver_QuerySQL(t *testing.T) {
	type fields struct {
		alias   string
		connStr string
		engine  string
		Db      *sql.DB
		closed  bool
		Config  cu.IM
	}
	type args struct {
		sqlString string
		params    []any
		trans     any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "query",
			fields: fields{
				engine: "postgres",
				Db: func() *sql.DB {
					db, _ := sql.Open("sqltest", "test")
					return db
				}(),
			},
			args: args{
				sqlString: "select *, 3 as count from test where id=? and name=?",
				params:    []any{"1", "test"},
			},
			wantErr: false,
		},
		{
			name: "too many parameters",
			fields: fields{
				engine: "postgres",
				Db: func() *sql.DB {
					db, _ := sql.Open("sqltest", "test")
					return db
				}(),
			},
			args: args{
				sqlString: "select *, 3 as count from test where id=? and name=? and id=?",
				params:    []any{"1", "test"},
			},
			wantErr: true,
		},
		{
			name: "invalid_trans",
			args: args{
				sqlString: "",
				params:    []any{},
				trans:     "trans",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &SQLDriver{
				alias:   tt.fields.alias,
				connStr: tt.fields.connStr,
				engine:  tt.fields.engine,
				Db:      tt.fields.Db,
				closed:  tt.fields.closed,
				Config:  tt.fields.Config,
			}
			_, err := ds.QuerySQL(tt.args.sqlString, tt.args.params, tt.args.trans)
			if (err != nil) != tt.wantErr {
				t.Errorf("SQLDriver.QuerySQL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

type testResult struct {
}

func (tr testResult) LastInsertId() (int64, error) {
	return 0, errors.New("error")
}

func (tr testResult) RowsAffected() (int64, error) {
	return 0, nil
}

func TestSQLDriver_lastInsertID(t *testing.T) {
	type fields struct {
		alias   string
		connStr string
		engine  string
		Db      *sql.DB
		closed  bool
		Config  cu.IM
	}
	type args struct {
		model  string
		result sql.Result
		trans  any
	}
	db, _ := sql.Open("sqltest", "test")
	trans, _ := db.Begin()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "postgres",
			fields: fields{
				engine: "postgres",
				Db:     db,
			},
			args: args{
				model:  "table",
				result: testResult{},
				trans:  trans,
			},
			wantErr: true,
		},
		{
			name: "mssql",
			fields: fields{
				engine: "mssql",
				Db:     db,
			},
			args: args{
				model:  "table",
				result: testResult{},
			},
			wantErr: true,
		},
		{
			name: "mysql",
			fields: fields{
				engine: "mysql",
				Db:     db,
			},
			args: args{
				model:  "table",
				result: testResult{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &SQLDriver{
				alias:   tt.fields.alias,
				connStr: tt.fields.connStr,
				engine:  tt.fields.engine,
				Db:      tt.fields.Db,
				closed:  tt.fields.closed,
				Config:  tt.fields.Config,
			}
			_, err := ds.lastInsertID(tt.args.model, tt.args.result, tt.args.trans)
			if (err != nil) != tt.wantErr {
				t.Errorf("SQLDriver.lastInsertID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSQLDriver_Update(t *testing.T) {
	type fields struct {
		alias   string
		connStr string
		engine  string
		Db      *sql.DB
		closed  bool
		Config  cu.IM
	}
	type args struct {
		options md.Update
	}
	db, _ := sql.Open("sqltest", "test")
	trans, _ := db.Begin()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "insert",
			fields: fields{
				engine: "sqlite",
				Db:     db,
			},
			args: args{
				options: md.Update{
					Model: "test",
					IDKey: int64(0),
					Values: cu.IM{
						"value": "insert",
					},
					Trans: func() any {
						db, _ := sql.Open("sqltest", "rows_affected")
						trans, _ := db.Begin()
						return trans
					}(),
				},
			},
			wantErr: false,
		},
		{
			name: "update",
			fields: fields{
				engine: "sqlite",
				Db:     db,
			},
			args: args{
				options: md.Update{
					Model: "test",
					IDKey: int64(1),
					Values: cu.IM{
						"value": "update",
					},
					Trans: func() any {
						db, _ := sql.Open("sqltest", "rows_affected")
						trans, _ := db.Begin()
						return trans
					}(),
				},
			},
			wantErr: false,
		},
		{
			name: "update",
			fields: fields{
				engine: "sqlite",
				Db:     db,
			},
			args: args{
				options: md.Update{
					Model: "test",
					IDKey: int64(5),
					Values: cu.IM{
						"value": "update",
					},
					Trans: trans,
				},
			},
			wantErr: false,
		},
		{
			name: "delete",
			fields: fields{
				engine: "sqlite",
				Db: func() *sql.DB {
					db, _ := sql.Open("sqltest", "exec_error")
					return db
				}(),
			},
			args: args{
				options: md.Update{
					Model:  "test2",
					IDKey:  int64(1),
					Values: cu.IM{},
				},
			},
			wantErr: true,
		},
		{
			name: "invalid_trans",
			fields: fields{
				engine: "sqlite",
				Db:     db,
			},
			args: args{
				options: md.Update{
					Model:  "test2",
					IDKey:  int64(1),
					Values: cu.IM{},
					Trans:  "trans",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &SQLDriver{
				alias:   tt.fields.alias,
				connStr: tt.fields.connStr,
				engine:  tt.fields.engine,
				Db:      tt.fields.Db,
				closed:  tt.fields.closed,
				Config:  tt.fields.Config,
			}
			_, err := ds.Update(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("SQLDriver.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSQLDriver_BeginTransaction(t *testing.T) {
	type fields struct {
		alias   string
		connStr string
		engine  string
		Db      *sql.DB
		closed  bool
		Config  cu.IM
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "begin",
			fields: fields{
				Db: func() *sql.DB {
					db, _ := sql.Open("sqltest", "test")
					return db
				}(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &SQLDriver{
				alias:   tt.fields.alias,
				connStr: tt.fields.connStr,
				engine:  tt.fields.engine,
				Db:      tt.fields.Db,
				closed:  tt.fields.closed,
				Config:  tt.fields.Config,
			}
			_, err := ds.BeginTransaction()
			if (err != nil) != tt.wantErr {
				t.Errorf("SQLDriver.BeginTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSQLDriver_CommitTransaction(t *testing.T) {
	type fields struct {
		alias   string
		connStr string
		engine  string
		Db      *sql.DB
		closed  bool
		Config  cu.IM
	}
	type args struct {
		trans any
	}
	db, _ := sql.Open("sqltest", "test")
	trans, _ := db.Begin()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "commit_ok",
			fields: fields{
				Db: db,
			},
			args: args{
				trans: trans,
			},
			wantErr: false,
		},
		{
			name: "invalid_trans",
			fields: fields{
				Db: db,
			},
			args: args{
				trans: "trans",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &SQLDriver{
				alias:   tt.fields.alias,
				connStr: tt.fields.connStr,
				engine:  tt.fields.engine,
				Db:      tt.fields.Db,
				closed:  tt.fields.closed,
				Config:  tt.fields.Config,
			}
			if err := ds.CommitTransaction(tt.args.trans); (err != nil) != tt.wantErr {
				t.Errorf("SQLDriver.CommitTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSQLDriver_RollbackTransaction(t *testing.T) {
	type fields struct {
		alias   string
		connStr string
		engine  string
		Db      *sql.DB
		closed  bool
		Config  cu.IM
	}
	type args struct {
		trans any
	}
	db, _ := sql.Open("sqltest", "test")
	trans, _ := db.Begin()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "rb_ok",
			fields: fields{
				Db: db,
			},
			args: args{
				trans: trans,
			},
			wantErr: false,
		},
		{
			name: "invalid_trans",
			fields: fields{
				Db: db,
			},
			args: args{
				trans: "trans",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &SQLDriver{
				alias:   tt.fields.alias,
				connStr: tt.fields.connStr,
				engine:  tt.fields.engine,
				Db:      tt.fields.Db,
				closed:  tt.fields.closed,
				Config:  tt.fields.Config,
			}
			if err := ds.RollbackTransaction(tt.args.trans); (err != nil) != tt.wantErr {
				t.Errorf("SQLDriver.RollbackTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSQLDriver_UpdateSQL(t *testing.T) {
	type fields struct {
		alias   string
		connStr string
		engine  string
		Db      *sql.DB
		closed  bool
		Config  cu.IM
	}
	type args struct {
		sqlString   string
		transaction any
	}
	db, _ := sql.Open("sqltest", "test")
	trans, _ := db.Begin()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "update trans",
			fields: fields{
				engine: "sqlite",
				Db:     db,
			},
			args: args{
				sqlString:   "update test set value = 'update' where id = 1",
				transaction: trans,
			},
			wantErr: false,
		},
		{
			name: "update",
			fields: fields{
				engine: "sqlite",
				Db:     db,
			},
			args: args{
				sqlString:   "update test set value = 'update' where id = 1",
				transaction: nil,
			},
			wantErr: false,
		},
		{
			name: "invalid_trans",
			fields: fields{
				engine: "sqlite",
				Db:     db,
			},
			args: args{
				sqlString:   "update test set value = 'update' where id = 1",
				transaction: "trans",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &SQLDriver{
				alias:   tt.fields.alias,
				connStr: tt.fields.connStr,
				engine:  tt.fields.engine,
				Db:      tt.fields.Db,
				closed:  tt.fields.closed,
				Config:  tt.fields.Config,
			}
			if err := ds.UpdateSQL(tt.args.sqlString, tt.args.transaction); (err != nil) != tt.wantErr {
				t.Errorf("SQLDriver.UpdateSQL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
