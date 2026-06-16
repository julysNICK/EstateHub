package main

import (
	"context"
	lead "estatehub-api/internal/lead"
	"estatehub-api/internal/platform/config"
	"estatehub-api/internal/platform/database"
	platformhttp "estatehub-api/internal/platform/http"
	"estatehub-api/internal/platform/logger"
	property "estatehub-api/internal/property"
	user "estatehub-api/internal/user"
	"log"
	"net/http"
	nethttp "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	logg := logger.New(cfg.AppEnv)

	db, err := database.Open(cfg.DatabaseUrl)

	if err != nil {
		logg.Error("database connection error", "error", err)
		os.Exit(1)
	}

	defer db.Close()

	migrationCtx, migrationCancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer migrationCancel()

	if err := database.RunMigration(migrationCtx, db, "migrations"); err != nil {
		logg.Error("database migration error", "error", err)
		os.Exit(1)
	}

	leadRepository := lead.NewRepository(db)
	leadService := lead.NewService(leadRepository)
	leadHandler := lead.NewHandler(leadService)

	userRepository := user.NewRepository(db)
	userService := user.NewUserService(userRepository)
	userHandler := user.NewUserHandler(userService)

	propertyRepository := property.NewPropertyRepository(db)
	propertyService := property.NewPropertyService(propertyRepository)
	propertyHandler := property.NewPropertyHandler(propertyService)

	healthHandler := platformhttp.NewHealthHandler(db, cfg.ReadinessTimeout())

	router := platformhttp.NewRouter(platformhttp.RouterConfig{
		Logger:          logg,
		HealthHandler:   healthHandler,
		LeadHandler:     leadHandler,
		UserHandler:     userHandler,
		PropertyHandler: propertyHandler,
	})

	log.Printf("%s running in %s mode on port %s\n", cfg.AppName, cfg.AppEnv, cfg.AppPort)

	server := &http.Server{
		Addr:         cfg.Addr(),
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		logg.Info("server started",
			"app", cfg.AppName,
			"env", cfg.AppEnv,
			"port", cfg.AppPort)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logg.Error("server error", "error", err)
			os.Exit(1)
		}
	}()

	waitForShutdown(logg, server)

}

func waitForShutdown(logg interface {
	Info(msg string, args ...any)
	Error(msg string, args ...any)
}, server *nethttp.Server) {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	logg.Info("server shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logg.Error("graceful shutdown failed", "error", err)
		return
	}

	logg.Info("server stopped")
}
