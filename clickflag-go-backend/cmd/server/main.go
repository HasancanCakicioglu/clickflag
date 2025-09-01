package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"clickflag-go-backend/cache"
	"clickflag-go-backend/config"
	"clickflag-go-backend/database"
	"clickflag-go-backend/handlers"
	"clickflag-go-backend/middleware"
	"clickflag-go-backend/processor"
	"clickflag-go-backend/utils"

	"github.com/gofiber/fiber/v2"
)

// go run ./cmd/server/main.go
func main() {
	// Load configuration first
	cfg := config.Load()

	// Initialize logger with environment
	if err := utils.InitLogger(cfg.Environment); err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer utils.CloseLogger()

	utils.AppLogger.Info("Starting server with configuration: %+v", cfg)

	// Initialize database
	if err := database.InitDatabase(cfg.DatabasePath); err != nil {
		utils.AppLogger.Critical("Failed to initialize database: %v", err)
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.CloseDatabase()

	// Initialize cache
	cacheInstance := cache.NewCache()

	// Load initial data from database to cache
	log.Println("Loading initial data from database...")
	countries, err := database.GetAllCountries()
	if err != nil {
		utils.AppLogger.Critical("Failed to load initial countries: %v", err)
		log.Fatalf("Failed to load initial countries: %v", err)
	}
	cacheInstance.RefreshCountries(countries)
	log.Printf("Loaded %d countries into cache", len(countries))

	// Initialize background processor with cron job (every 5 seconds)
	bgProcessor := processor.NewBackgroundProcessor(cacheInstance, "*/5 * * * * *")
	bgProcessor.Start()
	defer bgProcessor.Stop()

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		AppName:      "ClickFlag Go Backend",
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		IdleTimeout:  3 * time.Second,
	})

	// Setup middleware
	middleware.SetupMiddleware(app)

	// Initialize handlers
	countryHandler := handlers.NewCountryHandler(cacheInstance)

	// Setup routes
	setupRoutes(app, countryHandler)

	// Start server in a goroutine
	go func() {
		log.Printf("Server starting on port %s", cfg.Port)
		if err := app.Listen(":" + cfg.Port); err != nil {
			utils.AppLogger.Critical("Failed to start server: %v", err)
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Printf("Error during server shutdown: %v", err)
	}

	log.Println("Server stopped gracefully")
}

// setupRoutes sets up all application routes
func setupRoutes(app *fiber.App, countryHandler *handlers.CountryHandler) {
	// Health check endpoint
	app.Get("/health", middleware.HealthCheckMiddleware, countryHandler.HealthCheck)

	// API routes
	api := app.Group("/api/v1")

	// Country routes
	countries := api.Group("/countries")
	countries.Get("/", countryHandler.GetCountries)
	countries.Post("/", countryHandler.AddCountry)

	// Root endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Click Flag API",
			"version": "1.0.0",
			"endpoints": fiber.Map{
				"health":      "/health",
				"countries":   "/api/v1/countries",
				"add_country": "/api/v1/countries (POST)",
			},
		})
	})
}
