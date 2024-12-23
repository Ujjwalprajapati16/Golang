package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ujjwalprajapati16/students-api/internal/config"
	"github.com/ujjwalprajapati16/students-api/internal/http/handlers/student"
	"github.com/ujjwalprajapati16/students-api/internal/storage/sqlite"
)

func main() {
	//load config
	cfg := config.MustLoad()
	//load db
	storage, err := sqlite.New(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Connected to database", slog.String("address", cfg.StoragePath), slog.String("Version", "1.0.0"))
	//setup router
	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.New(storage))

	//setup server
	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	slog.Info("Starting server", "address", cfg.Address)

	// wait for interrupt signal
	done := make(chan os.Signal, 1)

	// notify when signal is received
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Server failed to start", err)
		}
	}()

	<-done

	slog.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	// graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Server failed to shutdown", "error", err.Error())
	}

	slog.Info("Server shutdown successfully")
}
