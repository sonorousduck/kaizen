package goalentries

import (
	"backend/controllers"
	"backend/middleware"
	"backend/models"
	"backend/services"
	"errors"
	"fmt"
	"net/http"
	"time"

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

// @Summary Get goal entries
// @Description Gets goal entries based on either its id, start/end date, etc.
// @ID get-goal-entries
// @Produce json
// @Param goal_id path string true "Goal ID"
// @Param limit query integer false "Limit number of results"
// @Param offset query integer false "Offset for pagination"
// @Param startDate query date false "Optional filter for start date"
// @Param endDate query date false "Optional filter for end date"
// @Success 200 {object} []models.GoalEntry
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /goals/{goal_id}/goalEntries
func GetGoalEntriesController(goalEntryService *services.GoalEntryService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := middleware.UserFromContext(ctx).ID

		goalId, err := uuid.Parse(ctx.Param("goal_id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid goal_id"})
			return
		}

		paginationFilter, err := controllers.GetPaginationFilterFromContext(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		startDate, err := parseOptionalDateQuery(ctx, "startDate")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid start date"})
			return
		}

		endDate, err := parseOptionalDateQuery(ctx, "endDate")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid end date"})
			return
		}

		goalEntryFilter := models.GoalEntryFilter{
			UserID:           userId,
			GoalID:           &goalId,
			StartDate:        startDate,
			EndDate:          endDate,
			PaginationFilter: *paginationFilter,
		}

		goalEntries, err := goalEntryService.GetGoalEntries(ctx, goalEntryFilter)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, goalEntries)
	}
}

func parseOptionalDateQuery(ctx *gin.Context, key string) (*time.Time, error) {
	dateQueryParameter := ctx.Query(key)
	if dateQueryParameter == "" {
		return nil, nil
	}

	parsed, err := time.Parse("2006-01-02", dateQueryParameter)
	if err != nil {
		return nil, err
	}

	return &parsed, nil
}

// @Summary Update goal entry
// @Description Updates a goal entry. Send all fields each time
// @ID update-goal-entry-by-id
// @Accept json
// @Produce json
// @Param goal_id path string true "Goal ID"
// @Param goal_entry_id path string true "Goal entry ID"
// @Param body body models.UpdateGoalEntry true "Updated goal entry data"
// @Success 202
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failture 500 {string} string
// @Router /goals/{goal_id}/goalEntries/{goal_entry_id} [put]
// @Security BearerAuth
func UpdateGoalEntryController(goalEntryService *services.GoalEntryService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var updateGoalEntry models.UpdateGoalEntry

		goalEntryId, err := uuid.Parse(ctx.Param("goal_entry_id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid goal entry id"})
			return
		}

		if err := ctx.ShouldBindJSON(&updateGoalEntry); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = goalEntryService.UpdateGoalEntry(ctx, goalEntryId, updateGoalEntry)

		if err != nil {
			if errors.Is(err, services.ErrGoalEntryNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "goal entry not found"})
				return
			}

			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.Status(http.StatusAccepted)
	}
}

// @Summary Deletes a goal entry
// @Description Deletes a goal entry by id
// @ID delete-goal-entry-by-id
// @Param goal_id path string true "Goal ID"
// @Param goal_entry_id path string true "Goal entry ID"
// @Success 204
// @Failure 400 {string} string
// @Failture 404 {string} string
// @Failure 500 {string} string
// @Router /goals/{goal_id}/goalEntries/{goal_entry_id} [delete]
// @Security BearerAuth
func DeleteGoalEntryController(goalEntryService *services.GoalEntryService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		goalEntryId, err := uuid.Parse(ctx.Param("goal_entry_id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid goal entry id"})
			return
		}

		err = goalEntryService.DeleteGoalEntry(ctx, goalEntryId)

		if err != nil {
			if errors.Is(err, services.ErrGoalEntryNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Goal entry id not found"})
				return
			}

			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.Status(http.StatusNoContent)
	}
}
