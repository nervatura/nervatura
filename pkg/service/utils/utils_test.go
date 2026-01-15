package utils

import (
	"net"
	"reflect"
	"testing"

	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	td "github.com/nervatura/nervatura/v6/test/driver"
)

func TestSMToJS(t *testing.T) {
	type args struct {
		sm cu.SM
	}
	tests := []struct {
		name string
		args args
		want cu.IM
	}{
		{
			name: "test",
			args: args{sm: cu.SM{"a": "1", "b": "2", "c": "true", "d": "false", "e": "{}", "err": "[}"}},
			want: cu.IM{"a": float64(1), "b": float64(2), "c": true, "d": false, "e": cu.IM{}, "err": "[}"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SMToJS(tt.args.sm); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SMToJS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertToType(t *testing.T) {
	type args struct {
		data   interface{}
		result any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				data:   cu.SM{"a": "1", "b": "2", "c": "true", "d": "false", "e": "{}", "err": "[}"},
				result: cu.IM{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ConvertToType(tt.args.data, tt.args.result); (err != nil) != tt.wantErr {
				t.Errorf("ConvertToType() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIntArrayToString(t *testing.T) {
	type args struct {
		arr    []int64
		prefix bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test_true",
			args: args{[]int64{1, 2, 3}, true},
			want: "{1,2,3}",
		},
		{
			name: "test_false",
			args: args{[]int64{1, 2, 3}, false},
			want: "1,2,3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntArrayToString(tt.args.arr, tt.args.prefix); got != tt.want {
				t.Errorf("IntArrayToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToIntArray(t *testing.T) {
	type args struct {
		arr interface{}
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int64
	}{
		{
			name: "int64",
			args: args{
				arr: []int64{1, 2, 3},
			},
			wantResult: []int64{1, 2, 3},
		},
		{
			name: "csv",
			args: args{
				arr: "1,2,3",
			},
			wantResult: []int64{1, 2, 3},
		},
		{
			name: "string",
			args: args{
				arr: []string{"1", "2", "3"},
			},
			wantResult: []int64{1, 2, 3},
		},
		{
			name: "interface{}",
			args: args{
				arr: []interface{}{"1", "2", "3"},
			},
			wantResult: []int64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := ToIntArray(tt.args.arr); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ToIntArray() = %v, want %v", gotResult, tt.wantResult)
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
			name: "shutdown_signal",
			args: args{
				key: "shutdown_signal",
			},
			want: "received shut down signal",
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

func TestGetMessages(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetMessages()
		})
	}
}

func TestGetDataField(t *testing.T) {
	type args struct {
		data     any
		JSONName string
	}
	tests := []struct {
		name           string
		args           args
		wantFieldName  string
		wantFieldValue interface{}
	}{
		{
			name: "ok",
			args: args{
				data: md.Auth{
					UserGroup: md.UserGroupUser,
				},
				JSONName: "user_group",
			},
			wantFieldName:  "UserGroup",
			wantFieldValue: md.UserGroupUser,
		},
		{
			name: "missing",
			args: args{
				data: md.Auth{
					UserGroup: md.UserGroupUser,
				},
				JSONName: "missing",
			},
			wantFieldName:  "",
			wantFieldValue: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFieldName, gotFieldValue := GetDataField(tt.args.data, tt.args.JSONName)
			if gotFieldName != tt.wantFieldName {
				t.Errorf("GetDataField() gotFieldName = %v, want %v", gotFieldName, tt.wantFieldName)
			}
			if !reflect.DeepEqual(gotFieldValue, tt.wantFieldValue) {
				t.Errorf("GetDataField() gotFieldValue = %v, want %v", gotFieldValue, tt.wantFieldValue)
			}
		})
	}
}

func TestGetSessionID(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetSessionID()
		})
	}
}

func TestSmtpClient(t *testing.T) {
	type args struct {
		conn net.Conn
		host string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				conn: td.NewTestConn(),
				host: "localhost",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := SmtpClient(tt.args.conn, tt.args.host)
			if (err != nil) != tt.wantErr {
				t.Errorf("SmtpClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestConvertByteToIMValue(t *testing.T) {
	type args struct {
		data      any
		initValue any
		imap      cu.IM
		key       string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				data:      nil,
				initValue: 1,
				imap:      cu.IM{},
				key:       "key",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ConvertByteToIMValue(tt.args.data, tt.args.initValue, tt.args.imap, tt.args.key)
		})
	}
}

func TestConvertByteToIMData(t *testing.T) {
	type args struct {
		data any
		imap cu.IM
		key  string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				data: []byte("test"),
				imap: cu.IM{},
				key:  "key",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ConvertByteToIMData(tt.args.data, tt.args.imap, tt.args.key)
		})
	}
}

func TestToStringArray(t *testing.T) {
	type args struct {
		arr interface{}
	}
	tests := []struct {
		name       string
		args       args
		wantResult []string
	}{
		{
			name: "success",
			args: args{
				arr: []string{"a", "b", "c"},
			},
			wantResult: []string{"a", "b", "c"},
		},
		{
			name: "interface{}",
			args: args{
				arr: []interface{}{"a", "b", "c"},
			},
			wantResult: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := ToStringArray(tt.args.arr); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ToStringArray() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestToBoolMap(t *testing.T) {
	type args struct {
		im       interface{}
		defValue map[string]bool
	}
	tests := []struct {
		name       string
		args       args
		wantResult map[string]bool
	}{
		{
			name: "nil defvalue",
			args: args{
				im:       nil,
				defValue: map[string]bool{"a": true, "b": false},
			},
			wantResult: map[string]bool{"a": true, "b": false},
		},
		{
			name: "success",
			args: args{
				im:       map[string]bool{"a": true, "b": false},
				defValue: map[string]bool{"a": true, "b": false},
			},
			wantResult: map[string]bool{"a": true, "b": false},
		},
		{
			name: "interface{}",
			args: args{
				im:       map[string]interface{}{"a": true, "b": false},
				defValue: map[string]bool{"a": true, "b": false},
			},
			wantResult: map[string]bool{"a": true, "b": false},
		},
		{
			name: "invalid",
			args: args{
				im:       "invalid",
				defValue: map[string]bool{"a": true, "b": false},
			},
			wantResult: map[string]bool{"a": true, "b": false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := ToBoolMap(tt.args.im, tt.args.defValue); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ToBoolMap() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestSortIMData(t *testing.T) {
	type args struct {
		data      []map[string]interface{}
		sortField string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				data:      []map[string]interface{}{{"a": 1, "b": 2}, {"a": 3, "b": 4}},
				sortField: "a",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SortIMData(tt.args.data, tt.args.sortField)
		})
	}
}

func TestToTagList(t *testing.T) {
	type args struct {
		tags []string
	}
	tests := []struct {
		name string
		args args
		want []cu.IM
	}{
		{
			name: "success",
			args: args{
				tags: []string{"a", "b", "c"},
			},
			want: []cu.IM{{"tag": "a"}, {"tag": "b"}, {"tag": "c"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToTagList(tt.args.tags); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToTagList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMetaName(t *testing.T) {
	type args struct {
		mp  cu.IM
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{
				mp:  cu.IM{"a_meta": 1, "b_metaxa": 2},
				key: "meta",
			},
			want: "a_meta",
		},
		{
			name: "not_found",
			args: args{
				mp:  cu.IM{"a": 1, "b": 2},
				key: "meta",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MetaName(tt.args.mp, tt.args.key); got != tt.want {
				t.Errorf("MetaName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnyPointer(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				v: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AnyPointer(tt.args.v)
		})
	}
}

func TestToAnyArray(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		arr  any
		want []any
	}{
		{
			name: "string",
			arr:  []string{"a", "b", "c"},
			want: []any{"a", "b", "c"},
		},
		{
			name: "int64",
			arr:  []int64{1, 2, 3},
			want: []any{int64(1), int64(2), int64(3)},
		},
		{
			name: "float64",
			arr:  []float64{1.1, 2.2, 3.3},
			want: []any{1.1, 2.2, 3.3},
		},
		{
			name: "any",
			arr:  []any{"a", "b", "c"},
			want: []any{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToAnyArray(tt.arr)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToAnyArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
