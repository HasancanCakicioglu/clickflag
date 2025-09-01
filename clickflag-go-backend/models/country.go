package models

import "clickflag-go-backend/constants"

// Country represents a country code with its associated value
type Country struct {
	ID          int    `json:"id" db:"id"`
	CountryCode string `json:"country_code" db:"country_code"`
	Value       int    `json:"value" db:"value"`
}

// CountryAPI represents a country for API responses (without ID)
type CountryAPI struct {
	CountryCode string `json:"country_code"`
	Value       int    `json:"value"`
}

// CountryRequest represents the request body for creating/updating a country
type CountryRequest struct {
	CountryCode string `json:"country_code" validate:"required"`
}

// CountryResponse represents the response for country operations
type CountryResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// IsValidCountryCode checks if the given country code is valid using constants
func IsValidCountryCode(code string) bool {
	return constants.IsValidCountryCode(code)
}
