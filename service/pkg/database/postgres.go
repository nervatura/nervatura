//+build postgres all

package nervatura

import _ "github.com/lib/pq" // postgres driver

func init() {
	registerDriver("postgres")
}
