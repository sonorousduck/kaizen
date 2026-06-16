package main

import (
	"backend/initializers"
	"backend/internal"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	app *internal.App
)

func init() {
	initializers.LoadEnvVariables()

	var err error
	pool, err := initializers.ConnectToDatabase()
	if err != nil {
		log.Fatalf("Failed to connnect to database: %v", err)
	}

	app = &internal.App{DB: pool, Router: gin.Default()}

	config := cors.DefaultConfig()
	config.AllowAllOrigins = false
	config.AllowOrigins = []string{"http://localhost:8080", "http://localhost:5173"}
	config.AllowCredentials = true
	config.MaxAge = 72 * time.Hour

	app.Router.Use(cors.New(config))

}

func main() {
	port := os.Getenv("PORT")

	app.Router.SetTrustedProxies([]string{"localhost", "127.0.0.1"})
	api := app.Router.Group("/")
	{
		api.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"status": "healthy"})
		})
	}

	if err := app.Router.Run(":" + port); err != nil {
		log.Fatal(err)
	}

}
