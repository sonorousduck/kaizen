package services

import (
	"backend/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type GoalService struct {
	db *pgxpool.Pool
}

func NewGoalService(db *pgxpool.Pool) *GoalService {
	return &GoalService{db: db}
}

func scanGoal(row pgx.Row, goal *models.Goal) error {
	return row.Scan(
		&goal.ID,
		&goal.UserId,
		&goal.ParentGoalId,
		&goal.Title,
		&goal.Description,
		&goal.StartingValue,
		&goal.TargetValue,
		&goal.Unit,
		&goal.FrequencyInterval,
		&goal.Frequency,
		&goal.GoalType,
		&goal.DueDate,
		&goal.DeletedAt,
		&goal.CreatedAt,
		&goal.UpdatedAt,
	)
}

func (service *GoalService) CreateGoal(ctx context.Context, createGoal models.CreateGoal) (*models.Goal, error) {
	goal := &models.Goal{}

	err := scanGoal(service.db.QueryRow(ctx,
		`INSERT INTO goals (user_id, parent_goal_id, title, description, starting_value, target_value, unit, frequency_interval, frequency, goal_type, due_date)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id, user_id, parent_goal_id, title, description, starting_value, target_value, unit, frequency_interval, frequency, goal_type, due_date, deleted_at, created_at, updated_at`,
		createGoal.UserId,
		createGoal.ParentGoalId,
		createGoal.Title,
		createGoal.Description,
		createGoal.StartingValue,
		createGoal.TargetValue,
		createGoal.Unit,
		createGoal.FrequencyInterval,
		createGoal.Frequency,
		createGoal.GoalType,
		createGoal.DueDate,
	), goal)

	if err != nil {
		return nil, fmt.Errorf("failed to create goal: %w", err)
	}

	return goal, nil
}
