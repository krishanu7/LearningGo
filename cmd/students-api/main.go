package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/krishanu7/students-api/internal/config"
)

func main() {
	// Load config
	cfg := config.MustLoadConfig()
	fmt.Printf("Loaded Config: %+v\n", cfg)

	// Setup router
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("My First Go web server is up and running!"))
	})

	// Setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	fmt.Println("Starting server on", cfg.Addr)
	slog.Info("Starting server on", slog.String("addr", cfg.Addr))
	// Create Channel to listen for OS signals
	signalChan := make(chan os.Signal, 1);

	signal.Notify(signalChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Start server
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server:", err)
		}
	}()

	// Wait for OS signal
	<-signalChan

	slog.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)

	if err != nil {
		slog.Error("Failed to shutdown server:", slog.String("error", err.Error()))
	}
	slog.Info("Server shutdown successfully")
}
