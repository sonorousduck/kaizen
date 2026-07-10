package permissions

import (
	"backend/models"
	"backend/services"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GoalPermissionChecker struct {
	goalService *services.GoalService
}

func NewGoalPermissionChecker(goalService *services.GoalService) *GoalPermissionChecker {
	return &GoalPermissionChecker{goalService: goalService}
}

func (goalPermissionChecker *GoalPermissionChecker) AuthUserHasPermissionOnGoal(ctx *gin.Context, authUser *models.User, goalId uuid.UUID) (bool, error) {
	if authUser == nil {
		return false, fmt.Errorf("unauthorized")
	}

	return goalPermissionChecker.goalService.UserOwnsGoal(ctx, goalId, authUser.ID)

}
