package api

import (
	"errors"
	"io/fs"
	"os"
	"testing"
	"time"

	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	st "github.com/nervatura/nervatura/v6/pkg/static"
)

func TestSessionService_saveFileSession(t *testing.T) {
	type fields struct {
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
	type args struct {
		fileName string
		data     any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "create_dir_err",
			fields: fields{
				Config: SessionConfig{
					FileDir: "dir",
				},
				FileStat: func(name string) (fs.FileInfo, error) {
					return nil, os.ErrNotExist
				},
				CreateDir: func(name string, perm fs.FileMode) error {
					return errors.New("error")
				},
			},
			args: args{
				fileName: "name",
				data:     cu.IM{},
			},
			wantErr: true,
		},
		{
			name: "create",
			fields: fields{
				Config: SessionConfig{
					FileDir: "dir",
				},
				FileStat: func(name string) (fs.FileInfo, error) {
					return nil, nil
				},
				CreateDir: func(name string, perm fs.FileMode) error {
					return nil
				},
				CreateFile: func(name string) (*os.File, error) {
					file := os.NewFile(os.Stdin.Fd(), "name")
					return file, nil
				},
				ConvertToByte: func(data interface{}) ([]byte, error) {
					return []byte{}, nil
				},
			},
			args: args{
				fileName: "name",
				data:     cu.IM{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ses := &SessionService{
				method:          tt.fields.method,
				memSession:      tt.fields.memSession,
				cleaningStamp:   tt.fields.cleaningStamp,
				Config:          tt.fields.Config,
				Conn:            tt.fields.Conn,
				GetMessage:      tt.fields.GetMessage,
				CreateDir:       tt.fields.CreateDir,
				CreateFile:      tt.fields.CreateFile,
				ReadDir:         tt.fields.ReadDir,
				ReadFile:        tt.fields.ReadFile,
				FileStat:        tt.fields.FileStat,
				ConvertToByte:   tt.fields.ConvertToByte,
				ConvertFromByte: tt.fields.ConvertFromByte,
				RemoveFile:      tt.fields.RemoveFile,
			}
			if err := ses.saveFileSession(tt.args.fileName, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("SessionService.saveFileSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSessionService_loadFileSession(t *testing.T) {
	type fields struct {
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
	type args struct {
		fileName string
		data     any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				Config: SessionConfig{
					FileDir: "dir",
				},
				ReadFile: func(name string) ([]byte, error) {
					return []byte{}, nil
				},
				ConvertFromByte: func(data []byte, result interface{}) error {
					return nil
				},
			},
			args: args{
				fileName: "name", data: cu.IM{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ses := &SessionService{
				method:          tt.fields.method,
				memSession:      tt.fields.memSession,
				cleaningStamp:   tt.fields.cleaningStamp,
				Config:          tt.fields.Config,
				Conn:            tt.fields.Conn,
				GetMessage:      tt.fields.GetMessage,
				CreateDir:       tt.fields.CreateDir,
				CreateFile:      tt.fields.CreateFile,
				ReadDir:         tt.fields.ReadDir,
				ReadFile:        tt.fields.ReadFile,
				FileStat:        tt.fields.FileStat,
				ConvertToByte:   tt.fields.ConvertToByte,
				ConvertFromByte: tt.fields.ConvertFromByte,
				RemoveFile:      tt.fields.RemoveFile,
			}
			if err := ses.loadFileSession(tt.args.fileName, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("SessionService.loadFileSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSessionService_checkSessionTable(t *testing.T) {
	type fields struct {
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
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "missing",
			fields: fields{
				Config: SessionConfig{},
			},
			wantErr: true,
		},
		{
			name: "sqlite",
			fields: fields{
				Config: SessionConfig{
					DbConn: "test",
				},
				Conn: &md.TestDriver{
					Config: cu.IM{
						"Connection": func() struct {
							Alias     string
							Connected bool
							Engine    string
						} {
							return struct {
								Alias     string
								Connected bool
								Engine    string
							}{
								Alias:     "test",
								Connected: false,
								Engine:    "sqlite",
							}
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "mysql",
			fields: fields{
				Config: SessionConfig{
					DbConn: "test",
				},
				Conn: &md.TestDriver{
					Config: cu.IM{
						"Connection": func() struct {
							Alias     string
							Connected bool
							Engine    string
						} {
							return struct {
								Alias     string
								Connected bool
								Engine    string
							}{
								Alias:     "test",
								Connected: false,
								Engine:    "mysql",
							}
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ses := &SessionService{
				method:          tt.fields.method,
				memSession:      tt.fields.memSession,
				cleaningStamp:   tt.fields.cleaningStamp,
				Config:          tt.fields.Config,
				Conn:            tt.fields.Conn,
				GetMessage:      tt.fields.GetMessage,
				CreateDir:       tt.fields.CreateDir,
				CreateFile:      tt.fields.CreateFile,
				ReadDir:         tt.fields.ReadDir,
				ReadFile:        tt.fields.ReadFile,
				FileStat:        tt.fields.FileStat,
				ConvertToByte:   tt.fields.ConvertToByte,
				ConvertFromByte: tt.fields.ConvertFromByte,
				RemoveFile:      tt.fields.RemoveFile,
			}
			if err := ses.checkSessionTable(); (err != nil) != tt.wantErr {
				t.Errorf("SessionService.checkSessionTable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSessionService_saveDbSession(t *testing.T) {
	type fields struct {
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
	type args struct {
		sessionID string
		data      any
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
				Config: SessionConfig{},
				Conn:   &md.TestDriver{Config: cu.IM{}},
				ConvertToByte: func(data interface{}) ([]byte, error) {
					return []byte{}, nil
				},
			},
			args: args{
				sessionID: "id",
				data:      cu.IM{},
			},
			wantErr: false,
		},
		{
			name: "update",
			fields: fields{
				Config: SessionConfig{},
				Conn: &md.TestDriver{Config: cu.IM{
					"QuerySQL": func(sqlString string) ([]cu.IM, error) {
						return []cu.IM{
							{"id": int64(1)},
						}, nil
					},
				}},
				ConvertToByte: func(data interface{}) ([]byte, error) {
					return []byte{}, nil
				},
			},
			args: args{
				sessionID: "id",
				data:      cu.IM{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ses := &SessionService{
				method:          tt.fields.method,
				memSession:      tt.fields.memSession,
				cleaningStamp:   tt.fields.cleaningStamp,
				Config:          tt.fields.Config,
				Conn:            tt.fields.Conn,
				GetMessage:      tt.fields.GetMessage,
				CreateDir:       tt.fields.CreateDir,
				CreateFile:      tt.fields.CreateFile,
				ReadDir:         tt.fields.ReadDir,
				ReadFile:        tt.fields.ReadFile,
				FileStat:        tt.fields.FileStat,
				ConvertToByte:   tt.fields.ConvertToByte,
				ConvertFromByte: tt.fields.ConvertFromByte,
				RemoveFile:      tt.fields.RemoveFile,
			}
			if err := ses.saveDbSession(tt.args.sessionID, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("SessionService.saveDbSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSessionService_loadDbSession(t *testing.T) {
	type fields struct {
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
	type args struct {
		sessionID string
		data      any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "nodata",
			fields: fields{
				Config: SessionConfig{},
				Conn:   &md.TestDriver{Config: cu.IM{}},
				GetMessage: func(lang, key string) string {
					return key
				},
			},
			args:    args{},
			wantErr: true,
		},
		{
			name: "ok",
			fields: fields{
				Config: SessionConfig{},
				Conn: &md.TestDriver{Config: cu.IM{
					"QuerySQL": func(sqlString string) ([]cu.IM, error) {
						return []cu.IM{
							{"id": int64(1)},
						}, nil
					},
				}},
				ConvertFromByte: func(data []byte, result interface{}) error {
					return nil
				},
			},
			args: args{
				sessionID: "id",
				data:      cu.IM{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ses := &SessionService{
				method:          tt.fields.method,
				memSession:      tt.fields.memSession,
				cleaningStamp:   tt.fields.cleaningStamp,
				Config:          tt.fields.Config,
				Conn:            tt.fields.Conn,
				GetMessage:      tt.fields.GetMessage,
				CreateDir:       tt.fields.CreateDir,
				CreateFile:      tt.fields.CreateFile,
				ReadDir:         tt.fields.ReadDir,
				ReadFile:        tt.fields.ReadFile,
				FileStat:        tt.fields.FileStat,
				ConvertToByte:   tt.fields.ConvertToByte,
				ConvertFromByte: tt.fields.ConvertFromByte,
				RemoveFile:      tt.fields.RemoveFile,
			}
			if err := ses.loadDbSession(tt.args.sessionID, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("SessionService.loadDbSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSessionService_SaveSession(t *testing.T) {
	type fields struct {
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
	type args struct {
		sessionKey string
		data       any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "file_err",
			fields: fields{
				Config: SessionConfig{
					FileDir: "dir",
				},
				FileStat: func(name string) (fs.FileInfo, error) {
					return nil, nil
				},
				CreateFile: func(name string) (*os.File, error) {
					return nil, errors.New("error")
				},
			},
			args: args{
				sessionKey: "key",
				data:       cu.IM{},
			},
		},
		{
			name: "db_err",
			fields: fields{
				Config: SessionConfig{
					DbConn: "test",
				},
				Conn: &md.TestDriver{Config: cu.IM{}},
				ConvertToByte: func(data interface{}) ([]byte, error) {
					return []byte{}, errors.New("error")
				},
			},
			args: args{
				sessionKey: "key",
				data:       cu.IM{},
			},
		},
		{
			name: "mem1",
			fields: fields{
				Config: SessionConfig{},
			},
			args: args{
				sessionKey: "key",
				data:       cu.IM{},
			},
		},
		{
			name: "mem2",
			fields: fields{
				Config: SessionConfig{
					Method: SessionMethodMemory,
				},
			},
			args: args{
				sessionKey: "key",
				data:       cu.IM{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ses := &SessionService{
				method:          tt.fields.method,
				memSession:      tt.fields.memSession,
				cleaningStamp:   tt.fields.cleaningStamp,
				Config:          tt.fields.Config,
				Conn:            tt.fields.Conn,
				GetMessage:      tt.fields.GetMessage,
				CreateDir:       tt.fields.CreateDir,
				CreateFile:      tt.fields.CreateFile,
				ReadDir:         tt.fields.ReadDir,
				ReadFile:        tt.fields.ReadFile,
				FileStat:        tt.fields.FileStat,
				ConvertToByte:   tt.fields.ConvertToByte,
				ConvertFromByte: tt.fields.ConvertFromByte,
				RemoveFile:      tt.fields.RemoveFile,
			}
			ses.SaveSession(tt.args.sessionKey, tt.args.data)
		})
	}
}

func TestSessionService_LoadSession(t *testing.T) {
	type fields struct {
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
	type args struct {
		sessionKey string
		data       any
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult any
		wantErr    bool
	}{
		{
			name: "nodata",
			fields: fields{
				GetMessage: func(lang, key string) string {
					return key
				},
			},
			args: args{
				sessionKey: "key",
				data:       cu.IM{},
			},
			wantErr: true,
		},
		{
			name: "mem",
			fields: fields{
				memSession: map[string]memStore{
					"key": {},
				},
			},
			args: args{
				sessionKey: "key",
				data:       cu.IM{},
			},
			wantErr: false,
		},
		{
			name: "file",
			fields: fields{
				Config: SessionConfig{
					FileDir: "dir",
					Method:  SessionMethodFile,
				},
				ReadFile: func(name string) ([]byte, error) {
					return []byte{}, nil
				},
				ConvertFromByte: func(data []byte, result interface{}) error {
					return nil
				},
			},
			args: args{
				sessionKey: "key",
				data:       cu.IM{},
			},
			wantErr: false,
		},
		{
			name: "db",
			fields: fields{
				Config: SessionConfig{
					Method: SessionMethodDatabase,
				},
				Conn: &md.TestDriver{Config: cu.IM{
					"QuerySQL": func(sqlString string) ([]cu.IM, error) {
						return []cu.IM{
							{"id": int64(1)},
						}, nil
					},
				}},
				ConvertFromByte: func(data []byte, result interface{}) error {
					return nil
				},
			},
			args: args{
				sessionKey: "key",
				data:       cu.IM{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ses := &SessionService{
				method:          tt.fields.method,
				memSession:      tt.fields.memSession,
				cleaningStamp:   tt.fields.cleaningStamp,
				Config:          tt.fields.Config,
				Conn:            tt.fields.Conn,
				GetMessage:      tt.fields.GetMessage,
				CreateDir:       tt.fields.CreateDir,
				CreateFile:      tt.fields.CreateFile,
				ReadDir:         tt.fields.ReadDir,
				ReadFile:        tt.fields.ReadFile,
				FileStat:        tt.fields.FileStat,
				ConvertToByte:   tt.fields.ConvertToByte,
				ConvertFromByte: tt.fields.ConvertFromByte,
				RemoveFile:      tt.fields.RemoveFile,
			}
			_, err := ses.LoadSession(tt.args.sessionKey, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("SessionService.LoadSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSessionService_deleteFileSession(t *testing.T) {
	type fields struct {
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
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "delete",
			fields: fields{
				RemoveFile: func(name string) error {
					return nil
				},
			},
			args: args{
				fileName: "fileName.json",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ses := &SessionService{
				method:          tt.fields.method,
				memSession:      tt.fields.memSession,
				cleaningStamp:   tt.fields.cleaningStamp,
				Config:          tt.fields.Config,
				Conn:            tt.fields.Conn,
				GetMessage:      tt.fields.GetMessage,
				CreateDir:       tt.fields.CreateDir,
				CreateFile:      tt.fields.CreateFile,
				ReadDir:         tt.fields.ReadDir,
				ReadFile:        tt.fields.ReadFile,
				FileStat:        tt.fields.FileStat,
				ConvertToByte:   tt.fields.ConvertToByte,
				ConvertFromByte: tt.fields.ConvertFromByte,
				RemoveFile:      tt.fields.RemoveFile,
			}
			if err := ses.deleteFileSession(tt.args.fileName); (err != nil) != tt.wantErr {
				t.Errorf("SessionService.deleteFileSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSessionService_deleteDbSession(t *testing.T) {
	type fields struct {
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
	type args struct {
		sessionID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "delete",
			fields: fields{
				Config: SessionConfig{},
				Conn: &md.TestDriver{Config: cu.IM{
					"QuerySQL": func(sqlString string) ([]cu.IM, error) {
						return []cu.IM{
							{"id": int64(1)},
						}, nil
					},
				}},
				ConvertToByte: func(data interface{}) ([]byte, error) {
					return []byte{}, nil
				},
			},
			args: args{
				sessionID: "SES012345",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ses := &SessionService{
				method:          tt.fields.method,
				memSession:      tt.fields.memSession,
				cleaningStamp:   tt.fields.cleaningStamp,
				Config:          tt.fields.Config,
				Conn:            tt.fields.Conn,
				GetMessage:      tt.fields.GetMessage,
				CreateDir:       tt.fields.CreateDir,
				CreateFile:      tt.fields.CreateFile,
				ReadDir:         tt.fields.ReadDir,
				ReadFile:        tt.fields.ReadFile,
				FileStat:        tt.fields.FileStat,
				ConvertToByte:   tt.fields.ConvertToByte,
				ConvertFromByte: tt.fields.ConvertFromByte,
				RemoveFile:      tt.fields.RemoveFile,
			}
			if err := ses.deleteDbSession(tt.args.sessionID); (err != nil) != tt.wantErr {
				t.Errorf("SessionService.deleteDbSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSessionService_DeleteSession(t *testing.T) {
	type fields struct {
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
	type args struct {
		sessionKey string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "file",
			fields: fields{
				Config: SessionConfig{
					Method: SessionMethodFile,
				},
				RemoveFile: func(name string) error {
					return nil
				},
			},
			args: args{
				sessionKey: "SES012345",
			},
			wantErr: false,
		},
		{
			name: "db",
			fields: fields{
				Config: SessionConfig{
					Method: SessionMethodDatabase,
				},
				Conn: &md.TestDriver{Config: cu.IM{
					"QuerySQL": func(sqlString string) ([]cu.IM, error) {
						return []cu.IM{
							{"id": int64(1)},
						}, nil
					},
				}},
				ConvertToByte: func(data interface{}) ([]byte, error) {
					return []byte{}, nil
				},
			},
			args: args{
				sessionKey: "SES012345",
			},
			wantErr: false,
		},
		{
			name: "mem",
			fields: fields{
				Config: SessionConfig{
					Method: SessionMethodMemory,
				},
				memSession: map[string]memStore{
					"key": {},
				},
			},
			args: args{
				sessionKey: "SES012345",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ses := &SessionService{
				method:          tt.fields.method,
				memSession:      tt.fields.memSession,
				cleaningStamp:   tt.fields.cleaningStamp,
				Config:          tt.fields.Config,
				Conn:            tt.fields.Conn,
				GetMessage:      tt.fields.GetMessage,
				CreateDir:       tt.fields.CreateDir,
				CreateFile:      tt.fields.CreateFile,
				ReadDir:         tt.fields.ReadDir,
				ReadFile:        tt.fields.ReadFile,
				FileStat:        tt.fields.FileStat,
				ConvertToByte:   tt.fields.ConvertToByte,
				ConvertFromByte: tt.fields.ConvertFromByte,
				RemoveFile:      tt.fields.RemoveFile,
			}
			if err := ses.DeleteSession(tt.args.sessionKey); (err != nil) != tt.wantErr {
				t.Errorf("SessionService.DeleteSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSessionService_cleaningFileSession(t *testing.T) {
	type fields struct {
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
	type args struct {
		exp time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "clean",
			fields: fields{
				Config: SessionConfig{
					FileDir: "dir",
				},
				ReadDir: func(name string) ([]fs.DirEntry, error) {
					return st.Static.ReadDir(".")
				},
				RemoveFile: func(name string) error {
					return nil
				},
			},
			args: args{
				exp: time.Now(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ses := &SessionService{
				method:          tt.fields.method,
				memSession:      tt.fields.memSession,
				cleaningStamp:   tt.fields.cleaningStamp,
				Config:          tt.fields.Config,
				Conn:            tt.fields.Conn,
				GetMessage:      tt.fields.GetMessage,
				CreateDir:       tt.fields.CreateDir,
				CreateFile:      tt.fields.CreateFile,
				ReadDir:         tt.fields.ReadDir,
				ReadFile:        tt.fields.ReadFile,
				FileStat:        tt.fields.FileStat,
				ConvertToByte:   tt.fields.ConvertToByte,
				ConvertFromByte: tt.fields.ConvertFromByte,
				RemoveFile:      tt.fields.RemoveFile,
			}
			if err := ses.cleaningFileSession(tt.args.exp); (err != nil) != tt.wantErr {
				t.Errorf("SessionService.cleaningFileSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSessionService_cleaningDbSession(t *testing.T) {
	type fields struct {
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
	type args struct {
		exp time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "clean",
			fields: fields{
				Conn: &md.TestDriver{Config: cu.IM{}},
			},
			args: args{
				exp: time.Now(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ses := &SessionService{
				memSession:      tt.fields.memSession,
				method:          tt.fields.method,
				cleaningStamp:   tt.fields.cleaningStamp,
				Config:          tt.fields.Config,
				Conn:            tt.fields.Conn,
				GetMessage:      tt.fields.GetMessage,
				CreateDir:       tt.fields.CreateDir,
				CreateFile:      tt.fields.CreateFile,
				ReadDir:         tt.fields.ReadDir,
				ReadFile:        tt.fields.ReadFile,
				FileStat:        tt.fields.FileStat,
				ConvertToByte:   tt.fields.ConvertToByte,
				ConvertFromByte: tt.fields.ConvertFromByte,
				RemoveFile:      tt.fields.RemoveFile,
			}
			if err := ses.cleaningDbSession(tt.args.exp); (err != nil) != tt.wantErr {
				t.Errorf("SessionService.cleaningDbSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSessionService_cleaningMemSession(t *testing.T) {
	type fields struct {
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
	type args struct {
		exp time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "clean",
			fields: fields{
				memSession: map[string]memStore{
					"SES0123": {},
				},
			},
			args: args{
				exp: time.Now(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ses := &SessionService{
				memSession:      tt.fields.memSession,
				method:          tt.fields.method,
				cleaningStamp:   tt.fields.cleaningStamp,
				Config:          tt.fields.Config,
				Conn:            tt.fields.Conn,
				GetMessage:      tt.fields.GetMessage,
				CreateDir:       tt.fields.CreateDir,
				CreateFile:      tt.fields.CreateFile,
				ReadDir:         tt.fields.ReadDir,
				ReadFile:        tt.fields.ReadFile,
				FileStat:        tt.fields.FileStat,
				ConvertToByte:   tt.fields.ConvertToByte,
				ConvertFromByte: tt.fields.ConvertFromByte,
				RemoveFile:      tt.fields.RemoveFile,
			}
			if err := ses.cleaningMemSession(tt.args.exp); (err != nil) != tt.wantErr {
				t.Errorf("SessionService.cleaningMemSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSessionService_CleaningSession(t *testing.T) {
	type fields struct {
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
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "mem",
			fields: fields{
				Config: SessionConfig{
					Cleaning: true,
				},
			},
			wantErr: false,
		},
		{
			name: "file",
			fields: fields{
				Config: SessionConfig{
					Cleaning: true,
					Method:   SessionMethodFile,
				},
			},
			wantErr: false,
		},
		{
			name: "db",
			fields: fields{
				Config: SessionConfig{
					Cleaning: true,
					Method:   SessionMethodDatabase,
				},
				Conn: &md.TestDriver{Config: cu.IM{}},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ses := &SessionService{
				memSession:      tt.fields.memSession,
				method:          tt.fields.method,
				cleaningStamp:   tt.fields.cleaningStamp,
				Config:          tt.fields.Config,
				Conn:            tt.fields.Conn,
				GetMessage:      tt.fields.GetMessage,
				CreateDir:       tt.fields.CreateDir,
				CreateFile:      tt.fields.CreateFile,
				ReadDir:         tt.fields.ReadDir,
				ReadFile:        tt.fields.ReadFile,
				FileStat:        tt.fields.FileStat,
				ConvertToByte:   tt.fields.ConvertToByte,
				ConvertFromByte: tt.fields.ConvertFromByte,
				RemoveFile:      tt.fields.RemoveFile,
			}
			if err := ses.CleaningSession(); (err != nil) != tt.wantErr {
				t.Errorf("SessionService.CleaningSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewSession(t *testing.T) {
	type args struct {
		config cu.IM
		method string
		alias  string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				config: cu.IM{},
				method: SessionMethodFile,
				alias:  "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewSession(tt.args.config, tt.args.method, tt.args.alias)
		})
	}
}
