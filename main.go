package main

import (
	"goblogart/inits"
	"goblogart/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	inits.LoadEnv()
	inits.DBInit()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	postGroup := r.Group("/posts")
	routes.SetupPost(postGroup)

	userGroup := r.Group("/users")
	routes.SetupUser(userGroup)

	r.Run()
}
