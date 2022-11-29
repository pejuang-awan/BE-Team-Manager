package router

import (
	"github.com/gin-gonic/gin"

	"github.com/pejuang-awan/BE-Team-Manager/controller/api"
	"github.com/pejuang-awan/BE-Team-Manager/controller/middleware"
)

func InitializeRouter() (router *gin.Engine) {
	router = gin.Default()

	apiRoute := router.Group("/api")
	apiRoute.Use(
		middleware.CorsMiddleware,
	)
	{
		apiRoute.GET("/participants/:id", api.Participants)
		apiRoute.GET("/tournaments/:id", api.Tournaments)
		apiRoute.POST("/join", api.Join)
		apiRoute.POST("/leave", api.Leave)
	}

	return router
}
