package utils

import (
	"testing"
)

func TestGetHash(t *testing.T) {
	type args struct {
		text string
		code string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "sha_hash_ok",
			args: args{
				text: "hello",
				code: "sha25",
			},
			want: "2CF24DBA5FB0A30E26E83B2AC5B9E29E1B161E5C1FA7425E73043362938B9824",
		},
		{
			name: "sha3_hash_ok",
			args: args{
				text: "hello",
				code: "sha512",
			},
			want: "9B71D224BD62F3785D96D46AD3EA3D73319BFBC2890CAADAE2DFF72519673CA72323C3D99BA5C11D7C7ACC6E14B8C5DA0C4663475C2E5C3ADEF46F73BCDEC043",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetHash(tt.args.text, tt.args.code); got != tt.want {
				t.Errorf("GetHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreatePasswordHash(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "password_hash_ok",
			args: args{
				password: "password",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CreatePasswordHash(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreatePasswordHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestComparePasswordAndHash(t *testing.T) {
	type args struct {
		password string
		hash     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "password_hash_ok",
			args: args{
				password: "password",
				hash:     "password",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ComparePasswordAndHash(tt.args.password, tt.args.hash); (err != nil) != tt.wantErr {
				t.Errorf("ComparePasswordAndHash() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
