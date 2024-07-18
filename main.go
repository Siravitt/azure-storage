package main

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Siravitt/azure-storage/handler"
	"github.com/Siravitt/azure-storage/repository"
	"github.com/Siravitt/azure-storage/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	e := echo.New()

	e.Logger.SetLevel(log.INFO)

	db, err := initDatabase()
	if err != nil {
		log.Fatalf("connect to database error: %s", err)
	}
	_ = db

	repo := repository.NewRepository(db)
	srv := service.NewService(repo)
	handler := handler.NewHandler(srv)

	e.GET("/health", handler.Health)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Start server
	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func initDatabase() (*sql.DB, error) {
	// os.Remove("./database/user.sqlite")

	db, err := sql.Open("sqlite3", "./database/user.sqlite")
	if err != nil {
		return nil, err
	}
	return db, nil

}
