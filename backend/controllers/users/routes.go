package users

import (
	"backend/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, userService *services.UserService) {
	rg.POST("/users/logout", LogoutUserController())
	rg.GET("/users/me", GetCurrentUserController(userService))
}
