package mcp

import (
	"bytes"
	"context"
	"crypto/tls"
	"errors"
	"log/slog"
	"net"
	"testing"
	"time"

	"github.com/modelcontextprotocol/go-sdk/auth"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
)

func Test_emailSendHandler(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		req        *mcp.CallToolRequest
		parameters cu.IM
		wantErr    bool
		ds         *api.DataStore
	}{
		{
			name: "success",
			req: &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_report_query"},
				Extra: &mcp.RequestExtra{
					TokenInfo: &auth.TokenInfo{Expiration: time.Now().Add(time.Duration(1) * time.Hour)},
				}},
			parameters: cu.IM{
				"code": "CUS1731101982N123", "report_key": "", "recipients": []string{"test@example.com"}, "subject": "Test Email",
			},
			wantErr: true,
			ds: &api.DataStore{
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
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{
								{"id": 1, "name": "test", "report_key": "ntr_customer_en", "data": cu.IM{"template": ""}},
							}, nil
						},
						"QuerySQL": func(sqlString string) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "report_key": "ntr_customer_en", "data": cu.IM{"file_type": "FILE_PDF"}}}, nil
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
		},
		{
			name: "invalid code",
			parameters: cu.IM{
				"code": "XXX1731101982N123", "report_key": "", "recipients": []string{"test@example.com"}, "subject": "Test Email",
			},
			wantErr: true,
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{}, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
		},
		{
			name: "missing subject",
			parameters: cu.IM{
				"code": "XXX1731101982N123", "report_key": "", "recipients": []string{"test@example.com"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			_, _, gotErr := emailSendHandler(ctx, tt.req, tt.parameters)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("emailSendHandler() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("emailSendHandler() succeeded unexpectedly")
			}
		})
	}
}
