package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/automaxprocs/maxprocs"
	"go.uber.org/zap"
	"lawise-go/config"
	"lawise-go/internal/domain/user/repository"
	"lawise-go/internal/domain/user/service"
	userHandler "lawise-go/internal/entrypoint/http/user"
	"lawise-go/pkg/aws/credential"
	"lawise-go/pkg/database"
	middlewares "lawise-go/pkg/http.router.middlewares"
	"lawise-go/pkg/logger"
	"lawise-go/pkg/server"
	"os"
	"os/signal"
	"time"
)

func main() {

	fmt.Println("App Lawise starting...")
	defer fmt.Println("App Lawise shutdown")

	if err := run(); err != nil {
		fmt.Println("startup", "ERROR", err)
		os.Exit(1)
	}
}

func run() error {

	l := logger.Get()

	ctx := context.Background()

	// Set the maximum number of processes
	if _, err := maxprocs.Set(maxprocs.Logger(func(format string, args ...interface{}) {
		l.Info(format, zap.Any("args", args))
	})); err != nil {
		return fmt.Errorf("failed to set maxprocs: %w", err)
	}

	cfg := config.GetConfig()

	// Initialize AWS credentials
	if _, err := credential.NewCredential(cfg.AWSCredential, ctx); err != nil {
		return fmt.Errorf("failed to create AWS credential: %w", err)
	}

	// Initialize database connection
	db, err := database.New(cfg.LawiseDb)
	if err != nil {
		l.Fatal("Error connecting to database:", zap.Error(err))
	}
	defer db.Close()

	// Initialize repositories
	usrRepo := repository.NewUserRepository(db)

	// Initialize services
	usrService := service.NewService(usrRepo)

	// Initialize handlers
	usrHandler := userHandler.NewUserHandler(usrService)

	// Initialize and start the web server
	apiServer := server.NewHttpServerApi(&cfg.HttpServer)
	router := apiServer.GetRouter()

	// Global middlewares
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.AddMiddleware(gin.Recovery())

	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.AddMiddleware(middlewares.Logging("/swagger", "/healthcheck", "/openApi"))

	//configure handlers
	router.ConfigureRoutes(usrHandler)

	go apiServer.Run(ctx)

	l.Info("Server started successfully")

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	l.Info("Shutting down server...")

	// Gracefully shutdown the server with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		l.Fatal("Server forced to shutdown:", zap.Error(err))
		return err
	}

	return nil
}
