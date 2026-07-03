package models

import "github.com/google/uuid"

type GoalCategory struct {
	ID       uuid.UUID `json:"id" binding:"required"`
	UserId   uuid.UUID `json:"userId" binding:"required"`
	Title    string    `json:"title" binding:"required"`
	MaxGoals int32     `json:"maxGoals"`
	Icon     string    `json:"icon" binding:"required"`
	Color    string    `json:"color" binding:"required"`
}

type CreateGoalCategory struct {
	UserId   uuid.UUID `json:"-"`
	Title    string    `json:"title" binding:"required"`
	MaxGoals int32     `json:"maxGoals"`
	Icon     string    `json:"icon" binding:"required"`
	Color    string    `json:"color" binding:"required"`
}

type UpdateGoalCategory struct {
	Title    string `json:"title" binding:"required"`
	MaxGoals int32  `json:"maxGoals"`
	Icon     string `json:"icon" binding:"required"`
	Color    string `json:"color" binding:"required"`
}
