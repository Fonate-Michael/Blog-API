package routes

import (
	"app/controller"
	"app/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegistedRoute(r *gin.Engine) {
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.GET("/posts", controller.GetPosts)
	r.GET("/posts/:id/comment", controller.GetComment)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	protectedRoutes := r.Group("/")

	protectedRoutes.Use(middleware.AuthMiddleWare())

	{
		protectedRoutes.POST("/posts", controller.AddPost)

		protectedRoutes.POST("/posts/:id/comment", controller.AddComment)
	}
}
