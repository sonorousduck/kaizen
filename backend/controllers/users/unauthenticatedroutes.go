package users

import (
	"github.com/gin-gonic/gin"
)

func RegisterUnauthenticatedRoutes(rg *gin.RouterGroup) {
	rg.POST("/users", CreateUserController(userService))
	rg.POST("/users/login", LoginUserController(userService))
}
