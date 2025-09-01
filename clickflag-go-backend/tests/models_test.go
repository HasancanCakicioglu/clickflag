package tests

import (
	"clickflag-go-backend/models"
	"testing"
)

// TestCountryStruct tests the Country struct
func TestCountryStruct(t *testing.T) {
	country := models.Country{
		ID:          1,
		CountryCode: "tr",
		Value:       100,
	}

	if country.ID != 1 {
		t.Errorf("Expected ID 1, got %d", country.ID)
	}

	if country.CountryCode != "tr" {
		t.Errorf("Expected CountryCode 'tr', got %s", country.CountryCode)
	}

	if country.Value != 100 {
		t.Errorf("Expected Value 100, got %d", country.Value)
	}
}

// TestCountryRequestStruct tests the CountryRequest struct
func TestCountryRequestStruct(t *testing.T) {
	request := models.CountryRequest{
		CountryCode: "us",
	}

	if request.CountryCode != "us" {
		t.Errorf("Expected CountryCode 'us', got %s", request.CountryCode)
	}
}

// TestCountryResponseStruct tests the CountryResponse struct
func TestCountryResponseStruct(t *testing.T) {
	response := models.CountryResponse{
		Success: true,
		Message: "Success",
		Data:    "test data",
	}

	if !response.Success {
		t.Error("Expected Success to be true")
	}

	if response.Message != "Success" {
		t.Errorf("Expected Message 'Success', got %s", response.Message)
	}

	if response.Data != "test data" {
		t.Errorf("Expected Data 'test data', got %v", response.Data)
	}
}

// TestIsValidCountryCode tests the IsValidCountryCode function
func TestIsValidCountryCode(t *testing.T) {
	// Test valid codes
	validCodes := []string{"tr", "us", "gb", "de", "fr", "jp", "cn", "in", "br", "ru"}
	for _, code := range validCodes {
		if !models.IsValidCountryCode(code) {
			t.Errorf("Country code %s should be valid", code)
		}
	}

	// Test invalid codes
	invalidCodes := []string{"xx", "yy", "zz", "aa", "cc"}
	for _, code := range invalidCodes {
		if models.IsValidCountryCode(code) {
			t.Errorf("Country code %s should not be valid", code)
		}
	}
}

// TestIsValidCountryCodeEdgeCases tests edge cases for country code validation
func TestIsValidCountryCodeEdgeCases(t *testing.T) {
	// Test empty string
	if models.IsValidCountryCode("") {
		t.Error("Empty string should not be a valid country code")
	}

	// Test single character
	if models.IsValidCountryCode("a") {
		t.Error("Single character should not be a valid country code")
	}

	// Test three characters
	if models.IsValidCountryCode("abc") {
		t.Error("Three characters should not be a valid country code")
	}

	// Test uppercase (should be case sensitive)
	if models.IsValidCountryCode("TR") {
		t.Error("Uppercase should not be a valid country code")
	}

	// Test mixed case
	if models.IsValidCountryCode("Tr") {
		t.Error("Mixed case should not be a valid country code")
	}
}
