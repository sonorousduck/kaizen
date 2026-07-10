package middleware

import (
	"backend/permissions"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequireGoalPermission(goalPermissionChecker *permissions.GoalPermissionChecker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		goalId, err := uuid.Parse(ctx.Param("goal_id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid goal id"})
			ctx.Abort()
			return
		}

		authUser := UserFromContext(ctx)
		hasPermission, err := goalPermissionChecker.AuthUserHasPermissionOnGoal(ctx, authUser, goalId)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			ctx.Abort()
			return
		}

		if !hasPermission {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
