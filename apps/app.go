package apps

import "github.com/lib/pq"

type App struct {
	db *pq.Driver
}
