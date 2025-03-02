// Package oracle contains go-ora oracle driver registration for xk6-sql.
package oracle

import (
	"github.com/grafana/xk6-sql/sql"
	// Blank import required for initialization of driver.
	_ "github.com/sijms/go-ora/v2"
)

func init() {
	sql.RegisterModule("oracle")
}
