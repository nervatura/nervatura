package nervatura

import (
	"errors"
	"io/fs"
	"os"
	"testing"
)

func TestSessionService_saveFileSession(t *testing.T) {
	type fields struct {
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
				Config: IM{
					"NT_SESSION_DIR": "dir",
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
				data:     IM{},
			},
			wantErr: true,
		},
		{
			name: "create",
			fields: fields{
				Config: IM{
					"NT_SESSION_DIR": "dir",
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
				data:     IM{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ses := &SessionService{
				Config:          tt.fields.Config,
				Conn:            tt.fields.Conn,
				Method:          tt.fields.Method,
				MemSession:      tt.fields.MemSession,
				CreateDir:       tt.fields.CreateDir,
				CreateFile:      tt.fields.CreateFile,
				ReadFile:        tt.fields.ReadFile,
				FileStat:        tt.fields.FileStat,
				ConvertToByte:   tt.fields.ConvertToByte,
				ConvertFromByte: tt.fields.ConvertFromByte,
			}
			if err := ses.saveFileSession(tt.args.fileName, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("SessionService.saveFileSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSessionService_loadFileSession(t *testing.T) {
	type fields struct {
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
				Config: IM{
					"NT_SESSION_DIR": "dir",
				},
				ReadFile: func(name string) ([]byte, error) {
					return []byte{}, nil
				},
				ConvertFromByte: func(data []byte, result interface{}) error {
					return nil
				},
			},
			args: args{
				fileName: "name", data: IM{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ses := &SessionService{
				Config:          tt.fields.Config,
				Conn:            tt.fields.Conn,
				Method:          tt.fields.Method,
				MemSession:      tt.fields.MemSession,
				CreateDir:       tt.fields.CreateDir,
				CreateFile:      tt.fields.CreateFile,
				ReadFile:        tt.fields.ReadFile,
				FileStat:        tt.fields.FileStat,
				ConvertToByte:   tt.fields.ConvertToByte,
				ConvertFromByte: tt.fields.ConvertFromByte,
			}
			if err := ses.loadFileSession(tt.args.fileName, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("SessionService.loadFileSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSessionService_checkSessionTable(t *testing.T) {
	type fields struct {
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
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "missing",
			fields: fields{
				Config: IM{},
			},
			wantErr: true,
		},
		{
			name: "sqlite",
			fields: fields{
				Config: IM{
					"NT_SESSION_DB": "test",
				},
				Conn: &testDriver{
					Config: IM{
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
				Config: IM{
					"NT_SESSION_DB": "test",
				},
				Conn: &testDriver{
					Config: IM{
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
				Config:          tt.fields.Config,
				Conn:            tt.fields.Conn,
				Method:          tt.fields.Method,
				MemSession:      tt.fields.MemSession,
				CreateDir:       tt.fields.CreateDir,
				CreateFile:      tt.fields.CreateFile,
				ReadFile:        tt.fields.ReadFile,
				FileStat:        tt.fields.FileStat,
				ConvertToByte:   tt.fields.ConvertToByte,
				ConvertFromByte: tt.fields.ConvertFromByte,
			}
			if err := ses.checkSessionTable(); (err != nil) != tt.wantErr {
				t.Errorf("SessionService.checkSessionTable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSessionService_saveDbSession(t *testing.T) {
	type fields struct {
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
				Config: IM{},
				Conn:   &testDriver{Config: IM{}},
				ConvertToByte: func(data interface{}) ([]byte, error) {
					return []byte{}, nil
				},
			},
			args: args{
				sessionID: "id",
				data:      IM{},
			},
			wantErr: false,
		},
		{
			name: "update",
			fields: fields{
				Config: IM{},
				Conn: &testDriver{Config: IM{
					"QuerySQL": func(sqlString string) ([]IM, error) {
						return []IM{
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
				data:      IM{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ses := &SessionService{
				Config:          tt.fields.Config,
				Conn:            tt.fields.Conn,
				Method:          tt.fields.Method,
				MemSession:      tt.fields.MemSession,
				CreateDir:       tt.fields.CreateDir,
				CreateFile:      tt.fields.CreateFile,
				ReadFile:        tt.fields.ReadFile,
				FileStat:        tt.fields.FileStat,
				ConvertToByte:   tt.fields.ConvertToByte,
				ConvertFromByte: tt.fields.ConvertFromByte,
			}
			if err := ses.saveDbSession(tt.args.sessionID, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("SessionService.saveDbSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSessionService_loadDbSession(t *testing.T) {
	type fields struct {
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
				Config: IM{},
				Conn:   &testDriver{Config: IM{}},
			},
			args:    args{},
			wantErr: true,
		},
		{
			name: "ok",
			fields: fields{
				Config: IM{},
				Conn: &testDriver{Config: IM{
					"QuerySQL": func(sqlString string) ([]IM, error) {
						return []IM{
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
				data:      IM{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ses := &SessionService{
				Config:          tt.fields.Config,
				Conn:            tt.fields.Conn,
				Method:          tt.fields.Method,
				MemSession:      tt.fields.MemSession,
				CreateDir:       tt.fields.CreateDir,
				CreateFile:      tt.fields.CreateFile,
				ReadFile:        tt.fields.ReadFile,
				FileStat:        tt.fields.FileStat,
				ConvertToByte:   tt.fields.ConvertToByte,
				ConvertFromByte: tt.fields.ConvertFromByte,
			}
			if err := ses.loadDbSession(tt.args.sessionID, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("SessionService.loadDbSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSessionService_SaveSession(t *testing.T) {
	type fields struct {
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
				Config: IM{
					"NT_SESSION_DIR": "dir",
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
				data:       IM{},
			},
		},
		{
			name: "db_err",
			fields: fields{
				Config: IM{
					"NT_SESSION_DB": "test",
				},
				Conn: &testDriver{Config: IM{}},
				ConvertToByte: func(data interface{}) ([]byte, error) {
					return []byte{}, errors.New("error")
				},
			},
			args: args{
				sessionKey: "key",
				data:       IM{},
			},
		},
		{
			name: "mem1",
			fields: fields{
				Config: IM{},
			},
			args: args{
				sessionKey: "key",
				data:       IM{},
			},
		},
		{
			name: "mem2",
			fields: fields{
				Config: IM{},
				Method: "mem",
			},
			args: args{
				sessionKey: "key",
				data:       IM{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ses := &SessionService{
				Config:          tt.fields.Config,
				Conn:            tt.fields.Conn,
				Method:          tt.fields.Method,
				MemSession:      tt.fields.MemSession,
				CreateDir:       tt.fields.CreateDir,
				CreateFile:      tt.fields.CreateFile,
				ReadFile:        tt.fields.ReadFile,
				FileStat:        tt.fields.FileStat,
				ConvertToByte:   tt.fields.ConvertToByte,
				ConvertFromByte: tt.fields.ConvertFromByte,
			}
			ses.SaveSession(tt.args.sessionKey, tt.args.data)
		})
	}
}

func TestSessionService_LoadSession(t *testing.T) {
	type fields struct {
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
	}
	type args struct {
		sessionKey string
		data       any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "nodata",
			fields: fields{},
			args: args{
				sessionKey: "key",
				data:       IM{},
			},
			wantErr: true,
		},
		{
			name: "mem",
			fields: fields{
				MemSession: IM{
					"key": IM{},
				},
			},
			args: args{
				sessionKey: "key",
				data:       IM{},
			},
			wantErr: false,
		},
		{
			name: "file",
			fields: fields{
				Config: IM{
					"NT_SESSION_DIR": "dir",
				},
				ReadFile: func(name string) ([]byte, error) {
					return []byte{}, nil
				},
				ConvertFromByte: func(data []byte, result interface{}) error {
					return nil
				},
				Method: "file",
			},
			args: args{
				sessionKey: "key",
				data:       IM{},
			},
			wantErr: false,
		},
		{
			name: "db",
			fields: fields{
				Config: IM{},
				Conn: &testDriver{Config: IM{
					"QuerySQL": func(sqlString string) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
				}},
				ConvertFromByte: func(data []byte, result interface{}) error {
					return nil
				},
				Method: "db",
			},
			args: args{
				sessionKey: "key",
				data:       IM{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ses := &SessionService{
				Config:          tt.fields.Config,
				Conn:            tt.fields.Conn,
				Method:          tt.fields.Method,
				MemSession:      tt.fields.MemSession,
				CreateDir:       tt.fields.CreateDir,
				CreateFile:      tt.fields.CreateFile,
				ReadFile:        tt.fields.ReadFile,
				FileStat:        tt.fields.FileStat,
				ConvertToByte:   tt.fields.ConvertToByte,
				ConvertFromByte: tt.fields.ConvertFromByte,
			}
			_, err := ses.LoadSession(tt.args.sessionKey, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("SessionService.LoadSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
