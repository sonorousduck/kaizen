package users

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup) {
	rg.POST("/users/logout", LogoutUserController())
	rg.GET("/users/me", GetCurrentUserController(userService))
}
