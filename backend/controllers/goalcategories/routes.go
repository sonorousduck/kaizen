package goalcategories

import (
	"backend/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, goalCategoryService *services.GoalCategoryService) {
	rg.POST("/goalCategories", CreateGoalCategoryController(goalCategoryService))
	rg.GET("/goalCategories/:goalCategoryId")
	rg.GET("/goalCategories/me", GetGoalCategoriesForUser(goalCategoryService))
	rg.PUT("/goalCategories/:goalCategoryId")
	rg.DELETE("/goalCategories/:goalCategoryId")

}
