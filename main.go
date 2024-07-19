package main

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/joho/godotenv"

	hdl "github.com/Siravitt/azure-storage/handler"
	repo "github.com/Siravitt/azure-storage/repository"
	srv "github.com/Siravitt/azure-storage/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("load .env error: %v", err)
	}

	e := echo.New()

	e.Logger.SetLevel(log.INFO)

	// sqlite database
	db := connectDatabase()
	client := createAzureClient()

	repo := repo.NewRepository(db)
	srv := srv.NewService(repo, client)
	handler := hdl.NewHandler(srv)

	e.GET("/health", handler.Health)
	e.POST("/signed-url-upload", handler.GenerateSASUpload)
	e.POST("/signed-url-read", handler.GenerateSASRead)

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

func connectDatabase() *sql.DB {
	// os.Remove("./database/user.sqlite")

	db, err := sql.Open("sqlite3", "./database/user.sqlite")
	if err != nil {
		log.Fatalf("connect to database error: %v", err)
	}
	return db
}

func createAzureClient() *azblob.Client {
	connectionStr := os.Getenv("AZURE_STORAGE_CONNECTION")

	client, err := azblob.NewClientFromConnectionString(connectionStr, nil)
	if err != nil {
		log.Fatalf("NewClientFromConnectionString error: %s", err)
	}
	return client
}
