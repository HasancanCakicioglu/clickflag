package tests

import (
	"clickflag-go-backend/cache"
	"clickflag-go-backend/models"
	"sync"
	"testing"
	"unsafe"
)

// TestCountryCounterSize tests that CountryCounter is exactly 64 bytes
func TestCountryCounterSize(t *testing.T) {
	size := unsafe.Sizeof(cache.CountryCounter{})
	if size != 64 {
		t.Errorf("CountryCounter size should be 64 bytes, got %d", size)
	}
}

// TestCacheLineAlignment tests that CountryCounter is properly aligned
func TestCacheLineAlignment(t *testing.T) {
	counter := &cache.CountryCounter{}
	offset := unsafe.Offsetof(counter.Counter)
	if offset != 0 {
		t.Errorf("Counter should be at offset 0, got %d", offset)
	}
}

// TestConcurrentUpdates tests that multiple goroutines can update different countries
func TestConcurrentUpdates(t *testing.T) {
	c := cache.NewCache()

	// Test countries
	countries := []string{"tr", "us", "in", "jp", "de"}

	// Number of goroutines per country
	numGoroutines := 10
	updatesPerGoroutine := 100

	var wg sync.WaitGroup

	// Start goroutines for each country
	for _, country := range countries {
		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func(countryCode string) {
				defer wg.Done()
				for j := 0; j < updatesPerGoroutine; j++ {
					c.AddPendingUpdate(countryCode)
				}
			}(country)
		}
	}

	wg.Wait()

	// Get pending updates
	pending := c.GetPendingUpdates()

	// Check that each country has the expected number of updates
	expectedUpdates := numGoroutines * updatesPerGoroutine
	for _, country := range countries {
		if count, exists := pending[country]; !exists {
			t.Errorf("Country %s should have pending updates", country)
		} else if count != int32(expectedUpdates) {
			t.Errorf("Country %s should have %d updates, got %d", country, expectedUpdates, count)
		}
	}
}

// TestCacheRefresh tests that cache refresh works correctly
func TestCacheRefresh(t *testing.T) {
	c := cache.NewCache()

	// Create test countries
	testCountries := []models.Country{
		{ID: 1, CountryCode: "tr", Value: 100},
		{ID: 2, CountryCode: "us", Value: 200},
		{ID: 3, CountryCode: "in", Value: 300},
	}

	// Refresh cache
	c.RefreshCountries(testCountries)

	// Get countries from cache
	countries := c.GetCountries()

	// Check that all countries are present
	for _, expected := range testCountries {
		if country, exists := countries[expected.CountryCode]; !exists {
			t.Errorf("Country %s should exist in cache", expected.CountryCode)
		} else if country.Value != expected.Value {
			t.Errorf("Country %s should have value %d, got %d", expected.CountryCode, expected.Value, country.Value)
		}
	}
}

// TestGetCountryByCode tests individual country retrieval
func TestGetCountryByCode(t *testing.T) {
	c := cache.NewCache()

	// Create test countries
	testCountries := []models.Country{
		{ID: 1, CountryCode: "tr", Value: 100},
		{ID: 2, CountryCode: "us", Value: 200},
	}

	// Refresh cache
	c.RefreshCountries(testCountries)

	// Test existing country
	if country, exists := c.GetCountryByCode("tr"); !exists {
		t.Error("Country 'tr' should exist")
	} else if country.Value != 100 {
		t.Errorf("Country 'tr' should have value 100, got %d", country.Value)
	}

	// Test non-existing country
	if _, exists := c.GetCountryByCode("xx"); exists {
		t.Error("Country 'xx' should not exist")
	}
}

// TestHasPendingUpdates tests the HasPendingUpdates function
func TestHasPendingUpdates(t *testing.T) {
	c := cache.NewCache()

	// Initially should have no pending updates
	if c.HasPendingUpdates() {
		t.Error("Should have no pending updates initially")
	}

	// Add some updates
	c.AddPendingUpdate("tr")
	c.AddPendingUpdate("us")

	// Should have pending updates now
	if !c.HasPendingUpdates() {
		t.Error("Should have pending updates after adding them")
	}

	// Get pending updates (this clears them)
	c.GetPendingUpdates()

	// Should have no pending updates again
	if c.HasPendingUpdates() {
		t.Error("Should have no pending updates after clearing them")
	}
}

// BenchmarkConcurrentUpdates benchmarks concurrent update performance
func BenchmarkConcurrentUpdates(b *testing.B) {
	c := cache.NewCache()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		countries := []string{"tr", "us", "in", "jp", "de", "fr", "gb", "ca", "au", "br"}
		i := 0
		for pb.Next() {
			c.AddPendingUpdate(countries[i%len(countries)])
			i++
		}
	})
}

// BenchmarkGetCountries benchmarks GetCountries performance
func BenchmarkGetCountries(b *testing.B) {
	c := cache.NewCache()

	// Setup test data
	testCountries := make([]models.Country, 196)
	for i := 0; i < 196; i++ {
		testCountries[i] = models.Country{
			ID:          i + 1,
			CountryCode: "test",
			Value:       i * 10,
		}
	}
	c.RefreshCountries(testCountries)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.GetCountries()
	}
}
