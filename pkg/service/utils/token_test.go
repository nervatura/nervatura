package utils

import (
	"testing"

	cu "github.com/nervatura/component/pkg/util"
)

func TestCreateLoginToken(t *testing.T) {
	type args struct {
		code     string
		userName string
		alias    string
		config   cu.IM
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "missing_fieldname",
			args: args{
				code:     "",
				userName: "",
				alias:    "",
				config:   cu.IM{},
			},
			wantErr: true,
		},
		{
			name: "create",
			args: args{
				code:     "code",
				userName: "user",
				alias:    "alias",
				config:   cu.IM{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CreateLoginToken(tt.args.code, tt.args.userName, tt.args.alias, tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateLoginToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestCreateToken(t *testing.T) {
	type args struct {
		subject string
		claims  cu.IM
		config  cu.IM
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "create",
			args: args{
				subject: "",
				claims:  cu.IM{},
				config:  cu.IM{},
			},
			wantErr: false,
		},
		{
			name: "alg_err",
			args: args{
				subject: "",
				claims:  cu.IM{},
				config: cu.IM{
					"NT_TOKEN_ALG": "MISSING",
				},
			},
			wantErr: true,
		},
		{
			name: "key_err",
			args: args{
				subject: "",
				claims:  cu.IM{},
				config: cu.IM{
					"NT_TOKEN_ALG":         "RS512",
					"NT_TOKEN_PRIVATE_KEY": "MARHASAG",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CreateToken(tt.args.subject, tt.args.claims, tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestTokenDecode(t *testing.T) {
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "decode_error",
			args: args{
				tokenString: "eyJhbGciOiJIUzI1NiIsImtpZCI6ImFmNTY2ZmQ4MDEyMTEyMzhhMTgxYzEzMWEw",
			},
			wantErr: true,
		},
		{
			name: "decode_ok",
			args: args{
				tokenString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJJRDAxMjMiLCJlbWFpbCI6InVzZXJAbWFpbC5jb20iLCJpc3MiOiJlb25leCJ9.1TKb3k2xYslwhPDQD50AuSsFqRlIxUB7kErrQqaHVIg",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := TokenDecode(tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("TokenDecode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_parsePEM(t *testing.T) {
	type args struct {
		method string
		stype  string
		value  []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "RSA_private",
			args: args{
				method: "RSA",
				stype:  "private",
				value:  []byte("ABABABA"),
			},
			wantErr: true,
		},
		{
			name: "RSA_public",
			args: args{
				method: "RSA",
				stype:  "public",
				value:  []byte("ABABABA"),
			},
			wantErr: true,
		},
		{
			name: "ECDSA_private",
			args: args{
				method: "ECDSA",
				stype:  "private",
				value:  []byte("ABABABA"),
			},
			wantErr: true,
		},
		{
			name: "ECDSA_public",
			args: args{
				method: "ECDSA",
				stype:  "public",
				value:  []byte("ABABABA"),
			},
			wantErr: true,
		},
		{
			name: "EdDSA_private",
			args: args{
				method: "EdDSA",
				stype:  "private",
				value:  []byte("ABABABA"),
			},
			wantErr: true,
		},
		{
			name: "EdDSA_public",
			args: args{
				method: "EdDSA",
				stype:  "public",
				value:  []byte("ABABABA"),
			},
			wantErr: true,
		},
		{
			name: "not_found",
			args: args{
				method: "BSA",
				stype:  "public",
				value:  []byte("ABABABA"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := parsePEM(tt.args.method, tt.args.stype, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("parsePEM() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestParseToken(t *testing.T) {
	type args struct {
		tokenString string
		keyMap      []cu.SM
		config      cu.IM
	}
	validToken, _ := CreateToken("ID0123", map[string]interface{}{"email": "user@mail.com"}, map[string]interface{}{
		"NT_TOKEN_PRIVATE_KEY": "SECRET_KEY",
	})
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "invalid_token",
			args: args{
				tokenString: "eyJhbGciOiJIUzI1NiIsImtpZCI6ImFmNTY2ZmQ4MDEyMTEyMzhhMTgxYzEzMWEwNzI1MzM2MTU3Y",
				keyMap:      make([]map[string]string, 0),
				config:      make(map[string]interface{}),
			},
			wantErr: true,
		},
		{
			name: "token_ok",
			args: args{
				tokenString: validToken,
				keyMap: []map[string]string{
					{"type": "private", "value": "SECRET_KEY"},
				},
				config: map[string]interface{}{
					"NT_TOKEN_PRIVATE_KEY": "SECRET_KEY",
				},
			},
			wantErr: false,
		},
		{
			name: "token_expired",
			args: args{
				tokenString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJJRDAxMjMiLCJlbWFpbCI6InVzZXJAbWFpbC5jb20iLCJpc3MiOiJlb25leCIsImV4cCI6MTY5MTY4OTk1Nn0.hUyXG_mFWlx_cuLJk1McsyN-1ddJHJ7iD8VM3-fnvA4",
				keyMap: []map[string]string{
					{"type": "private", "value": "SECRET_KEY"},
				},
				config: map[string]interface{}{
					"NT_TOKEN_PRIVATE_KEY": "SECRET_KEY",
				},
			},
			wantErr: true,
		},
		{
			name: "public_key",
			args: args{
				tokenString: `eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6IlRFU1QifQ.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTc0NzE1ODQwMCwiZXhwIjoxNzQ3MTYyMDAwfQ.097g1tByUgYwXLyjUa5p7rmQH1pShL4LKmr8Vr8XHv1vJf2QiP46kXxTaMl_KVVserbxnRaA1IXY1FIBeeNffBR-UE3HCpgKKXv-Fe4eT_8ElmsP1z6uFCjWaw_s-42xdKwua2LnZxCTqlcyA7777auOiraW07xoxLs_7PjSZbBdpcDRUQsDS3l_0JzUJXryotXHEFYy6429TANkGlLUlpz5nunSv0bcV-wMDGQriVc46oF_FpLC4SgFbtdh3LXiOfoieKXo2V8xN5igupiA6VGhno0b9TTGQ4a-7VhwRz25F69Se8ei6sTp1DFL7bHp9zPJ_Lm7uoYFJVZ2elmfcg`,
				keyMap: []map[string]string{
					{"type": "public", "kid": "TEST",
						"value": `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA6S7asUuzq5Q/3U9rbs+P
kDVIdjgmtgWreG5qWPsC9xXZKiMV1AiV9LXyqQsAYpCqEDM3XbfmZqGb48yLhb/X
qZaKgSYaC/h2DjM7lgrIQAp9902Rr8fUmLN2ivr5tnLxUUOnMOc2SQtr9dgzTONY
W5Zu3PwyvAWk5D6ueIUhLtYzpcB+etoNdL3Ir2746KIy/VUsDwAM7dhrqSK8U2xF
CGlau4ikOTtvzDownAMHMrfE7q1B6WZQDAQlBmxRQsyKln5DIsKv6xauNsHRgBAK
ctUxZG8M4QJIx3S6Aughd3RZC4Ca5Ae9fd8L8mlNYBCrQhOZ7dS0f4at4arlLcaj
twIDAQAB
-----END PUBLIC KEY-----`},
				},
				config: map[string]interface{}{
					"NT_TOKEN_PRIVATE_KEY": "SECRET_KEY",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid_key",
			args: args{
				tokenString: validToken,
				keyMap: []map[string]string{
					{"type": "private", "value": "INVALID_KEY"},
				},
				config: map[string]interface{}{
					"NT_TOKEN_PRIVATE_KEY": "SECRET_KEY",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid_alg",
			args: args{
				tokenString: "eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJJRDAxMjMiLCJlbWFpbCI6InVzZXJAbWFpbC5jb20iLCJpc3MiOiJlb25leCJ9.EbaxBapncFNTbFjm9Ti2AqRuR3I8_DAYN4p_9qFQoTM",
				keyMap:      make([]map[string]string, 0),
				config: map[string]interface{}{
					"NT_TOKEN_PRIVATE_KEY": "SECRET_KEY",
				},
			},
			wantErr: true,
		},
		{
			name: "empty_keymap",
			args: args{
				tokenString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJJRDAxMjMiLCJlbWFpbCI6InVzZXJAbWFpbC5jb20iLCJpc3MiOiJlb25leCJ9.1TKb3k2xYslwhPDQD50AuSsFqRlIxUB7kErrQqaHVIg",
				keyMap:      []map[string]string{},
				config: map[string]interface{}{
					"NT_TOKEN_PRIVATE_KEY": "SECRET_KEY",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseToken(tt.args.tokenString, tt.args.keyMap, tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
