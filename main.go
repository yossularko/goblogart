package main

import (
	"goblogart/controllers"
	"goblogart/inits"

	"github.com/gin-gonic/gin"
)

func init() {

	inits.LoadEnv()
	inits.DBInit()
}

func main() {
	r := gin.Default()

	r.POST("/", controllers.CreatePost)
	r.GET("/", controllers.GetPosts)
	r.GET("/:id", controllers.GetPost)
	r.PUT("/:id", controllers.UpdatePost)
	r.DELETE("/:id", controllers.DeletePost)

	r.Run()
}
