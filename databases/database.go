package databases

import (
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Test fun
type Test struct {
	gorm.Model
	Name string `gorm:"varchar(20) unique"`
	Age  int    `gorm:"unique"`
}

//InitDb fun
func InitDb() {
	defer glog.Flush()
	//connStr := os.Getenv("DBCONNSTR")
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=lixiaojun dbname=mydb password=123456 sslmode=disable")
	db.SingularTable(true)
	if err != nil {
		glog.Error("db connect error with", err)
	}
	// db.CreateTable(&Test{})
	user := Test{Name: "lixiaojun", Age: 456}
	// db.Create(&user)
	// m := db.NewRecord(user)
	// fmt.Println(m)
	db.Delete(&user)
}
