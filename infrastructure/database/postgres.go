package database

import (
	"context"
	"fmt"
	"tests/infrastructure/env"
	"time"

	pgx "github.com/jackc/pgx/v5"
)

func NewConnection() (*pgx.Conn, error) {
	envs, _ := env.GetEnvs()

	connConfig, err := pgx.ParseConfig(envs.DatabaseUrl)
	if err != nil {
		return nil, fmt.Errorf("unable to parse database connection string: %v\n", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	conn, err := pgx.Connect(ctx, connConfig.ConnString())
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v\n", err)
	}

	return conn, nil
}
