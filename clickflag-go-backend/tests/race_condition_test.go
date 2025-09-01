package tests

import (
	"sync"
	"testing"
	"time"

	"clickflag-go-backend/cache"
)

// TestCacheRaceCondition tests race conditions in cache operations
func TestCacheRaceCondition(t *testing.T) {
	// Create cache instance
	cacheInstance := cache.NewCache()

	// Test parameters - MUCH HIGHER NUMBERS
	countryCode := "TR"
	totalOperations := 10000     // 10,000 operations
	concurrentOperations := 1000 // 1,000 concurrent

	t.Logf("Starting cache race condition test: %d operations, %d concurrent", totalOperations, concurrentOperations)

	var wg sync.WaitGroup
	successfulOperations := 0
	var successMutex sync.Mutex

	// Start concurrent cache operations
	for i := 0; i < totalOperations; i++ {
		wg.Add(1)
		go func(operationID int) {
			defer wg.Done()

			// Add pending update to cache
			cacheInstance.AddPendingUpdate(countryCode)

			successMutex.Lock()
			successfulOperations++
			successMutex.Unlock()
		}(i)

		// Limit concurrent operations
		if (i+1)%concurrentOperations == 0 {
			time.Sleep(1 * time.Millisecond)
		}
	}

	wg.Wait()

	t.Logf("Completed %d successful cache operations", successfulOperations)

	// Check if cache has pending updates
	hasPending := cacheInstance.HasPendingUpdates()
	if !hasPending {
		t.Error("Cache should have pending updates after concurrent operations")
	}

	// Get pending updates
	pendingUpdates := cacheInstance.GetPendingUpdates()
	actualCount := pendingUpdates[countryCode]

	if int(actualCount) != successfulOperations {
		t.Errorf("Cache race condition detected! Expected: %d, Got: %d", successfulOperations, actualCount)
		t.Logf("This indicates a race condition in the cache operations")
	} else {
		t.Logf("✅ No cache race condition detected! Pending updates: %d", actualCount)
	}
}

// TestMultipleCountriesRace tests race conditions with multiple countries
func TestMultipleCountriesRace(t *testing.T) {
	// Create cache instance
	cacheInstance := cache.NewCache()

	// Test multiple countries simultaneously - HIGHER NUMBERS
	countries := []string{"TR", "US", "GB", "DE", "FR", "JP", "CN", "IN", "BR", "RU"}
	operationsPerCountry := 5000 // 5,000 operations per country

	t.Logf("Starting multi-country race test: %d countries, %d operations each", len(countries), operationsPerCountry)

	// Debug: Check initial cache state
	t.Logf("Initial cache state - Has pending updates: %v", cacheInstance.HasPendingUpdates())

	var wg sync.WaitGroup
	results := make(map[string]int)
	var resultsMutex sync.Mutex

	// Start operations for each country
	for _, countryCode := range countries {
		wg.Add(1)
		go func(code string) {
			defer wg.Done()

			countryWg := sync.WaitGroup{}
			countryOperations := 0
			var countryMutex sync.Mutex

			// Send operations for this country
			for i := 0; i < operationsPerCountry; i++ {
				countryWg.Add(1)
				go func(operationID int) {
					defer countryWg.Done()

					// Add pending update to cache
					cacheInstance.AddPendingUpdate(code)

					countryMutex.Lock()
					countryOperations++
					countryMutex.Unlock()
				}(i)
			}

			countryWg.Wait()

			resultsMutex.Lock()
			results[code] = countryOperations
			resultsMutex.Unlock()

			t.Logf("Country %s: Completed %d operations", code, countryOperations)
		}(countryCode)
	}

	wg.Wait()

	// Debug: Check cache state after operations
	t.Logf("After operations - Has pending updates: %v", cacheInstance.HasPendingUpdates())

	// Get all pending updates first
	allPendingUpdates := cacheInstance.GetPendingUpdates()
	t.Logf("All pending updates: %v", allPendingUpdates)

	// Check results for each country
	for _, countryCode := range countries {
		expectedValue := results[countryCode]

		// Get pending updates for this country
		actualCount := allPendingUpdates[countryCode]

		t.Logf("Country %s: Expected %d, Got %d", countryCode, expectedValue, actualCount)

		if int(actualCount) != expectedValue {
			t.Errorf("Multi-country race test failed for %s! Expected: %d, Got: %d",
				countryCode, expectedValue, actualCount)
		} else {
			t.Logf("✅ %s: %d operations processed correctly", countryCode, actualCount)
		}
	}
}

// TestCacheInitialization tests if cache is properly initialized
func TestCacheInitialization(t *testing.T) {
	cacheInstance := cache.NewCache()

	// Test if cache has pending updates initially
	hasPending := cacheInstance.HasPendingUpdates()
	t.Logf("Initial cache state - Has pending updates: %v", hasPending)

	// Test adding updates for different countries
	testCountries := []string{"TR", "US", "GB", "DE", "FR", "JP", "CN"}

	for _, countryCode := range testCountries {
		// Add one update
		cacheInstance.AddPendingUpdate(countryCode)

		// Check if cache has pending updates
		hasPending = cacheInstance.HasPendingUpdates()
		t.Logf("After adding %s - Has pending updates: %v", countryCode, hasPending)

		if !hasPending {
			t.Errorf("Cache should have pending updates after adding %s", countryCode)
		}

		// Get pending updates
		pendingUpdates := cacheInstance.GetPendingUpdates()
		count := pendingUpdates[countryCode]
		t.Logf("Pending updates for %s: %d", countryCode, count)

		if count != 1 {
			t.Errorf("Expected 1 pending update for %s, got %d", countryCode, count)
		}
	}
}

// TestHighConcurrencyLoad tests with very high concurrency
func TestHighConcurrencyLoad(t *testing.T) {
	// Create cache instance
	cacheInstance := cache.NewCache()

	// Test parameters - EXTREMELY HIGH LOAD
	countryCode := "US"
	totalOperations := 100000     // 100,000 operations
	concurrentOperations := 10000 // 10,000 concurrent

	t.Logf("Starting high concurrency test: %d operations, %d concurrent", totalOperations, concurrentOperations)

	startTime := time.Now()

	var wg sync.WaitGroup
	successfulOperations := 0
	var successMutex sync.Mutex

	// Start concurrent operations
	for i := 0; i < totalOperations; i++ {
		wg.Add(1)
		go func(operationID int) {
			defer wg.Done()

			// Add pending update to cache
			cacheInstance.AddPendingUpdate(countryCode)

			successMutex.Lock()
			successfulOperations++
			successMutex.Unlock()
		}(i)
	}

	wg.Wait()

	duration := time.Since(startTime)

	t.Logf("Completed %d successful operations in %v", successfulOperations, duration)
	t.Logf("Average operation time: %v", duration/time.Duration(successfulOperations))
	t.Logf("Operations per second: %.2f", float64(successfulOperations)/duration.Seconds())

	// Check final value
	pendingUpdates := cacheInstance.GetPendingUpdates()
	actualCount := pendingUpdates[countryCode]

	if int(actualCount) != successfulOperations {
		t.Errorf("High concurrency test failed! Expected: %d, Got: %d", successfulOperations, actualCount)
	} else {
		t.Logf("✅ High concurrency test passed! Pending updates: %d", actualCount)
	}
}

// TestBackgroundProcessorSimulation simulates how background processor would handle race conditions
func TestBackgroundProcessorSimulation(t *testing.T) {
	// Create cache instance
	cacheInstance := cache.NewCache()

	// Test parameters - HIGHER NUMBERS
	countryCode := "TR"
	totalOperations := 50000     // 50,000 operations
	concurrentOperations := 5000 // 5,000 concurrent

	t.Logf("Starting background processor simulation test: %d operations, %d concurrent", totalOperations, concurrentOperations)

	var wg sync.WaitGroup
	successfulOperations := 0
	var successMutex sync.Mutex

	// Start concurrent cache operations
	for i := 0; i < totalOperations; i++ {
		wg.Add(1)
		go func(operationID int) {
			defer wg.Done()

			// Add pending update to cache
			cacheInstance.AddPendingUpdate(countryCode)

			successMutex.Lock()
			successfulOperations++
			successMutex.Unlock()
		}(i)

		// Limit concurrent operations
		if (i+1)%concurrentOperations == 0 {
			time.Sleep(1 * time.Millisecond)
		}
	}

	wg.Wait()

	t.Logf("Completed %d successful cache operations", successfulOperations)

	// Simulate background processor getting pending updates
	pendingUpdates := cacheInstance.GetPendingUpdates()
	actualCount := pendingUpdates[countryCode]

	t.Logf("Background processor would process %d updates for %s", actualCount, countryCode)

	if int(actualCount) != successfulOperations {
		t.Errorf("Background processor simulation failed! Expected: %d, Got: %d", successfulOperations, actualCount)
		t.Logf("This indicates a race condition in the cache operations")
	} else {
		t.Logf("✅ Background processor simulation successful! Would process %d updates", actualCount)
	}

	// Check if cache is cleared after getting pending updates
	hasPending := cacheInstance.HasPendingUpdates()
	if hasPending {
		t.Logf("Cache still has pending updates after processing: %v", hasPending)
	} else {
		t.Logf("✅ Cache cleared after processing pending updates")
	}
}

// TestConcurrentBackgroundProcessorSimulation tests multiple background processor cycles
func TestConcurrentBackgroundProcessorSimulation(t *testing.T) {
	// Create cache instance
	cacheInstance := cache.NewCache()

	// Test parameters - HIGHER NUMBERS
	countryCode := "TR"
	cycles := 10                // 10 cycles
	operationsPerCycle := 10000 // 10,000 operations per cycle

	t.Logf("Starting concurrent background processor simulation: %d cycles, %d operations each", cycles, operationsPerCycle)

	for cycle := 0; cycle < cycles; cycle++ {
		t.Logf("Starting cycle %d", cycle+1)

		var wg sync.WaitGroup
		successfulOperations := 0
		var successMutex sync.Mutex

		// Start concurrent operations for this cycle
		for i := 0; i < operationsPerCycle; i++ {
			wg.Add(1)
			go func(operationID int) {
				defer wg.Done()

				// Add pending update to cache
				cacheInstance.AddPendingUpdate(countryCode)

				successMutex.Lock()
				successfulOperations++
				successMutex.Unlock()
			}(i)
		}

		wg.Wait()

		t.Logf("Cycle %d: Completed %d operations", cycle+1, successfulOperations)

		// Simulate background processor processing
		pendingUpdates := cacheInstance.GetPendingUpdates()
		actualCount := pendingUpdates[countryCode]

		t.Logf("Cycle %d: Background processor would process %d updates", cycle+1, actualCount)

		if int(actualCount) != successfulOperations {
			t.Errorf("Cycle %d failed! Expected: %d, Got: %d", cycle+1, successfulOperations, actualCount)
		} else {
			t.Logf("✅ Cycle %d successful! Processed %d updates", cycle+1, actualCount)
		}

		// Small delay between cycles
		time.Sleep(5 * time.Millisecond)
	}
}
