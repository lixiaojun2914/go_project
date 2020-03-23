package main

import (
	"flag"
	"test/databases"
	"test/router"

	"github.com/gin-gonic/gin"
)

func main() {
	flag.Parse()
	r := gin.Default()
	router.InitRouter(r)
	databases.InitDb()
	// r.Run()
}
