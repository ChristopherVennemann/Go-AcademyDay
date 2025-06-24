package main

import (
	"github.com/ChristopherVennemann/Go-AcademyDay/internal"
	"github.com/ChristopherVennemann/Go-AcademyDay/internal/database"
	"github.com/ChristopherVennemann/Go-AcademyDay/internal/handler"
	"github.com/ChristopherVennemann/Go-AcademyDay/internal/repository"
	"github.com/ChristopherVennemann/Go-AcademyDay/internal/service"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

func main() {
	cfg := internal.CreateConfig()
	conn, err := database.NewConnection(
		cfg.DbConfig.Address,
		cfg.DbConfig.MaxOpenConnections,
		cfg.DbConfig.MaxIdleConnections,
		cfg.DbConfig.MaxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}
	defer func(c *database.Database) {
		err := c.Connection.Close()
		if err != nil {
			log.Panic(err)
		}
	}(conn)
	log.Println("db connection pool established")

	useRepository := repository.NewRepository(conn)
	useService := service.NewService(useRepository)
	useHandler := handler.NewRouter(useService)

	server := &http.Server{
		Addr:         cfg.Port,
		Handler:      useHandler,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	log.Printf("server has started at %s", cfg.Port)
	log.Fatal(server.ListenAndServe())
}
