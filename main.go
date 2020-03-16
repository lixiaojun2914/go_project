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
	r.Run("127.0.0.1:8080")
}
