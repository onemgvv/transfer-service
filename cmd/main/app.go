package main

import (
	"context"
	"github.com/joho/godotenv"
	natslib "github.com/nats-io/nats.go"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"transaction-service/internal/app"
	"transaction-service/internal/config"
	"transaction-service/internal/repository"
	"transaction-service/internal/service"
	"transaction-service/pkg/database/postgres"
	"transaction-service/pkg/nats"
)

const configPath = "configs"

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("[ENV LOAD ERROR]: %s\n", err.Error())
	}

	cfg, err := config.Init(configPath)
	if err != nil {
		log.Fatalf("[CONFIG ERROR]: %s\n", err.Error())
	}

	db, err := postgres.Init(cfg)
	if err != nil {
		log.Fatalf("[POSTGRES DB ERROR]: %s\n", err.Error())
	}

	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)

	streaming := nats.NewStreaming()
	if err != nil {
		log.Fatalf("[NATS CONNECTING ERROR]: %s", err.Error())
	}

	log.Println("Connected to " + natslib.DefaultURL)
	mainApp := app.NewApp(streaming, services)

	if err = mainApp.Start(); err != nil {
		log.Fatalf("[Application Starting Error]")
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	const timeout = 5 * time.Second
	_, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := mainApp.Stop(); err != nil {
		log.Println(err)
	}
	// Closing nats connection
	streaming.NC.Close()

	if err = mainApp.Stop(); err != nil {
		log.Fatalf("[MAIN SERVICE CLOSING] || [FAILED]: %s", err.Error())
	}

	if err = db.Close(); err != nil {
		log.Fatalf("[DATABASE CONN CLOSE] || [FAILED]: %s", err.Error())
	}

	log.Print("Application stopped")
}
