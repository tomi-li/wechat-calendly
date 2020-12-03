package server

import (
	"calendly/controllers"
	"calendly/middlewares"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("user")
		{
			user := new(controllers.UserController)
			userGroup.GET(":id", user.Retrieve)
			userGroup.GET("", user.RetrieveAll)
		}

		taskGroup := v1.Group("task")
		taskGroup.Use(middlewares.AuthMiddleware())
		{
			task := new(controllers.UserController)
			taskGroup.GET(":id", task.Retrieve)
		}
	}
	return router
}
