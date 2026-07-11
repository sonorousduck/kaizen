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

type GoalEntryService struct {
	db *pgxpool.Pool
}

func NewGoalEntryService(db *pgxpool.Pool) *GoalEntryService {
	return &GoalEntryService{db: db}
}

var ErrGoalEntryNotFound = errors.New("goal entry not found")

func scanGoalEntry(row pgx.Row, goalEntry *models.GoalEntry) error {
	return row.Scan(
		&goalEntry.ID,
		&goalEntry.GoalId,
		&goalEntry.Value,
		&goalEntry.Note,
		&goalEntry.Date,
		&goalEntry.CreatedAt,
	)
}

func (service *GoalEntryService) CreateGoalEntry(ctx context.Context, createGoalEntry models.CreateGoalEntry) (*models.GoalEntry, error) {
	goalEntry := &models.GoalEntry{}

	err := scanGoalEntry(service.db.QueryRow(ctx,
		`INSERT INTO goal_entries (goal_id, value, note, date)
		VALUES ($1, $2, $3, $4)
		RETURNING id, goal_id, value, note, date, created_at
		`,
		createGoalEntry.GoalId,
		createGoalEntry.Value,
		createGoalEntry.Note,
		createGoalEntry.Date,
	), goalEntry)

	if err != nil {
		return nil, fmt.Errorf("failed to create goal entry: %w", err)
	}

	return goalEntry, nil
}

func (service *GoalEntryService) GetGoalEntries(ctx context.Context, filter models.GoalEntryFilter) ([]*models.GoalEntry, error) {
	var goalEntries []*models.GoalEntry

	rows, err := service.db.Query(ctx,
		`SELECT goalEntry.id, goalEntry.goal_id, goalEntry.value, goalEntry.note, goalEntry.date, goalEntry.created_at
		FROM goal_entries goalEntry
		JOIN goals ON goals.id = goalEntry.goal_id
		WHERE goals.user_id = $1
			AND ($2::uuid IS NULL OR goalEntry.goal_id = $2)
			AND ($3::timestamptz IS NULL OR goalEntry.date >= $3)
			AND ($4::timestamptz IS NULL OR goalEntry.date <= $4)
		ORDER BY goalEntry.date DESC
		LIMIT NULLIF($5, 0) OFFSET $6`,
		filter.UserID,
		filter.GoalID,
		filter.StartDate,
		filter.EndDate,
		filter.PaginationFilter.Limit,
		filter.PaginationFilter.Offset,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get goal entries for goal by date: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		goalEntry := &models.GoalEntry{}

		if err := scanGoalEntry(rows, goalEntry); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}

		goalEntries = append(goalEntries, goalEntry)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate all the goal entries: %w", err)
	}

	return goalEntries, nil
}

func (service *GoalEntryService) UpdateGoalEntry(ctx context.Context, goalEntryId uuid.UUID, updateGoalEntry models.UpdateGoalEntry) error {
	commandTag, err := service.db.Exec(ctx,
		`UPDATE goal_entries SET value = $1, note = $2, date = $3 WHERE id = $4`,
		updateGoalEntry.Value,
		updateGoalEntry.Note,
		updateGoalEntry.Date,
		goalEntryId,
	)

	if err != nil {
		return fmt.Errorf("failed to update goal entry: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return ErrGoalEntryNotFound
	}

	return nil
}

func (service *GoalEntryService) DeleteGoalEntry(ctx context.Context, goalEntryId uuid.UUID) error {
	var deletedId uuid.UUID

	err := service.db.QueryRow(ctx,
		`DELETE FROM goal_entries
	WHERE id = $1
	RETURNING id`,
		goalEntryId,
	).Scan(&deletedId)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrGoalEntryNotFound
		}

		return fmt.Errorf("failed to delete goal entry: %w", err)
	}

	return nil
}
