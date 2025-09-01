package tests

import (
	"clickflag-go-backend/constants"
	"testing"
)

// TestCountryCount tests that we have exactly 195 countries
func TestCountryCount(t *testing.T) {
	count := constants.GetCountryCount()
	if count != 195 {
		t.Errorf("Expected 195 countries, got %d", count)
	}
}

// TestAllCountryCodesLength tests that AllCountryCodes has 195 elements
func TestAllCountryCodesLength(t *testing.T) {
	if len(constants.AllCountryCodes) != 195 {
		t.Errorf("Expected 195 country codes, got %d", len(constants.AllCountryCodes))
	}
}

// TestValidCountryCodes tests that all country codes in the list are valid
func TestValidCountryCodes(t *testing.T) {
	for _, code := range constants.AllCountryCodes {
		if !constants.IsValidCountryCode(code) {
			t.Errorf("Country code %s should be valid", code)
		}
	}
}

// TestSpecificCountryCodes tests specific important country codes
func TestSpecificCountryCodes(t *testing.T) {
	importantCodes := []string{"tr", "us", "gb", "de", "fr", "jp", "cn", "in", "br", "ru"}

	for _, code := range importantCodes {
		if !constants.IsValidCountryCode(code) {
			t.Errorf("Important country code %s should be valid", code)
		}
	}
}

// TestInvalidCountryCodes tests that invalid codes return false
func TestInvalidCountryCodes(t *testing.T) {
	invalidCodes := []string{"xx", "yy", "zz", "aa", "cc"}

	for _, code := range invalidCodes {
		if constants.IsValidCountryCode(code) {
			t.Errorf("Invalid country code %s should not be valid", code)
		}
	}
}

// TestPalestineIncluded tests that Palestine (ps) is included
func TestPalestineIncluded(t *testing.T) {
	if !constants.IsValidCountryCode("ps") {
		t.Error("Palestine (ps) should be included in the country codes")
	}
}

// TestTaiwanIncluded tests that Taiwan (tw) is included
func TestTaiwanIncluded(t *testing.T) {
	if !constants.IsValidCountryCode("tw") {
		t.Error("Taiwan (tw) should be included in the country codes")
	}
}

// TestCountryCodesUnique tests that all country codes are unique
func TestCountryCodesUnique(t *testing.T) {
	seen := make(map[string]bool)

	for _, code := range constants.AllCountryCodes {
		if seen[code] {
			t.Errorf("Duplicate country code found: %s", code)
		}
		seen[code] = true
	}
}

// TestCountryCodesFormat tests that all country codes are 2 characters
func TestCountryCodesFormat(t *testing.T) {
	for _, code := range constants.AllCountryCodes {
		if len(code) != 2 {
			t.Errorf("Country code %s should be 2 characters long", code)
		}
	}
}

// BenchmarkIsValidCountryCode benchmarks the validation function
func BenchmarkIsValidCountryCode(b *testing.B) {
	testCodes := []string{"tr", "us", "gb", "de", "fr", "jp", "cn", "in", "br", "ru"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		code := testCodes[i%len(testCodes)]
		constants.IsValidCountryCode(code)
	}
}
