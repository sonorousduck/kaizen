package goals

import (
	"backend/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, goalService *services.GoalService) {
	// Get goal by id
	rg.GET("/goals/:goal_id", GetGoalByIdController(goalService))

	// Get goal for user
	rg.GET("/goals/me", GetGoalsForUserController(goalService))

	// Create goal
	rg.POST("/goals", CreateGoalController(goalService))

	// Update goal by id
	rg.PUT("/goals/:goal_id", UpdateGoalController(goalService))

	// Delete goal by id
	rg.DELETE("/goals/:goal_id", DeleteGoalIdController(goalService))

	// Delete goal by parent id
	rg.DELETE("/goals/parentGoals/:parent_goal_id", DeleteGoalsByParentIdController(goalService))

}
