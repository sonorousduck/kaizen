package models

import (
	"time"

	"github.com/google/uuid"
)

type GoalEntry struct {
	ID        uuid.UUID `json:"id" binding:"required"`
	GoalId    uuid.UUID `json:"goalId" binding:"required"`
	Value     float32   `json:"value" binding:"required"`
	Note      *string   `json:"note"`
	Date      time.Time `json:"date" binding:"required"`
	CreatedAt time.Time `json:"createdAt" binding:"required"`
}

type CreateGoalEntry struct {
	GoalId uuid.UUID `json:"-"`
	Value  float32   `json:"value" binding:"required"`
	Note   *string   `json:"note"`
	Date   time.Time `json:"date" binding:"required"`
}

type UpdateGoalEntry struct {
	Value float32   `json:"value" binding:"required"`
	Note  *string   `json:"note"`
	Date  time.Time `json:"date" binding:"required"`
}
