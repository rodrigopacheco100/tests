package database

import (
	"fmt"
	"tests/infrastructure/env"

	"github.com/jackc/pgx"
)

func NewConnection() (*pgx.Conn, error) {
	envs, _ := env.GetEnvs()

	connConfig, err := pgx.ParseConnectionString(envs.DatabaseUrl)
	if err != nil {
		return nil, fmt.Errorf("unable to parse database connection string: %v\n", err)
	}

	conn, err := pgx.Connect(connConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v\n", err)
	}

	return conn, nil
}
