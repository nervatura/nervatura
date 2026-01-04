package api

import (
	"bytes"
	"crypto/tls"
	"errors"
	"io"
	"log/slog"
	"net"
	"net/smtp"
	"path"
	"testing"

	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	st "github.com/nervatura/nervatura/v6/pkg/static"
)

// Mock SMTP Client
type MockSMTPClient struct {
	auth           smtp.Auth
	rcptCalls      []string
	mailCalls      []string
	mockAuthError  error
	mockMailError  error
	mockRcptError  error
	mockDataError  error
	mockCloseError error
	mockQuitError  error
	mockWriteError error
}

func (c *MockSMTPClient) Auth(a smtp.Auth) error {
	c.auth = a
	return c.mockAuthError
}

func (c *MockSMTPClient) Mail(from string) error {
	c.mailCalls = append(c.mailCalls, from)
	return c.mockMailError
}

func (c *MockSMTPClient) Rcpt(to string) error {
	c.rcptCalls = append(c.rcptCalls, to)
	return c.mockRcptError
}

func (c *MockSMTPClient) Data() (io.WriteCloser, error) {
	return &mockWriter{mockError: c.mockWriteError}, c.mockDataError
}

func (c *MockSMTPClient) Close() error {
	return c.mockCloseError
}

func (c *MockSMTPClient) Quit() error {
	return c.mockQuitError
}

// Mock Writer
type mockWriter struct {
	written   []byte
	mockError error
}

func (w *mockWriter) Write(p []byte) (n int, err error) {
	if w.mockError != nil {
		return 0, w.mockError
	}
	w.written = append(w.written, p...)
	return len(p), nil
}

func (w *mockWriter) Close() error {
	return w.mockError
}

func TestDataStore_SendEmail(t *testing.T) {
	type fields struct {
		Db                     DataDriver
		Alias                  string
		Config                 cu.IM
		AppLog                 *slog.Logger
		ReadAll                func(r io.Reader) ([]byte, error)
		ConvertToByte          func(data interface{}) ([]byte, error)
		ConvertFromByte        func(data []byte, result interface{}) error
		ConvertFromReader      func(data io.Reader, result interface{}) error
		ConvertToType          func(data interface{}, result any) (err error)
		GetDataField           func(data any, JSONName string) (fieldName string, fieldValue interface{})
		CreateLoginToken       func(params cu.SM, config cu.IM) (result string, err error)
		ParseToken             func(token string, keyMap []cu.SM, config cu.IM) (cu.IM, error)
		CreatePasswordHash     func(password string) (hash string, err error)
		ComparePasswordAndHash func(password string, hash string) (err error)
		ReadFile               func(name string) ([]byte, error)
		NewSmtpClient          func(conn net.Conn, host string) (md.SmtpClient, error)
	}
	type args struct {
		options cu.IM
	}
	sample_json, _ := st.Report.ReadFile(path.Join("template", "ntr_customer_en.json"))
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				Config: cu.IM{
					"NT_SMTP_USER":            "test@example.com",
					"NT_SMTP_PASSWORD":        "test",
					"NT_SMTP_HOST":            "localhost",
					"NT_SMTP_PORT":            25,
					"NT_SMTP_TLS_MIN_VERSION": tls.VersionTLS12,
					"NT_SMTP_CONN":            "test",
					"NT_SMTP_AUTH":            "auth",
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				NewSmtpClient: func(conn net.Conn, host string) (md.SmtpClient, error) {
					return &MockSMTPClient{}, nil
				},
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{
								{"id": 1, "name": "test", "data": cu.IM{"template": string(sample_json)}},
							}, nil
						},
						"QuerySQL": func(sqlString string) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "data": cu.IM{"file_type": "FILE_PDF"}}}, nil
						},
					},
				},
				ReadFile: func(name string) ([]byte, error) {
					return []byte(`{"id": 1, "name": "test"}`), nil
				},
				ConvertFromByte: func(data []byte, result interface{}) error {
					return cu.ConvertFromByte(data, result)
				},
				ConvertToByte: func(data interface{}) ([]byte, error) {
					return []byte(`{"meta": {"report_key": "test", "report_name": "test", "report_type": "test", "file_type": "FILE_PDF"}}`), nil
				},
			},
			args: args{
				options: cu.IM{
					"email": cu.IM{
						"from":       "test@example.com",
						"recipients": []interface{}{cu.IM{"email": "recipient@example.com"}},
						"subject":    "Test Email",
						"text":       "Test content",
						"attachments": []interface{}{cu.IM{
							"report_key": "ntr_customer_en",
							"code":       "CUS0000000000N2",
							"filename":   "test.pdf",
						}},
					},
					"provider": "smtp",
				},
			},
			wantErr: false,
		},
		{
			name: "Missing Required Field email",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			args: args{
				options: cu.IM{},
			},
			wantErr: true,
		},
		{
			name: "Missing Required Field recipients",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			args: args{
				options: cu.IM{
					"email": cu.IM{"from": "test@example.com", "subject": "Test Email", "text": "Test content"},
				},
			},
			wantErr: true,
		},
		{
			name: "Invalid Provider",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			args: args{
				options: cu.IM{
					"email":    cu.IM{"from": "test@example.com", "recipients": []interface{}{cu.IM{"email": "recipient@example.com"}}, "subject": "Test Email", "text": "Test content"},
					"provider": "invalid",
				},
			},
			wantErr: true,
		},
		{
			name: "net connection error",
			fields: fields{
				Config: cu.IM{
					"NT_SMTP_USER":            "test@example.com",
					"NT_SMTP_PASSWORD":        "test",
					"NT_SMTP_HOST":            "localhost",
					"NT_SMTP_PORT":            25,
					"NT_SMTP_TLS_MIN_VERSION": tls.VersionTLS12,
					"NT_SMTP_CONN":            "net",
					"NT_SMTP_AUTH":            "auth",
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				NewSmtpClient: func(conn net.Conn, host string) (md.SmtpClient, error) {
					return nil, errors.New("net connection error")
				},
			},
			args: args{
				options: cu.IM{
					"email": cu.IM{
						"from":       "test@example.com",
						"recipients": []interface{}{cu.IM{"email": "recipient@example.com"}},
						"subject":    "Test Email",
						"text":       "Test content",
					},
					"provider": "smtp",
				},
			},
			wantErr: true,
		},
		{
			name: "tls connection error",
			fields: fields{
				Config: cu.IM{
					"NT_SMTP_USER":            "test@example.com",
					"NT_SMTP_PASSWORD":        "test",
					"NT_SMTP_HOST":            "localhost",
					"NT_SMTP_PORT":            25,
					"NT_SMTP_TLS_MIN_VERSION": tls.VersionTLS12,
					"NT_SMTP_CONN":            "tls",
					"NT_SMTP_AUTH":            "auth",
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			args: args{
				options: cu.IM{
					"email": cu.IM{
						"from":       "test@example.com",
						"recipients": []interface{}{cu.IM{"email": "recipient@example.com"}},
						"subject":    "Test Email",
						"text":       "Test content",
					},
					"provider": "smtp",
				},
			},
			wantErr: true,
		},
		{
			name: "client error",
			fields: fields{
				Config: cu.IM{
					"NT_SMTP_USER":            "test@example.com",
					"NT_SMTP_PASSWORD":        "test",
					"NT_SMTP_HOST":            "localhost",
					"NT_SMTP_PORT":            25,
					"NT_SMTP_TLS_MIN_VERSION": tls.VersionTLS12,
					"NT_SMTP_CONN":            "test",
					"NT_SMTP_AUTH":            "auth",
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				NewSmtpClient: func(conn net.Conn, host string) (md.SmtpClient, error) {
					return nil, errors.New("client error")
				},
			},
			args: args{
				options: cu.IM{
					"email": cu.IM{
						"from":       "test@example.com",
						"recipients": []interface{}{cu.IM{"email": "recipient@example.com"}},
						"subject":    "Test Email",
						"text":       "Test content",
					},
					"provider": "smtp",
				},
			},
			wantErr: true,
		},
		{
			name: "auth error",
			fields: fields{
				Config: cu.IM{
					"NT_SMTP_USER":            "test@example.com",
					"NT_SMTP_PASSWORD":        "test",
					"NT_SMTP_HOST":            "localhost",
					"NT_SMTP_PORT":            25,
					"NT_SMTP_TLS_MIN_VERSION": tls.VersionTLS12,
					"NT_SMTP_CONN":            "test",
					"NT_SMTP_AUTH":            "auth",
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				NewSmtpClient: func(conn net.Conn, host string) (md.SmtpClient, error) {
					return &MockSMTPClient{mockAuthError: errors.New("auth error")}, nil
				},
			},
			args: args{
				options: cu.IM{
					"email": cu.IM{
						"from":       "test@example.com",
						"recipients": []interface{}{cu.IM{"email": "recipient@example.com"}},
						"subject":    "Test Email",
						"text":       "Test content",
					},
					"provider": "smtp",
				},
			},
			wantErr: true,
		},
		{
			name: "client mail error",
			fields: fields{
				Config: cu.IM{
					"NT_SMTP_USER":            "test@example.com",
					"NT_SMTP_PASSWORD":        "test",
					"NT_SMTP_HOST":            "localhost",
					"NT_SMTP_PORT":            25,
					"NT_SMTP_TLS_MIN_VERSION": tls.VersionTLS12,
					"NT_SMTP_CONN":            "test",
					"NT_SMTP_AUTH":            "none",
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				NewSmtpClient: func(conn net.Conn, host string) (md.SmtpClient, error) {
					return &MockSMTPClient{mockMailError: errors.New("mail error")}, nil
				},
			},
			args: args{
				options: cu.IM{
					"email": cu.IM{
						"from":       "test@example.com",
						"recipients": []interface{}{cu.IM{"email": "recipient@example.com"}},
						"subject":    "Test Email",
						"text":       "Test content",
					},
					"provider": "smtp",
				},
			},
			wantErr: true,
		},
		{
			name: "client rcpt error",
			fields: fields{
				Config: cu.IM{
					"NT_SMTP_USER":            "test@example.com",
					"NT_SMTP_PASSWORD":        "test",
					"NT_SMTP_HOST":            "localhost",
					"NT_SMTP_PORT":            25,
					"NT_SMTP_TLS_MIN_VERSION": tls.VersionTLS12,
					"NT_SMTP_CONN":            "test",
					"NT_SMTP_AUTH":            "none",
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				NewSmtpClient: func(conn net.Conn, host string) (md.SmtpClient, error) {
					return &MockSMTPClient{mockRcptError: errors.New("rcpt error")}, nil
				},
			},
			args: args{
				options: cu.IM{
					"email": cu.IM{
						"from":       "test@example.com",
						"recipients": []interface{}{cu.IM{"email": "recipient@example.com"}},
						"subject":    "Test Email",
						"text":       "Test content",
					},
					"provider": "smtp",
				},
			},
			wantErr: true,
		},
		{
			name: "client data error",
			fields: fields{
				Config: cu.IM{
					"NT_SMTP_USER":            "test@example.com",
					"NT_SMTP_PASSWORD":        "test",
					"NT_SMTP_HOST":            "localhost",
					"NT_SMTP_PORT":            25,
					"NT_SMTP_TLS_MIN_VERSION": tls.VersionTLS12,
					"NT_SMTP_CONN":            "test",
					"NT_SMTP_AUTH":            "none",
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				NewSmtpClient: func(conn net.Conn, host string) (md.SmtpClient, error) {
					return &MockSMTPClient{mockDataError: errors.New("data error")}, nil
				},
			},
			args: args{
				options: cu.IM{
					"email": cu.IM{
						"from":       "test@example.com",
						"recipients": []interface{}{cu.IM{"email": "recipient@example.com"}},
						"subject":    "Test Email",
						"text":       "Test content",
					},
					"provider": "smtp",
				},
			},
			wantErr: true,
		},
		{
			name: "attachments error",
			fields: fields{
				Config: cu.IM{
					"NT_SMTP_USER":            "test@example.com",
					"NT_SMTP_PASSWORD":        "test",
					"NT_SMTP_HOST":            "localhost",
					"NT_SMTP_PORT":            25,
					"NT_SMTP_TLS_MIN_VERSION": tls.VersionTLS12,
					"NT_SMTP_CONN":            "test",
					"NT_SMTP_AUTH":            "none",
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				NewSmtpClient: func(conn net.Conn, host string) (md.SmtpClient, error) {
					return &MockSMTPClient{}, nil
				},
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{
								{"id": 1, "name": "test", "data": cu.IM{"template": string(sample_json)}},
							}, nil
						},
						"QuerySQL": func(sqlString string) ([]cu.IM, error) {
							return []cu.IM{}, nil
						},
					},
				},
				ReadFile: func(name string) ([]byte, error) {
					return []byte(`{"id": 1, "name": "test"}`), nil
				},
				ConvertFromByte: func(data []byte, result interface{}) error {
					return cu.ConvertFromByte(data, result)
				},
				ConvertToByte: func(data interface{}) ([]byte, error) {
					return []byte(`{"meta": {"report_key": "test", "report_name": "test", "report_type": "test", "file_type": "FILE_PDF"}}`), nil
				},
			},
			args: args{
				options: cu.IM{
					"email": cu.IM{
						"from":       "test@example.com",
						"recipients": []interface{}{cu.IM{"email": "recipient@example.com"}},
						"subject":    "Test Email",
						"text":       "Test content",
						"attachments": []interface{}{cu.IM{
							"report_key": "ntr_customer_en",
							"code":       "CUS0000000000N2",
							"filename":   "test.pdf",
						}},
					},
					"provider": "smtp",
				},
			},
			wantErr: true,
		},
		{
			name: "client write error",
			fields: fields{
				Config: cu.IM{
					"NT_SMTP_USER":            "test@example.com",
					"NT_SMTP_PASSWORD":        "test",
					"NT_SMTP_HOST":            "localhost",
					"NT_SMTP_PORT":            25,
					"NT_SMTP_TLS_MIN_VERSION": tls.VersionTLS12,
					"NT_SMTP_CONN":            "test",
					"NT_SMTP_AUTH":            "none",
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				NewSmtpClient: func(conn net.Conn, host string) (md.SmtpClient, error) {
					return &MockSMTPClient{mockWriteError: errors.New("write error")}, nil
				},
			},
			args: args{
				options: cu.IM{
					"email": cu.IM{
						"from":       "test@example.com",
						"recipients": []interface{}{cu.IM{"email": "recipient@example.com"}},
						"subject":    "Test Email",
						"text":       "Test content",
					},
					"provider": "smtp",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &DataStore{
				Db:                     tt.fields.Db,
				Alias:                  tt.fields.Alias,
				Config:                 tt.fields.Config,
				AppLog:                 tt.fields.AppLog,
				ReadAll:                tt.fields.ReadAll,
				ConvertToByte:          tt.fields.ConvertToByte,
				ConvertFromByte:        tt.fields.ConvertFromByte,
				ConvertFromReader:      tt.fields.ConvertFromReader,
				ConvertToType:          tt.fields.ConvertToType,
				GetDataField:           tt.fields.GetDataField,
				CreateLoginToken:       tt.fields.CreateLoginToken,
				ParseToken:             tt.fields.ParseToken,
				CreatePasswordHash:     tt.fields.CreatePasswordHash,
				ComparePasswordAndHash: tt.fields.ComparePasswordAndHash,
				ReadFile:               tt.fields.ReadFile,
				NewSmtpClient:          tt.fields.NewSmtpClient,
			}
			_, err := ds.SendEmail(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataStore.SendEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
