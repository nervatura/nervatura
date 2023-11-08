package nervatura

import (
	"database/sql"
	"errors"
	"reflect"
	"testing"

	nt "github.com/nervatura/nervatura/service/pkg/nervatura"
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

func TestSQLDriver_decodeEngine(t *testing.T) {
	type fields struct {
		alias   string
		connStr string
		engine  string
		db      *sql.DB
		Config  IM
	}
	type args struct {
		sqlStr string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "sqlite",
			fields: fields{
				engine: "sqlite",
			},
			args: args{
				sqlStr: "select * from table",
			},
			want: "select * from table",
		},
		{
			name: "mysql",
			fields: fields{
				engine: "mysql",
			},
			args: args{
				sqlStr: "select * from table",
			},
			want: "select * from table",
		},
		{
			name: "mssql",
			fields: fields{
				engine: "mssql",
			},
			args: args{
				sqlStr: "select * from table",
			},
			want: "select * from table",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &SQLDriver{
				alias:   tt.fields.alias,
				connStr: tt.fields.connStr,
				engine:  tt.fields.engine,
				db:      tt.fields.db,
				Config:  tt.fields.Config,
			}
			if got := ds.decodeEngine(tt.args.sqlStr); got != tt.want {
				t.Errorf("SQLDriver.decodeEngine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSQLDriver_getDataType(t *testing.T) {
	type fields struct {
		alias   string
		connStr string
		engine  string
		db      *sql.DB
		Config  IM
	}
	type args struct {
		dtype string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "found",
			fields: fields{
				engine: "postgres",
			},
			args: args{
				dtype: "float",
			},
			want: "FLOAT8",
		},
		{
			name: "base",
			fields: fields{
				engine: "postgres",
			},
			args: args{
				dtype: "text",
			},
			want: "TEXT",
		},
		{
			name: "not_found",
			fields: fields{
				engine: "postgres",
			},
			args: args{
				dtype: "int",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &SQLDriver{
				alias:   tt.fields.alias,
				connStr: tt.fields.connStr,
				engine:  tt.fields.engine,
				db:      tt.fields.db,
				Config:  tt.fields.Config,
			}
			if got := ds.getDataType(tt.args.dtype); got != tt.want {
				t.Errorf("SQLDriver.getDataType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setRefID(t *testing.T) {
	type args struct {
		refID     IM
		mname     string
		keyFields []string
		values    IM
		id        int64
	}
	tests := []struct {
		name string
		args args
		want IM
	}{
		{
			name: "case_1",
			args: args{
				refID:     IM{},
				mname:     "currency",
				keyFields: []string{"curr"},
				values: IM{
					"curr": "EUR",
				},
				id: int64(1),
			},
			want: IM{
				"currency": IM{
					"EUR": "1",
				},
			},
		},
		{
			name: "case_1_refID",
			args: args{
				refID:     IM{},
				mname:     "currency",
				keyFields: []string{"curr"},
				values:    IM{},
				id:        int64(1),
			},
			want: IM{
				"currency": IM{},
			},
		},
		{
			name: "case_2",
			args: args{
				refID:     IM{},
				mname:     "groups",
				keyFields: []string{"groupname", "groupvalue"},
				values: IM{
					"groupname":  "name",
					"groupvalue": "value",
				},
				id: int64(1),
			},
			want: IM{
				"groups": IM{
					"name": IM{
						"value": "1",
					},
				},
			},
		},
		{
			name: "case_2_refID",
			args: args{
				refID:     IM{},
				mname:     "groups",
				keyFields: []string{"groupname", "groupvalue"},
				values: IM{
					"groupname": "name",
				},
				id: int64(1),
			},
			want: IM{
				"groups": IM{},
			},
		},
		{
			name: "case_3",
			args: args{
				refID:     IM{},
				mname:     "fieldvalue",
				keyFields: []string{"fieldname", "refnumber", "rownumber"},
				values: IM{
					"fieldname": "name",
					"refnumber": "ref",
					"rownumber": "row",
				},
				id: int64(1),
			},
			want: IM{
				"fieldvalue": IM{
					"name": IM{
						"ref": IM{
							"row": "1",
						},
					},
				},
			},
		},
		{
			name: "case_3_refID",
			args: args{
				refID:     IM{},
				mname:     "fieldvalue",
				keyFields: []string{"fieldname", "refnumber", "rownumber"},
				values: IM{
					"fieldname": "name",
					"refnumber": "ref",
				},
				id: int64(1),
			},
			want: IM{
				"fieldvalue": IM{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setRefID(tt.args.refID, tt.args.mname, tt.args.keyFields, tt.args.values, tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setRefID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRefID(t *testing.T) {
	type args struct {
		refID IM
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "case_0",
			args: args{
				refID: IM{},
				value: []string{"customer"},
			},
			want: "0",
		},
		{
			name: "case_1",
			args: args{
				refID: IM{
					"customer": "1",
				},
				value: "customer",
			},
			want: "customer",
		},
		{
			name: "case_2",
			args: args{
				refID: IM{
					"customer": IM{
						"HOME": "1",
					},
				},
				value: []string{"customer", "HOME"},
			},
			want: "1",
		},
		{
			name: "case_3",
			args: args{
				refID: IM{
					"groups": IM{
						"custtype": IM{
							"own": "115",
						},
					},
				},
				value: []string{"groups", "custtype", "own"},
			},
			want: "115",
		},
		{
			name: "case_4",
			args: args{
				refID: IM{
					"key1": IM{
						"key2": IM{
							"key3": IM{
								"key4": "1",
							},
						},
					},
				},
				value: []string{"key1", "key2", "key3", "key4"},
			},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRefID(tt.args.refID, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRefID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSQLDriver_Properties(t *testing.T) {
	type fields struct {
		alias   string
		connStr string
		engine  string
		db      *sql.DB
		Config  IM
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
				db:      tt.fields.db,
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
		db      *sql.DB
		Config  IM
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
				db:     nil,
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
				db:      tt.fields.db,
				Config:  tt.fields.Config,
			}
			if got := ds.Connection(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SQLDriver.Connection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSQLDriver_CreateConnection(t *testing.T) {
	type fields struct {
		alias   string
		connStr string
		engine  string
		db      *sql.DB
		Config  IM
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
			name: "sqlite",
			fields: fields{
				db: func() *sql.DB {
					db, _ := sql.Open("sqlite", "file::memory:")
					return db
				}(),
			},
			args: args{
				alias:   "test",
				connStr: "sqlite://file::memory:",
			},
			wantErr: false,
		},
		{
			name:   "mysql",
			fields: fields{},
			args: args{
				alias:   "test",
				connStr: "mysql://test:test@tcp(localhost:1234)/nervatura",
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
				db:      tt.fields.db,
				Config:  tt.fields.Config,
			}
			if err := ds.CreateConnection(tt.args.alias, tt.args.connStr); (err != nil) != tt.wantErr {
				t.Errorf("SQLDriver.CreateConnection() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSQLDriver_getPrmString(t *testing.T) {
	type fields struct {
		alias   string
		connStr string
		engine  string
		db      *sql.DB
		Config  IM
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
				db:      tt.fields.db,
				Config:  tt.fields.Config,
			}
			if got := ds.getPrmString(tt.args.index); got != tt.want {
				t.Errorf("SQLDriver.getPrmString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSQLDriver_CheckHashtable(t *testing.T) {
	type fields struct {
		alias   string
		connStr string
		engine  string
		db      *sql.DB
		Config  IM
	}
	type args struct {
		hashtable string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "sqlite",
			fields: fields{
				engine: "sqlite",
				db: func() *sql.DB {
					db, _ := sql.Open("sqlite", "file::memory:?cache=shared")
					return db
				}(),
			},
			args: args{
				hashtable: "testhash",
			},
			wantErr: false,
		},
		{
			name: "create_error",
			fields: fields{
				engine: "mssql",
				db: func() *sql.DB {
					db, _ := sql.Open("sqlite", "file::memory:?cache=shared")
					return db
				}(),
			},
			args: args{
				hashtable: "testhash",
			},
			wantErr: true,
		},
		{
			name: "missing_driver",
			fields: fields{
				engine: "postgres",
			},
			args: args{
				hashtable: "testhash",
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
				db:      tt.fields.db,
				Config:  tt.fields.Config,
			}
			if err := ds.CheckHashtable(tt.args.hashtable); (err != nil) != tt.wantErr {
				t.Errorf("SQLDriver.CheckHashtable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSQLDriver_UpdateHashtable(t *testing.T) {
	type fields struct {
		alias   string
		connStr string
		engine  string
		db      *sql.DB
		Config  IM
	}
	type args struct {
		hashtable string
		refname   string
		value     string
	}
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
				db: func() *sql.DB {
					db, _ := sql.Open("sqlite", "file::memory:?cache=shared")
					return db
				}(),
			},
			args: args{
				hashtable: "testhash",
				refname:   "refname",
				value:     "value",
			},
			wantErr: false,
		},
		{
			name: "update",
			fields: fields{
				engine: "sqlite",
				db: func() *sql.DB {
					db, _ := sql.Open("sqlite", "file::memory:?cache=shared")
					_, _ = db.Exec("CREATE TABLE testhash ( refname CHAR(255), value TEXT);")
					_, _ = db.Exec("insert into testhash(value, refname) values('refname','value')")
					return db
				}(),
			},
			args: args{
				hashtable: "testhash",
				refname:   "refname",
				value:     "value",
			},
			wantErr: false,
		},
		{
			name:   "check_error",
			fields: fields{},
			args: args{
				hashtable: "testhash",
				refname:   "refname",
				value:     "value",
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
				db:      tt.fields.db,
				Config:  tt.fields.Config,
			}
			if err := ds.UpdateHashtable(tt.args.hashtable, tt.args.refname, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("SQLDriver.UpdateHashtable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSQLDriver_tableName(t *testing.T) {
	type fields struct {
		alias   string
		connStr string
		engine  string
		db      *sql.DB
		Config  IM
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "base",
			fields: fields{
				engine: "sqlite",
			},
			args: args{
				name: "table",
			},
			want: "table",
		},
		{
			name: "mysql",
			fields: fields{
				engine: "mysql",
			},
			args: args{
				name: "table",
			},
			want: "`table`",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &SQLDriver{
				alias:   tt.fields.alias,
				connStr: tt.fields.connStr,
				engine:  tt.fields.engine,
				db:      tt.fields.db,
				Config:  tt.fields.Config,
			}
			if got := ds.tableName(tt.args.name); got != tt.want {
				t.Errorf("SQLDriver.tableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSQLDriver_dropData(t *testing.T) {
	type fields struct {
		alias   string
		connStr string
		engine  string
		db      *sql.DB
		Config  IM
	}
	type args struct {
		logData []SM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "drop_all",
			fields: fields{
				engine: "sqlite",
				db: func() *sql.DB {
					db, _ := sql.Open("sqlite", "file::memory:")
					return db
				}(),
			},
			args: args{
				logData: []SM{},
			},
			wantErr: false,
		},
		{
			name: "Begin_error",
			fields: fields{
				engine: "sqlite",
				db: func() *sql.DB {
					db, _ := sql.Open("sqlite", "sqlite://file::memory:?cache=shared")
					return db
				}(),
			},
			args: args{
				logData: []SM{},
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
				db:      tt.fields.db,
				Config:  tt.fields.Config,
			}
			_, err := ds.dropData(tt.args.logData)
			if (err != nil) != tt.wantErr {
				t.Errorf("SQLDriver.dropData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSQLDriver_createTableFields(t *testing.T) {
	type fields struct {
		alias   string
		connStr string
		engine  string
		db      *sql.DB
		Config  IM
	}
	type args struct {
		sqlString string
		fieldname string
		indexName string
		field     nt.MF
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "NotNull",
			fields: fields{
				engine: "sqlite",
			},
			args: args{
				sqlString: "",
				fieldname: "id",
				indexName: "groups",
				field:     nt.MF{Type: "id", NotNull: true, Length: 3, Default: 0},
			},
			want: "id INTEGER PRIMARY KEY AUTOINCREMENT(3) NOT NULL DEFAULT 0, ",
		},
		{
			name: "References",
			fields: fields{
				engine: "sqlite",
			},
			args: args{
				sqlString: "",
				fieldname: "method",
				indexName: "ui_menu",
				field:     nt.MF{References: SL{"groups", "RESTRICT", "NO ACTION"}, NotNull: true, Requires: IM{"method": SL{}}},
			},
			want: "method INTEGER REFERENCES groups(id) ON DELETE RESTRICT, ",
		},
		{
			name: "References_mssql",
			fields: fields{
				engine: "mssql",
			},
			args: args{
				sqlString: "",
				fieldname: "method",
				indexName: "ui_menu",
				field:     nt.MF{References: SL{"groups", "RESTRICT", "NO ACTION"}, NotNull: true, Requires: IM{"method": SL{}}},
			},
			want: "method INT NULL, CONSTRAINT ui_menu__method__constraint FOREIGN KEY (method) REFERENCES groups(id) ON DELETE NO ACTION, ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &SQLDriver{
				alias:   tt.fields.alias,
				connStr: tt.fields.connStr,
				engine:  tt.fields.engine,
				db:      tt.fields.db,
				Config:  tt.fields.Config,
			}
			if got := ds.createTableFields(tt.args.sqlString, tt.args.fieldname, tt.args.indexName, tt.args.field); got != tt.want {
				t.Errorf("SQLDriver.createTableFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSQLDriver_createTable(t *testing.T) {
	type fields struct {
		alias   string
		connStr string
		engine  string
		db      *sql.DB
		Config  IM
	}
	type args struct {
		logData []SM
		trans   *sql.Tx
	}
	db1, _ := sql.Open("sqlite", "file::memory:")
	trans1, _ := db1.Begin()
	db2, _ := sql.Open("sqlite", "file::memory:")
	trans2, _ := db2.Begin()
	trans2.Rollback()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "create_ok",
			fields: fields{
				engine: "sqlite",
				db:     db1,
			},
			args: args{
				logData: []SM{},
				trans:   trans1,
			},
			wantErr: false,
		},
		{
			name: "create_error",
			fields: fields{
				engine: "sqlite",
				db:     db2,
			},
			args: args{
				logData: []SM{},
				trans:   trans2,
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
				db:      tt.fields.db,
				Config:  tt.fields.Config,
			}
			_, err := ds.createTable(tt.args.logData, tt.args.trans)
			if (err != nil) != tt.wantErr {
				t.Errorf("SQLDriver.createTable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSQLDriver_createIndex(t *testing.T) {
	type fields struct {
		alias   string
		connStr string
		engine  string
		db      *sql.DB
		Config  IM
	}
	type args struct {
		logData []SM
		trans   *sql.Tx
	}
	db1, _ := sql.Open("sqlite", "file::memory:")
	trans1, _ := db1.Begin()
	_, _ = (&SQLDriver{db: db1}).createTable([]SM{}, trans1)
	db2, _ := sql.Open("sqlite", "file::memory:")
	trans2, _ := db2.Begin()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "index_ok",
			fields: fields{
				engine: "sqlite",
				db:     db1,
			},
			args: args{
				logData: []SM{},
				trans:   trans1,
			},
			wantErr: false,
		},
		{
			name: "index_error",
			fields: fields{
				engine: "sqlite",
				db:     db2,
			},
			args: args{
				logData: []SM{},
				trans:   trans2,
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
				db:      tt.fields.db,
				Config:  tt.fields.Config,
			}
			_, err := ds.createIndex(tt.args.logData, tt.args.trans)
			if (err != nil) != tt.wantErr {
				t.Errorf("SQLDriver.createIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSQLDriver_insertData(t *testing.T) {
	type fields struct {
		alias   string
		connStr string
		engine  string
		db      *sql.DB
		Config  IM
	}
	type args struct {
		logData []SM
		trans   *sql.Tx
	}
	db1, _ := sql.Open("sqlite", "file::memory:")
	trans1, _ := db1.Begin()
	_, _ = (&SQLDriver{db: db1}).createTable([]SM{}, trans1)
	db2, _ := sql.Open("sqlite", "file::memory:")
	trans2, _ := db2.Begin()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "insert_ok",
			fields: fields{
				engine: "sqlite",
				db:     db1,
			},
			args: args{
				logData: []SM{},
				trans:   trans1,
			},
			wantErr: false,
		},
		{
			name: "insert_error",
			fields: fields{
				engine: "mssql",
				db:     db2,
			},
			args: args{
				logData: []SM{},
				trans:   trans2,
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
				db:      tt.fields.db,
				Config:  tt.fields.Config,
			}
			_, err := ds.insertData(tt.args.logData, tt.args.trans)
			if (err != nil) != tt.wantErr {
				t.Errorf("SQLDriver.insertData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_rollBackTrans(t *testing.T) {
	type args struct {
		trans *sql.Tx
		err   error
	}
	db, _ := sql.Open("sqlite", "file::memory:")
	trans, _ := db.Begin()
	trans.Rollback()
	tests := []struct {
		name string
		args args
	}{
		{
			name: "rollback_err",
			args: args{
				trans: trans,
				err:   errors.New("error"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rollBackTrans(tt.args.trans, tt.args.err)
		})
	}
}

func TestSQLDriver_CreateDatabase(t *testing.T) {
	type fields struct {
		alias   string
		connStr string
		engine  string
		db      *sql.DB
		Config  IM
	}
	type args struct {
		logData []SM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "create_ok",
			fields: fields{
				engine: "sqlite",
				db: func() *sql.DB {
					db, _ := sql.Open("sqlite", "file::memory:")
					return db
				}(),
			},
			args: args{
				logData: []SM{},
			},
			wantErr: false,
		},
		{
			name: "postgres",
			fields: fields{
				engine: "postgres",
				db: func() *sql.DB {
					db, _ := sql.Open("sqlite", "file::memory:")
					return db
				}(),
			},
			args: args{
				logData: []SM{},
			},
			wantErr: true,
		},
		{
			name: "db_nil",
			fields: fields{
				engine: "sqlite",
			},
			args: args{
				logData: []SM{},
			},
			wantErr: true,
		},
		{
			name: "drop_err",
			fields: fields{
				engine: "sqlite",
				db: func() *sql.DB {
					db, _ := sql.Open("sqlite", "file::memory:")
					db.Close()
					return db
				}(),
			},
			args: args{
				logData: []SM{},
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
				db:      tt.fields.db,
				Config:  tt.fields.Config,
			}
			_, err := ds.CreateDatabase(tt.args.logData)
			if (err != nil) != tt.wantErr {
				t.Errorf("SQLDriver.CreateDatabase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSQLDriver_getFilterString(t *testing.T) {
	type fields struct {
		alias   string
		connStr string
		engine  string
		db      *sql.DB
		Config  IM
	}
	type args struct {
		filter    nt.Filter
		start     bool
		sqlString string
		params    []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  []interface{}
	}{
		{
			name:   "start_eq",
			fields: fields{},
			args: args{
				start:     true,
				sqlString: "",
				params:    []interface{}{},
				filter: nt.Filter{
					Field: "fieldname",
					Comp:  "==",
					Value: "value",
				},
			},
			want:  "(fieldname=?)",
			want1: []interface{}{"value"},
		},
		{
			name:   "and_like",
			fields: fields{},
			args: args{
				sqlString: "",
				params:    []interface{}{},
				filter: nt.Filter{
					Or:    false,
					Field: "fieldname",
					Comp:  "like",
					Value: "value",
				},
			},
			want:  " and (fieldname like ?)",
			want1: []interface{}{"value"},
		},
		{
			name:   "or_is",
			fields: fields{},
			args: args{
				sqlString: "",
				params:    []interface{}{},
				filter: nt.Filter{
					Or:    true,
					Field: "fieldname",
					Comp:  "is",
					Value: "null",
				},
			},
			want:  " or fieldname is null",
			want1: []interface{}{},
		},
		{
			name:   "and_in",
			fields: fields{},
			args: args{
				sqlString: "",
				params:    []interface{}{},
				filter: nt.Filter{
					Or:    false,
					Field: "fieldname",
					Comp:  "in",
					Value: "value1,value2,value3",
				},
			},
			want:  " and (fieldname in(?,?,?))",
			want1: []interface{}{"value1", "value2", "value3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &SQLDriver{
				alias:   tt.fields.alias,
				connStr: tt.fields.connStr,
				engine:  tt.fields.engine,
				db:      tt.fields.db,
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
		db      *sql.DB
		Config  IM
	}
	type args struct {
		queries []nt.Query
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  []interface{}
	}{
		{
			name:   "queries",
			fields: fields{},
			args: args{
				queries: []nt.Query{
					{
						Fields: []string{"field1", "field2"},
						From:   "table",
						Filters: []nt.Filter{
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
						Filters: []nt.Filter{
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
			},
			want:  "select field1,field2 from table where (field1=?) and (field2 in(?,?,?)) order by field1 union select field1,field2 from table where (field1=?) and (field2 in(?,?,?)) order by field1",
			want1: []interface{}{"value", "1", "2", "3", "value", "1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &SQLDriver{
				alias:   tt.fields.alias,
				connStr: tt.fields.connStr,
				engine:  tt.fields.engine,
				db:      tt.fields.db,
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
		db      *sql.DB
		Config  IM
	}
	type args struct {
		queries []nt.Query
		trans   interface{}
	}
	db, _ := sql.Open("sqlite", "file::memory:")
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
				engine: "sqlite",
				db:     db,
			},
			args: args{
				queries: []nt.Query{
					{
						Fields: []string{"field1", "field2"},
						From:   "table",
						Filters: []nt.Filter{
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
				db:      tt.fields.db,
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
		db      *sql.DB
		Config  IM
	}
	type args struct {
		sqlString string
		params    []interface{}
		trans     interface{}
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
				engine: "sqlite",
				db: func() *sql.DB {
					db, _ := sql.Open("sqlite", "file::memory:")
					_, _ = db.Exec(`CREATE TABLE test(id INTEGER PRIMARY KEY AUTOINCREMENT,
						int_field INTEGER DEFAULT 0,
						float_field DOUBLE DEFAULT 0,bool_field BOOLEAN DEFAULT FALSE,
						datetime_field TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
						date_field DATE,
						string_field TEXT DEFAULT '')`)
					_, _ = db.Exec("INSERT INTO test(float_field,string_field,date_field) VALUES(1.1,'value','2021-12-24')")
					_, _ = db.Exec("INSERT INTO test(float_field,string_field,date_field,bool_field,int_field) VALUES(null,null,null,null,null)")
					return db
				}(),
			},
			args: args{
				sqlString: "select *, 3 as count from test",
				params:    []interface{}{},
			},
			wantErr: false,
		},
		{
			name: "invalid_trans",
			args: args{
				sqlString: "",
				params:    []interface{}{},
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
				db:      tt.fields.db,
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
		db      *sql.DB
		Config  IM
	}
	type args struct {
		model  string
		result sql.Result
		trans  interface{}
	}
	db, _ := sql.Open("sqlite", "file::memory:")
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
				db:     db,
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
				db:     db,
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
				db:     db,
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
				db:      tt.fields.db,
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
		db      *sql.DB
		Config  IM
	}
	type args struct {
		options nt.Update
	}
	db, _ := sql.Open("sqlite", "file::memory:")
	_, _ = db.Exec(`CREATE TABLE test(id INTEGER PRIMARY KEY,value TEXT)`)
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
				db:     db,
			},
			args: args{
				options: nt.Update{
					Model: "test",
					IDKey: int64(0),
					Values: IM{
						"value": "insert",
					},
					Trans: trans,
				},
			},
			wantErr: false,
		},
		{
			name: "update",
			fields: fields{
				engine: "sqlite",
				db:     db,
			},
			args: args{
				options: nt.Update{
					Model: "test",
					IDKey: int64(1),
					Values: IM{
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
				db:     db,
			},
			args: args{
				options: nt.Update{
					Model:  "test2",
					IDKey:  int64(1),
					Values: IM{},
				},
			},
			wantErr: true,
		},
		{
			name: "invalid_trans",
			fields: fields{
				engine: "sqlite",
				db:     db,
			},
			args: args{
				options: nt.Update{
					Model:  "test2",
					IDKey:  int64(1),
					Values: IM{},
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
				db:      tt.fields.db,
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
		db      *sql.DB
		Config  IM
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "begin",
			fields: fields{
				db: func() *sql.DB {
					db, _ := sql.Open("sqlite", "file::memory:")
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
				db:      tt.fields.db,
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
		db      *sql.DB
		Config  IM
	}
	type args struct {
		trans interface{}
	}
	db, _ := sql.Open("sqlite", "file::memory:")
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
				db: db,
			},
			args: args{
				trans: trans,
			},
			wantErr: false,
		},
		{
			name: "invalid_trans",
			fields: fields{
				db: db,
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
				db:      tt.fields.db,
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
		db      *sql.DB
		Config  IM
	}
	type args struct {
		trans interface{}
	}
	db, _ := sql.Open("sqlite", "file::memory:")
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
				db: db,
			},
			args: args{
				trans: trans,
			},
			wantErr: false,
		},
		{
			name: "invalid_trans",
			fields: fields{
				db: db,
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
				db:      tt.fields.db,
				Config:  tt.fields.Config,
			}
			if err := ds.RollbackTransaction(tt.args.trans); (err != nil) != tt.wantErr {
				t.Errorf("SQLDriver.RollbackTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSQLDriver_getID2Refnumber(t *testing.T) {
	type fields struct {
		alias   string
		connStr string
		engine  string
		db      *sql.DB
		Config  IM
	}
	type args struct {
		options IM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "address_refId",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "address", "refId": 1,
					"useDeleted": true, "retfield": ""},
			},
			wantErr: false,
		},
		{
			name:   "address",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "address",
					"useDeleted": false, "retfield": ""},
			},
			wantErr: false,
		},
		{
			name:   "fieldvalue_refId",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "fieldvalue", "refId": 1,
					"useDeleted": false, "retfield": ""},
			},
			wantErr: false,
		},
		{
			name:   "fieldvalue",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "fieldvalue",
					"useDeleted": false, "retfield": ""},
			},
			wantErr: false,
		},
		{
			name:   "item_refId",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "item", "refId": 1,
					"useDeleted": false, "retfield": ""},
			},
			wantErr: false,
		},
		{
			name:   "item",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "item",
					"useDeleted": false, "retfield": ""},
			},
			wantErr: false,
		},
		{
			name:   "price",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "price",
					"useDeleted": false, "retfield": ""},
			},
			wantErr: false,
		},
		{
			name:   "link",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "link",
					"useDeleted": false, "retfield": ""},
			},
			wantErr: false,
		},
		{
			name:   "rate",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "rate",
					"useDeleted": false, "retfield": ""},
			},
			wantErr: false,
		},
		{
			name:   "log",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "log",
					"useDeleted": false, "retfield": ""},
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
				db:      tt.fields.db,
				Config:  tt.fields.Config,
			}
			_, _, err := ds.getID2Refnumber(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("SQLDriver.getID2Refnumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSQLDriver_getRefnumber2ID(t *testing.T) {
	type fields struct {
		alias   string
		connStr string
		engine  string
		db      *sql.DB
		Config  IM
	}
	type args struct {
		options IM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "barcode",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "barcode", "refnumber": "refnumber",
					"useDeleted": false, "extraInfo": false},
			},
			wantErr: false,
		},
		{
			name:   "customer",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "customer", "refnumber": "refnumber",
					"useDeleted": false, "extraInfo": false},
			},
			wantErr: false,
		},
		{
			name:   "event",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "event", "refnumber": "refnumber",
					"useDeleted": false, "extraInfo": false},
			},
			wantErr: false,
		},
		{
			name:   "groups",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "groups", "refnumber": "refnumber",
					"useDeleted": false, "extraInfo": false},
			},
			wantErr: false,
		},
		{
			name:   "deffield",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "deffield", "refnumber": "refnumber",
					"useDeleted": false, "extraInfo": false},
			},
			wantErr: false,
		},
		{
			name:   "fieldvalue",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "fieldvalue", "refnumber": "refnumber",
					"useDeleted": false, "extraInfo": false},
			},
			wantErr: false,
		},
		{
			name:   "item",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "item", "refnumber": "refnumber",
					"useDeleted": false, "extraInfo": false},
			},
			wantErr: false,
		},
		{
			name:   "payment",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "payment", "refnumber": "refnumber",
					"useDeleted": false, "extraInfo": false},
			},
			wantErr: false,
		},
		{
			name:   "movement",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "movement", "refnumber": "refnumber",
					"useDeleted": false, "extraInfo": false},
			},
			wantErr: false,
		},
		{
			name:   "price",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "price", "refnumber": "refnumber",
					"useDeleted": false, "extraInfo": false},
			},
			wantErr: false,
		},
		{
			name:   "product",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "product", "refnumber": "refnumber",
					"useDeleted": false, "extraInfo": false},
			},
			wantErr: false,
		},
		{
			name:   "place",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "place", "refnumber": "refnumber",
					"useDeleted": false, "extraInfo": false},
			},
			wantErr: false,
		},
		{
			name:   "tax",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "tax", "refnumber": "refnumber",
					"useDeleted": false, "extraInfo": false},
			},
			wantErr: false,
		},
		{
			name:   "trans",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "trans", "refnumber": "refnumber",
					"useDeleted": false, "extraInfo": false},
			},
			wantErr: false,
		},
		{
			name:   "setting",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "setting", "refnumber": "refnumber",
					"useDeleted": false, "extraInfo": false},
			},
			wantErr: false,
		},
		{
			name:   "link",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "link", "refnumber": "refnumber",
					"useDeleted": false, "extraInfo": false},
			},
			wantErr: false,
		},
		{
			name:   "rate",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "rate", "refnumber": "refnumber",
					"useDeleted": false, "extraInfo": false},
			},
			wantErr: false,
		},
		{
			name:   "rate_planumber",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "rate", "refnumber": "refnumber",
					"planumber": "planumber"},
			},
			wantErr: false,
		},
		{
			name:   "log",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "log", "refnumber": "refnumber",
					"useDeleted": false, "extraInfo": false},
			},
			wantErr: false,
		},
		{
			name:   "ui_audit",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "ui_audit", "refnumber": "refnumber",
					"useDeleted": false, "extraInfo": false},
			},
			wantErr: false,
		},
		{
			name:   "ui_audit_transType_trans",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "ui_audit", "refnumber": "refnumber",
					"transType": "invoice", "refType": "trans"},
			},
			wantErr: false,
		},
		{
			name:   "ui_audit_transType",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "ui_audit", "refnumber": "refnumber",
					"transType": "invoice", "refType": "refType"},
			},
			wantErr: false,
		},
		{
			name:   "ui_menufields",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "ui_menufields", "refnumber": "refnumber",
					"useDeleted": false, "extraInfo": false},
			},
			wantErr: false,
		},
		{
			name:   "employee",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "employee", "refnumber": "refnumber",
					"useDeleted": false, "extraInfo": false},
			},
			wantErr: false,
		},
		{
			name:   "address",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "address", "refnumber": "refnumber",
					"useDeleted": false, "extraInfo": false},
			},
			wantErr: false,
		},
		{
			name:   "missing",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "missing", "refnumber": "refnumber",
					"useDeleted": false, "extraInfo": false},
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
				db:      tt.fields.db,
				Config:  tt.fields.Config,
			}
			_, _, err := ds.getRefnumber2ID(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("SQLDriver.getRefnumber2ID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSQLDriver_getUpdateDeffields(t *testing.T) {
	type fields struct {
		alias   string
		connStr string
		engine  string
		db      *sql.DB
		Config  IM
	}
	type args struct {
		options IM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "fieldname",
			fields: fields{},
			args: args{
				options: IM{"fieldname": "fieldname"},
			},
			wantErr: false,
		},
		{
			name:   "nervatype_ref_id",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "nervatype", "ref_id": ""},
			},
			wantErr: false,
		},
		{
			name:   "nervatype",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "nervatype"},
			},
			wantErr: false,
		},
		{
			name:   "missing",
			fields: fields{},
			args: args{
				options: IM{},
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
				db:      tt.fields.db,
				Config:  tt.fields.Config,
			}
			_, _, err := ds.getUpdateDeffields(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("SQLDriver.getUpdateDeffields() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSQLDriver_getIntegrity(t *testing.T) {
	type fields struct {
		alias   string
		connStr string
		engine  string
		db      *sql.DB
		Config  IM
	}
	type args struct {
		options IM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "currency",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "currency", "ref_id": 1},
			},
			wantErr: false,
		},
		{
			name:   "customer",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "customer", "ref_id": 1},
			},
			wantErr: false,
		},
		{
			name:   "deffield",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "deffield", "ref_id": 1},
			},
			wantErr: false,
		},
		{
			name:   "employee",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "employee", "ref_id": 1},
			},
			wantErr: false,
		},
		{
			name:   "groups",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "groups", "ref_id": 1},
			},
			wantErr: false,
		},
		{
			name:   "place",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "place", "ref_id": 1},
			},
			wantErr: false,
		},
		{
			name:   "product",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "product", "ref_id": 1},
			},
			wantErr: false,
		},
		{
			name:   "project",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "project", "ref_id": 1},
			},
			wantErr: false,
		},
		{
			name:   "tax",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "tax", "ref_id": 1},
			},
			wantErr: false,
		},
		{
			name:   "tool",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "tool", "ref_id": 1},
			},
			wantErr: false,
		},
		{
			name:   "trans",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "trans", "ref_id": 1},
			},
			wantErr: false,
		},
		{
			name:   "missing",
			fields: fields{},
			args: args{
				options: IM{"nervatype": "missing", "ref_id": 1},
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
				db:      tt.fields.db,
				Config:  tt.fields.Config,
			}
			_, _, err := ds.getIntegrity(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("SQLDriver.getIntegrity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSQLDriver_QueryKey(t *testing.T) {
	type fields struct {
		alias   string
		connStr string
		engine  string
		db      *sql.DB
		Config  IM
	}
	type args struct {
		options IM
		trans   interface{}
	}

	db, _ := sql.Open("sqlite", "file::memory:")
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "user",
			fields: fields{
				engine: "sqlite",
				db:     db,
			},
			args: args{
				options: IM{"qkey": "user", "username": "admin"},
			},
			wantErr: true,
		},
		{
			name: "metadata",
			fields: fields{
				engine: "sqlite",
				db:     db,
			},
			args: args{
				options: IM{"qkey": "metadata", "nervatype": "customer", "ids": "1,2"},
			},
			wantErr: true,
		},
		{
			name: "post_transtype",
			fields: fields{
				engine: "sqlite",
				db:     db,
			},
			args: args{
				options: IM{"qkey": "post_transtype"},
			},
			wantErr: true,
		},
		{
			name: "default_report",
			fields: fields{
				engine: "sqlite",
				db:     db,
			},
			args: args{
				options: IM{"qkey": "default_report", "nervatype": "trans"},
			},
			wantErr: true,
		},
		{
			name: "listprice",
			fields: fields{
				engine: "sqlite",
				db:     db,
			},
			args: args{
				options: IM{
					"qkey":        "listprice",
					"curr":        "EUR",
					"product_id":  123,
					"vendorprice": false,
					"posdate":     "2021-12-24",
					"qty":         1,
					"customer_id": 1},
			},
			wantErr: true,
		},
		{
			name: "grouprice",
			fields: fields{
				engine: "sqlite",
				db:     db,
			},
			args: args{
				options: IM{
					"qkey":        "grouprice",
					"curr":        "EUR",
					"product_id":  123,
					"vendorprice": false,
					"posdate":     "2021-12-24",
					"qty":         1,
					"customer_id": 1},
			},
			wantErr: true,
		},
		{
			name: "custprice",
			fields: fields{
				engine: "sqlite",
				db:     db,
			},
			args: args{
				options: IM{
					"qkey":        "custprice",
					"curr":        "EUR",
					"product_id":  123,
					"vendorprice": false,
					"posdate":     "2021-12-24",
					"qty":         1,
					"customer_id": 1},
			},
			wantErr: true,
		},
		{
			name: "data_audit",
			fields: fields{
				engine: "sqlite",
				db:     db,
			},
			args: args{
				options: IM{"qkey": "data_audit", "id": 1},
			},
			wantErr: true,
		},
		{
			name: "object_audit",
			fields: fields{
				engine: "sqlite",
				db:     db,
			},
			args: args{
				options: IM{"qkey": "object_audit", "usergroup": 1, "subtypeIn": "1,2,3"},
			},
			wantErr: true,
		},
		{
			name: "delete_deffields",
			fields: fields{
				engine: "sqlite",
				db:     db,
			},
			args: args{
				options: IM{"qkey": "delete_deffields", "nervatype": "customer", "ref_id": 1},
			},
			wantErr: true,
		},
		{
			name: "id->refnumber",
			fields: fields{
				engine: "sqlite",
				db:     db,
			},
			args: args{
				options: IM{"qkey": "id->refnumber", "nervatype": "", "id": "",
					"useDeleted": false, "retfield": ""},
			},
			wantErr: true,
		},
		{
			name: "id->refnumber",
			fields: fields{
				engine: "sqlite",
				db:     db,
			},
			args: args{
				options: IM{"qkey": "refnumber->id", "nervatype": "", "refnumber": "",
					"useDeleted": false, "extraInfo": false},
			},
			wantErr: true,
		},
		{
			name: "integrity",
			fields: fields{
				engine: "sqlite",
				db:     db,
			},
			args: args{
				options: IM{"qkey": "integrity", "nervatype": "customer", "ref_id": 1},
			},
			wantErr: true,
		},
		{
			name: "update_deffields",
			fields: fields{
				engine: "sqlite",
				db:     db,
			},
			args: args{
				options: IM{"qkey": "update_deffields", "fieldname": "fieldname"},
			},
			wantErr: true,
		},
		{
			name: "missing_fieldname",
			fields: fields{
				engine: "sqlite",
				db:     db,
			},
			args: args{
				options: IM{"qkey": "missing"},
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
				db:      tt.fields.db,
				Config:  tt.fields.Config,
			}
			_, err := ds.QueryKey(tt.args.options, tt.args.trans)
			if (err != nil) != tt.wantErr {
				t.Errorf("SQLDriver.QueryKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSQLDriver_CloseConnection(t *testing.T) {
	type fields struct {
		alias   string
		connStr string
		engine  string
		closed  bool
		db      *sql.DB
		Config  IM
	}
	db, _ := sql.Open("sqlite", "file::memory:")
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "nil",
			fields:  fields{},
			wantErr: false,
		},
		{
			name: "close",
			fields: fields{
				alias:   "test",
				connStr: "test",
				closed:  false,
				db:      db,
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
				closed:  tt.fields.closed,
				db:      tt.fields.db,
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
		closed  bool
		db      *sql.DB
		Config  IM
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "check",
			fields: fields{
				alias:   "test",
				connStr: "test",
				closed:  true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &SQLDriver{
				alias:   tt.fields.alias,
				connStr: tt.fields.connStr,
				engine:  tt.fields.engine,
				closed:  tt.fields.closed,
				db:      tt.fields.db,
				Config:  tt.fields.Config,
			}
			ds.checkConnection()
		})
	}
}
