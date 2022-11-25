package main

import (
	"fmt"
	"go-bitly/pkg/db"
	"go-bitly/pkg/models"
	"go-bitly/pkg/routes"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const (
	port = ":5000"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Init()
	defer db.DB.AutoMigrate(&models.Bitly{})
}

func main() {
	gin.ForceConsoleColor()
	r := gin.Default()

	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}, // can update this to be more inclusive
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
		// AllowCredentials: true,
		// MaxAge:           12 * time.Hour,
	}))

	routes.SetupRoutes(r)

	err := r.Run(port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
		os.Exit(0)
	}
}
