package goalentries

import (
	"backend/middleware"
	"backend/permissions"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, goalEntriesService *services.GoalEntryService, goalPermissionChecker permissions.GoalPermissionChecker) {
	// Create goal
	rg.POST("/goals/:goal_id/goalEntries", middleware.RequireGoalPermission(&goalPermissionChecker), CreateGoalEntryController(goalEntriesService))

	// Get goals by goal id
	rg.GET("/goals/:goal_id/goalEntries", middleware.RequireGoalPermission(&goalPermissionChecker), GetGoalEntriesController(goalEntriesService))

	// Update goal entry
	rg.PUT("/goals/:goal_id/goalEntries/:goal_entry_id", middleware.RequireGoalPermission(&goalPermissionChecker), UpdateGoalEntryController(goalEntriesService))

	// Delete goal entry
	rg.DELETE("/goals/:goal_id/goalEntries/:goal_entry_id", middleware.RequireGoalPermission(&goalPermissionChecker))
}
