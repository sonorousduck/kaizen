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

var ErrGoalNotFound = errors.New("goal not found")

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
		WHERE id = $1 AND deleted_at IS NULL`,
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

func (service *GoalService) GetGoalsForUser(ctx context.Context, userId uuid.UUID, filter models.PaginationFilter) ([]*models.Goal, error) {
	var goals []*models.Goal

	rows, err := service.db.Query(ctx,
		`SELECT id, user_id, parent_goal_id, title, description, starting_value, target_value, unit, frequency_interval, frequency, goal_type, due_date, deleted_at, created_at, updated_at
		FROM goals
		WHERE user_id = $1 AND deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT NULLIF($2, 0) OFFSET $3
		`,
		userId, filter.Limit, filter.Offset,
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

// Update endpoints

func (service *GoalService) UpdateGoal(ctx context.Context, userId uuid.UUID, goalId uuid.UUID, updateGoal models.UpdateGoal) error {
	commandTag, err := service.db.Exec(ctx,
		`UPDATE goals SET parent_goal_id = $1, title = $2, description = $3, starting_value = $4, target_value = $5, unit = $6, frequency_interval = $7, frequency = $8, goal_type = $9, due_date = $10, updated_at = NOW()
	WHERE id = $11 AND user_id = $12`,
		updateGoal.ParentGoalId,
		updateGoal.Title,
		updateGoal.Description,
		updateGoal.StartingValue,
		updateGoal.TargetValue,
		updateGoal.Unit,
		updateGoal.FrequencyInterval,
		updateGoal.Frequency,
		updateGoal.GoalType,
		updateGoal.DueDate,
		goalId,
		userId,
	)

	if err != nil {
		return fmt.Errorf("Failed to update goal: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return ErrGoalNotFound
	}

	return nil
}

// Delete endpoints

// Delete by id
func (service *GoalService) DeleteGoalById(ctx context.Context, goalId uuid.UUID) error {
	var deletedId uuid.UUID

	err := service.db.QueryRow(ctx,
		`UPDATE goals
		SET deleted_at = NOW()
		WHERE id = $1
		RETURNING id`,
		goalId,
	).Scan(&deletedId)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return fmt.Errorf("Goal not found: %s", goalId.String())
		}

		return fmt.Errorf("failed to delete goal: %w", err)
	}

	return nil
}

// Delete by parent
func (service *GoalService) DeleteGoalsByParentId(ctx context.Context, userId uuid.UUID, parentGoalId uuid.UUID) error {

	updatedGoals, err := service.db.Exec(ctx,
		`UPDATE goals
	SET deleted_at = NOW()
	WHERE user_id = $1 AND parent_goal_id = $2 AND deleted_at IS NULL
	`, userId, parentGoalId,
	)

	if err != nil {
		return fmt.Errorf("Error deleting goals by parent id: %w", err)
	}

	if updatedGoals.RowsAffected() == 0 {
		return fmt.Errorf("No goals found for parent id: %s", parentGoalId.String())
	}

	return nil
}

// Delete by category
