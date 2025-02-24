package main

import (
	"fmt"
	"log"
	"net/http"

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
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
