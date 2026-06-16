package services

import (
	"context"
	"errors"
	"fmt"
	"kaizen/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserService struct {
	db *pgxpool.Pool
}

func NewUserService(db *pgxpool.Pool) *UserService {
	return &UserService{db: db}
}

func (s *UserService) CreateUser(ctx context.Context, input models.CreateUser) (*models.User, error) {
	user := &models.User{}

	err := s.db.QueryRow(ctx,
		`INSERT INTO users (first_name, last_name, email, password)
		VALUES ($1, $2, $3, $4)
		RETURNING id, first_name, last_name, email, deleted_at, created_at, updated_at`,
		input.FirstName,
		input.LastName,
		input.Email,
		input.Password,
	).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DeletedAt, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

func (s *UserService) CreateUserWithDefaultTeam(ctx context.Context, input models.CreateUser) (*models.User, error) {
	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	user := &models.User{}
	err = tx.QueryRow(ctx,
		`INSERT INTO users (first_name, last_name, email, password)
		VALUES ($1, $2, $3, $4)
		RETURNING id, first_name, last_name, email, deleted_at, created_at, updated_at`,
		input.FirstName,
		input.LastName,
		input.Email,
		input.Password,
	).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DeletedAt, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return user, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	user := &models.User{}

	err := s.db.QueryRow(ctx,
		`SELECT id, first_name, last_name, email, deleted_at, created_at, updated_at
		FROM users WHERE id = $1`,
		id,
	).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DeletedAt, &user.CreatedAt, &user.UpdatedAt)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("user not found")
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, loginUser models.LoginUser) (*models.User, error) {
	user := &models.User{}

	err := s.db.QueryRow(ctx,
		`SELECT id, first_name, last_name, password, email, deleted_at, created_at, updated_at
		FROM users WHERE email = $1`,
		loginUser.Email,
	).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Password, &user.Email, &user.DeletedAt, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to login user: %w", err)
	}

	return user, nil
}
