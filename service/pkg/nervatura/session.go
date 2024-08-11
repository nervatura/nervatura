package nervatura

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	ut "github.com/nervatura/nervatura/service/pkg/utils"
)

// SessionService implements the session service
type SessionService struct {
	Config          IM
	Conn            DataDriver
	Method          string
	MemSession      IM
	CreateDir       func(name string, perm fs.FileMode) error
	CreateFile      func(name string) (*os.File, error)
	ReadFile        func(name string) ([]byte, error)
	FileStat        func(name string) (fs.FileInfo, error)
	ConvertToByte   func(data interface{}) ([]byte, error)
	ConvertFromByte func(data []byte, result interface{}) error
	RemoveFile      func(name string) error
}

func (ses *SessionService) saveFileSession(fileName string, data any) (err error) {
	filePath := fileName + ".json"
	sessionDir := ut.ToString(ses.Config["NT_SESSION_DIR"], "")
	if sessionDir != "" {
		if _, err = ses.FileStat(sessionDir); errors.Is(err, os.ErrNotExist) {
			err = ses.CreateDir(sessionDir, os.ModePerm)
			if err != nil {
				return err
			}
		}
		filePath = filepath.FromSlash(fmt.Sprintf(`%s/%s.json`, sessionDir, fileName))
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
	sessionDir := ut.ToString(ses.Config["NT_SESSION_DIR"], "")
	if sessionDir != "" {
		filePath = filepath.FromSlash(fmt.Sprintf(`%s/%s.json`, sessionDir, fileName))
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

func (ses *SessionService) checkSessionTable() (err error) {
	sessionDb := ut.ToString(ses.Config["NT_SESSION_DB"], "")
	if sessionDb == "" {
		return errors.New(ut.GetMessage("missing_driver"))
	}
	sessionTable := ut.ToString(ses.Config["NT_SESSION_TABLE"], "session")
	var sqlString string
	if err = ses.Conn.CreateConnection("session", sessionDb); err == nil {
		var rows []IM
		engine := ses.Conn.Connection().Engine
		if engine == "sqlite" {
			sqlString = fmt.Sprintf("select name from sqlite_master where name = '%s' ", sessionTable)
		} else {
			sqlString = fmt.Sprintf("select table_name from information_schema.tables where table_name = '%s' ", sessionTable)
		}
		rows, err = ses.Conn.QuerySQL(sqlString, []interface{}{}, nil)
		if err == nil && len(rows) == 0 {
			jsonType := SM{
				"sqlite": "JSON", "postgres": "JSONB", "mysql": "JSON", "mssql": "NVARCHAR(MAX)",
			}
			if engine == "mysql" {
				sqlString = fmt.Sprintf(
					"CREATE TABLE %s ( id VARCHAR(255) NOT NULL, value JSON, stamp VARCHAR(255), PRIMARY KEY (id) );",
					sessionTable)
			} else {
				sqlString = fmt.Sprintf(
					"CREATE TABLE %s ( id VARCHAR(255) NOT NULL PRIMARY KEY, value %s, stamp VARCHAR(255) );",
					sessionTable, jsonType[engine])
			}
			_, err = ses.Conn.QuerySQL(sqlString, []interface{}{}, nil)
		}
	}
	return err
}

func (ses *SessionService) getDbRows(rowID string) (rows []IM, err error) {
	sessionTable := ut.ToString(ses.Config["NT_SESSION_TABLE"], "session")
	sqlString := fmt.Sprintf("SELECT * FROM %s WHERE id='%s'", sessionTable, rowID)
	return ses.Conn.QuerySQL(sqlString, []interface{}{}, nil)
}

func (ses *SessionService) saveDbSession(sessionID string, data any) (err error) {
	sessionDb := ut.ToString(ses.Config["NT_SESSION_DB"], "")
	sessionTable := ut.ToString(ses.Config["NT_SESSION_TABLE"], "session")
	if err = ses.Conn.CreateConnection("session", sessionDb); err == nil {
		var bin []byte
		bin, err = ses.ConvertToByte(data)
		if err == nil {
			var sqlString string = fmt.Sprintf(
				"INSERT INTO %s(id, value, stamp) VALUES('%s', '%s', '%s')",
				sessionTable, sessionID, bin, time.Now().Format("2006-01-02T15:04:05-0700"))
			var rows []IM
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
	sessionDb := ut.ToString(ses.Config["NT_SESSION_DB"], "")
	sessionTable := ut.ToString(ses.Config["NT_SESSION_TABLE"], "session")
	if err = ses.Conn.CreateConnection("session", sessionDb); err == nil {
		sqlString := fmt.Sprintf(
			"DELETE FROM %s WHERE ID='%s'", sessionTable, sessionID)
		_, err = ses.Conn.QuerySQL(sqlString, []interface{}{}, nil)
	}
	return err
}

func (ses *SessionService) loadDbSession(sessionID string, data any) (err error) {
	sessionDb := ut.ToString(ses.Config["NT_SESSION_DB"], "")
	if err = ses.Conn.CreateConnection("session", sessionDb); err == nil {
		var rows []IM
		if rows, err = ses.getDbRows(sessionID); err == nil {
			if len(rows) == 0 {
				return errors.New(ut.GetMessage("nodata"))
			}
			jdata := []byte(ut.ToString(rows[0]["value"], ""))
			err = ses.ConvertFromByte(jdata, &data)
		}
	}
	return err
}

func (ses *SessionService) SaveSession(sessionKey string, data any) {
	if ses.MemSession == nil {
		ses.MemSession = make(IM)
	}
	switch ses.Method {
	case "file":
		if err := ses.saveFileSession(sessionKey, data); err != nil {
			ses.MemSession[sessionKey] = data
			ses.Method = "mem"
		}

	case "db":
		if err := ses.saveDbSession(sessionKey, data); err != nil {
			ses.MemSession[sessionKey] = data
			ses.Method = "mem"
		}

	case "mem":
		ses.MemSession[sessionKey] = data

	default:
		if ut.ToString(ses.Config["NT_SESSION_DB"], "") != "" {
			if err := ses.checkSessionTable(); err == nil {
				ses.Method = "db"
				ses.SaveSession(sessionKey, data)
				return
			}
		}
		if ut.ToString(ses.Config["NT_SESSION_DIR"], "") != "" {
			ses.Method = "file"
			ses.SaveSession(sessionKey, data)
			return
		}
		ses.Method = "mem"
		ses.MemSession[sessionKey] = data
	}
}

func (ses *SessionService) LoadSession(sessionKey string, data any) (result any, err error) {
	switch ses.Method {
	case "file":
		err = ses.loadFileSession(sessionKey, data)
		return data, err
	case "db":
		err = ses.loadDbSession(sessionKey, &data)
		return data, err
	default:
		if result, found := ses.MemSession[sessionKey]; found {
			return result, err
		}
	}
	return result, errors.New(ut.GetMessage("nodata"))
}

func (ses *SessionService) DeleteSession(sessionKey string) (err error) {
	switch ses.Method {
	case "file":
		return ses.deleteFileSession(sessionKey)
	case "db":
		return ses.deleteDbSession(sessionKey)
	default:
		delete(ses.MemSession, sessionKey)
	}
	return nil
}
