package services

import (
	"backend/models"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
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

func (service *GoalService) GetGoalById(ctx context.Context, goalId uuid.UUID) (*models.Goal, error) {
	goal := &models.Goal{}

	err := scanGoal(service.db.QueryRow(ctx,
		`SELECT id, user_id, parent_goal_id, title, description, starting_value, target_value, unit, frequency_interval, frequency, goal_type, due_date, deleted_at, created_at, updated_at
		FROM goals
		WHERE id = $1`,
		goalId,
	), goal)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("No goal found with this id")
	}

	if err != nil {
		return nil, fmt.Errorf("Failed to get goa: %w", err)
	}

	return goal, nil
}

func (service *GoalService) GetGoalsForUser(ctx context.Context, userId uuid.UUID) ([]*models.Goal, error) {
	var goals []*models.Goal

	rows, err := service.db.Query(ctx,
		`SELECT id, user_id, parent_goal_id, title, description, starting_value, target_value, unit, frequency_interval, frequency, goal_type, due_date, deleted_at, created_at, updated_at
		FROM goals
		WHERE user_id = $1`,
		userId,
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to get goals for user: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		goal := &models.Goal{}

		if err := scanGoal(rows, goal); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}

		goals = append(goals, goal)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate all the goals: %w", err)
	}

	return goals, nil
}

// func (service *GoalService) GetGoalsByCategory

// Delete endpoints

// Delete by id
// Delete by parent
// Delete by category
