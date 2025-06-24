package testutils

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"os"
)

func SetupTestPostgres(ctx context.Context) (*sql.DB, func(), error) {
	const pgPort = "5432"
	req := testcontainers.ContainerRequest{
		Image:        "postgres:15",
		Env:          map[string]string{"POSTGRES_PASSWORD": "pass", "POSTGRES_DB": "testdb"},
		ExposedPorts: []string{pgPort + "/tcp"},
		WaitingFor:   wait.ForListeningPort(pgPort + "/tcp"),
	}

	postgresC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, nil, err
	}

	host, _ := postgresC.Host(ctx)
	port, _ := postgresC.MappedPort(ctx, pgPort)
	dsn := fmt.Sprintf("host=%s port=%s user=postgres password=pass dbname=testdb sslmode=disable", host, port.Port())

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		if err := db.Close(); err != nil {
			log.Printf("couldnt close db connection: %s", err)
		}
		if err := postgresC.Terminate(ctx); err != nil {
			log.Printf("couldnt terminate testcontainer: %s", err)
		}
	}

	schemaBytes, err := os.ReadFile("../database/init.pg.sql")
	if err != nil {
		cleanup()
		return nil, nil, err
	}

	_, err = db.Exec(string(schemaBytes))
	if err != nil {
		cleanup()
		return nil, nil, err
	}

	return db, cleanup, nil
}
