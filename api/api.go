package api

import (
	"github.com/gin-gonic/gin"
	"rowing-market-api/api/middlewares"
	"rowing-market-api/api/routes/health"
	"rowing-market-api/api/routes/post"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		//HEALTH
		api.GET("/health", health.CheckHealth)
		api.GET("/health/report", health.CheckHealthReport)

		api.POST("/create-post", middlewares.Language(), post.RegisterPost)
		//api.POST("/login", middlewares.Language(), user.Login)
		//api.POST("/register-user", middlewares.Language(), user.RegisterUser)
	}
}
