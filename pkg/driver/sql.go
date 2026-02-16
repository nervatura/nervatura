package database

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
)

// Drivers -> Nervatura linked database drivers
var Drivers []string

func init() {
	registerDriver("")
}

func registerDriver(name string) {
	if name != "" {
		Drivers = append(Drivers, name)
	}
}

// SQLDriver a go database/sql DataDriver
type SQLDriver struct {
	alias   string
	connStr string
	engine  string
	Db      *sql.DB
	closed  bool
	Config  cu.IM
}

// Properties - DataDriver features
func (ds *SQLDriver) Properties() struct{ SQL, Transaction bool } {
	return struct{ SQL, Transaction bool }{SQL: true, Transaction: true}
}

// Connection - returns the database connection
func (ds *SQLDriver) Connection() struct {
	Alias     string
	Connected bool
	Engine    string
} {
	return struct {
		Alias     string
		Connected bool
		Engine    string
	}{
		Alias:     ds.alias,
		Connected: (ds.Db != nil),
		Engine:    ds.engine,
	}
}

// CreateConnection create a new database connection
func (ds *SQLDriver) CreateConnection(alias, connStr string) error {
	if ds.Db != nil {
		if err := ds.Db.Close(); err != nil {
			return err
		}
	}
	engine := strings.Split(connStr, "://")[0]
	conn := connStr
	conn = strings.ReplaceAll(connStr, "sqlite://", "")
	if engine == "mysql" {
		conn = strings.TrimPrefix(connStr, engine+"://")
		ccs := "?"
		if strings.Contains(conn, "?") {
			ccs = "&"
		}
		if !strings.Contains(conn, "multiStatements") {
			conn += ccs + "multiStatements=true"
		}
	}
	if engine == "mssql" {
		conn = strings.ReplaceAll(connStr, "mssql", "sqlserver")
	}
	db, err := sql.Open(engine, conn)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	db.SetMaxOpenConns(int(cu.ToInteger(ds.Config["SQL_MAX_OPEN_CONNS"], 10)))
	db.SetMaxIdleConns(int(cu.ToInteger(ds.Config["SQL_MAX_IDLE_CONNS"], 3)))
	db.SetConnMaxLifetime(time.Minute * time.Duration(int(cu.ToInteger(ds.Config["SQL_CONN_MAX_LIFETIME"], 15))))
	if slices.Contains([]string{"sqlite", "sqltest"}, engine) {
		db.Exec("PRAGMA foreign_keys = ON;")
	}
	ds.Db = db
	ds.alias = alias
	ds.engine = engine
	ds.connStr = connStr
	ds.closed = false
	return nil
}

func (ds *SQLDriver) CloseConnection() error {
	if ds.Db != nil && !ds.closed && !strings.Contains(ds.connStr, "memory") {
		ds.closed = true
		return ds.Db.Close()
	}
	return nil
}

func (ds *SQLDriver) checkConnection() {
	reconnect := (ds.closed && ds.alias != "" && ds.connStr != "" && !strings.Contains(ds.connStr, "memory"))
	if reconnect {
		ds.CreateConnection(ds.alias, ds.connStr)
	}
}

// getPrmString - get database parameter string
func (ds *SQLDriver) getPrmString(index int) string {
	if ds.engine == "postgres" {
		return "$" + strconv.Itoa(index)
	}
	return "?"
}

func (ds *SQLDriver) getFilterString(filter md.Filter, start bool, sqlString string, params []any) (string, []any) {
	if start {
		sqlString += "("
	} else if !filter.Or {
		sqlString += " and ("
	} else {
		sqlString += " or "
	}
	sqlString += filter.Field
	switch filter.Comp {
	case "==":
		params = append(params, filter.Value)
		sqlString += "=" + ds.getPrmString(len(params))
	case "like", "!=", "<", "<=", ">", ">=":
		params = append(params, filter.Value)
		sqlString += " " + filter.Comp + " " + ds.getPrmString(len(params))
	case "is":
		sqlString += " " + filter.Comp + " " + cu.ToString(filter.Value, "")
	case "in":
		if filterValue, valid := filter.Value.(string); valid {
			values := strings.Split(filterValue, ",")
			prmStr := make([]string, 0)
			for _, value := range values {
				params = append(params, value)
				prmStr = append(prmStr, ds.getPrmString(len(params)))
			}
			sqlString += " in(" + strings.Join(prmStr, ",") + ")"
		}
		if filterValue, valid := filter.Value.(md.Query); valid {
			queryString, queryParams := ds.decodeSQL([]md.Query{filterValue})
			sqlString += " in(" + queryString + ")"
			params = append(params, queryParams...)
		}
	}
	if !filter.Or {
		sqlString += ")"
	}
	return sqlString, params
}

func (ds *SQLDriver) decodeSQL(queries []md.Query) (string, []any) {
	sqlString := ""
	params := make([]any, 0)
	for qi := 0; qi < len(queries); qi++ {
		query := queries[qi]
		if qi > 0 {
			sqlString += " union select "
		} else {
			sqlString += "select "
		}
		sqlString += strings.Join(query.Fields, ",") + " from " + query.From
		if len(query.Filters) > 0 || query.Filter != "" {
			sqlString += " where "
		}
		for wi := 0; wi < len(query.Filters); wi++ {
			sqlString, params = ds.getFilterString(query.Filters[wi], (wi == 0), sqlString, params)
		}
		sqlString += query.Filter
		order := strings.Join(query.OrderBy, ",")
		if order != "" {
			sqlString += " order by " + order
		}
		if query.Limit > 0 {
			if ds.engine == "mssql" {
				sqlString = strings.ReplaceAll(sqlString, "select ", "select top("+cu.ToString(query.Limit, "")+") ")
			} else {
				sqlString += " limit " + cu.ToString(query.Limit, "")
			}
		}
		if query.Offset > 0 {
			sqlString += " offset " + cu.ToString(query.Offset, "")
		}
	}
	return strings.Trim(sqlString, " "), params
}

// Query is a basic nosql friendly queries the database
func (ds *SQLDriver) Query(queries []md.Query, trans any) ([]cu.IM, error) {
	sqlString, params := ds.decodeSQL(queries)
	return ds.QuerySQL(sqlString, params, trans)
}

func initQueryCols(engine string, cols []*sql.ColumnType) ([]any, []string, []string) {
	values := make([]any, len(cols))
	fields := make([]string, len(cols))
	dbtypes := make([]string, len(cols))
	for i := range cols {
		fields[i] = cols[i].Name()
		dbType := cols[i].DatabaseTypeName()
		if dbType == "" && cols[i].Name() == "count" {
			dbType = "INTEGER"
		}
		dbtypes[i] = dbType
		//println(dbType)
		values[i] = new(sql.NullString)
		switch dbType {
		case "BOOL", "BOOLEAN", "BIT", "TINYINT":
			values[i] = new(sql.NullBool)
		case "INTEGER", "SERIAL", "INT", "INT4", "INT8":
			values[i] = new(sql.NullInt64)
		case "DOUBLE", "FLOAT8", "DECIMAL(19,4)", "DECIMAL", "NUMERIC":
			values[i] = new(sql.NullFloat64)
		case "DATETIME", "TIMESTAMP", "DATE":
			if slices.Contains([]string{"postgres", "sqlite", "sqltest"}, engine) {
				values[i] = new(sql.NullTime)
			}
		}
	}
	return values, fields, dbtypes
}

func isJSON(dbtype, fname, value string) bool {
	return dbtype == "JSONB" || dbtype == "JSON" ||
		((dbtype == "" || dbtype == "NVARCHAR") && json.Valid([]byte(value)) &&
			!strings.Contains(fname, "_object"))
}

func getQueryRowValue(value any, dbtype, fname string) any {
	var vresult any = nil
	switch v := value.(type) {
	case *sql.NullBool:
		if v.Valid {
			return v.Bool
		}

	case *sql.NullInt64:
		if v.Valid {
			return v.Int64
		}

	case *sql.NullFloat64:
		if v.Valid {
			return v.Float64
		}

	case *sql.NullTime:
		if v.Valid {
			if dbtype == "DATE" {
				return v.Time.Format(time.DateOnly)
			}
			return v.Time
		}

	case *sql.NullString:
		if v.Valid {
			if isJSON(dbtype, fname, v.String) {
				return md.JSONBMap(fname, v.String)
			}
			return v.String
		}
	}
	return vresult
}

func (ds *SQLDriver) checkParams(sqlString string, params []any) (string, []any, error) {
	if len(params) > 0 && ds.engine == "postgres" {
		prmCount := 0
		for {
			regex := regexp.MustCompile(`\?`)
			index := regex.FindStringIndex(sqlString)
			if index == nil {
				break
			}
			if prmCount >= len(params) {
				return sqlString, params, errors.New("too many parameters in sqlString")
			}
			prmCount++
			sqlString = sqlString[:index[0]] + "$" + strconv.Itoa(prmCount) + sqlString[index[1]:]
		}
	}
	return sqlString, params, nil
}

// QuerySQL executes a SQL query
func (ds *SQLDriver) QuerySQL(sqlString string, params []any, trans any) (result []cu.IM, err error) {
	if sqlString, params, err = ds.checkParams(sqlString, params); err != nil {
		return result, err
	}
	result = make([]cu.IM, 0)
	var rows *sql.Rows
	if trans != nil {
		switch trans.(type) {
		case *sql.Tx:
		default:
			return result, errors.New("invalid transaction")
		}
	}

	ds.checkConnection()
	if trans == nil {
		defer ds.CloseConnection()
	}
	//println(sqlString)
	if trans != nil {
		rows, err = trans.(*sql.Tx).Query(sqlString, params...)
	} else {
		rows, err = ds.Db.Query(sqlString, params...)
	}
	if err != nil {
		return result, err
	}
	defer rows.Close()

	if cols, err := rows.ColumnTypes(); err == nil {
		values, fields, dbtypes := initQueryCols(ds.engine, cols)

		for rows.Next() {
			if err = rows.Scan(values...); err == nil {
				row := make(cu.IM)
				for index, value := range values {
					row[fields[index]] = getQueryRowValue(value, dbtypes[index], fields[index])
				}
				result = append(result, row)
			}
		}
	}
	return result, err
}

func (ds *SQLDriver) lastInsertID(model string, result sql.Result, trans any) (int64, error) {
	var sqlString string
	resid, err := result.LastInsertId()
	if err != nil {
		switch ds.engine {
		case "postgres":
			sqlString = fmt.Sprintf("select currval('%s_id_seq') as id", model)
		case "mssql":
			sqlString = fmt.Sprintf("select ident_current('%s') as id", model)
		default:
			return -1, err
		}
		if trans != nil {
			err = trans.(*sql.Tx).QueryRow(sqlString).Scan(&resid)
		} else {
			err = ds.Db.QueryRow(sqlString).Scan(&resid)
		}
		if err != nil {
			return -1, err
		}
	}
	return resid, nil
}

// Update is a basic nosql friendly update/insert/delete and returns the update/insert id
func (ds *SQLDriver) Update(options md.Update) (int64, error) {
	sqlString := ""
	id := options.IDKey
	params := make([]any, 0)
	fields := make([]string, 0)
	values := make([]string, 0)
	sets := make([]string, 0)
	for fieldname, value := range options.Values {
		params = append(params, value)
		fields = append(fields, fieldname)
		values = append(values, ds.getPrmString(len(params)))
		sets = append(sets, fmt.Sprintf("%s=%s", fieldname, ds.getPrmString(len(params))))
	}
	if id <= 0 {
		sqlString += fmt.Sprintf(
			"insert into %s (%s) values (%s)",
			options.Model, strings.Join(fields, ","), strings.Join(values, ","))
	} else if len(options.Values) == 0 {
		params = append(params, id)
		sqlString += fmt.Sprintf(
			"delete from %s where id=%s", options.Model, ds.getPrmString(len(params)))
	} else {
		params = append(params, id)
		sqlString += fmt.Sprintf(
			"update %s set %s where id=%s", options.Model, strings.Join(sets, ","), ds.getPrmString(len(params)))
	}
	if options.Trans != nil {
		switch options.Trans.(type) {
		case *sql.Tx:
		default:
			return id, errors.New("invalid transaction")
		}
	}

	ds.checkConnection()
	if options.Trans == nil {
		defer ds.CloseConnection()
	}
	//println(sqlString)
	var result sql.Result
	var err error
	if options.Trans != nil {
		result, err = options.Trans.(*sql.Tx).Exec(sqlString, params...)
	} else {
		result, err = ds.Db.Exec(sqlString, params...)
	}
	if err != nil {
		return id, err
	}
	rc, err := result.RowsAffected()
	if rc == 0 {
		return 0, err
	}
	if id <= 0 {
		return ds.lastInsertID(options.Model, result, options.Trans)
	}
	return id, nil
}

// UpdateSQL executes a SQL query string
func (ds *SQLDriver) UpdateSQL(sqlString string, transaction any) (err error) {
	ds.checkConnection()
	if transaction != nil {
		if trans, ok := transaction.(*sql.Tx); ok {
			_, err = trans.Exec(sqlString)
		} else {
			return errors.New("invalid transaction")
		}
	} else {
		defer ds.CloseConnection()
		_, err = ds.Db.Exec(sqlString)
	}
	return err
}

// BeginTransaction begins a transaction and returns an *sql.Tx
func (ds *SQLDriver) BeginTransaction() (any, error) {
	ds.checkConnection()
	return ds.Db.Begin()
}

// CommitTransaction commit a *sql.Tx transaction
func (ds *SQLDriver) CommitTransaction(trans any) error {
	switch trans.(type) {
	case *sql.Tx:
	default:
		return errors.New("invalid transaction")
	}
	return trans.(*sql.Tx).Commit()
}

// RollbackTransaction rollback a *sql.Tx transaction
func (ds *SQLDriver) RollbackTransaction(trans any) error {
	switch trans.(type) {
	case *sql.Tx:
	default:
		return errors.New("invalid transaction")
	}
	return trans.(*sql.Tx).Rollback()
}
