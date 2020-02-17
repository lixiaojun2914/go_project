package router

import (
	"github.com/gin-gonic/gin"
	"test/apps/test"
)

func InitRouter(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/get", test.Get)
		v1.POST("/post", test.Post)
	}
	//v2 := r.Group("/v2"){
	//	v1.GET("/get", testGet)
	//	v1.POST("/post", testPost)
	//}
}
