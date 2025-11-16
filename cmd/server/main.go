// PR Service for Avito backend trainee assignment 2025
package main

import (
	"AvitoPRService/internal/app"
	"AvitoPRService/internal/config"
	"AvitoPRService/internal/config/logger"
	"AvitoPRService/internal/router"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {
	logger.LogInit()
	logger.Logger.Info("Starting Avito-PR backend server")

	if err := godotenv.Load(); err != nil {
		logger.Logger.Fatalf("Error loading .env file: %s", err.Error())
	}
	config := config.NewConfig()
	app := app.NewApp(config)

	r := router.NewRouter(app)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.ServerPort),
		Handler: r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		log.Println("timeout of 5 seconds.")
	}

	log.Println("Server exiting")
}
