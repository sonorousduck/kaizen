package users

import (
	"backend/middleware"
	"backend/models"
	"backend/services"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// @Summary Create a new user
// @Description Register a new user with email and password
// @ID create-user
// @Accept json
// @Produce json
// @Param body body models.CreateUser true "User registration data"
// @Success 201 {object} models.UserResponse
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /users [post]
func CreateUserController(userService *services.UserService, secret []byte) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input models.CreateUser

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Hash the password
		hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to hash password",
			})

			return
		}

		input.Password = string(hash)
		user, err := userService.CreateUser(ctx, input)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
			return
		}

		generateJwtAndSetCookies(user, ctx, secret)
		ctx.JSON(http.StatusCreated, user.ToResponse())
	}
}

func generateJwtAndSetCookies(user *models.User, ctx *gin.Context, secret []byte) {

	// Generate JWT token with 30 day expiration
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID.String(),
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString(secret)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create jwt token"})
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	secure := os.Getenv("ENV") == "production"
	ctx.SetCookie("token", tokenString, 3600*24*30, "/", "", secure, true)
}

// @Summary Login user
// @Description Login with email and password, returns JWT token in cookie
// @ID login-user
// @Accept json
// @Produce json
// @Param body body models.LoginUser true "Login credentials"
// @Success 200 {object} models.UserResponse
// @Failure 400 {string} string
// @Failure 403 {string} string
// @Failure 500 {string} string
// @Router /users/login [post]
func LoginUserController(login *services.UserService, secret []byte) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get the email and password off the request body
		var loginUser models.LoginUser

		if err := ctx.ShouldBindJSON(&loginUser); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Look up request user
		user, err := login.GetUserByEmail(ctx.Request.Context(), loginUser)

		if err != nil {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "incorrect email or password"})
			return
		}

		// Compare sent in password with saved user password hash
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password))

		if err != nil {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "incorrect email or password"})
			return
		}

		generateJwtAndSetCookies(user, ctx, secret)
		ctx.JSON(http.StatusOK, user.ToResponse())
	}
}

// @Summary Get current user
// @Description Get the current authenticated user's information
// @ID get-current-user
// @Produce json
// @Success 200 {object} models.UserResponse
// @Failure 401 {string} string
// @Router /users/me [get]
// @Security BearerAuth
func GetCurrentUserController(userService *services.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := middleware.UserFromContext(ctx)

		if user == nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		ctx.JSON(http.StatusOK, user.ToResponse())
	}
}

// @Summary Logout user
// @Description Clear the authentication cookie to logout
// @ID logout-user
// @Produce json
// @Success 200 {string} string
// @Router /users/logout [post]
func LogoutUserController() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Clear the authorization cookie by setting maxAge to 0
		c.SetCookie("token", "", 0, "/", "", false, true)
		c.JSON(http.StatusOK, gin.H{"message": "logged out"})
	}
}
