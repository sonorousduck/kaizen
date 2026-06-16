package initializers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectToDatabase() (*pgxpool.Pool, error) {
	host := os.Getenv("HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("SSL_MODE")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

	config, err := pgxpool.ParseConfig(connectionString)

	if err != nil {
		return nil, fmt.Errorf("Unable to parse connection string: %w", err)
	}

	config.MinConns = 2                       // Minimum idle connections
	config.MaxConns = 10                      // Maximum connections in pool
	config.MaxConnLifetime = time.Hour        // Maximum lifetime of a connection
	config.MaxConnIdleTime = 30 * time.Minute // Maximum idle time before closing
	config.HealthCheckPeriod = time.Minute    // How often to check connection health

	ctx := context.Background()

	pool, err := pgxpool.NewWithConfig(ctx, config)

	if err != nil {
		return nil, fmt.Errorf("Unable to create connection pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("Unable to ping database: %w", err)
	}

	log.Println("Successfully connected to PostgreSQL database")
	return pool, nil
}
