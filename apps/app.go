package apps

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type App struct {
	sql.ColumnType
}
