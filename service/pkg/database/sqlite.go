//+build sqlite all

package nervatura

import _ "modernc.org/sqlite" //sqlite driver
//import _ "github.com/mattn/go-sqlite3" //sqlite driver

func init() {
	registerDriver("sqlite")
}
