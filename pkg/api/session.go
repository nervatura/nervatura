package api

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	cu "github.com/nervatura/component/pkg/util"
	drv "github.com/nervatura/nervatura/v6/pkg/driver"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

const (
	SessionMethodMemory   = "mem"
	SessionMethodFile     = "file"
	SessionMethodDatabase = "db"

	SessionErrorDriver = "missing_driver"
	SessionErrorData   = "nodata"
)

type memStore struct {
	Session any
	Stamp   time.Time
}

type SessionConfig struct {
	FileDir     string
	DbConn      string
	DbTable     string
	Cleaning    bool
	CleanExp    int64
	CleanPeriod int64
	Method      string
}

// SessionService implements the session service
type SessionService struct {
	memSession      map[string]memStore
	method          string
	cleaningStamp   time.Time
	Config          SessionConfig
	Conn            DataDriver
	GetMessage      func(lang string, key string) string
	CreateDir       func(name string, perm fs.FileMode) error
	CreateFile      func(name string) (*os.File, error)
	ReadDir         func(name string) ([]fs.DirEntry, error)
	ReadFile        func(name string) ([]byte, error)
	FileStat        func(name string) (fs.FileInfo, error)
	ConvertToByte   func(data interface{}) ([]byte, error)
	ConvertFromByte func(data []byte, result interface{}) error
	RemoveFile      func(name string) error
}

func NewSession(config cu.IM, method, alias string) *SessionService {
	return &SessionService{
		Config: SessionConfig{
			FileDir:     cu.ToString(config["NT_SESSION_DIR"], ""),
			DbConn:      cu.ToString(config["NT_ALIAS_"+strings.ToUpper(alias)], os.Getenv("NT_ALIAS_"+strings.ToUpper(alias))),
			DbTable:     cu.ToString(config["NT_SESSION_TABLE"], ""),
			Cleaning:    true,
			CleanExp:    cu.ToInteger(config["NT_SESSION_EXP"], 0),
			CleanPeriod: 6,
			Method:      method,
		},
		Conn:            &drv.SQLDriver{Config: config},
		GetMessage:      ut.GetMessage,
		CreateDir:       os.Mkdir,
		CreateFile:      os.Create,
		ReadDir:         os.ReadDir,
		ReadFile:        os.ReadFile,
		FileStat:        os.Stat,
		ConvertToByte:   cu.ConvertToByte,
		ConvertFromByte: cu.ConvertFromByte,
		RemoveFile:      os.Remove,
	}
}

func (ses *SessionService) sessionMethod() string {
	ses.method = cu.ToString(ses.method, ses.Config.Method)
	if ses.method == "" {
		ses.method = SessionMethodMemory
		if ses.Config.DbConn != "" {
			ses.method = SessionMethodDatabase
		}
		if ses.Config.FileDir != "" {
			ses.method = SessionMethodFile
		}
	}
	return ses.method
}

func (ses *SessionService) saveFileSession(fileName string, data any) (err error) {
	filePath := fileName + ".json"
	if ses.Config.FileDir != "" {
		if _, err = ses.FileStat(ses.Config.FileDir); errors.Is(err, os.ErrNotExist) {
			err = ses.CreateDir(ses.Config.FileDir, os.ModePerm)
			if err != nil {
				return err
			}
		}
		filePath = filepath.FromSlash(fmt.Sprintf(`%s/%s.json`, ses.Config.FileDir, fileName))
	}

	var sessionFile *os.File
	sessionFile, err = ses.CreateFile(filePath)
	if err == nil {
		var bin []byte
		bin, err = ses.ConvertToByte(data)
		if err == nil {
			sessionFile.Write(bin)
		}
	}
	defer sessionFile.Close()
	return err
}

func (ses *SessionService) getFilePath(fileName string) (filePath string) {
	filePath = fileName + ".json"
	if ses.Config.FileDir != "" {
		filePath = filepath.FromSlash(fmt.Sprintf(`%s/%s.json`, ses.Config.FileDir, fileName))
	}
	return filePath
}

func (ses *SessionService) loadFileSession(fileName string, data any) (err error) {
	filePath := ses.getFilePath(fileName)
	var sessionFile []byte
	sessionFile, err = ses.ReadFile(filePath)
	if err == nil {
		err = ses.ConvertFromByte(sessionFile, &data)
	}
	return err
}

func (ses *SessionService) deleteFileSession(fileName string) (err error) {
	filePath := ses.getFilePath(fileName)
	return ses.RemoveFile(filePath)
}

func (ses *SessionService) cleaningFileSession(exp time.Time) (err error) {
	var files []fs.DirEntry
	if ses.Config.FileDir != "" {
		if files, err = ses.ReadDir(ses.Config.FileDir); err == nil {
			for _, f := range files {
				var info fs.FileInfo
				if info, err = f.Info(); err == nil && info.ModTime().Before(exp) {
					ses.deleteFileSession(f.Name())
				}
			}
		}
	}
	return err
}

func (ses *SessionService) message(lang string, key string) string {
	if ses.GetMessage != nil {
		return ses.GetMessage(lang, key)
	}
	return key
}

func (ses *SessionService) checkSessionTable() (err error) {
	if ses.Config.DbConn == "" {
		return errors.New(ses.message("en", SessionErrorDriver))
	}
	sessionTable := cu.ToString(ses.Config.DbTable, "session")
	var sqlString string
	if err = ses.Conn.CreateConnection("session", ses.Config.DbConn); err == nil {
		var rows []cu.IM
		engine := ses.Conn.Connection().Engine
		if engine == "sqlite" {
			sqlString = fmt.Sprintf("select name from sqlite_master where name = '%s' ", sessionTable)
		} else {
			sqlString = fmt.Sprintf("select table_name from information_schema.tables where table_name = '%s' ", sessionTable)
		}
		rows, err = ses.Conn.QuerySQL(sqlString, []interface{}{}, nil)
		if err == nil && len(rows) == 0 {
			jsonType := cu.SM{
				"sqlite": "TEXT", "postgres": "TEXT", "mysql": "TEXT", "mssql": "NVARCHAR(MAX)",
			}
			intType := cu.SM{
				"sqlite": "INTEGER", "postgres": "INTEGER", "mysql": "INTEGER", "mssql": "INT",
			}
			if engine == "mysql" {
				sqlString = fmt.Sprintf(
					"CREATE TABLE %s ( id VARCHAR(255) NOT NULL, value TEXT, stamp INTEGER, PRIMARY KEY (id) );",
					sessionTable)
			} else {
				sqlString = fmt.Sprintf(
					"CREATE TABLE %s ( id VARCHAR(255) NOT NULL PRIMARY KEY, value %s, stamp %s );",
					sessionTable, jsonType[engine], intType[engine])
			}
			_, err = ses.Conn.QuerySQL(sqlString, []interface{}{}, nil)
		}
	}
	return err
}

func (ses *SessionService) getDbRows(rowID string) (rows []cu.IM, err error) {
	sessionTable := cu.ToString(ses.Config.DbTable, "session")
	sqlString := fmt.Sprintf("SELECT * FROM %s WHERE id='%s'", sessionTable, rowID)
	return ses.Conn.QuerySQL(sqlString, []interface{}{}, nil)
}

func (ses *SessionService) saveDbSession(sessionID string, data any) (err error) {
	sessionTable := cu.ToString(ses.Config.DbTable, "session")
	if err = ses.Conn.CreateConnection("session", ses.Config.DbConn); err == nil {
		var bin []byte
		bin, err = ses.ConvertToByte(data)
		if err == nil {
			var sqlString string = fmt.Sprintf(
				"INSERT INTO %s(id, value, stamp) VALUES('%s', '%s', %d)",
				sessionTable, sessionID, bin, time.Now().Unix())
			var rows []cu.IM
			if rows, err = ses.getDbRows(sessionID); err == nil && (len(rows) > 0) {
				sqlString = fmt.Sprintf(
					"UPDATE %s SET value='%s' WHERE id='%s'", sessionTable, bin, sessionID)
			}
			_, err = ses.Conn.QuerySQL(sqlString, []interface{}{}, nil)
		}
	}
	return err
}

func (ses *SessionService) deleteDbSession(sessionID string) (err error) {
	sessionTable := cu.ToString(ses.Config.DbTable, "session")
	if err = ses.Conn.CreateConnection("session", ses.Config.DbConn); err == nil {
		sqlString := fmt.Sprintf(
			"DELETE FROM %s WHERE ID='%s'", sessionTable, sessionID)
		_, err = ses.Conn.QuerySQL(sqlString, []interface{}{}, nil)
	}
	return err
}

func (ses *SessionService) cleaningDbSession(exp time.Time) (err error) {
	sessionTable := cu.ToString(ses.Config.DbTable, "session")
	if err = ses.Conn.CreateConnection("session", ses.Config.DbConn); err == nil {
		sqlString := fmt.Sprintf(
			"DELETE FROM %s WHERE stamp<%d", sessionTable, exp.Unix())
		_, err = ses.Conn.QuerySQL(sqlString, []interface{}{}, nil)
	}
	return err
}

func (ses *SessionService) loadDbSession(sessionID string, data any) (err error) {
	if err = ses.Conn.CreateConnection("session", ses.Config.DbConn); err == nil {
		var rows []cu.IM
		if rows, err = ses.getDbRows(sessionID); err == nil {
			if len(rows) == 0 {
				return errors.New(ses.message("en", SessionErrorData))
			}
			jdata := []byte(cu.ToString(rows[0]["value"], ""))
			err = ses.ConvertFromByte(jdata, &data)
		}
	}
	return err
}

func (ses *SessionService) loadMemSession(sessionID string) (result any, err error) {
	if result, found := ses.memSession[sessionID]; found {
		return result.Session, nil
	}
	return result, errors.New(ses.message("en", SessionErrorData))
}

func (ses *SessionService) saveMemSession(sessionID string, data any) (err error) {
	if ses.memSession == nil {
		ses.memSession = make(map[string]memStore)
	}
	ses.memSession[sessionID] = memStore{Session: data, Stamp: time.Now()}
	return nil
}

func (ses *SessionService) deleteMemSession(sessionID string) (err error) {
	delete(ses.memSession, sessionID)
	return nil
}

func (ses *SessionService) cleaningMemSession(exp time.Time) (err error) {
	for key := range ses.memSession {
		if _, found := ses.memSession[key]; found && ses.memSession[key].Stamp.Before(exp) {
			ses.deleteMemSession(key)
		}
	}
	return nil
}

func (ses *SessionService) SaveSession(sessionID string, data any) {
	save := map[string]func(){
		SessionMethodFile: func() {
			if err := ses.saveFileSession(sessionID, data); err != nil {
				ses.method = SessionMethodMemory
				ses.saveMemSession(sessionID, data)
			}
		},
		SessionMethodDatabase: func() {
			var err error
			if err = ses.checkSessionTable(); err == nil {
				err = ses.saveDbSession(sessionID, data)
			}
			if err != nil {
				ses.method = SessionMethodMemory
				ses.saveMemSession(sessionID, data)
			}
		},
		SessionMethodMemory: func() {
			ses.saveMemSession(sessionID, data)
		},
	}
	save[ses.sessionMethod()]()
	ses.CleaningSession()
}

func (ses *SessionService) LoadSession(sessionID string, data any) (result any, err error) {
	switch ses.sessionMethod() {
	case SessionMethodFile:
		err = ses.loadFileSession(sessionID, data)
		return data, err
	case SessionMethodDatabase:
		err = ses.loadDbSession(sessionID, &data)
		return data, err
	default:
		return ses.loadMemSession(sessionID)
	}
}

func (ses *SessionService) DeleteSession(sessionID string) (err error) {
	switch ses.sessionMethod() {
	case SessionMethodFile:
		return ses.deleteFileSession(sessionID)
	case SessionMethodDatabase:
		return ses.deleteDbSession(sessionID)
	default:
		return ses.deleteMemSession(sessionID)
	}
}

func (ses *SessionService) CleaningSession() (err error) {
	if ses.Config.Cleaning {
		period := cu.ToFloat(ses.Config.CleanPeriod, 6)
		if ses.cleaningStamp.IsZero() || time.Now().Before(ses.cleaningStamp.Add(time.Duration(period)*time.Hour)) {
			exp := time.Now().Add((-time.Duration(cu.ToFloat(ses.Config.CleanExp, 1))*time.Hour + 1))
			methodCl := map[string]func() error{
				SessionMethodFile: func() error {
					return ses.cleaningFileSession(exp)
				},
				SessionMethodDatabase: func() error {
					return ses.cleaningDbSession(exp)
				},
				SessionMethodMemory: func() error {
					return ses.cleaningMemSession(exp)
				},
			}
			if err = methodCl[ses.sessionMethod()](); err == nil {
				ses.cleaningStamp = time.Now()
			}
		}
	}
	return nil
}
