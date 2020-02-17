package main

import (
	"github.com/gin-gonic/gin"
	"test/router"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	r.Run("127.0.0.1:8080")
}
