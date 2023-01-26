//go:build sqlite || all
// +build sqlite all

package nervatura

//import _ "modernc.org/sqlite" //sqlite driver
//import _ "github.com/mattn/go-sqlite3" //sqlite driver
import _ "github.com/glebarez/go-sqlite" //sqlite driver

func init() {
	registerDriver("sqlite")
}
