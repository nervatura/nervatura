//go:build postgres || all
// +build postgres all

package database

import _ "github.com/lib/pq" // postgres driver

func init() {
	registerDriver("postgres")
}
