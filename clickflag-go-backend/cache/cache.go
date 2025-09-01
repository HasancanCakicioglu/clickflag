package cache

import (
	"clickflag-go-backend/constants"
	"clickflag-go-backend/models"
	"maps"
	"sync/atomic"
)

// CountryCache handles country data operations (read-optimized)
type CountryCache struct {
	// Atomic value for countries data
	countries atomic.Value // map[string]*models.Country
}

// NewCountryCache creates a new country cache instance
func NewCountryCache() *CountryCache {
	initialCountries := make(map[string]*models.Country)

	cc := &CountryCache{}

	// Initialize atomic value
	cc.countries.Store(initialCountries)

	return cc
}

// GetCountries returns all countries from cache (lock-free read)
func (cc *CountryCache) GetCountries() map[string]*models.Country {
	// Atomic load of countries
	countries := cc.countries.Load().(map[string]*models.Country)

	// Create a shallow copy to avoid race conditions (modern approach)
	result := make(map[string]*models.Country, len(countries))
	maps.Copy(result, countries)
	return result
}

// RefreshCountries updates the cache with fresh data from database (atomic swap)
func (cc *CountryCache) RefreshCountries(countries []models.Country) {
	// Create new countries map
	newCountries := make(map[string]*models.Country, len(countries))

	for i := range countries {
		country := countries[i]
		newCountries[country.CountryCode] = &country
	}

	// Atomic swap of the countries
	cc.countries.Store(newCountries)
}

// GetCountryByCode returns a specific country from cache (lock-free read)
func (cc *CountryCache) GetCountryByCode(countryCode string) (*models.Country, bool) {
	// Atomic load to ensure we get consistent view
	countries := cc.countries.Load().(map[string]*models.Country)

	country, exists := countries[countryCode]
	if !exists {
		return nil, false
	}

	// Return a copy to avoid race conditions
	countryCopy := *country
	return &countryCopy, true
}

// CountryCounter represents a single country's counter with cache-line padding
// This ensures each country counter is on its own cache line (64 bytes)
type CountryCounter struct {
	// Counter value (4 bytes)
	Counter int32

	// Padding to fill the rest of the cache line (60 bytes)
	// Total: 4 + 60 = 64 bytes (exact cache line size)
	_ [60]byte
}

// PendingUpdatesCache handles pending updates operations (write-optimized)
// Each country has its own cache line to prevent false sharing
type PendingUpdatesCache struct {
	// Each country gets its own cache line
	counters atomic.Value // map[string]*CountryCounter
}

// NewPendingUpdatesCache creates a new pending updates cache instance
func NewPendingUpdatesCache() *PendingUpdatesCache {
	puc := &PendingUpdatesCache{}

	// Initialize atomic counters for all countries from constants
	// Each counter is allocated separately to ensure cache line isolation
	counters := make(map[string]*CountryCounter)
	for _, code := range constants.AllCountryCodes {
		// Allocate each counter separately to ensure cache line isolation
		counter := &CountryCounter{
			Counter: 0,
		}
		counters[code] = counter
	}

	// Store the counters map atomically
	puc.counters.Store(counters)

	return puc
}

// AddPendingUpdate adds a country code to pending updates using atomic operations
func (puc *PendingUpdatesCache) AddPendingUpdate(countryCode string) {
	counters := puc.counters.Load().(map[string]*CountryCounter)
	if counter, exists := counters[countryCode]; exists {
		atomic.AddInt32(&counter.Counter, 1)
	}
}

// GetPendingUpdates returns all pending updates and clears them atomically
func (puc *PendingUpdatesCache) GetPendingUpdates() map[string]int32 {
	result := make(map[string]int32)
	counters := puc.counters.Load().(map[string]*CountryCounter)

	for code, counter := range counters {
		// Atomic exchange to get current value and reset to 0
		value := atomic.SwapInt32(&counter.Counter, 0)
		if value > 0 {
			result[code] = value
		}
	}

	return result
}

// HasPendingUpdates checks if there are any pending updates (atomic read)
func (puc *PendingUpdatesCache) HasPendingUpdates() bool {
	counters := puc.counters.Load().(map[string]*CountryCounter)
	for _, counter := range counters {
		if atomic.LoadInt32(&counter.Counter) > 0 {
			return true
		}
	}
	return false
}

// Cache represents the in-memory cache system with atomic operations
// This is the main cache that combines both CountryCache and PendingUpdatesCache
type Cache struct {
	countries      *CountryCache
	pendingUpdates *PendingUpdatesCache
}

// NewCache creates a new cache instance
func NewCache() *Cache {
	return &Cache{
		countries:      NewCountryCache(),
		pendingUpdates: NewPendingUpdatesCache(),
	}
}

// Backward compatibility methods - these delegate to the appropriate sub-cache

// GetCountries returns all countries from cache (lock-free read)
func (c *Cache) GetCountries() map[string]*models.Country {
	return c.countries.GetCountries()
}

// AddPendingUpdate adds a country code to pending updates using atomic operations
func (c *Cache) AddPendingUpdate(countryCode string) {
	c.pendingUpdates.AddPendingUpdate(countryCode)
}

// GetPendingUpdates returns all pending updates and clears them atomically
func (c *Cache) GetPendingUpdates() map[string]int32 {
	return c.pendingUpdates.GetPendingUpdates()
}

// RefreshCountries updates the cache with fresh data from database (atomic swap)
func (c *Cache) RefreshCountries(countries []models.Country) {
	c.countries.RefreshCountries(countries)
}

// GetCountryByCode returns a specific country from cache (lock-free read)
func (c *Cache) GetCountryByCode(countryCode string) (*models.Country, bool) {
	return c.countries.GetCountryByCode(countryCode)
}

// HasPendingUpdates checks if there are any pending updates (atomic read)
func (c *Cache) HasPendingUpdates() bool {
	return c.pendingUpdates.HasPendingUpdates()
}
