package main

import (
	"context"
	"fmt"
	"go-xray-config/internal/config"
	"go-xray-config/internal/handlers"
	"go-xray-config/internal/middleware"
	"go-xray-config/internal/repository"
	"go-xray-config/internal/services"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	if cfg.GinMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	xrayRepository := repository.NewXrayRepository(
		cfg.XRayConfigPath, cfg.XRayHost, cfg.XrayPublicKey, cfg.XRayServerName,
	)
	xrayService := services.NewXrayService(xrayRepository)
	xrayHandler := handlers.NewXRayHandler(xrayService)

	router := gin.New()

	router.Use(
		gin.Recovery(),
		middleware.AuthMiddleware(cfg.Token),
	)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":   "OK",
			"database": "healthy",
			"time":     time.Now().Format(time.RFC3339),
		})
	})

	api := router.Group("/api")
	xrayHandler.RegisterRoutes(api)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.ServerPort),
		Handler: router,
	}

	go func() {
		slog.Info("Starting server", "port", cfg.ServerPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Failed to start server", "error", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Server forced to shutdown", "error", err)
	}
	slog.Info("Server exited properly")
}
