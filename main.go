package main

import (
	"flag"
	"test/router"

	"github.com/gin-gonic/gin"
)

func main() {
	flag.Parse()
	r := gin.Default()
	router.InitRouter(r)
	// database.InitDb()
	r.Run()
}
