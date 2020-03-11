package database

import (
	"database/sql"
	"fmt"
	"test/status/error_code"

	"github.com/golang/glog"
	_ "github.com/lib/pq"
)

type Test struct {
	Aaa int
	Bbb string
}

func InitDb() {
	defer glog.Flush()
	//connStr := os.Getenv("DBCONNSTR")
	cc := "host=127.0.0.1 port=5432 user=lixiaojun dbname=mydb password='123456' sslmode=disable"
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
	test := Test{}
	for rows.Next() {
		if err := rows.Scan(&test); err != nil {
			glog.Errorln("error")
			break
		}
		fmt.Println(test.Aaa, test.Bbb)
	}
}
