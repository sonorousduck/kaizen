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

type GoalCategoryService struct {
	db *pgxpool.Pool
}

var ErrGoalCategoryNotFound = errors.New("goal category not found")

func NewGoalCategoryService(db *pgxpool.Pool) *GoalCategoryService {
	return &GoalCategoryService{db: db}
}

func scanGoalCategory(row pgx.Row, goalCategory *models.GoalCategory) error {
	return row.Scan(
		&goalCategory.ID,
		&goalCategory.UserId,
		&goalCategory.Title,
		&goalCategory.MaxGoals,
		&goalCategory.Color,
		&goalCategory.Icon,
	)
}

// CREATE

func (service *GoalCategoryService) CreateGoalCategory(ctx context.Context, createGoalCategory models.CreateGoalCategory) (*models.GoalCategory, error) {
	goalCategory := &models.GoalCategory{}

	err := scanGoalCategory(service.db.QueryRow(ctx,
		`INSERT INTO goal_categories (user_id, title, max_goals, icon, color)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, user_id, title, max_goals, icon, color`,
		createGoalCategory.UserId,
		createGoalCategory.Title,
		createGoalCategory.MaxGoals,
		createGoalCategory.Color,
		createGoalCategory.Icon,
	), goalCategory)

	if err != nil {
		return nil, fmt.Errorf("failed to create goal category: %w", err)
	}

	return goalCategory, nil
}

// GET
func (service *GoalCategoryService) GetGoalCategoriesByUserId(ctx context.Context, userId uuid.UUID) ([]*models.GoalCategory, error) {
	var goalCategories []*models.GoalCategory

	rows, err := service.db.Query(ctx,
		`SELECT id, user_id, title, max_goals, icon, color 
		FROM goal_categories
		WHERE user_id = $1`,
		userId,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get goal categories for user: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		goalCategory := &models.GoalCategory{}

		if err := scanGoalCategory(rows, goalCategory); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}

		goalCategories = append(goalCategories, goalCategory)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over all the goal categories: %w", err)
	}

	return goalCategories, nil
}

func (service *GoalCategoryService) GetGoalCategoriesById(ctx context.Context, id uuid.UUID) (*models.GoalCategory, error) {
	goalCategory := &models.GoalCategory{}

	err := scanGoalCategory(service.db.QueryRow(ctx,
		`SELECT id, user_id, title, max_goals, color, icon 
		FROM goal_categories
		WHERE id = $1`,
		id,
	), goalCategory)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrGoalCategoryNotFound
	}

	if err != nil {
		return nil, fmt.Errorf("Failed to get goal category: %w", err)
	}

	return goalCategory, nil
}

// UPDATE

func (service *GoalCategoryService) UpdateGoalCategory(ctx context.Context, userId uuid.UUID, goalCategoryId uuid.UUID, updateGoalCategory models.UpdateGoalCategory) error {
	commandTag, err := service.db.Exec(ctx,
		`UPDATE goal_categories SET title = $1, max_goals = $2, icon = $3, color = $4 WHERE id = $5 AND user_id = $6`,
		updateGoalCategory.Title,
		updateGoalCategory.MaxGoals,
		updateGoalCategory.Icon,
		updateGoalCategory.Color,
		goalCategoryId,
		userId,
	)

	if err != nil {
		return fmt.Errorf("Failed to update goal category: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return ErrGoalCategoryNotFound
	}

	return nil
}

// DELETE
func (service *GoalCategoryService) DeleteGoalCategoryById(ctx context.Context, goalCategoryId uuid.UUID, userId uuid.UUID) error {
	var deletedId uuid.UUID

	err := service.db.QueryRow(ctx,
		`DELETE FROM goal_categories
	WHERE id = $1 AND user_id = $2
	RETURNING id
	`, goalCategoryId, userId,
	).Scan(&deletedId)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrGoalCategoryNotFound
		}

		return fmt.Errorf("failed to delete goal category: %w", err)
	}

	return nil
}
