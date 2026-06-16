package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	DB     *pgxpool.Pool
	Router *gin.Engine
}
