package nervatura

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	nt "github.com/nervatura/nervatura/service/pkg/nervatura"
	ut "github.com/nervatura/nervatura/service/pkg/utils"
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

//IM is a map[string]interface{} type short alias
type IM = map[string]interface{}

//IL is a []interface{} type short alias
type IL = []interface{}

//SM is a map[string]string type short alias
type SM = map[string]string

//SL is a []string type short alias
type SL = []string

//SQLDriver a go database/sql DataDriver
type SQLDriver struct {
	alias   string
	connStr string
	engine  string
	db      *sql.DB
	Config  IM
}

var dropList = []string{
	"pattern", "movement", "payment", "item", "trans", "barcode", "price", "tool", "product", "tax", "rate",
	"place", "currency", "project", "customer", "event", "contact", "address", "numberdef", "log", "fieldvalue",
	"deffield", "ui_audit", "link", "ui_userconfig", "ui_printqueue", "employee",
	"ui_report", "ui_message", "ui_menufields", "ui_menu", "groups"}

var createList = []string{
	"groups", "ui_menu", "ui_menufields", "ui_message", "ui_report",
	"employee", "ui_printqueue", "ui_userconfig", "link", "ui_audit", "deffield", "fieldvalue", "log", "numberdef",
	"address", "contact", "event", "customer", "project", "currency", "place", "rate", "tax", "product", "tool", "price",
	"barcode", "trans", "item", "payment", "movement", "pattern"}

func (ds *SQLDriver) decodeEngine(sqlStr string) string {
	const (
		dcCasInt      = "{CAS_INT}"
		dcCaeInt      = "{CAE_INT}"
		dcCasFloat    = "{CAS_FLOAT}"
		dcCaeFloat    = "{CAE_FLOAT}"
		dcCasDate     = "{CAS_DATE}"
		dcCaeDate     = "{CAE_DATE}"
		dcFmsDate     = "{FMS_DATE}"
		dcFmeDate     = "{FME_DATE}"
		dcFmsDateTime = "{FMS_DATETIME}"
		dcFmeDateTime = "{FME_DATETIME}"
	)

	switch ds.engine {
	case "sqlite", "sqlite3", "postgres":
		sqlStr = strings.ReplaceAll(sqlStr, dcCasInt, "cast(")
		sqlStr = strings.ReplaceAll(sqlStr, dcCaeInt, " as integer)")
		sqlStr = strings.ReplaceAll(sqlStr, dcCasFloat, "cast(")
		sqlStr = strings.ReplaceAll(sqlStr, dcCaeFloat, " as float8)")
		sqlStr = strings.ReplaceAll(sqlStr, dcCasDate, "cast(")
		sqlStr = strings.ReplaceAll(sqlStr, dcCaeDate, " "+"as date)")
		sqlStr = strings.ReplaceAll(sqlStr, dcFmsDate, "to_char(")
		sqlStr = strings.ReplaceAll(sqlStr, dcFmeDate, ", 'YYYY-MM-DD')")
		sqlStr = strings.ReplaceAll(sqlStr, dcFmsDateTime, "to_char(")
		sqlStr = strings.ReplaceAll(sqlStr, dcFmeDateTime, ", 'YYYY-MM-DD HH24:MI')")
	case "mysql":
		sqlStr = strings.ReplaceAll(sqlStr, dcCasInt, "cast(")
		sqlStr = strings.ReplaceAll(sqlStr, dcCaeInt, " as signed)")
		sqlStr = strings.ReplaceAll(sqlStr, dcCasFloat, "cast(")
		sqlStr = strings.ReplaceAll(sqlStr, dcCaeFloat, " as decimal)")
		sqlStr = strings.ReplaceAll(sqlStr, dcCasDate, "cast(")
		sqlStr = strings.ReplaceAll(sqlStr, dcCaeDate, " as date)")
		sqlStr = strings.ReplaceAll(sqlStr, dcFmsDate, "date_format(")
		sqlStr = strings.ReplaceAll(sqlStr, dcFmeDate, ", '%Y-%m-%d')")
		sqlStr = strings.ReplaceAll(sqlStr, dcFmsDateTime, "date_format(")
		sqlStr = strings.ReplaceAll(sqlStr, dcFmeDateTime, ", '%Y-%m-%d %H:%i')")
		sqlStr = strings.ReplaceAll(sqlStr, " groups", " `groups`")
	case "mssql":
		sqlStr = strings.ReplaceAll(sqlStr, dcCasInt, "cast(")
		sqlStr = strings.ReplaceAll(sqlStr, dcCaeInt, " as int)")
		sqlStr = strings.ReplaceAll(sqlStr, dcCasFloat, "cast(")
		sqlStr = strings.ReplaceAll(sqlStr, dcCaeFloat, " as float)")
		sqlStr = strings.ReplaceAll(sqlStr, dcCasDate, "cast(")
		sqlStr = strings.ReplaceAll(sqlStr, dcCaeDate, " as date)")
		sqlStr = strings.ReplaceAll(sqlStr, dcFmsDate, "convert(varchar(10),")
		sqlStr = strings.ReplaceAll(sqlStr, dcFmeDate, ", 120)")
		sqlStr = strings.ReplaceAll(sqlStr, dcFmsDateTime, "convert(varchar(16),")
		sqlStr = strings.ReplaceAll(sqlStr, dcFmeDateTime, ", 120)")
	}
	//print(sqlStr)
	return sqlStr
}

func (ds *SQLDriver) getDataType(dtype string) string {
	types := IM{
		"id": SM{
			"base":     "INTEGER PRIMARY KEY AUTOINCREMENT",
			"postgres": "SERIAL PRIMARY KEY",
			"mysql":    "INT AUTO_INCREMENT NOT NULL, PRIMARY KEY (id)",
			"mssql":    "INT IDENTITY PRIMARY KEY"},
		"integer": SM{
			"base":  "INTEGER",
			"mysql": "INT",
			"mssql": "INT"},
		"float": SM{
			"base":     "DOUBLE",
			"postgres": "FLOAT8",
			"mssql":    "DECIMAL(19,4)"},
		"string": SM{
			"base":     "CHAR",
			"postgres": "VARCHAR",
			"mysql":    "VARCHAR",
			"mssql":    "VARCHAR"},
		"password": SM{
			"base":     "CHAR",
			"postgres": "VARCHAR",
			"mysql":    "VARCHAR",
			"mssql":    "VARCHAR"},
		"text": SM{
			"base":  "TEXT",
			"mysql": "LONGTEXT",
			"mssql": "VARCHAR(max)"},
		"date": SM{
			"base": "DATE"},
		"datetime": SM{
			"base":  "TIMESTAMP",
			"mysql": "DATETIME",
			"mssql": "DATETIME2"},
		"reference": SM{
			"base":     "INTEGER REFERENCES foreign_key ON DELETE ",
			"postgres": "INTEGER REFERENCES foreign_key ON DELETE ",
			"mysql":    "INT, INDEX index_name (field_name), FOREIGN KEY (field_name) REFERENCES foreign_key ON DELETE ",
			"mssql":    "INT NULL, CONSTRAINT constraint_name FOREIGN KEY (field_name) REFERENCES foreign_key ON DELETE "}}
	if _, found := types[dtype]; found {
		if _, found := types[dtype].(SM)[ds.engine]; found {
			return types[dtype].(SM)[ds.engine]
		}
		return types[dtype].(SM)["base"]
	}
	return ""
}

func setRefID(refID IM, mname string, keyFields []string, values IM, id int64) IM {
	if _, found := refID[mname]; !found {
		refID[mname] = IM{}
	}
	if _, valid := values[keyFields[0]].(string); !valid {
		return refID
	}
	kf1 := values[keyFields[0]].(string)
	switch len(keyFields) {
	case 1:
		refID[mname].(IM)[kf1] = strconv.FormatInt(id, 10)

	case 2:
		if _, valid := values[keyFields[1]].(string); !valid {
			return refID
		}
		kf2 := values[keyFields[1]].(string)
		if _, found := refID[mname].(IM)[kf1]; !found {
			refID[mname].(IM)[kf1] = IM{}
		}
		refID[mname].(IM)[kf1].(IM)[kf2] = strconv.FormatInt(id, 10)

	case 3:
		_, valid_1 := values[keyFields[1]].(string)
		_, valid_2 := values[keyFields[2]].(string)
		if !valid_1 || !valid_2 {
			return refID
		}
		kf2 := values[keyFields[1]].(string)
		kf3 := values[keyFields[2]].(string)
		if _, found := refID[mname].(IM)[kf1]; !found {
			refID[mname].(IM)[kf1] = IM{}
		}
		if _, found := refID[mname].(IM)[kf1].(IM)[kf2]; !found {
			refID[mname].(IM)[kf1].(IM)[kf2] = IM{}
		}
		refID[mname].(IM)[kf1].(IM)[kf2].(IM)[kf3] = strconv.FormatInt(id, 10)
	}
	return refID
}

func getRefID(refID IM, value interface{}) interface{} {
	if keyFields, valid := value.([]string); valid {
		switch len(keyFields) {
		case 2:
			return refID[keyFields[0]].(IM)[keyFields[1]]
		case 3:
			return refID[keyFields[0]].(IM)[keyFields[1]].(IM)[keyFields[2]]
		case 4:
			return refID[keyFields[0]].(IM)[keyFields[1]].(IM)[keyFields[2]].(IM)[keyFields[3]]
		default:
			return "0"
		}
	}
	return value

}

//Properties - DataDriver features
func (ds *SQLDriver) Properties() struct{ SQL, Transaction bool } {
	return struct{ SQL, Transaction bool }{SQL: true, Transaction: true}
}

//Connection - returns the database connection
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
		Connected: (ds.db != nil),
		Engine:    ds.engine,
	}
}

//CreateConnection create a new database connection
func (ds *SQLDriver) CreateConnection(alias, connStr string) error {
	if ds.db != nil {
		if err := ds.db.Close(); err != nil {
			return err
		}
	}
	engine := strings.Split(connStr, "://")[0]
	if engine == "sqlite" {
		connStr = strings.ReplaceAll(connStr, "sqlite://", "")
	}
	if engine == "mysql" {
		connStr = strings.TrimPrefix(connStr, engine+"://")
	}
	db, err := sql.Open(engine, connStr)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	db.SetMaxOpenConns(int(ut.ToInteger(ds.Config["SQL_MAX_OPEN_CONNS"], 10)))
	db.SetMaxIdleConns(int(ut.ToInteger(ds.Config["SQL_MAX_IDLE_CONNS"], 3)))
	db.SetConnMaxLifetime(time.Minute * time.Duration(int(ut.ToInteger(ds.Config["SQL_CONN_MAX_LIFETIME"], 15))))
	ds.db = db
	ds.alias = alias
	ds.engine = engine
	ds.connStr = connStr
	return nil
}

// getPrmString - get database parameter string
func (ds *SQLDriver) getPrmString(index int) string {
	if ds.engine == "postgres" {
		return "$" + strconv.Itoa(index)
	}
	return "?"
}

// CheckHashtable - check/create a password ref. table
func (ds *SQLDriver) CheckHashtable(hashtable string) error {

	if ds.db == nil {
		return errors.New(ut.GetMessage("missing_driver"))
	}
	var name string
	sqlString := ""
	if ds.engine == "sqlite" {
		sqlString = fmt.Sprintf(
			"select name from sqlite_master where name = %s ", ds.getPrmString(1))
	} else {
		sqlString = fmt.Sprintf(
			"select table_name from information_schema.tables where table_name = %s ", ds.getPrmString(1))
	}
	err := ds.db.QueryRow(sqlString, hashtable).Scan(&name)
	if err != nil {
		stringType := ds.getDataType("string") + "(255)"
		textType := ds.getDataType("text")
		sqlString = fmt.Sprintf("CREATE TABLE %s ( refname %s, value %s);", hashtable, stringType, textType)
		_, err = ds.db.Exec(sqlString)
		if err != nil {
			return err
		}
		sqlString = fmt.Sprintf("CREATE UNIQUE INDEX %s_refname_idx ON %s (refname);", hashtable, hashtable)
		_, err = ds.db.Exec(sqlString)
	}

	return err
}

// UpdateHashtable - set a password
func (ds *SQLDriver) UpdateHashtable(hashtable, refname, value string) error {
	err := ds.CheckHashtable(hashtable)
	if err != nil {
		return err
	}
	sqlString := fmt.Sprintf(
		"select value from %s where refname = %s", hashtable, ds.getPrmString(1))
	var hash string
	err = ds.db.QueryRow(sqlString, refname).Scan(&hash)
	if err != nil {
		sqlString = fmt.Sprintf(
			"insert into %s(value, refname) values(%s,%s)",
			hashtable, ds.getPrmString(1), ds.getPrmString(2))
	} else {
		sqlString = fmt.Sprintf(
			"update %s set value=%s where refname=%s",
			hashtable, ds.getPrmString(1), ds.getPrmString(2))
	}
	_, err = ds.db.Exec(sqlString, value, refname)
	return err
}

func (ds *SQLDriver) tableName(name string) string {
	if ds.engine == "mysql" {
		return "`" + name + "`"
	}
	return name
}

//dropData - drop all tables if exist
func (ds *SQLDriver) dropData(logData []SM) ([]SM, error) {

	trans, err := ds.db.Begin()
	if err != nil {
		logData = append(logData, SM{
			"stamp":   time.Now().Format(nt.TimeLayout),
			"state":   "err",
			"message": err.Error()})
		return logData, err
	}
	defer rollBackTrans(trans, err)

	logData = append(logData, SM{
		"stamp":   time.Now().Format(nt.TimeLayout),
		"state":   "create",
		"message": ut.GetMessage("log_drop_table")})

	dropList = append(dropList, ut.ToString(os.Getenv("NT_HASHTABLE"), "ref17890714"))
	for index := 0; index < len(dropList); index++ {
		sqlString := ""
		if ds.engine == "mssql" {
			sqlString = "DROP TABLE " + dropList[index] + ";"
		} else {
			sqlString = "DROP TABLE IF EXISTS " + ds.tableName(dropList[index]) + ";"
		}
		if _, err := trans.Exec(sqlString); err != nil {
			logData = append(logData, SM{
				"stamp":   time.Now().Format(nt.TimeLayout),
				"state":   "err",
				"message": err.Error()})
			return logData, err
		}
	}
	if err = trans.Commit(); err != nil {
		logData = append(logData, SM{
			"stamp":   time.Now().Format(nt.TimeLayout),
			"state":   "err",
			"message": err.Error()})
		return logData, err
	}
	return logData, nil
}

func (ds *SQLDriver) createTableFields(sqlString, fieldname, indexName string, field nt.MF) string {
	sqlString += fieldname
	if field.References != nil {
		reference := ds.getDataType("reference")
		reference = strings.ReplaceAll(reference, "foreign_key", ds.tableName(field.References[0])+"(id)")
		reference = strings.ReplaceAll(reference, "field_name", fieldname)
		reference = strings.ReplaceAll(reference, "index_name", fieldname+"__idx")
		reference = strings.ReplaceAll(reference, "constraint_name", indexName+"__"+fieldname+"__constraint")
		if (ds.engine == "mssql") && (len(field.References) == 3) {
			reference += field.References[2]
		} else {
			reference += field.References[1]
		}
		sqlString += " " + reference
	} else {
		sqlString += " " + ds.getDataType(field.Type)
		if field.Length > 0 {
			sqlString += "(" + strconv.Itoa(field.Length) + ")"
		}
	}
	if field.NotNull && field.References == nil {
		sqlString += " NOT NULL"
	}
	if field.Default != nil && field.Default != "nextnumber" {
		sqlString += " " + "DEFAULT" + " " + ut.ToString(field.Default, "")
	}
	sqlString += ", "
	return sqlString
}

//createTable - create all tables
func (ds *SQLDriver) createTable(logData []SM, trans *sql.Tx) ([]SM, error) {

	logData = append(logData, SM{
		"stamp":   time.Now().Format(nt.TimeLayout),
		"state":   "create",
		"message": ut.GetMessage("log_create_table")})
	model := nt.DataModel()["model"].(IM)

	for index := 0; index < len(createList); index++ {
		sqlString := "CREATE TABLE " + ds.tableName(createList[index]) + "("
		for fld := 0; fld < len(model[createList[index]].(IM)["_fields"].(SL)); fld++ {
			fieldname := model[createList[index]].(IM)["_fields"].(SL)[fld]
			field := model[createList[index]].(IM)[fieldname].(nt.MF)
			sqlString = ds.createTableFields(sqlString, fieldname, createList[index], field)
		}
		sqlString += ");"
		sqlString = strings.ReplaceAll(sqlString, ", );", ");")
		//println(sqlString)
		_, err := trans.Exec(sqlString)
		if err != nil {
			logData = append(logData, SM{
				"stamp":   time.Now().Format(nt.TimeLayout),
				"state":   "err",
				"message": err.Error()})
			return logData, err
		}
	}

	return logData, nil
}

// createIndex - create indexes
func (ds *SQLDriver) createIndex(logData []SM, trans *sql.Tx) ([]SM, error) {

	logData = append(logData, SM{
		"stamp":   time.Now().Format(nt.TimeLayout),
		"state":   "create",
		"message": ut.GetMessage("log_create_index")})
	indexRows := nt.DataModel()["index"].(map[string]nt.MI)
	for ikey, ifield := range indexRows {
		sqlString := "CREATE INDEX "
		if ifield.Unique {
			sqlString = "CREATE UNIQUE INDEX "
		}
		sqlString += ikey + " ON " + ds.tableName(ifield.Model) + "("
		for index := 0; index < len(ifield.Fields); index++ {
			sqlString += ifield.Fields[index] + ", "
		}
		sqlString += ");"
		sqlString = strings.ReplaceAll(sqlString, ", );", ");")
		//println(sqlString)
		_, err := trans.Exec(sqlString)
		if err != nil {
			logData = append(logData, SM{
				"stamp":   time.Now().Format(nt.TimeLayout),
				"state":   "err",
				"message": err.Error()})
			return logData, err
		}
	}

	return logData, nil
}

// insertData - insert data
func (ds *SQLDriver) insertData(logData []SM, trans *sql.Tx) ([]SM, error) {

	model := nt.DataModel()["model"].(IM)
	logData = append(logData, SM{
		"stamp":   time.Now().Format(nt.TimeLayout),
		"state":   "create",
		"message": ut.GetMessage("log_init_data")})
	dataRows := nt.DataModel()["data"].(map[string][]IM)
	dataList := []string{
		"groups", "currency", "customer", "employee", "address", "contact", "place", "tax", "product",
		"numberdef", "deffield", "fieldvalue"}
	refID := IM{}
	for index := 0; index < len(dataList); index++ {
		mname := dataList[index]
		keyFields := model[mname].(IM)["_key"].([]string)
		insertID := ""
		if ds.engine == "mssql" {
			insertID = fmt.Sprintf("SET IDENTITY_INSERT %s ON; ", mname)
		}
		for index := 0; index < len(dataRows[mname]); index++ {
			refID = setRefID(refID, mname, keyFields, dataRows[mname][index], int64(index+1))
			params := []interface{}{index + 1}
			fldNames := []string{"id"}
			fldValues := []string{ds.getPrmString(1)}
			for field, value := range dataRows[mname][index] {
				params = append(params, getRefID(refID, value))
				fldNames = append(fldNames, field)
				fldValues = append(fldValues, ds.getPrmString(len(params)))
			}
			sqlString := insertID + fmt.Sprintf("INSERT INTO %s (%s) VALUES(%s);",
				ds.tableName(mname), strings.Join(fldNames, ","), strings.Join(fldValues, ","))
			//println(sqlString)
			_, err := trans.Exec(sqlString, params...)
			if err != nil {
				logData = append(logData, SM{
					"stamp":   time.Now().Format(nt.TimeLayout),
					"state":   "err",
					"message": err.Error()})
				return logData, err
			}
		}
	}

	return logData, nil
}

func rollBackTrans(trans *sql.Tx, err error) {
	pe := recover()
	if trans != nil {
		if err != nil || pe != nil {
			if rb_err := trans.Rollback(); rb_err != nil {
				return
			}
		}
	}
	if pe != nil {
		panic(pe)
	}
}

// CreateDatabase - create a Nervatura Database
func (ds *SQLDriver) CreateDatabase(logData []SM) ([]SM, error) {
	var err error
	if ds.db == nil {
		logData = append(logData, SM{
			"stamp":   time.Now().Format(nt.TimeLayout),
			"state":   "err",
			"message": ut.GetMessage("missing_driver")})
		return logData, errors.New(ut.GetMessage("missing_driver"))
	}

	logData = append(logData, SM{
		"database": ds.alias,
		"stamp":    time.Now().Format(nt.TimeLayout),
		"state":    "create",
		"message":  ut.GetMessage("log_start_process")})

	if logData, err = ds.dropData(logData); err != nil {
		return logData, err
	}

	trans, err := ds.db.Begin()
	if err != nil {
		logData = append(logData, SM{
			"stamp":   time.Now().Format(nt.TimeLayout),
			"state":   "err",
			"message": err.Error()})
		return logData, err
	}
	defer rollBackTrans(trans, err)

	if logData, err = ds.createTable(logData, trans); err != nil {
		return logData, err
	}

	if logData, err = ds.createIndex(logData, trans); err != nil {
		return logData, err
	}

	if logData, err = ds.insertData(logData, trans); err != nil {
		return logData, err
	}

	switch ds.engine {
	case "postgres":
		//update all sequences
		sqlString := ""
		for index := 0; index < len(createList); index++ {
			sqlString += fmt.Sprintf("SELECT setval('%s_id_seq', (SELECT max(id) FROM %s));",
				createList[index], createList[index])
		}
		if _, err := trans.Exec(sqlString); err != nil {
			logData = append(logData, SM{
				"stamp":   time.Now().Format(nt.TimeLayout),
				"state":   "err",
				"message": err.Error()})
			return logData, err
		}
	}

	sqlString := fmt.Sprintf(
		"insert into fieldvalue(fieldname, value) values('%s','%s')",
		"version", ut.ToString(ds.Config["version"], ""))
	if _, err := trans.Exec(sqlString); err != nil {
		logData = append(logData, SM{
			"stamp":   time.Now().Format(nt.TimeLayout),
			"state":   "err",
			"message": err.Error()})
		return logData, err
	}

	err = trans.Commit()
	if err != nil {
		logData = append(logData, SM{
			"stamp":   time.Now().Format(nt.TimeLayout),
			"state":   "err",
			"message": err.Error()})
		return logData, err
	}

	//compact
	logData = append(logData, SM{
		"stamp":   time.Now().Format(nt.TimeLayout),
		"state":   "create",
		"message": ut.GetMessage("log_rebuilding")})
	switch ds.engine {
	case "postgres", "sqlite", "sqlite3":
		sqlString := "vacuum"
		_, err = ds.db.Exec(sqlString)
		if err != nil {
			logData = append(logData, SM{
				"stamp":   time.Now().Format(nt.TimeLayout),
				"state":   "err",
				"message": err.Error()})
			return logData, err
		}
	}

	return logData, nil
}

func (ds *SQLDriver) getFilterString(filter nt.Filter, start bool, sqlString string, params []interface{}) (string, []interface{}) {
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
		sqlString += " " + filter.Comp + " " + filter.Value.(string)
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
	}
	if !filter.Or {
		sqlString += ")"
	}
	return sqlString, params
}

func (ds *SQLDriver) decodeSQL(queries []nt.Query) (string, []interface{}) {
	sqlString := ""
	params := make([]interface{}, 0)
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
	}
	return strings.Trim(sqlString, " "), params
}

//Query is a basic nosql friendly queries the database
func (ds *SQLDriver) Query(queries []nt.Query, trans interface{}) ([]IM, error) {
	sqlString, params := ds.decodeSQL(queries)
	return ds.QuerySQL(sqlString, params, trans)
}

func initQueryCols(engine string, cols []*sql.ColumnType) ([]interface{}, []string, []string) {
	values := make([]interface{}, len(cols))
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
		case "BOOL", "BOOLEAN", "BIT":
			values[i] = new(sql.NullBool)
		case "INTEGER", "SERIAL", "INT", "INT4", "INT8":
			values[i] = new(sql.NullInt64)
		case "DOUBLE", "FLOAT8", "DECIMAL(19,4)", "DECIMAL", "NUMERIC":
			values[i] = new(sql.NullFloat64)
		case "DATETIME", "TIMESTAMP", "DATE":
			if engine == "postgres" || engine == "sqlite" {
				values[i] = new(sql.NullTime)
			}
		}
	}
	return values, fields, dbtypes
}

func getQueryRowValue(value interface{}, dbtype string) interface{} {
	switch v := value.(type) {
	case *sql.NullBool:
		if v.Valid {
			return v.Bool
		}
		return nil

	case *sql.NullInt64:
		if v.Valid {
			return v.Int64
		}
		return nil

	case *sql.NullFloat64:
		if v.Valid {
			return v.Float64
		}
		return nil

	case *sql.NullTime:
		if v.Valid {
			if dbtype == "DATE" {
				return v.Time.Format("2006-01-02")
			}
			return v.Time
		}
		return nil

	case *sql.NullString:
		if v.Valid {
			return v.String
		}
		return nil

	}
	return value
}

//QuerySQL executes a SQL query
func (ds *SQLDriver) QuerySQL(sqlString string, params []interface{}, trans interface{}) ([]IM, error) {
	result := make([]IM, 0)
	var rows *sql.Rows
	var err error
	if trans != nil {
		switch trans.(type) {
		case *sql.Tx:
		default:
			return result, errors.New(ut.GetMessage("invalid_trans"))
		}
	}

	//println(ds.decodeEngine(sqlString))
	if trans != nil {
		rows, err = trans.(*sql.Tx).Query(ds.decodeEngine(sqlString), params...)
	} else {
		rows, err = ds.db.Query(ds.decodeEngine(sqlString), params...)
	}
	if err != nil {
		return result, err
	}
	defer rows.Close()

	cols, err := rows.ColumnTypes()
	if err != nil {
		return result, err
	}
	values, fields, dbtypes := initQueryCols(ds.engine, cols)

	for rows.Next() {
		err = rows.Scan(values...)
		if err != nil {
			return result, err
		}
		row := make(IM)
		for index, value := range values {
			row[fields[index]] = getQueryRowValue(value, dbtypes[index])
		}
		result = append(result, row)
	}
	return result, nil
}

func (ds *SQLDriver) lastInsertID(model string, result sql.Result, trans interface{}) (int64, error) {
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
			err = ds.db.QueryRow(sqlString).Scan(&resid)
		}
		if err != nil {
			return -1, err
		}
	}
	return resid, nil
}

//Update is a basic nosql friendly update/insert/delete and returns the update/insert id
func (ds *SQLDriver) Update(options nt.Update) (int64, error) {
	sqlString := ""
	id := options.IDKey
	params := make([]interface{}, 0)
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
			return id, errors.New(ut.GetMessage("invalid_trans"))
		}
	}
	//println(sqlString)
	var result sql.Result
	var err error
	if options.Trans != nil {
		result, err = options.Trans.(*sql.Tx).Exec(ds.decodeEngine(sqlString), params...)
	} else {
		result, err = ds.db.Exec(ds.decodeEngine(sqlString), params...)
	}
	if err != nil {
		return id, err
	}
	if id <= 0 {
		return ds.lastInsertID(options.Model, result, options.Trans)
	}
	return id, nil
}

//BeginTransaction begins a transaction and returns an *sql.Tx
func (ds *SQLDriver) BeginTransaction() (interface{}, error) {
	return ds.db.Begin()
}

//CommitTransaction commit a *sql.Tx transaction
func (ds *SQLDriver) CommitTransaction(trans interface{}) error {
	switch trans.(type) {
	case *sql.Tx:
	default:
		return errors.New(ut.GetMessage("invalid_trans"))
	}
	return trans.(*sql.Tx).Commit()
}

//RollbackTransaction rollback a *sql.Tx transaction
func (ds *SQLDriver) RollbackTransaction(trans interface{}) error {
	switch trans.(type) {
	case *sql.Tx:
	default:
		return errors.New(ut.GetMessage("invalid_trans"))
	}
	return trans.(*sql.Tx).Rollback()
}

func (ds *SQLDriver) getQueryKeyOption(options IM, keys SL, sqlString string, params IL) (string, IL) {
	queryParams := make([]interface{}, 0)
	fieldParams := make([]interface{}, 0)
	for _, key := range keys {
		if _, found := options[key]; found {
			queryParams = append(queryParams, options[key])
			fieldParams = append(fieldParams, ds.getPrmString(len(params)+len(queryParams)))
		}
	}
	if len(queryParams) > 0 {
		params = append(params, queryParams...)
		sqlString = fmt.Sprintf(sqlString, fieldParams...)
		return sqlString, params
	}
	return "", params
}

func (ds *SQLDriver) splitInParams(inString string, start int) (IL, SL) {
	inValues := strings.Split(inString, ",")
	params := make([]interface{}, 0)
	inPrm := make([]string, 0)
	for _, value := range inValues {
		params = append(params, value)
		inPrm = append(inPrm, ds.getPrmString(len(params)+start))
	}
	return params, inPrm
}

func (ds *SQLDriver) getQueryKeySplit(options IM, key, sqlString string, params IL) (string, IL) {
	if _, found := options[key]; found {
		values, fieldParams := ds.splitInParams(options[key].(string), len(params))
		params = append(params, values...)
		sqlString = fmt.Sprintf(sqlString, strings.Join(fieldParams, ","))
		return sqlString, params
	}
	return "", params
}

func (ds *SQLDriver) getID2Refnumber(options IM) (string, IL, error) {
	var sqlString, whereString string
	params := make(IL, 0)

	const whereDelete = " and deleted = 0"
	useDeleted := func(value interface{}, whereString string) string {
		if !ut.ToBoolean(value, false) {
			return whereString
		}
		return ""
	}

	switch options["nervatype"] {
	case "address", "contact":
		if _, found := options["refId"]; found {
			sqlString = fmt.Sprintf(`select {CAS_INT}count(*){CAE_INT} as count from %s `, options["nervatype"])
			whereString, params = ds.getQueryKeyOption(options,
				SL{"refTypeId", "refId", "id"}, ` where nervatype = %s and ref_id = %s and id <= %s `, params)
			sqlString += whereString
			sqlString += useDeleted(options["useDeleted"], whereDelete)
			return sqlString, params, nil
		}
		sqlString = fmt.Sprintf(`select nt.groupvalue as head_nervatype, t.*
		  from %s t inner join groups nt on t.nervatype = nt.id`, options["nervatype"])
		whereString, params = ds.getQueryKeyOption(options,
			SL{"id"}, ` where t.id = %s `, params)
		sqlString += whereString
		sqlString += useDeleted(options["useDeleted"], " and t.deleted = 0")
		return sqlString, params, nil

	case "fieldvalue", "setting":
		if _, found := options["refId"]; found {
			sqlString = `select {CAS_INT}count(*){CAE_INT} as count from fieldvalue `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"fieldname", "refId", "id"}, ` where fieldname = %s and ref_id = %s and id <= %s `, params)
			sqlString += whereString
			sqlString += useDeleted(options["useDeleted"], whereDelete)
			return sqlString, params, nil
		}
		sqlString = `select fv.*, nt.groupvalue as head_nervatype 
			  from fieldvalue fv inner join deffield df on fv.fieldname = df.fieldname 
				inner join groups nt on df.nervatype = nt.id `
		whereString, params = ds.getQueryKeyOption(options,
			SL{"id"}, ` where fv.id = %s `, params)
		sqlString += whereString
		sqlString += useDeleted(options["useDeleted"], " and fv.deleted = 0")
		return sqlString, params, nil

	case "item", "payment", "movement":
		if _, found := options["refId"]; found {
			sqlString = fmt.Sprintf(`select {CAS_INT}count(*){CAE_INT} as count from %s `, options["nervatype"])
			whereString, params = ds.getQueryKeyOption(options,
				SL{"refId", "id"}, ` where trans_id = %s and id <= %s `, params)
			sqlString += whereString
			sqlString += useDeleted(options["useDeleted"], whereDelete)
			return sqlString, params, nil
		}
		sqlString = fmt.Sprintf(`select ti.*, t.transnumber, tt.groupvalue as transtype 
				  from %s ti inner join trans t on ti.trans_id = t.id 
					inner join groups tt on t.transtype = tt.id  `, options["nervatype"])
		whereString, params = ds.getQueryKeyOption(options,
			SL{"id"}, ` where ti.id = %s `, params)
		sqlString += whereString
		sqlString += useDeleted(options["useDeleted"],
			" and ( t.deleted = 0 or tt.groupvalue in ('cash', 'invoice', 'receipt'))")
		return sqlString, params, nil

	case "price":
		sqlString = `select pr.*, p.partnumber from price pr 
			  inner join product p on pr.product_id = p.id  `
		whereString, params = ds.getQueryKeyOption(options,
			SL{"id"}, ` where pr.id = %s `, params)
		sqlString += whereString
		sqlString += useDeleted(options["useDeleted"], " and p.deleted = 0")
		return sqlString, params, nil

	case "link":
		sqlString = `select l.*, nt1.groupvalue as nervatype1, nt2.groupvalue as nervatype2 
			  from link l inner join groups nt1 on l.nervatype_1 = nt1.id 
				inner join groups nt2 on l.nervatype_2 = nt2.id  `
		whereString, params = ds.getQueryKeyOption(options,
			SL{"id"}, ` where l.id = %s `, params)
		sqlString += whereString
		sqlString += useDeleted(options["useDeleted"], " and l.deleted = 0")
		return sqlString, params, nil

	case "rate":
		sqlString = `select r.*, rt.groupvalue as rate_type, p.planumber 
			  from rate r inner join groups rt on r.ratetype = rt.id 
				left join place p on r.place_id = p.id  `
		whereString, params = ds.getQueryKeyOption(options,
			SL{"id"}, ` where r.id = %s `, params)
		sqlString += whereString
		sqlString += useDeleted(options["useDeleted"], " and r.deleted = 0")
		return sqlString, params, nil

	case "log":
		sqlString = `select l.*, e.empnumber from log l 
			  inner join employee e on l.employee_id = e.id  `
		whereString, params = ds.getQueryKeyOption(options,
			SL{"id"}, ` where l.id = %s `, params)
		sqlString += whereString
		return sqlString, params, nil

	default:
		sqlString = fmt.Sprintf(`select * from %s `, options["nervatype"])
		whereString, params = ds.getQueryKeyOption(options,
			SL{"id"}, ` where id = %s `, params)
		sqlString += whereString
		sqlString += useDeleted(options["useDeleted"], whereDelete)
		return sqlString, params, nil
	}
}

func filterValue(filter, value interface{}, trueResult, falseResult string) string {
	if value == filter {
		return trueResult
	}
	return falseResult
}

func (ds *SQLDriver) getRefnumber2ID(options IM) (string, IL, error) {
	var sqlString, whereString string
	params := make(IL, 0)

	const whereDelete = " and deleted = 0"

	keys := map[string]func(options IM) (string, IL, error){
		"barcode": func(options IM) (string, IL, error) {
			sqlString = filterValue(false, options["useDeleted"],
				`select barcode.id from barcode 
				inner join product on barcode.product_id = product.id
				where product.deleted = 0 and `, `select id from barcode where `)
			whereString, params = ds.getQueryKeyOption(options,
				SL{"refnumber"}, ` code = %s `, params)
			sqlString += whereString
			return sqlString, params, nil
		},

		"customer": func(options IM) (string, IL, error) {
			sqlString = filterValue(true, options["extraInfo"],
				fmt.Sprintf(`select c.id as id, ct.groupvalue as custtype, c.terms as terms, c.custname as custname, 
						c.taxnumber as taxnumber, addr.zipcode as zipcode, addr.city as city, addr.street as street 
					from customer c 
					inner join groups ct on c.custtype = ct.id 
					left join ( 
						select * from address where id = ( select min(id) from address where deleted = 0 
							and nervatype = ( select id from groups where groupname = 'nervatype' and groupvalue = 'customer') 
							and ref_id = ( select min(c.id) from customer c inner join groups ct on c.custtype = ct.id and groupvalue = 'own' where c.deleted = 0))
					) addr on c.id = addr.ref_id 
					where c.id = ( select min(c.id) from customer c inner join groups ct on c.custtype = ct.id and groupvalue = 'own' where c.deleted = 0)  
					union select c.id as id, ct.groupvalue as custtype, c.terms as terms, c.custname as custname, 
						c.taxnumber as taxnumber, addr.zipcode as zipcode, addr.city as city, addr.street as street 
					from customer c 
					inner join groups ct on c.custtype = ct.id 
					left join ( 
						select * from address where id = ( select min(id) from address where deleted = 0 
							and nervatype = ( select id from groups where groupname = 'nervatype' and groupvalue = 'customer') 
							and ref_id = ( select id from customer where custnumber = '%s'))
					) addr on c.id = addr.ref_id `, options["refnumber"]),
				`select c.id as id, ct.groupvalue as custtype 
					from customer c inner join groups ct on c.custtype = ct.id `)
			whereString, params = ds.getQueryKeyOption(options,
				SL{"refnumber"}, ` where c.custnumber = %s `, params)
			sqlString += whereString
			sqlString += filterValue(false, options["useDeleted"], " and c.deleted = 0", "")
			return sqlString, params, nil
		},

		"event": func(options IM) (string, IL, error) {
			sqlString = `select e.id as id, ntype.groupvalue as ref_nervatype 
				from event e inner join groups ntype on e.nervatype = ntype.id  `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"refnumber"}, ` where e.calnumber = %s `, params)
			sqlString += whereString
			sqlString += filterValue(false, options["useDeleted"], " and e.deleted = 0", "")
			return sqlString, params, nil
		},

		"groups": func(options IM) (string, IL, error) {
			sqlString = `select id from groups  `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"refType", "refnumber"}, ` where groupname = %s and groupvalue = %s `, params)
			sqlString += whereString
			sqlString += filterValue(false, options["useDeleted"], " and deleted=0", "")
			return sqlString, params, nil
		},

		"deffield": func(options IM) (string, IL, error) {
			sqlString = `select df.id, nt.groupvalue as ref_nervatype from deffield df 
			  inner join groups nt on df.nervatype = nt.id  `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"refnumber"}, ` where df.fieldname = %s `, params)
			sqlString += whereString
			sqlString += filterValue(false, options["useDeleted"], " and df.deleted=0", "")
			return sqlString, params, nil
		},

		"fieldvalue": func(options IM) (string, IL, error) {
			sqlString = `select id from fieldvalue  `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"refID", "refnumber"}, ` where ref_id = %s and fieldname = %s `, params)
			sqlString += whereString
			sqlString += filterValue(false, options["useDeleted"], whereDelete, "")
			return sqlString, params, nil
		},

		"item": func(options IM) (string, IL, error) {
			sqlString = fmt.Sprintf(
				`select it.id as id, ttype.groupvalue as transtype, dir.groupvalue as direction, cu.digit as digit, 
					it.qty as qty, it.discount as discount, it.tax_id as tax_id, ta.rate as rate 
				from item it 
				inner join trans t on it.trans_id = t.id and t.transnumber = '%s' 
				inner join tax ta on it.tax_id = ta.id 
				inner join groups ttype on t.transtype = ttype.id 
				inner join groups dir on t.direction = dir.id 
				left join currency cu on t.curr = cu.curr `,
				options["refnumber"])
			sqlString += filterValue(false, options["useDeleted"], " where t.deleted = 0 and it.deleted = 0", "")
			return sqlString, params, nil
		},

		"payment": func(options IM) (string, IL, error) {
			sqlString = fmt.Sprintf(
				`select it.id as id, ttype.groupvalue as transtype, dir.groupvalue as direction 
				from payment it 
				inner join trans t on it.trans_id = t.id and t.transnumber = '%s' 
				inner join groups ttype on t.transtype = ttype.id 
				inner join groups dir on t.direction = dir.id `, options["refnumber"])
			sqlString += filterValue(false, options["useDeleted"], ` where t.deleted = 0 and it.deleted = 0`, "")
			return sqlString, params, nil
		},

		"movement": func(options IM) (string, IL, error) {
			sqlString = fmt.Sprintf(
				`select it.id as id, ttype.groupvalue as transtype, dir.groupvalue as direction, mt.groupvalue as movetype 
				from movement it 
				inner join groups mt on it.movetype = mt.id 
				inner join trans t on it.trans_id = t.id and t.transnumber = '%s' 
				inner join groups ttype on t.transtype = ttype.id 
				inner join groups dir on t.direction = dir.id `, options["refnumber"])
			sqlString += filterValue(false, options["useDeleted"], ` where t.deleted = 0 and it.deleted = 0`, "")
			return sqlString, params, nil
		},

		"price": func(options IM) (string, IL, error) {
			sqlString = `select pr.id as id from price pr 
			inner join product p on pr.product_id = p.id   `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"refnumber", "curr", "validfrom", "qty"},
				` where p.partnumber = %s and pr.curr = %s and pr.validfrom = %s and pr.qty = %s `, params)
			sqlString += whereString
			sqlString += filterValue("price", options["pricetype"],
				` and pr.discount is null`, ` and pr.discount is not null`)
			sqlString += filterValue(false, options["useDeleted"], " and p.deleted = 0 and pr.deleted = 0", "")
			return sqlString, params, nil
		},

		"product": func(options IM) (string, IL, error) {
			sqlString =
				`select p.id as id, p.description as description, p.unit as unit, p.tax_id as tax_id, t.rate as rate 
			from product p left join tax t on p.tax_id = t.id `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"refnumber"}, ` where p.partnumber = %s `, params)
			sqlString += whereString
			sqlString += filterValue(false, options["useDeleted"], ` and p.deleted = 0`, "")
			return sqlString, params, nil
		},

		"place": func(options IM) (string, IL, error) {
			sqlString =
				`select p.id as id, pt.groupvalue as placetype 
			  from place p inner join groups pt on p.placetype = pt.id `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"refnumber"}, ` where p.planumber = %s `, params)
			sqlString += whereString
			sqlString += filterValue(false, options["useDeleted"], ` and p.deleted = 0`, "")
			return sqlString, params, nil
		},

		"tax": func(options IM) (string, IL, error) {
			sqlString = `select id, rate from tax `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"refnumber"}, ` where taxcode = %s `, params)
			sqlString += whereString
			return sqlString, params, nil
		},

		"trans": func(options IM) (string, IL, error) {
			sqlString = filterValue(false, options["useDeleted"],
				`select t.id as id, ttype.groupvalue as transtype, dir.groupvalue as direction, cu.digit as digit 
				from trans t 
				inner join groups ttype on t.transtype = ttype.id 
				inner join groups dir on t.direction = dir.id 
				left join currency cu on t.curr = cu.curr 
				where t.transnumber = %s and ( t.deleted = 0 or ( ttype.groupvalue = 'invoice' and dir.groupvalue = 'out') 
					or ( ttype.groupvalue = 'receipt' and dir.groupvalue = 'out') or ( ttype.groupvalue = 'cash'))`,
				`select t.id as id, ttype.groupvalue as transtype, dir.groupvalue as direction, cu.digit as digit 
				from trans t 
				inner join groups ttype on t.transtype = ttype.id 
				inner join groups dir on t.direction = dir.id 
				left join currency cu on t.curr = cu.curr where t.transnumber = %s `)
			sqlString, params = ds.getQueryKeyOption(options, SL{"refnumber"}, sqlString, params)
			return sqlString, params, nil
		},

		"setting": func(options IM) (string, IL, error) {
			sqlString = `select id from fieldvalue `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"refnumber"}, ` where ref_id is null and fieldname = %s `, params)
			sqlString += whereString
			sqlString += filterValue(false, options["useDeleted"], ` and deleted = 0`, "")
			return sqlString, params, nil
		},

		"link": func(options IM) (string, IL, error) {
			sqlString = fmt.Sprintf(
				`select l.id as id from link l 
				inner join groups nt1 on l.nervatype_1 = nt1.id and nt1.groupname = 'nervatype' 
				  and nt1.groupvalue = '%s' 
				inner join groups nt2 on l.nervatype_2 = nt2.id and nt2.groupname = 'nervatype' 
				and nt2.groupvalue = '%s' `,
				options["refType1"], options["refType2"])
			whereString, params = ds.getQueryKeyOption(options,
				SL{"refID1", "refID2"}, ` where l.ref_id_1 = %s and l.ref_id_2 = %s `, params)
			sqlString += whereString
			sqlString += filterValue(false, options["useDeleted"], ` and l.deleted = 0`, "")
			return sqlString, params, nil
		},

		"rate": func(options IM) (string, IL, error) {
			sqlString = fmt.Sprintf(
				`select r.id as id from rate r 
			  inner join groups rt on r.ratetype = rt.id and rt.groupvalue = '%s' 
				left join place p on r.place_id = p.id `, options["ratetype"])
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ratedate", "curr"}, ` where r.ratedate = %s and r.curr = %s `, params)
			sqlString += whereString
			if _, found := options["planumber"]; found {
				whereString, params = ds.getQueryKeyOption(options,
					SL{"planumber"}, ` and p.planumber = %s `, params)
				sqlString += whereString
			} else {
				sqlString += ` and r.place_id is null `
			}
			sqlString += filterValue(false, options["useDeleted"], ` and r.deleted = 0`, "")
			return sqlString, params, nil
		},

		"log": func(options IM) (string, IL, error) {
			sqlString = `select l.id as id from log l inner join employee e on l.employee_id = e.id `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"refnumber", "crdate"}, ` where e.empnumber = %s and l.crdate = %s `, params)
			sqlString += whereString
			return sqlString, params, nil
		},

		"ui_audit": func(options IM) (string, IL, error) {
			sqlString = `select au.id as id from ui_audit au `
			whereString, params = ds.getQueryKeyOption(options, SL{"usergroup"},
				` inner join groups ug on au.usergroup = ug.id and ug.groupvalue = %s 
				inner join groups nt on au.nervatype = nt.id `, params)
			sqlString += whereString
			if _, found := options["transType"]; found {
				if options["refType"] == "trans" {
					whereString, params = ds.getQueryKeyOption(options, SL{"transType"},
						` inner join groups st on au.subtype = st.id and 
					st.groupvalue = %s where `, params)
					sqlString += whereString
				} else {
					whereString, params = ds.getQueryKeyOption(options, SL{"transType"},
						` inner join ui_report rp on au.subtype = rp.id and 
					rp.reportkey = %s where `, params)
					sqlString += whereString
				}
			} else {
				sqlString += ` where subtype is null and `
			}
			whereString, params = ds.getQueryKeyOption(options, SL{"refType"},
				` nt.groupvalue = %s `, params)
			sqlString += whereString
			return sqlString, params, nil
		},

		"ui_menufields": func(options IM) (string, IL, error) {
			sqlString =
				`select mf.id as id from ui_menufields mf 
			  inner join ui_menu m on mf.menu_id = m.id `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"refnumber", "fieldname"}, ` and m.menukey = %s where mf.fieldname = %s `, params)
			sqlString += whereString
			return sqlString, params, nil
		},
	}

	if _, found := keys[options["nervatype"].(string)]; found {
		return keys[options["nervatype"].(string)](options)
	}

	switch options["nervatype"] {

	case "employee", "pattern", "project", "tool", "currency", "numberdef", "ui_report", "ui_menu":

		sqlString = fmt.Sprintf(`select * from %s where %s `, options["nervatype"], options["refField"])
		whereString, params = ds.getQueryKeyOption(options,
			SL{"refnumber"}, ` = %s `, params)
		sqlString += whereString
		sqlString += filterValue(false, options["useDeleted"], whereDelete, "")
		return sqlString, params, nil

	case "address", "contact":
		sqlString = fmt.Sprintf(
			`select %s.id as id from %s 
				inner join groups nt on %s.nervatype = nt.id 
				  and nt.groupname = 'nervatype' and nt.groupvalue = '%s' 
				inner join %s on %s.ref_id = %s.id 
				  and %s.%s = '%s' `,
			options["nervatype"], options["nervatype"], options["nervatype"], options["refType"],
			options["refType"], options["nervatype"], options["refType"], options["refType"],
			options["refField"], options["refnumber"])
		whereString = fmt.Sprintf(`where %s.deleted = 0 and %s.deleted = 0`,
			options["refType"], options["nervatype"])
		sqlString += filterValue(false, options["useDeleted"], whereString, "")
		return sqlString, params, nil

	default:
		return sqlString, params, errors.New(ut.GetMessage("invalid_refnumber"))
	}

}

func (ds *SQLDriver) getUpdateDeffields(options IM) (string, IL, error) {
	var sqlString, whereString string
	params := make(IL, 0)

	if _, found := options["fieldname"]; found {
		sqlString = `select ft.groupvalue as fieldtype from deffield df
        inner join groups ft on (df.fieldtype=ft.id) `
		whereString, params = ds.getQueryKeyOption(options,
			SL{"fieldname"}, ` where df.fieldname = %s `, params)
		sqlString += whereString
		return sqlString, params, nil
	}
	if _, found := options["nervatype"]; found {
		if _, found := options["ref_id"]; found {
			if options["ref_id"] == "" {
				options["ref_id"] = nil
			}
			sqlString = `select df.fieldname as fieldname, fv.id as fieldvalue_id
						from deffield df
						inner join groups nt on ((df.nervatype=nt.id) `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"nervatype", "ref_id"}, ` and (nt.groupvalue=%s))
						left join fieldvalue fv on ((df.fieldname=fv.fieldname) and (fv.deleted = 0) and (fv.ref_id = %s)) `, params)
			sqlString += whereString
			sqlString += `union select 'fieldtype_string' as fieldname, id as fieldvalue_id 
						from groups where (groupname='fieldtype') and (groupvalue='string')
					union select 'nervatype_id' as fieldname, id as fieldvalue_id 
						from groups where (groupname='nervatype') `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"nervatype"}, ` and (groupvalue = %s) `, params)
			sqlString += whereString
			return sqlString, params, nil
		}
		sqlString = `select df.fieldname, ft.groupvalue as fieldtype, df.addnew, df.visible
						from deffield df
						inner join groups nt on (df.nervatype=nt.id) `
		whereString, params = ds.getQueryKeyOption(options,
			SL{"nervatype"}, ` and (nt.groupvalue = %s )
						inner join groups ft on (df.fieldtype=ft.id)
						where df.deleted=0 `, params)
		sqlString += whereString
		return sqlString, params, nil
	}
	return sqlString, params, errors.New(ut.GetMessage("missing_fieldname"))
}

func (ds *SQLDriver) getIntegrity(options IM) (string, IL, error) {
	var sqlString, whereString string
	params := make(IL, 0)

	keys := map[string]func(options IM) (string, IL, error){
		"currency": func(options IM) (string, IL, error) {
			//(link), place,price,rate,trans
			sqlString = `select {CAS_INT}sum(co){CAE_INT} as count from (
				select count(place.id) as co from place
				inner join currency on (place.curr=currency.curr)
				where ((place.deleted=0) `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` and (currency.id=%s)) `, params)
			sqlString += whereString
			sqlString += `union select count(price.id) as co from price
				inner join currency on (price.curr=currency.curr) 
				where ((price.deleted=0) `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` and (currency.id=%s)) `, params)
			sqlString += whereString
			sqlString += `union select count(rate.id) as co from rate 
				inner join currency on (rate.curr=currency.curr)
				where ((rate.deleted=0) `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` and (currency.id=%s)) `, params)
			sqlString += whereString
			sqlString += `union select count(trans.id) as co from trans
				inner join currency on (trans.curr=currency.curr)
				where ((trans.deleted=0) `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` and (currency.id=%s))) foo `, params)
			sqlString += whereString
			return sqlString, params, nil
		},

		"customer": func(options IM) (string, IL, error) {
			//(address,contact), event,project,trans,link
			sqlString = `select {CAS_INT}sum(co){CAE_INT} as count from (
				select {CAS_INT}count(*){CAE_INT} as co  from trans where `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` (customer_id=%s) `, params)
			sqlString += whereString
			sqlString += `union select {CAS_INT}count(*){CAE_INT} as co  from project where `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` (customer_id=%s) `, params)
			sqlString += whereString
			sqlString += `union select {CAS_INT}count(*){CAE_INT} as co  from event
				inner join groups nt on ((event.nervatype=nt.id) and (nt.groupvalue='customer'))
				where (event.deleted=0) `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` and (event.ref_id=%s) `, params)
			sqlString += whereString
			sqlString += `union select {CAS_INT}count(*){CAE_INT} as co  from link
				where nervatype_2=(
					select id  from groups 
					where (groupname='nervatype') and (groupvalue='customer') `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` and (ref_id_2=%s))
				) foo `, params)
			sqlString += whereString
			return sqlString, params, nil
		},

		"deffield": func(options IM) (string, IL, error) {
			//fieldvalue
			sqlString = `select {CAS_INT}sum(co){CAE_INT} as count from (
			  select count(fieldvalue.id) as co from fieldvalue
				inner join deffield on (deffield.fieldname=fieldvalue.fieldname)
				where (fieldvalue.deleted=0) `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` and (deffield.id=%s)) foo `, params)
			sqlString += whereString
			return sqlString, params, nil
		},

		"employee": func(options IM) (string, IL, error) {
			//(address,contact), event,trans,log,link,ui_printqueue,ui_userconfig
			sqlString = `select {CAS_INT}sum(co){CAE_INT} as count from (
				select {CAS_INT}count(*){CAE_INT} as co from trans where `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` (employee_id=%s) `, params)
			sqlString += whereString
			sqlString += `union select {CAS_INT}count(*){CAE_INT} as co from trans where `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` (cruser_id=%s) `, params)
			sqlString += whereString
			sqlString += `union select {CAS_INT}count(*){CAE_INT} as co from log where `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` (employee_id=%s) `, params)
			sqlString += whereString
			sqlString += `union select {CAS_INT}count(*){CAE_INT} as co from ui_printqueue where `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` (employee_id=%s) `, params)
			sqlString += whereString
			sqlString += `union select {CAS_INT}count(*){CAE_INT} as co from ui_userconfig where `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` (employee_id=%s) `, params)
			sqlString += whereString
			sqlString += `union select {CAS_INT}count(*){CAE_INT} as co from event
				  inner join groups nt on ((event.nervatype=nt.id) and (nt.groupvalue='employee'))
				  where (event.deleted=0) and `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` (event.ref_id=%s) `, params)
			sqlString += whereString
			sqlString += `union select {CAS_INT}count(*){CAE_INT} as co from link 
					where (nervatype_2=(select id from groups 
						where (groupname='nervatype') and (groupvalue='employee') `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` and (ref_id_2=%s)))) foo `, params)
			sqlString += whereString
			return sqlString, params, nil
		},

		"groups": func(options IM) (string, IL, error) {
			//barcode,deffield,employee,event,rate,tool,trans,link
			sqlString = `select {CAS_INT}sum(co){CAE_INT} as count from (
				select {CAS_INT}count(*){CAE_INT} as co from groups 
				where groupname in ('nervatype', 'custtype', 'fieldtype', 'logstate', 'movetype', 'transtype', 
					'placetype', 'calcmode', 'protype', 'ratetype', 'direction', 'transtate', 
					'inputfilter', 'filetype', 'wheretype', 'aggretype') `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` and id = %s  `, params)
			sqlString += whereString
			sqlString += `union select {CAS_INT}count(*){CAE_INT} as co from link 
				where nervatype_2 = ( select id from groups where groupname = 'nervatype' and groupvalue = 'groups') and `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` ref_id_2 = %s `, params)
			sqlString += whereString
			sqlString += `union select {CAS_INT}count(*){CAE_INT} from barcode where `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` barcode.barcodetype = %s `, params)
			sqlString += whereString
			sqlString += `union select {CAS_INT}count(*){CAE_INT} from deffield where deffield.deleted = 0 `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` and deffield.subtype = %s `, params)
			sqlString += whereString
			sqlString += `union select {CAS_INT}count(*){CAE_INT} from employee where employee.deleted = 0 `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` and employee.usergroup = %s `, params)
			sqlString += whereString
			sqlString += `union select {CAS_INT}count(*){CAE_INT} from employee where employee.deleted = 0 `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` and employee.department = %s `, params)
			sqlString += whereString
			sqlString += `union select {CAS_INT}count(*){CAE_INT} from event where event.deleted = 0 `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` and event.eventgroup = %s `, params)
			sqlString += whereString
			sqlString += `union select {CAS_INT}count(*){CAE_INT} from rate where rate.deleted = 0 `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` and rate.rategroup = %s `, params)
			sqlString += whereString
			sqlString += `union select {CAS_INT}count(*){CAE_INT} from tool where tool.deleted = 0 `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` and tool.toolgroup = %s `, params)
			sqlString += whereString
			sqlString += `union select {CAS_INT}count(*){CAE_INT} from trans where trans.deleted = 0 `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` and trans.department = %s ) foo `, params)
			sqlString += whereString
			return sqlString, params, nil
		},

		"place": func(options IM) (string, IL, error) {
			//"place":
			//(address,contact), event,movement,place,rate,trans,link
			sqlString = `select {CAS_INT}sum(co){CAE_INT} as count from (
				select {CAS_INT}count(*){CAE_INT} as co from event inner join groups nt on event.nervatype = nt.id and nt.groupvalue = 'place' 
				where event.deleted = 0 `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` and event.ref_id = %s `, params)
			sqlString += whereString
			sqlString += `union select {CAS_INT}count(*){CAE_INT} as co from link where nervatype_2 = ( 
					select id from groups where groupname = 'nervatype' and groupvalue = 'place') `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` and ref_id_2 = %s `, params)
			sqlString += whereString
			sqlString += `union select {CAS_INT}count(*){CAE_INT} as co from movement where movement.deleted = 0 `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` and movement.place_id = %s `, params)
			sqlString += whereString
			sqlString += `union select {CAS_INT}count(*){CAE_INT} as co from rate where rate.deleted = 0 `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` and rate.place_id = %s `, params)
			sqlString += whereString
			sqlString += `union select {CAS_INT}count(*){CAE_INT} as co from trans where trans.deleted = 0 `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` and trans.place_id = %s ) foo `, params)
			sqlString += whereString
			return sqlString, params, nil
		},

		"product": func(options IM) (string, IL, error) {
			//address,barcode,contact,event,item,movement,price,tool,link
			sqlString = `select {CAS_INT}sum(co){CAE_INT} as count from (
				select {CAS_INT}count(*){CAE_INT} as co from event inner join groups nt on event.nervatype = nt.id and nt.groupvalue = 'product' 
				where event.deleted = 0 `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` and event.ref_id = %s `, params)
			sqlString += whereString
			sqlString += `union select {CAS_INT}count(*){CAE_INT} as co from address inner join groups nt on address.nervatype = nt.id and nt.groupvalue = 'product' 
				where address.deleted = 0 `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` and address.ref_id = %s `, params)
			sqlString += whereString
			sqlString += `union select {CAS_INT}count(*){CAE_INT} as co from contact inner join groups nt on contact.nervatype = nt.id and nt.groupvalue = 'product' 
				where contact.deleted = 0 `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` and contact.ref_id = %s `, params)
			sqlString += whereString
			sqlString += `union select {CAS_INT}count(*){CAE_INT} as co from link where nervatype_2 = ( 
					select id from groups where groupname = 'nervatype' and groupvalue = 'product') `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` and ref_id_2 = %s `, params)
			sqlString += whereString
			sqlString += `union select {CAS_INT}count(*){CAE_INT} as co from barcode where `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` barcode.product_id = %s `, params)
			sqlString += whereString
			sqlString += `union select {CAS_INT}count(*){CAE_INT} as co from movement where movement.deleted = 0  `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` and movement.product_id = %s `, params)
			sqlString += whereString
			sqlString += `union select {CAS_INT}count(*){CAE_INT} as co from item where item.deleted = 0  `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` and item.product_id = %s `, params)
			sqlString += whereString
			sqlString += `union select {CAS_INT}count(*){CAE_INT} as co from price where price.deleted = 0  `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` and price.product_id = %s `, params)
			sqlString += whereString
			sqlString += `union select {CAS_INT}count(*){CAE_INT} as co from tool where tool.deleted = 0  `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` and tool.product_id = %s ) foo `, params)
			sqlString += whereString
			return sqlString, params, nil
		},

		"project": func(options IM) (string, IL, error) {
			//(address,contact), event,trans,link
			sqlString = `select {CAS_INT}sum(co){CAE_INT} as count from (
				select {CAS_INT}count(*){CAE_INT} as co from trans where `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` project_id = %s `, params)
			sqlString += whereString
			sqlString += `union select {CAS_INT}count(*){CAE_INT} as co from event inner join groups nt on event.nervatype = nt.id and nt.groupvalue = 'project' 
				where event.deleted = 0 `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` and event.ref_id = %s `, params)
			sqlString += whereString
			sqlString += `union select {CAS_INT}count(*){CAE_INT} as co from link where nervatype_2 = ( 
					select id from groups where groupname = 'nervatype' and groupvalue = 'project') `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` and ref_id_2 = %s ) foo `, params)
			sqlString += whereString
			return sqlString, params, nil
		},

		"tax": func(options IM) (string, IL, error) {
			//item,product
			sqlString = `select {CAS_INT}sum(co){CAE_INT} as count from (
				select {CAS_INT}count(*){CAE_INT} as co from item where item.deleted = 0 `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` and item.tax_id = %s `, params)
			sqlString += whereString
			sqlString += `union select {CAS_INT}count(*){CAE_INT} as co from product where product.deleted = 0 `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` and product.tax_id = %s) foo `, params)
			sqlString += whereString
			return sqlString, params, nil
		},

		"tool": func(options IM) (string, IL, error) {
			//(address,contact), event,movement,link
			sqlString = `select {CAS_INT}sum(co){CAE_INT} as count from (
				select {CAS_INT}count(*){CAE_INT} as co from movement where `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` tool_id = %s `, params)
			sqlString += whereString
			sqlString += `union select {CAS_INT}count(*){CAE_INT} as co from event inner join groups nt on event.nervatype = nt.id and nt.groupvalue = 'tool' 
				where event.deleted = 0 `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` and event.ref_id = %s `, params)
			sqlString += whereString
			sqlString += `union select {CAS_INT}count(*){CAE_INT} as co from link where nervatype_2 = ( 
					select id from groups where groupname = 'nervatype' and groupvalue = 'tool') `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` and ref_id_2 = %s) foo `, params)
			sqlString += whereString
			return sqlString, params, nil
		},

		"trans": func(options IM) (string, IL, error) {
			//(address,contact), event,link
			sqlString = `select {CAS_INT}sum(co){CAE_INT} as count from (
				select {CAS_INT}count(*){CAE_INT} as co from event inner join groups nt on event.nervatype = nt.id and nt.groupvalue = 'trans' 
				where event.deleted = 0 `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` and event.ref_id = %s `, params)
			sqlString += whereString
			sqlString += `union select {CAS_INT}count(*){CAE_INT} as co from link where nervatype_2 = ( 
					select id from groups where groupname = 'nervatype' and groupvalue = 'trans') `
			whereString, params = ds.getQueryKeyOption(options,
				SL{"ref_id"}, ` and ref_id_2 = %s) foo `, params)
			sqlString += whereString
			return sqlString, params, nil
		},
	}

	if _, found := keys[options["nervatype"].(string)]; found {
		if _, found := options["ref_id"]; found {
			return keys[options["nervatype"].(string)](options)
		}
	}
	return sqlString, params, errors.New(ut.GetMessage("integrity_error"))
}

//QueryKey - complex data queries
func (ds *SQLDriver) QueryKey(options IM, trans interface{}) (result []IM, err error) {
	result = []IM{}
	sqlString := ""
	params := make(IL, 0)

	keys := map[string]func(options IM) (string, IL, error){
		"user": func(options IM) (string, IL, error) {
			params := IL{options["username"], options["username"]}
			sqlString =
				`select e.id, e.username, e.empnumber, e.usergroup, ug.groupvalue as scope, 
				case when dg.groupvalue is null then '' else dg.groupvalue end as department
				from employee e
				inner join groups ug on e.usergroup = ug.id
				left join groups dg on e.department = dg.id
				where e.deleted = 0 and e.inactive = 0 and ((e.username = ` + ds.getPrmString(1) + `) or (e.registration_key = ` + ds.getPrmString(2) + `))`
			return sqlString, params, nil
		},

		"metadata": func(options IM) (string, IL, error) {
			params := IL{options["nervatype"]}
			values, inPrm := ds.splitInParams(options["ids"].(string), len(params))
			params = append(params, values...)
			sqlString =
				`select fv.*, ft.groupvalue as fieldtype from fieldvalue fv 
				inner join deffield df on fv.fieldname = df.fieldname
				inner join groups nt on df.nervatype = nt.id 
				inner join groups ft on df.fieldtype = ft.id 
				where fv.deleted = 0 and df.deleted = 0 and nt.groupvalue = ` + ds.getPrmString(1) + `
					and fv.ref_id in (` + strings.Join(inPrm, ",") + `) 
				order by fv.fieldname, fv.id `
			return sqlString, params, nil
		},

		"post_transtype": func(options IM) (string, IL, error) {
			params := IL{options["trans_id"], options["transtype_key"], options["transtype_id"],
				options["customer_id"], options["custnumber"]}
			sqlString =
				`select 'trans' as rtype, tt.groupvalue as transtype, c.custnumber
				from trans t
				inner join groups tt on t.transtype=tt.id
				left join customer c on t.customer_id=c.id
				where t.id = ` + ds.getPrmString(1) + `
				union select 'groups' as rtype, groupvalue as transtype, null 
				from groups
				where groupname='transtype' 
					and (groupvalue=` + ds.getPrmString(2) + ` or id=` + ds.getPrmString(3) + `)
				union select 'customer' as rtype, null as transtype, custnumber
				from customer
				where id=` + ds.getPrmString(4) + ` or custnumber=` + ds.getPrmString(5) + ` `
			return sqlString, params, nil
		},

		"default_report": func(options IM) (string, IL, error) {
			params := IL{}
			whereString := ""
			sqlString =
				`select r.*, ft.groupvalue as reptype
				from ui_report r
				inner join groups ft on r.filetype=ft.id `
			if _, found := options["nervatype"]; found {
				params = append(params, options["nervatype"])
				sqlString += ` inner join groups nt on r.nervatype=nt.id and nt.groupvalue=` + ds.getPrmString(len(params))
				if options["nervatype"] == "trans" {
					params = append(params, options["transtype"], options["direction"])
					sqlString += ` inner join groups tt on r.transtype=tt.id and tt.groupvalue=` + ds.getPrmString(len(params)) + `
            inner join groups dir on r.direction=dir.id and dir.groupvalue=` + ds.getPrmString(len(params))
				}
			}
			whereString, params = ds.getQueryKeyOption(options, SL{"reportkey"}, ` where r.reportkey=%s`, params)
			sqlString += whereString
			whereString, params = ds.getQueryKeyOption(options, SL{"report_id"}, ` where r.id=%s`, params)
			sqlString += whereString
			return sqlString, params, nil
		},

		"listprice": func(options IM) (string, IL, error) {
			params := IL{options["vendorprice"], options["product_id"], options["posdate"], options["posdate"], options["curr"], options["qty"]}
			sqlString =
				`select min(p.pricevalue) as mp 
				from price p 
				left join link l on l.ref_id_1 = p.id and l.nervatype_1 = ( 
					select id from groups 
					where groupname = 'nervatype' and groupvalue = 'price') and l.deleted = 0 
				where p.deleted = 0 and p.discount is null and p.pricevalue <> 0 
					and l.ref_id_2 is null and p.vendorprice = ` + ds.getPrmString(1) + ` and p.product_id = ` + ds.getPrmString(2) + ` 
					and p.validfrom <= ` + ds.getPrmString(3) + ` and ( p.validto >= ` + ds.getPrmString(4) + ` or 
					p.validto is null) and p.curr = ` + ds.getPrmString(5) + ` and p.qty <= ` + ds.getPrmString(6)
			return sqlString, params, nil
		},

		"custprice": func(options IM) (string, IL, error) {
			params := IL{options["vendorprice"], options["product_id"], options["posdate"], options["posdate"],
				options["curr"], options["qty"], options["customer_id"]}
			sqlString =
				`select min(p.pricevalue) as mp 
				from price p 
				inner join link l on l.ref_id_1 = p.id 
					and l.nervatype_1 = (select id from groups where groupname = 'nervatype' and groupvalue = 'price') 
					and l.nervatype_2 = (select id from groups where groupname = 'nervatype' and groupvalue = 'customer') 
					and l.deleted = 0 
				where p.deleted = 0 and p.discount is null and p.pricevalue <> 0 
					and p.vendorprice = ` + ds.getPrmString(1) + ` and p.product_id = ` + ds.getPrmString(2) + ` and p.validfrom <= ` + ds.getPrmString(3) + ` 
					and ( p.validto >= ` + ds.getPrmString(4) + ` or p.validto is null) and p.curr = ` + ds.getPrmString(5) + ` 
					and p.qty <= ` + ds.getPrmString(6) + ` and l.ref_id_2 = ` + ds.getPrmString(7)
			return sqlString, params, nil
		},

		"grouprice": func(options IM) (string, IL, error) {
			params := IL{options["customer_id"], options["vendorprice"], options["product_id"], options["posdate"], options["posdate"],
				options["curr"], options["qty"]}
			sqlString =
				`select min(p.pricevalue) as mp 
				from price p 
				inner join link l on l.ref_id_1 = p.id and l.deleted = 0 
					and l.nervatype_1 = (select id from groups where groupname = 'nervatype' and groupvalue = 'price') 
					and l.nervatype_2 = (select id from groups where groupname = 'nervatype' and groupvalue = 'groups') 
				inner join groups g on g.id = l.ref_id_2 
					and g.id in (select l.ref_id_2 from link l where l.deleted = 0 
					and l.nervatype_1 = (select id from groups where groupname = 'nervatype' and groupvalue = 'customer') 
					and l.nervatype_2 = (select id from groups where groupname = 'nervatype' and groupvalue = 'groups') 
					and l.ref_id_1 = ` + ds.getPrmString(1) + `) 
				where p.deleted = 0 and p.discount is null and p.pricevalue <> 0 and p.vendorprice = ` + ds.getPrmString(2) + ` 
					and p.product_id = ` + ds.getPrmString(3) + ` and p.validfrom <= ` + ds.getPrmString(4) + ` 
					and ( p.validto >= ` + ds.getPrmString(5) + ` or p.validto is null) 
					and p.curr = ` + ds.getPrmString(6) + ` and p.qty <= ` + ds.getPrmString(7)
			return sqlString, params, nil
		},

		"data_audit": func(options IM) (string, IL, error) {
			params := IL{options["id"]}
			sqlString =
				`select tt.groupvalue as transfilter 
				from employee e inner join link l on l.ref_id_1 = e.usergroup and l.deleted = 0 
				inner join groups nt1 on l.nervatype_1 = nt1.id and nt1.groupname = 'nervatype' and nt1.groupvalue = 'groups' 
				inner join groups nt2 on l.nervatype_2 = nt2.id and nt2.groupname = 'nervatype' and nt2.groupvalue = 'groups' 
				inner join groups tt on l.ref_id_2 = tt.id where e.id = ` + ds.getPrmString(1)
			return sqlString, params, nil
		},

		"object_audit": func(options IM) (string, IL, error) {
			whereString := ""
			params := IL{options["usergroup"]}
			sqlString =
				`select inf.groupvalue as inputfilter from ui_audit a 
				inner join groups inf on a.inputfilter = inf.id 
				inner join groups nt on a.nervatype = nt.id 
				left join groups st on a.subtype = st.id 
				where (a.usergroup = ` + ds.getPrmString(1) + `) `
			whereString, params = ds.getQueryKeyOption(options, SL{"subtype"}, ` and a.subtype=%s`, params)
			sqlString += whereString
			whereString, params = ds.getQueryKeySplit(options, "subtypeIn", ` and a.subtype in (%s) `, params)
			sqlString += whereString
			whereString, params = ds.getQueryKeyOption(options, SL{"transtype"}, ` and st.groupvalue=%s`, params)
			sqlString += whereString
			whereString, params = ds.getQueryKeySplit(options, "transtypeIn", ` and st.groupvalue in (%s) `, params)
			sqlString += whereString
			whereString, params = ds.getQueryKeyOption(options, SL{"nervatype"}, ` and a.nervatype=%s`, params)
			sqlString += whereString
			whereString, params = ds.getQueryKeySplit(options, "nervatypeIn", ` and a.nervatype in (%s) `, params)
			sqlString += whereString
			whereString, params = ds.getQueryKeyOption(options, SL{"groupvalue"}, ` and nt.groupvalue=%s`, params)
			sqlString += whereString
			whereString, params = ds.getQueryKeySplit(options, "groupvalueIn", ` and nt.groupvalue in (%s) `, params)
			sqlString += whereString
			return sqlString, params, nil
		},

		"delete_deffields": func(options IM) (string, IL, error) {
			sqlString =
				`select fv.id as id from deffield df 
          inner join groups nt on ((df.nervatype=nt.id) 
					  and (nt.groupvalue=%s))
          inner join fieldvalue fv on ((df.fieldname=fv.fieldname) 
					  and (fv.deleted=0) and (fv.ref_id=%s))`
			sqlString, params := ds.getQueryKeyOption(options, SL{"nervatype", "ref_id"}, sqlString, params)
			return sqlString, params, nil
		},

		"id->refnumber": func(options IM) (string, IL, error) {
			return ds.getID2Refnumber(options)
		},

		"refnumber->id": func(options IM) (string, IL, error) {
			return ds.getRefnumber2ID(options)
		},

		"integrity": func(options IM) (string, IL, error) {
			return ds.getIntegrity(options)
		},

		"update_deffields": func(options IM) (string, IL, error) {
			return ds.getUpdateDeffields(options)
		},
	}

	if _, found := keys[options["qkey"].(string)]; found {
		sqlString, params, err = keys[options["qkey"].(string)](options)
		if err != nil {
			return result, err
		}
	} else {
		return result, errors.New(ut.GetMessage("missing_fieldname"))
	}

	//print(sqlString)
	return ds.QuerySQL(sqlString, params, trans)
}
