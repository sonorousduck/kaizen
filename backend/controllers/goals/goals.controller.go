package goals

import (
	"backend/middleware"
	"backend/models"
	"backend/services"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary Create a new goal
// @Description Creates a new goal for the user
// @ID create-goal
// @Accept json
// @Produce json
// @Param body body models.CreateGoal true "Goal creation data"
// @Success 201 {object} models.Goal
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /goals [post]
// @Security BearerAuth
func CreateGoalController(goalService *services.GoalService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var createGoal models.CreateGoal

		if err := ctx.ShouldBindJSON(&createGoal); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		createGoal.UserId = middleware.UserFromContext(ctx).ID

		goal, err := goalService.CreateGoal(ctx, createGoal)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to create goal with error: %s", err.Error())})
			return
		}

		ctx.JSON(http.StatusCreated, goal)
	}
}

// @Summary Retrieve a goal by its id
// @Description Retrieves a goal by its id
// @ID get-goal-by-id
// @Produce json
// @Param goal_id path string true "Goal ID"
// @Success 200 {object} models.Goal
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /goals/{goal_id} [get]
// @Security BearerAuth
func GetGoalByIdController(goalService *services.GoalService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		goal_id, err := uuid.Parse(ctx.Param("goal_id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid goal id"})
			return
		}

		goal, err := goalService.GetGoalById(ctx, goal_id)

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "goal not found"})
			return
		}

		ctx.JSON(http.StatusOK, goal)
	}
}

// @Summary Retrieve goals by owner
// @Description Retrieves goals by the user
// @ID get-goals-by-user
// @Produce json
// @Param limit query integer false "Limit number of results"
// @Param offset query integer false "Offset for pagination"
// @Success 200 {object} []models.Goal
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /goals/me [get]
// @Security BearerAuth
func GetGoalsForUserController(goalService *services.GoalService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		filter := models.PaginationFilter{}

		if limitRaw := ctx.Query("limit"); limitRaw != "" {
			limit, err := strconv.Atoi(limitRaw)
			if err != nil || limit < 0 {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit query"})
				return
			}
			filter.Limit = limit
		}

		if offsetRaw := ctx.Query("offset"); offsetRaw != "" {
			offset, err := strconv.Atoi(offsetRaw)
			if err != nil || offset < 0 {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid offset query"})
				return
			}

			filter.Offset = offset
		}

		userId := middleware.UserFromContext(ctx).ID
		goals, err := goalService.GetGoalsForUser(ctx, userId, filter)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, goals)
	}
}

// @Summary Updates goal
// @Description Updates goals. Send ALL fields each time
// @ID update-goal-by-id
// @Accept json
// @Produce json
// @Param goal_id path string true "Goal ID"
// @Param body body models.UpdateGoal true "Updated goal data"
// @Success 202
// @Failure 400 {string} string
// @Failture 404 {string} string
// @Failure 500 {string} string
// @Router /goals/{goal_id} [put]
// @Security BearerAuth
func UpdateGoalController(goalService *services.GoalService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var updateGoal models.UpdateGoal

		goalId, err := uuid.Parse(ctx.Param("goal_id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid goal id"})
			return
		}

		if err := ctx.ShouldBindJSON(&updateGoal); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = goalService.UpdateGoal(ctx, middleware.UserFromContext(ctx).ID, goalId, updateGoal)

		if err != nil {
			if errors.Is(err, services.ErrGoalNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Goal not found"})
				return
			}

			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.Status(http.StatusAccepted)
	}
}

// @Summary Deletes a goal
// @Description Deletes a goal by id
// @ID delete-goal-by-id
// @Param goal_id path string true "Goal ID"
// @Success 204
// @Failure 400 {string} string
// @Failture 404 {string} string
// @Failure 500 {string} string
// @Router /goals/{goal_id} [delete]
// @Security BearerAuth
func DeleteGoalIdController(goalService *services.GoalService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		goalId, err := uuid.Parse(ctx.Param("goal_id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid goal id"})
			return
		}

		err = goalService.DeleteGoalById(ctx, goalId)

		if err != nil {
			if errors.Is(err, services.ErrGoalNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Goal id not found"})
				return
			}

			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.Status(http.StatusNoContent)
	}
}

// @Summary Deletes all goals by parent id
// @Description Deletes all goals with the specified parent id
// @ID delete-goals-by-parent-id
// @Param parent_goal_id path string true "Parent goal ID"
// @Success 204
// @Failure 400 {string} string
// @Failture 404 {string} string
// @Failure 500 {string} string
// @Router /goals/parentGoals/{parent_goal_id} [delete]
// @Security BearerAuth
func DeleteGoalsByParentIdController(goalService *services.GoalService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		parentGoalId, err := uuid.Parse(ctx.Param("parent_goal_id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid goal id"})
			return
		}

		err = goalService.DeleteGoalsByParentId(ctx, middleware.UserFromContext(ctx).ID, parentGoalId)

		if err != nil {
			if errors.Is(err, services.ErrGoalNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Parent goal id not found"})
				return
			}

			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.Status(http.StatusNoContent)
	}
}
