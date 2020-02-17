package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"test/database"
	"test/router"
)

func main() {
	flag.Parse()
	r := gin.Default()
	router.InitRouter(r)
	database.Init()
	//r.Run("127.0.0.1:8080")
}
