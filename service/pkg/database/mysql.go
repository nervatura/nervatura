//+build mysql all

package nervatura

import _ "github.com/go-sql-driver/mysql" // mysql driver

func init() {
	registerDriver("mysql")
}
