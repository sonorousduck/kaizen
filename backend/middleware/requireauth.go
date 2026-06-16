package middleware

import (
	"backend/models"
	"backend/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(userService *services.UserService, secret []byte) gin.HandlerFunc {
	if len(secret) == 0 {
		panic("JWT secret must not be empty (RequireAuth)")
	}

	if userService == nil {
		panic("User service must not be nil (RequireAuth)")
	}

	keyFunction := func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	}

	return func(ctx *gin.Context) {
		tokenString, err := ctx.Cookie("token")
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(
			tokenString,
			keyFunction,
			jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
			jwt.WithExpirationRequired(),
		)

		if err != nil || !token.Valid {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		userId, ok := claims["user_id"].(string)
		if !ok || userId == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		user, err := userService.GetUserByID(ctx.Request.Context(), userId)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("user", user)
		ctx.Next()
	}
}

func UserFromContext(ctx *gin.Context) *models.User {
	val, exists := ctx.Get("user")

	if !exists {
		return nil
	}

	user, _ := val.(*models.User)
	return user
}
