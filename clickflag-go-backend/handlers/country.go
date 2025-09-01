package handlers

import (
	"log"
	"time"

	"clickflag-go-backend/cache"
	"clickflag-go-backend/models"

	"github.com/gofiber/fiber/v2"
)

// CountryHandler handles country-related HTTP requests
type CountryHandler struct {
	cache *cache.Cache
}

// NewCountryHandler creates a new country handler
func NewCountryHandler(cache *cache.Cache) *CountryHandler {
	return &CountryHandler{
		cache: cache,
	}
}

// GetCountries returns all countries from cache
func (h *CountryHandler) GetCountries(c *fiber.Ctx) error {
	countries := h.cache.GetCountries()

	// Convert map to object format for JSON response
	countryMap := make(map[string]int, len(countries))
	for _, country := range countries {
		countryMap[country.CountryCode] = country.Value
	}

	return c.JSON(models.CountryResponse{
		Success: true,
		Message: "Countries retrieved successfully",
		Data:    countryMap,
	})
}

// AddCountry adds a country code to pending updates
func (h *CountryHandler) AddCountry(c *fiber.Ctx) error {
	var request models.CountryRequest

	if err := c.BodyParser(&request); err != nil {
		log.Printf("Error parsing request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.CountryResponse{
			Success: false,
			Message: "Invalid request body",
		})
	}

	// Validate country code
	if request.CountryCode == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.CountryResponse{
			Success: false,
			Message: "Country code is required",
		})
	}

	if !models.IsValidCountryCode(request.CountryCode) {
		return c.Status(fiber.StatusBadRequest).JSON(models.CountryResponse{
			Success: false,
			Message: "Invalid country code.",
		})
	}

	// Add to pending updates
	h.cache.AddPendingUpdate(request.CountryCode)

	log.Printf("Added country code %s to pending updates", request.CountryCode)

	return c.JSON(models.CountryResponse{
		Success: true,
		Message: "Country code added to pending updates successfully",
		Data: map[string]string{
			"country_code": request.CountryCode,
			"status":       "pending",
		},
	})
}

// HealthCheck returns health status
func (h *CountryHandler) HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":    "healthy",
		"timestamp": time.Now().UTC(),
		"cache": fiber.Map{
			"has_pending": h.cache.HasPendingUpdates(),
		},
	})
}
