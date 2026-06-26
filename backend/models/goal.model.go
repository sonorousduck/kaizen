package models

import (
	"time"

	"github.com/google/uuid"
)

type Frequency string

const (
	Once    Frequency = "once"
	Daily   Frequency = "daily"
	Weekly  Frequency = "weekly"
	Monthly Frequency = "monthly"
	Yearly  Frequency = "yearly"
)

type GoalType string

const (
	Habit   GoalType = "habit"
	Numeric GoalType = "numeric"
)

type Goal struct {
	ID                uuid.UUID  `json:"id" binding:"required"`
	UserId            uuid.UUID  `json:"userId" binding:"required"`
	ParentGoalId      *uuid.UUID `json:"parentGoalId"`
	Title             string     `json:"title" binding:"required"`
	Description       string     `json:"description"`
	StartingValue     *float32   `json:"startingValue"`
	TargetValue       *float32   `json:"targetValue"`
	Unit              string     `json:"unit"`
	FrequencyInterval int32      `json:"frequencyInterval"`
	Frequency         *Frequency `json:"frequency" binding:"required"`
	GoalType          GoalType   `json:"goalType" binding:"required"`
	DueDate           *time.Time `json:"dueDate"`
	DeletedAt         *time.Time `json:"deletedAt,omitempty"`
	CreatedAt         time.Time  `json:"createdAt"`
	UpdatedAt         time.Time  `json:"updatedAt"`
	// TODO: Create category in the db table
	CategoryId *uuid.UUID `json:"categoryId"`
}
