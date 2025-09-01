package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"clickflag-go-backend/models"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db   *sql.DB
	once sync.Once
)

// InitDatabase initializes the database connection and runs migrations
func InitDatabase(dbPath string) error {
	var err error
	once.Do(func() {
		// Ensure directory exists
		dir := filepath.Dir(dbPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Fatalf("Error creating database directory: %v", err)
		}

		// Open database connection
		db, err = sql.Open("sqlite3", dbPath)
		if err != nil {
			log.Fatalf("Error opening database: %v", err)
		}

		// Set connection pool settings
		db.SetMaxOpenConns(1) // SQLite only supports one writer at a time
		db.SetMaxIdleConns(1)
		db.SetConnMaxLifetime(time.Hour)

		// Test connection
		if err := db.Ping(); err != nil {
			log.Fatalf("Error pinging database: %v", err)
		}

		// Run migrations
		if err := runMigrations(); err != nil {
			log.Fatalf("Error running migrations: %v", err)
		}

		log.Println("Database initialized successfully")
	})

	return err
}

// GetDB returns the database instance
func GetDB() *sql.DB {
	return db
}

// CloseDatabase closes the database connection
func CloseDatabase() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

// runMigrations executes database migrations
func runMigrations() error {
	// Read migration file from migrations directory
	migrationPath := "migrations/001_create_countries_table.sql"
	migrationSQL, err := os.ReadFile(migrationPath)
	if err != nil {
		log.Fatalf("Error reading migration file %s: %v", migrationPath, err)
	}

	// Execute migration
	_, err = db.Exec(string(migrationSQL))
	if err != nil {
		log.Fatalf("Error executing migration: %v", err)
	}

	log.Println("Database migrations completed successfully")
	return nil
}

// GetAllCountries retrieves all countries from the database
func GetAllCountries() ([]models.Country, error) {
	query := `
		SELECT id, country_code, value
		FROM countries
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying countries: %w", err)
	}
	defer rows.Close()

	var countries []models.Country
	for rows.Next() {
		var country models.Country
		err := rows.Scan(
			&country.ID,
			&country.CountryCode,
			&country.Value,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning country: %w", err)
		}
		countries = append(countries, country)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating countries: %w", err)
	}

	return countries, nil
}

// IncrementCountryValueBy increments the value for a specific country code by given amount
func IncrementCountryValueBy(countryCode string, amount int32) error {
	query := `
		UPDATE countries 
		SET value = value + ?
		WHERE country_code = ?
	`

	result, err := db.Exec(query, amount, countryCode)
	if err != nil {
		return fmt.Errorf("error updating country value: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("country code %s not found", countryCode)
	}

	return nil
}
