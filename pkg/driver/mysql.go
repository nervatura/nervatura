//go:build mysql || all
// +build mysql all

package database

import _ "github.com/go-sql-driver/mysql" // mysql driver

func init() {
	registerDriver("mysql")
}
