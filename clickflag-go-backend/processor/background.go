package processor

import (
	"context"
	"log"

	"clickflag-go-backend/cache"
	"clickflag-go-backend/database"

	"github.com/robfig/cron/v3"
)

// BackgroundProcessor handles background processing tasks
type BackgroundProcessor struct {
	cache    *cache.Cache
	cronExpr string
	ctx      context.Context
	cancel   context.CancelFunc
	cron     *cron.Cron
}

// NewBackgroundProcessor creates a new background processor
func NewBackgroundProcessor(cache *cache.Cache, cronExpression string) *BackgroundProcessor {
	ctx, cancel := context.WithCancel(context.Background())

	return &BackgroundProcessor{
		cache:    cache,
		cronExpr: cronExpression,
		ctx:      ctx,
		cancel:   cancel,
	}
}

// Start starts the background processor
func (bp *BackgroundProcessor) Start() {
	log.Printf("Starting background processor with cron expression: %s", bp.cronExpr)

	// Create new cron instance with seconds enabled
	bp.cron = cron.New(cron.WithSeconds())

	// Add the job to cron
	entryID, err := bp.cron.AddFunc(bp.cronExpr, bp.processPendingUpdates)
	if err != nil {
		log.Printf("Error adding cron job: %v", err)
		return
	}

	log.Printf("Cron job added with entry ID: %d", entryID)

	// Process immediately on start
	bp.processPendingUpdates()

	// Start the cron scheduler
	bp.cron.Start()

	log.Println("Background processor started successfully")
}

// Stop stops the background processor
func (bp *BackgroundProcessor) Stop() {
	log.Println("Stopping background processor...")

	if bp.cron != nil {
		// Stop the cron scheduler
		ctx := bp.cron.Stop()
		<-ctx.Done()
		log.Println("Cron scheduler stopped")
	}

	bp.cancel()
	log.Println("Background processor stopped")
}

// processPendingUpdates processes all pending country code updates
func (bp *BackgroundProcessor) processPendingUpdates() {
	// Get pending updates from cache
	pendingUpdates := bp.cache.GetPendingUpdates()

	if len(pendingUpdates) == 0 {
		log.Println("No pending updates to process")
		return
	}

	log.Printf("Processing %d pending updates", len(pendingUpdates))

	// Process each pending update
	for countryCode, count := range pendingUpdates {
		log.Printf("Processing %d updates for country code: %s", count, countryCode)

		// Increment the value in database by the total count (more efficient)
		if err := database.IncrementCountryValueBy(countryCode, count); err != nil {
			log.Printf("Error incrementing value for country %s: %v", countryCode, err)
			continue
		}
	}

	// Refresh cache with updated data
	bp.refreshCache()
}

// refreshCache refreshes the cache with fresh data from database
func (bp *BackgroundProcessor) refreshCache() {
	countries, err := database.GetAllCountries()
	if err != nil {
		log.Printf("Error refreshing cache: %v", err)
		return
	}

	bp.cache.RefreshCountries(countries)
	log.Printf("Cache refreshed with %d countries", len(countries))
}
