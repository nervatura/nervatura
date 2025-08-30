package main

import (
	"C"

	"github.com/nervatura/nervatura/v6/pkg/app"
)

var (
	version = "shared"
)

//export Version
func Version() *C.char {
	return C.CString(version)
}

func respondData(app *app.App, err error) *C.char {
	if err != nil {
		return C.CString(`{"code":500,"data": "` + err.Error() + `"}`)
	}
	return C.CString(app.GetResults())
}

//export Database
func Database(options *C.char) *C.char {
	return respondData(app.New(version, map[string]string{
		"cmd": "database", "options": C.GoString(options),
	}))
}

//export Function
func Function(options *C.char) *C.char {
	return respondData(app.New(version, map[string]string{
		"cmd": "function", "options": C.GoString(options),
	}))
}

//export ResetPassword
func ResetPassword(options *C.char) *C.char {
	return respondData(app.New(version, map[string]string{
		"cmd": "reset", "options": C.GoString(options),
	}))
}

//export Create
func Create(model *C.char, options *C.char, data *C.char) *C.char {
	return respondData(app.New(version, map[string]string{
		"cmd": "create", "model": C.GoString(model), "options": C.GoString(options), "data": C.GoString(data),
	}))
}

//export Update
func Update(model *C.char, options *C.char, data *C.char) *C.char {
	return respondData(app.New(version, map[string]string{
		"cmd": "update", "model": C.GoString(model), "options": C.GoString(options), "data": C.GoString(data),
	}))
}

//export Delete
func Delete(model *C.char, options *C.char) *C.char {
	return respondData(app.New(version, map[string]string{
		"cmd": "delete", "model": C.GoString(model), "options": C.GoString(options),
	}))
}

//export Get
func Get(model *C.char, options *C.char) *C.char {
	return respondData(app.New(version, map[string]string{
		"cmd": "get", "model": C.GoString(model), "options": C.GoString(options),
	}))
}

//export Query
func Query(model *C.char, options *C.char) *C.char {
	return respondData(app.New(version, map[string]string{
		"cmd": "query", "model": C.GoString(model), "options": C.GoString(options),
	}))
}

//export View
func View(options *C.char) *C.char {
	return respondData(app.New(version, map[string]string{
		"cmd": "view", "options": C.GoString(options),
	}))
}

func main() {}
