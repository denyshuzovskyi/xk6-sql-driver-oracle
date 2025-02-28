// Package oracle contains godror oracle driver registration for xk6-sql.
package oracle

import (
	"github.com/grafana/xk6-sql/sql"

	// Blank import required for initialization of driver.
	_ "github.com/godror/godror"
)

func init() {
	sql.RegisterModule("godror")
}
