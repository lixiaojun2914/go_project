package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	name := c.Query("name")
	action := c.Query("action")
	c.JSON(200, gin.H{
		"name":   name,
		"action": action,
	})
}

type PostStruct struct {
	Name   string `json:"name"`
	Action string `json:"action"`
}

func Post(c *gin.Context) {
	var data PostStruct
	err := c.BindJSON(&data)
	if err != nil {
		fmt.Println("error")
	}
	c.JSON(200, gin.H{
		"name":   data.Name,
		"action": data.Action,
	})
}
