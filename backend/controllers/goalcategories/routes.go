package goalcategories

import (
	"backend/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, goalCategoryService *services.GoalCategoryService) {
	rg.POST("/goalCategories", CreateGoalCategoryController(goalCategoryService))
	rg.GET("/goalCategories/:goal_category_id", GetGoalCategoriesById(goalCategoryService))
	rg.GET("/goalCategories/me", GetGoalCategoriesForUser(goalCategoryService))
	rg.PUT("/goalCategories/:goal_category_id")
	rg.DELETE("/goalCategories/:goal_category_id")

}
