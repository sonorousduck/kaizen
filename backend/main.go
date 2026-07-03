// @title Kaizen API
// @version 1.0
// @description Goal planning + execution
// @host localhost:8080
// @BasePath /
// @schemes http https
package main

import (
	"backend/controllers/goalcategories"
	"backend/controllers/goals"
	"backend/controllers/users"
	"backend/docs"
	"backend/initializers"
	"backend/internal"
	"backend/middleware"
	"backend/services"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	app                 *internal.App
	userService         *services.UserService
	goalService         *services.GoalService
	goalCategoryService *services.GoalCategoryService
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

	userService = services.NewUserService(app.DB)
	goalService = services.NewGoalService(app.DB)
	goalCategoryService = services.NewGoalCategoryService(app.DB)
}

func main() {
	defer app.DB.Close()
	port := os.Getenv("PORT")
	secret := []byte(os.Getenv("JWT_SECRET"))

	docs.SwaggerInfo.BasePath = "/"
	app.Router.StaticFile("/docs/swagger", "./docs/swagger.json")

	app.Router.SetTrustedProxies([]string{"localhost", "127.0.0.1"})
	authenticatedApi := app.Router.Group("/")
	authenticatedApi.Use(middleware.RequireAuth(userService, secret))
	{
		users.RegisterRoutes(authenticatedApi, userService)
		goals.RegisterRoutes(authenticatedApi, goalService)
		goalcategories.RegisterRoutes(authenticatedApi, goalCategoryService)
	}

	unauthenticatedApi := app.Router.Group("/public")
	unauthenticatedApi.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})
	users.RegisterUnauthenticatedRoutes(unauthenticatedApi, userService, secret)

	if err := app.Router.Run(":" + port); err != nil {
		log.Fatal(err)
	}

}
