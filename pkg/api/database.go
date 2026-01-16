package api

import (
	"bytes"
	"fmt"
	"log/slog"
	"strings"
	"time"

	cu "github.com/nervatura/component/pkg/util"
	st "github.com/nervatura/nervatura/v6/static"
)

func createDatabaseReports(ds *DataStore, reportDir string, log []cu.SM) (logData []cu.SM) {
	logData = append(log, cu.SM{"stamp": time.Now().Format(time.RFC3339), "state": "info", "message": "installing reports"})
	ireports := []string{}
	if reports, err := ds.ReportList(reportDir, ""); err == nil {
		for _, report := range reports {
			reportKey := cu.ToString(report["report_key"], "")
			if _, err = ds.ReportInstall(reportKey, reportDir); err != nil {
				logData = append(logData, cu.SM{"stamp": time.Now().Format(time.RFC3339), "state": "error", "message": fmt.Sprintf("failed to install report %s", reportKey)})
				return logData
			}
			ireports = append(ireports, reportKey)
		}
		logData = append(logData, cu.SM{"stamp": time.Now().Format(time.RFC3339), "state": "info", "message": fmt.Sprintf("installed reports: %s", strings.Join(ireports, ", "))})
	}
	return logData
}

func createDatabaseUpdate(ds *DataStore, trans interface{}, fileName, engine, pathType string) (err error) {
	var sqlData []byte
	if sqlData, err = st.Store.ReadFile(fmt.Sprintf("%s/%s/%s.sql", pathType, engine, fileName)); err != nil {
		return err
	}

	sqlString := string(sqlData)
	if err = ds.Db.UpdateSQL(sqlString, trans); err != nil {
		ds.Db.RollbackTransaction(trans)
		return err
	}
	return nil
}

func CreateDatabase(options cu.IM, config cu.IM) (logData []cu.SM) {
	var err error
	alias := cu.ToString(options["alias"], "")
	demo := cu.ToBoolean(options["demo"], false)
	reportDisabled := cu.ToBoolean(options["report_disabled"], false)
	reportDir := cu.ToString(options["report_dir"], cu.ToString(config["NT_REPORT_DIR"], ""))
	logData = []cu.SM{
		{"stamp": time.Now().Format(time.RFC3339), "state": "info", "message": "start creating database"},
	}

	ds := NewDataStore(config, alias, slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)))
	if err = ds.checkConnection(); err != nil {
		logData = append(logData, cu.SM{"stamp": time.Now().Format(time.RFC3339), "state": "error", "message": "failed to check connection"})
		return logData
	}
	conn := ds.Db.Connection()
	defer ds.Db.CloseConnection()

	var trans interface{}
	if trans, err = ds.Db.BeginTransaction(); err != nil {
		logData = append(logData, cu.SM{"stamp": time.Now().Format(time.RFC3339), "state": "error", "message": "failed to begin transaction"})
		return logData
	}

	dbcs := []string{"drop", "schema", "data"}
	if demo {
		dbcs = append(dbcs, "demo")
	}
	msg := cu.SM{
		"drop_info":    "cleaning up database",
		"drop_error":   "failed to clean up database",
		"schema_info":  "creating schema",
		"schema_error": "failed to create schema",
		"data_info":    "creating data",
		"data_error":   "failed to create data",
		"demo_info":    "creating demo data",
		"demo_error":   "failed to create demo data",
	}

	for _, dbc := range dbcs {
		logData = append(logData, cu.SM{"stamp": time.Now().Format(time.RFC3339), "state": "info", "message": msg[dbc+"_info"]})
		if err = createDatabaseUpdate(ds, trans, dbc, conn.Engine, "store"); err != nil {
			logData = append(logData, cu.SM{"stamp": time.Now().Format(time.RFC3339), "state": "error", "message": msg[dbc+"_error"]})
			return logData
		}
	}

	logData = append(logData, cu.SM{"stamp": time.Now().Format(time.RFC3339), "state": "info", "message": "committing transaction"})
	ds.Db.CommitTransaction(trans)

	if !reportDisabled {
		logData = createDatabaseReports(ds, reportDir, logData)
	}

	logData = append(logData, cu.SM{"stamp": time.Now().Format(time.RFC3339), "state": "info", "message": "database successfully created"})

	return logData
}

func UpgradeDatabase(options cu.IM, config cu.IM) (logData []cu.SM) {
	var err error
	alias := cu.ToString(options["alias"], "")
	logData = []cu.SM{
		{"stamp": time.Now().Format(time.RFC3339), "state": "info", "message": "start creating database"},
	}

	ds := NewDataStore(config, alias, slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)))
	if err = ds.checkConnection(); err != nil {
		logData = append(logData, cu.SM{"stamp": time.Now().Format(time.RFC3339), "state": "error", "message": "failed to check connection"})
		return logData
	}
	conn := ds.Db.Connection()
	defer ds.Db.CloseConnection()

	var trans interface{}
	if trans, err = ds.Db.BeginTransaction(); err != nil {
		logData = append(logData, cu.SM{"stamp": time.Now().Format(time.RFC3339), "state": "error", "message": "failed to begin transaction"})
		return logData
	}

	dbcs := []string{"backup", "schema", "insert", "clean"}
	msg := cu.SM{
		"backup_info":  "rename and delete old database structure",
		"backup_error": "failed to rename and delete old database structure",
		"schema_info":  "create new database structure",
		"schema_error": "failed to create new database structure",
		"insert_info":  "transfer data from old to new database structure",
		"insert_error": "failed to transfer data from old to new database structure",
		"clean_info":   "clean up old database structure",
		"clean_error":  "failed to clean up old database structure",
	}
	pathType := func(pc string) string {
		if pc == "schema" {
			return "store"
		}
		return "upgrade"
	}

	for _, dbc := range dbcs {
		logData = append(logData, cu.SM{"stamp": time.Now().Format(time.RFC3339), "state": "info", "message": msg[dbc+"_info"]})
		if err = createDatabaseUpdate(ds, trans, dbc, conn.Engine, pathType(dbc)); err != nil {
			logData = append(logData, cu.SM{"stamp": time.Now().Format(time.RFC3339), "state": "error", "message": msg[dbc+"_error"]})
			return logData
		}
	}

	logData = append(logData, cu.SM{"stamp": time.Now().Format(time.RFC3339), "state": "info", "message": "committing transaction"})
	ds.Db.CommitTransaction(trans)

	logData = append(logData, cu.SM{"stamp": time.Now().Format(time.RFC3339), "state": "info", "message": "database successfully upgraded"})

	return logData
}
