package utils

import (
	"image/color"
	"io"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestGetHash(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "hash_ok",
			args: args{
				text: "hello",
			},
			want: "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetHash(tt.args.text); got != tt.want {
				t.Errorf("GetHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToString(t *testing.T) {
	type args struct {
		value    interface{}
		defValue string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "string",
			args: args{
				value:    "test",
				defValue: "",
			},
			want: "test",
		},
		{
			name: "empty",
			args: args{
				value:    "",
				defValue: "",
			},
			want: "",
		},
		{
			name: "bool",
			args: args{
				value:    true,
				defValue: "",
			},
			want: "true",
		},
		{
			name: "int",
			args: args{
				value:    int(1),
				defValue: "",
			},
			want: "1",
		},
		{
			name: "int32",
			args: args{
				value:    int32(1),
				defValue: "",
			},
			want: "1",
		},
		{
			name: "int64",
			args: args{
				value:    int64(1),
				defValue: "",
			},
			want: "1",
		},
		{
			name: "float32",
			args: args{
				value:    float32(1),
				defValue: "",
			},
			want: "1",
		},
		{
			name: "float64",
			args: args{
				value:    float64(1.1),
				defValue: "",
			},
			want: "1.1",
		},
		{
			name: "time",
			args: args{
				value:    time.Now(),
				defValue: "",
			},
			want: time.Now().Format("2006-01-02T15:04:05-07:00"),
		},
		{
			name: "default",
			args: args{
				value:    []string{},
				defValue: "default",
			},
			want: "default",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToString(tt.args.value, tt.args.defValue); got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToFloat(t *testing.T) {
	type args struct {
		value    interface{}
		defValue float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "float64",
			args: args{
				value:    float64(1.1),
				defValue: float64(0),
			},
			want: float64(1.1),
		},
		{
			name: "0",
			args: args{
				value:    float64(0),
				defValue: float64(0),
			},
			want: float64(0),
		},
		{
			name: "bool",
			args: args{
				value:    true,
				defValue: float64(0),
			},
			want: float64(1),
		},
		{
			name: "int",
			args: args{
				value:    int(1),
				defValue: float64(0),
			},
			want: float64(1),
		},
		{
			name: "int32",
			args: args{
				value:    int32(1),
				defValue: float64(0),
			},
			want: float64(1),
		},
		{
			name: "int64",
			args: args{
				value:    int64(1),
				defValue: float64(0),
			},
			want: float64(1),
		},
		{
			name: "float32",
			args: args{
				value:    float32(1),
				defValue: float64(0),
			},
			want: float64(1),
		},
		{
			name: "string",
			args: args{
				value:    "1",
				defValue: float64(0),
			},
			want: float64(1),
		},
		{
			name: "default",
			args: args{
				value:    []string{},
				defValue: float64(0),
			},
			want: float64(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToFloat(tt.args.value, tt.args.defValue); got != tt.want {
				t.Errorf("ToFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToRGBA(t *testing.T) {
	type args struct {
		value    interface{}
		defValue color.RGBA
	}
	tests := []struct {
		name string
		args args
		want color.RGBA
	}{
		{
			name: "rgba",
			args: args{
				value:    color.RGBA{R: 1, G: 1, B: 1, A: 0},
				defValue: color.RGBA{R: 0, G: 0, B: 0, A: 0},
			},
			want: color.RGBA{R: 1, G: 1, B: 1, A: 0},
		},
		{
			name: "hex_ok",
			args: args{
				value:    "#ffffff",
				defValue: color.RGBA{R: 0, G: 0, B: 0, A: 0},
			},
			want: color.RGBA{R: 255, G: 255, B: 255, A: 0},
		},
		{
			name: "hex color must be 7 characters",
			args: args{
				value:    "#fffff",
				defValue: color.RGBA{R: 0, G: 0, B: 0, A: 0},
			},
			want: color.RGBA{R: 0, G: 0, B: 0, A: 0},
		},
		{
			name: "red component invalid",
			args: args{
				value:    "#xxffff",
				defValue: color.RGBA{R: 0, G: 0, B: 0, A: 0},
			},
			want: color.RGBA{R: 0, G: 0, B: 0, A: 0},
		},
		{
			name: "green component invalid",
			args: args{
				value:    "#ffxxff",
				defValue: color.RGBA{R: 0, G: 0, B: 0, A: 0},
			},
			want: color.RGBA{R: 0, G: 0, B: 0, A: 0},
		},
		{
			name: "blue component invalid",
			args: args{
				value:    "#ffffxx",
				defValue: color.RGBA{R: 0, G: 0, B: 0, A: 0},
			},
			want: color.RGBA{R: 0, G: 0, B: 0, A: 0},
		},
		{
			name: "int_string",
			args: args{
				value:    "200",
				defValue: color.RGBA{R: 0, G: 0, B: 0, A: 0},
			},
			want: color.RGBA{R: 200, G: 200, B: 200, A: 0},
		},
		{
			name: "int",
			args: args{
				value:    int(200),
				defValue: color.RGBA{R: 0, G: 0, B: 0, A: 0},
			},
			want: color.RGBA{R: 200, G: 200, B: 200, A: 0},
		},
		{
			name: "int32",
			args: args{
				value:    int32(200),
				defValue: color.RGBA{R: 0, G: 0, B: 0, A: 0},
			},
			want: color.RGBA{R: 200, G: 200, B: 200, A: 0},
		},
		{
			name: "int64",
			args: args{
				value:    int64(200),
				defValue: color.RGBA{R: 0, G: 0, B: 0, A: 0},
			},
			want: color.RGBA{R: 200, G: 200, B: 200, A: 0},
		},
		{
			name: "float32",
			args: args{
				value:    float32(200),
				defValue: color.RGBA{R: 0, G: 0, B: 0, A: 0},
			},
			want: color.RGBA{R: 200, G: 200, B: 200, A: 0},
		},
		{
			name: "float64",
			args: args{
				value:    float64(200),
				defValue: color.RGBA{R: 0, G: 0, B: 0, A: 0},
			},
			want: color.RGBA{R: 200, G: 200, B: 200, A: 0},
		},
		{
			name: "default",
			args: args{
				value:    []string{"200"},
				defValue: color.RGBA{R: 0, G: 0, B: 0, A: 0},
			},
			want: color.RGBA{R: 0, G: 0, B: 0, A: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToRGBA(tt.args.value, tt.args.defValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToRGBA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInteger(t *testing.T) {
	type args struct {
		value    interface{}
		defValue int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "int64",
			args: args{
				value:    int64(1),
				defValue: int64(0),
			},
			want: int64(1),
		},
		{
			name: "0",
			args: args{
				value:    int64(0),
				defValue: int64(0),
			},
			want: int64(0),
		},
		{
			name: "bool",
			args: args{
				value:    true,
				defValue: int64(0),
			},
			want: int64(1),
		},
		{
			name: "int",
			args: args{
				value:    int(1),
				defValue: int64(0),
			},
			want: int64(1),
		},
		{
			name: "int32",
			args: args{
				value:    int32(1),
				defValue: int64(0),
			},
			want: int64(1),
		},
		{
			name: "float32",
			args: args{
				value:    float32(1),
				defValue: int64(0),
			},
			want: int64(1),
		},
		{
			name: "float64",
			args: args{
				value:    float64(1),
				defValue: int64(0),
			},
			want: int64(1),
		},
		{
			name: "string",
			args: args{
				value:    "1",
				defValue: int64(0),
			},
			want: int64(1),
		},
		{
			name: "default",
			args: args{
				value:    []string{},
				defValue: int64(0),
			},
			want: int64(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInteger(tt.args.value, tt.args.defValue); got != tt.want {
				t.Errorf("ToInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToIntPointer(t *testing.T) {
	type args struct {
		value    interface{}
		defValue int64
	}
	tests := []struct {
		name string
		args args
		want *int64
	}{
		{
			name: "nil",
			args: args{
				value:    nil,
				defValue: int64(0),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToIntPointer(tt.args.value, tt.args.defValue); got != tt.want {
				t.Errorf("ToIntPointer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToStringPointer(t *testing.T) {
	type args struct {
		value    interface{}
		defValue string
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		{
			name: "nil",
			args: args{
				value:    nil,
				defValue: "",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToStringPointer(tt.args.value, tt.args.defValue); got != tt.want {
				t.Errorf("ToStringPointer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToBoolean(t *testing.T) {
	type args struct {
		value    interface{}
		defValue bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "bool",
			args: args{
				value:    true,
				defValue: false,
			},
			want: true,
		},
		{
			name: "int",
			args: args{
				value:    int(1),
				defValue: false,
			},
			want: true,
		},
		{
			name: "int32",
			args: args{
				value:    int32(1),
				defValue: false,
			},
			want: true,
		},
		{
			name: "int64",
			args: args{
				value:    int64(1),
				defValue: false,
			},
			want: true,
		},
		{
			name: "float32",
			args: args{
				value:    float32(1),
				defValue: false,
			},
			want: true,
		},
		{
			name: "float64",
			args: args{
				value:    float64(1),
				defValue: false,
			},
			want: true,
		},
		{
			name: "string",
			args: args{
				value:    "true",
				defValue: false,
			},
			want: true,
		},
		{
			name: "default",
			args: args{
				value:    []string{},
				defValue: false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToBoolean(tt.args.value, tt.args.defValue); got != tt.want {
				t.Errorf("ToBoolean() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToDateTime(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "time1",
			args: args{
				value: "2006-01-02T15:04:05-07:00",
			},
			wantErr: false,
		},
		{
			name: "time2",
			args: args{
				value: "2006-01-02T15:04:05-0700",
			},
			wantErr: false,
		},
		{
			name: "time3",
			args: args{
				value: "2006-01-02T15:04:05",
			},
			wantErr: false,
		},
		{
			name: "time4",
			args: args{
				value: "2006-01-02T15:04:05Z",
			},
			wantErr: false,
		},
		{
			name: "time5",
			args: args{
				value: "2006-01-02 15:04:05",
			},
			wantErr: false,
		},
		{
			name: "time6",
			args: args{
				value: "2006-01-02 15:04",
			},
			wantErr: false,
		},
		{
			name: "time7",
			args: args{
				value: "2006-01-02",
			},
			wantErr: false,
		},
		{
			name: "time8",
			args: args{
				value: "2006-01",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := StringToDateTime(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringToDateTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestFind(t *testing.T) {
	type args struct {
		a []string
		x string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "found",
			args: args{
				a: []string{"abba", "baba", "aabb"},
				x: "baba",
			},
			want: 1,
		},
		{
			name: "not_found",
			args: args{
				a: []string{"abba", "baba", "aabb"},
				x: "bbaa",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Find(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	type args struct {
		a []string
		x string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "found",
			args: args{
				a: []string{"abba", "baba", "aabb"},
				x: "baba",
			},
			want: true,
		},
		{
			name: "not_found",
			args: args{
				a: []string{"abba", "baba", "aabb"},
				x: "bbaa",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertToByte(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "json",
			args: args{
				data: `[{"field":"value"}]`,
			},
			wantErr: false,
			want:    []byte{34, 91, 123, 92, 34, 102, 105, 101, 108, 100, 92, 34, 58, 92, 34, 118, 97, 108, 117, 101, 92, 34, 125, 93, 34},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertToByte(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertToByte() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertFromByte(t *testing.T) {
	type args struct {
		data   []byte
		result interface{}
	}
	var result string
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "json",
			args: args{
				data:   []byte{34, 91, 123, 92, 34, 102, 105, 101, 108, 100, 92, 34, 58, 92, 34, 118, 97, 108, 117, 101, 92, 34, 125, 93, 34},
				result: &result,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ConvertFromByte(tt.args.data, tt.args.result); (err != nil) != tt.wantErr {
				t.Errorf("ConvertFromByte() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConvertFromReader(t *testing.T) {
	type args struct {
		data   io.Reader
		result interface{}
	}
	var result []map[string]interface{}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "error",
			args: args{
				data:   strings.NewReader(`[{"field":"value"}]`),
				result: &result,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ConvertFromReader(tt.args.data, tt.args.result); (err != nil) != tt.wantErr {
				t.Errorf("ConvertFromReader() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetMessage(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "cli_usage",
			args: args{
				key: "cli_usage",
			},
			want: "Program usage",
		},
		{
			name: "not_found",
			args: args{
				key: "usage",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMessage(tt.args.key); got != tt.want {
				t.Errorf("GetMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateToken(t *testing.T) {
	type args struct {
		username string
		database string
		config   map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "create_ok",
			args: args{
				username: "admin",
				database: "demo",
				config:   map[string]interface{}{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CreateToken(tt.args.username, tt.args.database, tt.args.config)
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
	token, _ := CreateToken("admin", "demo", make(map[string]interface{}))
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "decode_ok",
			args: args{
				tokenString: token,
			},
			wantErr: false,
		},
		{
			name: "decode_error",
			args: args{
				tokenString: "eyJhbGciOiJIUzI1NiIsImtpZCI6ImFmNTY2ZmQ4MDEyMTEyMzhhMTgxYzEzMWEw",
			},
			wantErr: true,
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
		key map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "RSA_private",
			args: args{
				key: map[string]string{
					"ktype": "RSA",
					"type":  "private",
					"value": "ABABABA",
				},
			},
			wantErr: true,
		},
		{
			name: "ECP_private",
			args: args{
				key: map[string]string{
					"ktype": "ECP",
					"type":  "private",
					"value": "ABABABA",
				},
			},
			wantErr: true,
		},
		{
			name: "RSA_public",
			args: args{
				key: map[string]string{
					"ktype": "RSA",
					"type":  "public",
					"value": "ABABABA",
				},
			},
			wantErr: true,
		},
		{
			name: "ECP_public",
			args: args{
				key: map[string]string{
					"ktype": "ECP",
					"type":  "public",
					"value": "ABABABA",
				},
			},
			wantErr: true,
		},
		{
			name: "not_found",
			args: args{
				key: map[string]string{
					"ktype": "BSA",
					"type":  "public",
					"value": "ABABABA",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := parsePEM(tt.args.key)
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
		keyMap      map[string]map[string]string
		config      map[string]interface{}
	}
	token, _ := CreateToken("admin", "demo", make(map[string]interface{}))
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "token_username",
			args: args{
				tokenString: token,
				keyMap:      make(map[string]map[string]string),
				config:      make(map[string]interface{}),
			},
			wantErr: false,
		},
		{
			name: "invalid_token",
			args: args{
				tokenString: "eyJhbGciOiJIUzI1NiIsImtpZCI6ImFmNTY2ZmQ4MDEyMTEyMzhhMTgxYzEzMWEwNzI1MzM2MTU3Y",
				keyMap:      make(map[string]map[string]string),
				config:      make(map[string]interface{}),
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

func TestRandString(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "create",
			args: args{
				length: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandString(tt.args.length); len(got) != tt.args.length {
				t.Errorf("RandString() = %v, want %v", len(got), tt.args.length)
			}
		})
	}
}
