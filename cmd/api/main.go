package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/apk471/go-crud-api/internal/config"
)

func main() {
	fmt.Println("Welcome to CRUD API")
	cfg := config.MustLoad()

	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to CRUD API"))
	})

	server := http.Server{
		Addr: cfg.HttpServer.Addr,
		Handler: router,
	}
	// Start the server
	slog.Info("server is running", "address", cfg.HttpServer.Addr)

	done := make(chan os.Signal, 1)
	// Listen for interrupt and terminate signals
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("server error", "error", err)
		}
	}()
	
	<- done

	slog.Info("server is shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("server shutdown error", "error", err)
	}

	slog.Info("server is shutdown")
}