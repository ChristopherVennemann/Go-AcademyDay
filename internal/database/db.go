package database

import (
	"context"
	"database/sql"
	"time"
)

type Database struct {
	Connection *sql.DB
}

func NewConnection(addr string, maxOpenConns, maxIdleConns int, maxIdleTime string) (*Database, error) {
	db, err := sql.Open("postgres", addr)
	if err != nil {
		return nil, err
	}

	duration, err := time.ParseDuration(maxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(duration)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	return &Database{Connection: db}, nil
}
