package goalentries

import (
	"backend/models"
	"backend/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary Create goal entry
// @Description Creates a new goal entry for a goal
// @ID create-goal-entry
// @Accept json
// @Produce json
// @Param body body models.CreateGoalEntry true "Goal entry creation data"
// @Success 201 {object} models.GoalEntry
// @Faiure 400 {string} string
// @Failure 500 {string} string
// @Router /goals/{goal_id}/goalEntries [post]
// @Security BearerAuth
func CreateGoalEntryController(goalEntryService *services.GoalEntryService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var createGoalEntry models.CreateGoalEntry

		if err := ctx.ShouldBindJSON(&createGoalEntry); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		goalId, err := uuid.Parse(ctx.Param("goal_id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		createGoalEntry.GoalId = goalId
		goalEntry, err := goalEntryService.CreateGoalEntry(ctx, createGoalEntry)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to create goal entry with error: %s", err.Error())})
			return
		}

		ctx.JSON(http.StatusCreated, goalEntry)
	}
}

// Get goal entry

// Update goal entry

// Delete goal entry
