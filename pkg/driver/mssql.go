//go:build mssql || all
// +build mssql all

package database

//import _ "github.com/denisenkom/go-mssqldb" // mssql driver

func init() {
	registerDriver("mssql")
}
