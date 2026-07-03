package goalcategories

import (
	"backend/middleware"
	"backend/models"
	"backend/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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
