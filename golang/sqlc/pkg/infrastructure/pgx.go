package infrastructure

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	"github.com/mukezhz/learn/tree/main/golang/sqlc/pkg/framework"
)

type Database struct {
	*pgxpool.Pool
}

func NewDatabase(logger framework.Logger, env *framework.Env) *Database {
	url := fmt.Sprintf("%s&TimeZone=%s", env.DBURL, env.DBTimeZone)
	logger.Info("Connecting to database", "url", url)
	ctx := context.Background()

	// Parse the connection string and setup config
	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		panic(err)
	}

	// Add Tracer for logging if needed
	config.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger:   logger.GetTraceLogger(),
		LogLevel: logger.GetTraceLogLevel(logger.Level().String()),
	}

	config.ConnConfig.ConnectTimeout = 5 * time.Minute

	// Create a connection pool
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %v", err))
	}

	// Return a Database struct with the connection pool
	return &Database{
		Pool: pool,
	}
}

func (db *Database) Close() error {
	if db.Pool != nil {
		db.Pool.Close()
		return nil
	}
	return fmt.Errorf("pool is nil")
}
