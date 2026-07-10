package controllers

import (
	"backend/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPaginationFilterFromContext(ctx *gin.Context) (*models.PaginationFilter, error) {
	filter := &models.PaginationFilter{}

	if limitRaw := ctx.Query("limit"); limitRaw != "" {
		limit, err := strconv.Atoi(limitRaw)
		if err != nil || limit < 0 {
			return nil, fmt.Errorf("invalid limit query")
		}
		filter.Limit = limit
	}

	if offsetRaw := ctx.Query("offset"); offsetRaw != "" {
		offset, err := strconv.Atoi(offsetRaw)
		if err != nil || offset < 0 {
			return nil, fmt.Errorf("invalid offset query")
		}

		filter.Offset = offset
	}

	return filter, nil
}
