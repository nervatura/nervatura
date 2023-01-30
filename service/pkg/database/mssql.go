//go:build mssql || all
// +build mssql all

package nervatura

import _ "github.com/denisenkom/go-mssqldb" // mssql driver

func init() {
	registerDriver("mssql")
}
