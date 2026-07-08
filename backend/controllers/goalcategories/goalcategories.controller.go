package goalcategories

import (
	"backend/middleware"
	"backend/models"
	"backend/services"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary Create a new goal category
// @Description creates a new goal category (like Physical) for the user
// @ID create-goal-category
// @Accept json
// @Produce json
// @Param body body models.CreateGoalCategory true "Goal category creation data"
// @Success 201 {object} models.GoalCategory
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /goalCategories [post]
// @Security BearerAuth
func CreateGoalCategoryController(goalCategoryService *services.GoalCategoryService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var createGoalCategory models.CreateGoalCategory

		if err := ctx.ShouldBindJSON(&createGoalCategory); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		createGoalCategory.UserId = middleware.UserFromContext(ctx).ID

		goalCategory, err := goalCategoryService.CreateGoalCategory(ctx, createGoalCategory)

		if err != nil {
			if err == services.ErrGoalCategoryNotFound {
				ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Goal category now found: %w", err)})
				return
			}

			ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to create goal with error: %s", err.Error())})
			return
		}

		ctx.JSON(http.StatusCreated, goalCategory)
	}
}

// @Summary Retrieve goals categories by owner
// @Description Retrieves goal categories by the authenticated user
// @ID get-goal-categories-by-user
// @Produce json
// @Success 200 {object} []models.GoalCategory
// @Failure 500 {string} string
// @Router /goalCategories/me [get]
// @Security BearerAuth
func GetGoalCategoriesForUser(goalCategoryService *services.GoalCategoryService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := middleware.UserFromContext(ctx).ID

		goalCategories, err := goalCategoryService.GetGoalCategoriesByUserId(ctx, userId)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		ctx.JSON(http.StatusOK, goalCategories)
	}
}

// @Summary Retrieve goals categories by owner
// @Description Retrieves goal categories by the authenticated user
// @ID get-goal-categories-by-user
// @Param goal_category_id path string true "Goal category id"
// @Produce json
// @Success 200 {object} models.GoalCategory
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /goalCategories/{goal_category_id} [get]
// @Security BearerAuth
func GetGoalCategoriesById(goalCategoryService *services.GoalCategoryService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		goalCategoryId, err := uuid.Parse(ctx.Param("goal_category_id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		goalCategory, err := goalCategoryService.GetGoalCategoriesById(ctx, goalCategoryId)

		if err != nil {
			if err == services.ErrGoalCategoryNotFound {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "goal category not found"})
				return
			}

			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, goalCategory)
	}
}

// @Summary Updates goal category
// @Description Updates goal category. Send ALL fields each time
// @ID update-goal-category-by-id
// @Accept json
// @Produce json
// @Param goal_category_id path string true "Goal ID"
// @Param body body models.UpdateGoalCategory true "Updated goal category data"
// @Success 204
// @Failure 400 {string} string
// @Failture 404 {string} string
// @Failure 500 {string} string
// @Router /goalCategories/{goal_category_id} [put]
// @Security BearerAuth
func UpdateGoalCategoryById(goalCategoryService *services.GoalCategoryService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		goalCategoryId, err := uuid.Parse(ctx.Param("goal_category_id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userId := middleware.UserFromContext(ctx).ID

		var updateGoalCategory models.UpdateGoalCategory

		if err := ctx.ShouldBindJSON(&updateGoalCategory); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = goalCategoryService.UpdateGoalCategory(ctx, userId, goalCategoryId, updateGoalCategory)

		if err != nil {
			if errors.Is(err, services.ErrGoalNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Goal category not found"})
				return
			}

			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.Status(http.StatusNoContent)
	}
}

// @Summary Deletes goal category
// @Description Deletes goal category.
// @ID delete-goal-category-by-id
// @Produce json
// @Param goal_category_id path string true "Goal ID"
// @Success 204
// @Failure 400 {string} string
// @Failture 404 {string} string
// @Failure 500 {string} string
// @Router /goalCategories/{goal_category_id} [delete]
// @Security BearerAuth
func DeleteGoalCategoryById(goalCategoryService *services.GoalCategoryService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		goalCategoryId, err := uuid.Parse(ctx.Param("goal_category_id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userId := middleware.UserFromContext(ctx).ID
		err = goalCategoryService.DeleteGoalCategoryById(ctx, goalCategoryId, userId)

		if err != nil {
			if errors.Is(err, services.ErrGoalNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Goal category not found"})
				return
			}

			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.Status(http.StatusNoContent)
	}
}
