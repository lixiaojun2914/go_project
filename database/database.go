package database

import (
	"database/sql"
	"fmt"
	"github.com/golang/glog"
	_ "github.com/lib/pq"
	"test/status/error_code"
)

func Init() {
	defer glog.Flush()
	//connStr := os.Getenv("DBCONNSTR")
	cc := "host=127.0.0.1 port=5432 user=postgres dbname=testDB password=lxj sslmode=disable"
	db, err := sql.Open("postgres", cc)
	if err != nil {
		glog.Errorf("open db error with code: %d", error_code.DBOPENERR)
		glog.Errorln(err)
		return
	}
	rows, err := db.Query(`SELECT $1 FROM test`, "test")
	defer rows.Close()
	if err != nil {
		glog.Errorln(err)
	}
	fmt.Println(rows.Columns())
}
