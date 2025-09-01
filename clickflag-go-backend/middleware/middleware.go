package middleware

import (
	"context"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// SetupMiddleware sets up all middleware for the application
func SetupMiddleware(app *fiber.App) {
	// CORS middleware - Environment'a göre ayarla
	corsConfig := getCORSConfig()
	app.Use(cors.New(corsConfig))

	// Rate limiting middleware - 10 requests per second
	app.Use(limiter.New(limiter.Config{
		Max:        20,
		Expiration: 1 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"success": false,
				"message": "Rate limit exceeded. Please try again later.",
				"error":   "TOO_MANY_REQUESTS",
			})
		},
	}))

	// Logger middleware
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Local",
	}))

	// Recovery middleware
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	// Request timeout middleware
	app.Use(func(c *fiber.Ctx) error {
		// Set timeout for all requests
		ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
		defer cancel()

		c.Locals("ctx", ctx)
		return c.Next()
	})
}

// getCORSConfig returns CORS configuration based on environment
func getCORSConfig() cors.Config {
	environment := os.Getenv("ENVIRONMENT")

	switch environment {
	case "production":
		return cors.Config{
			AllowOrigins: "https://clickflag.com,https://www.clickflag.com",
			AllowMethods: "GET,POST,OPTIONS",
			AllowHeaders: "Origin, Content-Type, Accept",
		}
	case "staging":
		return cors.Config{
			AllowOrigins: "https://staging.clickflag.com,https://clickflag.com",
			AllowMethods: "GET,POST,OPTIONS",
			AllowHeaders: "Origin, Content-Type, Accept",
		}
	default: // development
		return cors.Config{
			AllowOrigins: "http://localhost:3000,http://localhost:8080,http://127.0.0.1:3000",
			AllowMethods: "GET,POST,OPTIONS",
			AllowHeaders: "Origin, Content-Type, Accept",
		}
	}
}

// HealthCheckMiddleware adds health check headers
func HealthCheckMiddleware(c *fiber.Ctx) error {
	c.Set("X-Health-Check", "true")
	return c.Next()
}
