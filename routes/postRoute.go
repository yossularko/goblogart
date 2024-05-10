package routes

import (
	"goblogart/controllers"
	"goblogart/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupPost(r *gin.RouterGroup) {
	r.POST("/", middlewares.RequireAuth, controllers.CreatePost)
	r.GET("/", middlewares.WithParamsTest("administrator"), controllers.GetPosts)
	r.GET("/:id", controllers.GetPost)
	r.GET("/my-post", middlewares.RequireAuth, controllers.GetMyPosts)
	r.PUT("/:id", middlewares.RequireAuth, controllers.UpdatePost)
	r.DELETE("/:id", middlewares.RequireAuth, controllers.DeletePost)
}
