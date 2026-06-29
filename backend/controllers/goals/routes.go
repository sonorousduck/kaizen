package goals

import (
	"backend/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, goalService *services.GoalService) {
	// Get goal by id
	rg.GET("/goals/:goal_id")

	// Get goal for user
	rg.GET("/goals/me")

	// Create goal
	rg.POST("/goals")

	// Update goal by id
	rg.PUT("/goals/:goal_id")

	// Delete goal by id
	rg.PATCH("/goals/:goal_id/delete")

	// Delete goal by parent id
	rg.PATCH("/goals/parentGoals/:parent_goal_id/delete")

}
