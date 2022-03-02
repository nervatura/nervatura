package main

import (
	"C"

	"github.com/nervatura/nervatura-service/app"
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
		return C.CString(`{"code":0,"error": "` + err.Error() + `"}`)
	}
	return C.CString(app.GetResults())
}

//export DatabaseCreate
func DatabaseCreate(key *C.char, options *C.char) *C.char {
	return respondData(app.New(version, map[string]string{
		"cmd": "DatabaseCreate", "options": C.GoString(options), "key": C.GoString(key),
	}))
}

//export UserLogin
func UserLogin(options *C.char) *C.char {
	return respondData(app.New(version, map[string]string{
		"cmd": "UserLogin", "options": C.GoString(options),
	}))
}

//export TokenLogin
func TokenLogin(token *C.char) *C.char {
	return respondData(app.New(version, map[string]string{
		"cmd": "TokenLogin", "token": C.GoString(token),
	}))
}

//export TokenRefresh
func TokenRefresh(token *C.char) *C.char {
	return respondData(app.New(version, map[string]string{
		"cmd": "TokenRefresh", "token": C.GoString(token),
	}))
}

//export TokenDecode
func TokenDecode(token *C.char) *C.char {
	return respondData(app.New(version, map[string]string{
		"cmd": "TokenDecode", "token": C.GoString(token),
	}))
}

//export UserPassword
func UserPassword(token *C.char, options *C.char) *C.char {
	return respondData(app.New(version, map[string]string{
		"cmd": "UserPassword", "options": C.GoString(options), "token": C.GoString(token),
	}))
}

//export Get
func Get(token *C.char, options *C.char) *C.char {
	return respondData(app.New(version, map[string]string{
		"cmd": "Get", "options": C.GoString(options), "token": C.GoString(token),
	}))
}

//export View
func View(token *C.char, data *C.char) *C.char {
	return respondData(app.New(version, map[string]string{
		"cmd": "View", "data": C.GoString(data), "token": C.GoString(token),
	}))
}

//export Function
func Function(token *C.char, options *C.char) *C.char {
	return respondData(app.New(version, map[string]string{
		"cmd": "Function", "options": C.GoString(options), "token": C.GoString(token),
	}))
}

//export Update
func Update(token *C.char, nervatype *C.char, data *C.char) *C.char {
	return respondData(app.New(version, map[string]string{
		"cmd": "Update", "nervatype": C.GoString(nervatype), "data": C.GoString(data), "token": C.GoString(token),
	}))
}

//export Delete
func Delete(token *C.char, options *C.char) *C.char {
	return respondData(app.New(version, map[string]string{
		"cmd": "Delete", "options": C.GoString(options), "token": C.GoString(token),
	}))
}

//export Report
func Report(token *C.char, options *C.char) *C.char {
	return respondData(app.New(version, map[string]string{
		"cmd": "Report", "options": C.GoString(options), "token": C.GoString(token),
	}))
}

//export ReportList
func ReportList(token *C.char, options *C.char) *C.char {
	return respondData(app.New(version, map[string]string{
		"cmd": "ReportList", "options": C.GoString(options), "token": C.GoString(token),
	}))
}

//export ReportDelete
func ReportDelete(token *C.char, options *C.char) *C.char {
	return respondData(app.New(version, map[string]string{
		"cmd": "ReportDelete", "options": C.GoString(options), "token": C.GoString(token),
	}))
}

//export ReportInstall
func ReportInstall(token *C.char, options *C.char) *C.char {
	return respondData(app.New(version, map[string]string{
		"cmd": "ReportInstall", "options": C.GoString(options), "token": C.GoString(token),
	}))
}

func main() {}
