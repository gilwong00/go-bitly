package routes

import (
	handler "go-bitly/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

	api.GET("/", handler.GetAllBitlys)
	api.GET("/g/:url", handler.Redirect)
	api.GET("/:id", handler.GetBitlyById)
	api.POST("/", handler.CreateBitly)
	api.PUT("/:id", handler.UpdateBitly)
	api.DELETE("/:id", handler.DeleteBitly)
}
