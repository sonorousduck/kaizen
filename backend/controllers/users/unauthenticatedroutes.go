package users

import (
	"backend/services"

	"github.com/gin-gonic/gin"
)

func RegisterUnauthenticatedRoutes(rg *gin.RouterGroup, userService *services.UserService, secret []byte) {
	rg.POST("/users", CreateUserController(userService, secret))
	rg.POST("/users/login", LoginUserController(userService, secret))
}
