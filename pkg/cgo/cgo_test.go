package main

import (
	"testing"

	"github.com/nervatura/nervatura/v6/pkg/app"
)

func TestDatabase(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		options *_Ctype_char
	}{
		{
			name:    "test_database",
			options: _Cfunc_CString("test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Database(tt.options)
		})
	}
}

func TestVersion(t *testing.T) {
	tests := []struct {
		name string // description of this test case
	}{
		{
			name: "test_version",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Version()
		})
	}
}

func TestFunction(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		options *_Ctype_char
	}{
		{
			name:    "test_function",
			options: _Cfunc_CString("test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Function(tt.options)
		})
	}
}

func TestResetPassword(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		options *_Ctype_char
	}{
		{
			name:    "test_reset_password",
			options: _Cfunc_CString("test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ResetPassword(tt.options)
		})
	}
}

func TestCreate(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		model   *_Ctype_char
		options *_Ctype_char
		data    *_Ctype_char
	}{
		{
			name:    "test_create",
			model:   _Cfunc_CString("test"),
			options: _Cfunc_CString("test"),
			data:    _Cfunc_CString("test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Create(tt.model, tt.options, tt.data)
		})
	}
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		model   *_Ctype_char
		options *_Ctype_char
		data    *_Ctype_char
	}{
		{
			name:    "test_update",
			model:   _Cfunc_CString("test"),
			options: _Cfunc_CString("test"),
			data:    _Cfunc_CString("test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Update(tt.model, tt.options, tt.data)
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		model   *_Ctype_char
		options *_Ctype_char
	}{
		{
			name:    "test_delete",
			model:   _Cfunc_CString("test"),
			options: _Cfunc_CString("test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Delete(tt.model, tt.options)
		})
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		model   *_Ctype_char
		options *_Ctype_char
	}{
		{
			name:    "test_get",
			model:   _Cfunc_CString("test"),
			options: _Cfunc_CString("test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Get(tt.model, tt.options)
		})
	}
}

func TestQuery(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		model   *_Ctype_char
		options *_Ctype_char
	}{
		{
			name:    "test_query",
			model:   _Cfunc_CString("test"),
			options: _Cfunc_CString("test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Query(tt.model, tt.options)
		})
	}
}

func TestView(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		options *_Ctype_char
	}{
		{
			name:    "test_view",
			options: _Cfunc_CString("test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			View(tt.options)
		})
	}
}

func Test_respondData(t *testing.T) {
	ap, _ := app.New("test", map[string]string{
		"cmd": "help",
	})
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		app *app.App
		err error
	}{
		{
			name: "test_respondData",
			app:  ap,
			err:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			respondData(tt.app, tt.err)
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string // description of this test case
	}{
		{
			name: "test_main",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
