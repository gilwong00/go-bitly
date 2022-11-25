package routes

import (
	handler "go-bitly/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

	api.GET("/g/:url", handler.Redirect)
	api.GET("/bitlies", handler.GetAllBitlys)
	api.GET("/bitlies/:id", handler.GetBitlyById)
	api.POST("/bitlies", handler.CreateBitly)
	api.PUT("/bitlies/:id", handler.UpdateBitly)
	api.DELETE("/bitlies/:id", handler.DeleteBitly)
}
