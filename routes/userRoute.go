package routes

import (
	"goblogart/controllers"
	"goblogart/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupUser(r *gin.RouterGroup) {
	r.POST("/register", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middlewares.RequireAuth, controllers.Validate)
	r.GET("/", middlewares.RequireAuth, controllers.GetUsers)
	r.GET("/with-posts", middlewares.RequireAuth, controllers.GetUsersWithPosts)
	r.GET("/logout", controllers.Logout)
}
