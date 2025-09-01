package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

// LogLevel represents the severity of a log message
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
	CRITICAL
)

// Logger handles application logging
type Logger struct {
	file        *os.File
	environment string
	logLevel    LogLevel
	filename    string
	maxSize     int64 // Maximum file size in bytes
	maxFiles    int   // Maximum number of log files to keep
	maxDays     int   // Maximum days to keep log files
}

// NewLogger creates a new logger instance
func NewLogger(filename, environment string) (*Logger, error) {
	// Create logs directory if it doesn't exist
	if err := os.MkdirAll("logs", 0755); err != nil {
		return nil, fmt.Errorf("failed to create logs directory: %w", err)
	}

	// Set max file size and retention based on environment
	var maxSize int64
	var maxFiles, maxDays int

	switch environment {
	case "production":
		maxSize = 10 * 1024 * 1024 // 10MB
		maxFiles = 10              // Keep only 10 log files
		maxDays = 7                // Keep logs for 7 days
	case "development":
		maxSize = 5 * 1024 * 1024 // 5MB
		maxFiles = 5              // Keep only 5 log files
		maxDays = 30              // Keep logs for 30 days
	default:
		maxSize = 10 * 1024 * 1024 // 10MB
		maxFiles = 10              // Keep only 10 log files
		maxDays = 7                // Keep logs for 7 days
	}

	// Open log file
	file, err := os.OpenFile("logs/"+filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	// Set log level based on environment
	var logLevel LogLevel
	switch environment {
	case "production":
		logLevel = INFO // Production'da DEBUG logları gösterme
	case "development":
		logLevel = DEBUG
	default:
		logLevel = INFO
	}

	return &Logger{
		file:        file,
		environment: environment,
		logLevel:    logLevel,
		filename:    filename,
		maxSize:     maxSize,
		maxFiles:    maxFiles,
		maxDays:     maxDays,
	}, nil
}

// Close closes the log file
func (l *Logger) Close() error {
	return l.file.Close()
}

// rotateLog rotates the log file if it exceeds max size
func (l *Logger) rotateLog() error {
	// Get current file info
	fileInfo, err := l.file.Stat()
	if err != nil {
		return err
	}

	// Check if file size exceeds max size
	if fileInfo.Size() < l.maxSize {
		return nil
	}

	// Close current file
	l.file.Close()

	// Create backup filename with timestamp
	timestamp := time.Now().Format("2006-01-02-15-04-05")
	backupName := fmt.Sprintf("%s.%s", l.filename, timestamp)
	backupPath := filepath.Join("logs", backupName)

	// Rename current file to backup
	if err := os.Rename(filepath.Join("logs", l.filename), backupPath); err != nil {
		return err
	}

	// Open new log file
	file, err := os.OpenFile(filepath.Join("logs", l.filename), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	l.file = file
	return nil
}

// cleanupOldLogs removes old log files based on age and count
func (l *Logger) cleanupOldLogs() error {
	logsDir := "logs"
	files, err := os.ReadDir(logsDir)
	if err != nil {
		return err
	}

	// Get all rotated log files
	var logFiles []struct {
		path    string
		modTime time.Time
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		// Check if it's a rotated log file
		if len(file.Name()) > len(l.filename) && file.Name()[:len(l.filename)] == l.filename {
			filePath := filepath.Join(logsDir, file.Name())
			fileInfo, err := os.Stat(filePath)
			if err != nil {
				continue
			}

			logFiles = append(logFiles, struct {
				path    string
				modTime time.Time
			}{
				path:    filePath,
				modTime: fileInfo.ModTime(),
			})
		}
	}

	// Sort by modification time (oldest first)
	for i := 0; i < len(logFiles)-1; i++ {
		for j := i + 1; j < len(logFiles); j++ {
			if logFiles[i].modTime.After(logFiles[j].modTime) {
				logFiles[i], logFiles[j] = logFiles[j], logFiles[i]
			}
		}
	}

	// Remove files based on age
	cutoffTime := time.Now().AddDate(0, 0, -l.maxDays)
	for _, logFile := range logFiles {
		if logFile.modTime.Before(cutoffTime) {
			os.Remove(logFile.path)
		}
	}

	// Remove files based on count (keep only maxFiles)
	if len(logFiles) > l.maxFiles {
		filesToRemove := len(logFiles) - l.maxFiles
		for i := 0; i < filesToRemove && i < len(logFiles); i++ {
			os.Remove(logFiles[i].path)
		}
	}

	return nil
}

// Log writes a log message with timestamp and level
func (l *Logger) Log(level LogLevel, message string, args ...interface{}) {
	// Check if we should log this level
	if level < l.logLevel {
		return
	}

	// Check if we need to rotate logs
	if err := l.rotateLog(); err != nil {
		log.Printf("Failed to rotate log: %v", err)
	}

	// Cleanup old logs periodically (every 1000 logs)
	if time.Now().Unix()%1000 == 0 {
		l.cleanupOldLogs()
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	levelStr := l.getLevelString(level)

	// Format message with args
	formattedMessage := fmt.Sprintf(message, args...)

	// Create log entry
	logEntry := fmt.Sprintf("[%s] %s: %s\n", timestamp, levelStr, formattedMessage)

	// Write to file
	if _, err := l.file.WriteString(logEntry); err != nil {
		log.Printf("Failed to write to log file: %v", err)
	}

	// Console output based on environment and level
	if l.environment == "development" || level >= ERROR {
		log.Printf("[%s] %s: %s", timestamp, levelStr, formattedMessage)
	}
}

// getLevelString returns the string representation of log level
func (l *Logger) getLevelString(level LogLevel) string {
	switch level {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case CRITICAL:
		return "CRITICAL"
	default:
		return "UNKNOWN"
	}
}

// Warning logs a warning message
func (l *Logger) Warning(message string, args ...interface{}) {
	l.Log(WARNING, message, args...)
}

// Error logs an error message
func (l *Logger) Error(message string, args ...interface{}) {
	l.Log(ERROR, message, args...)
}

// Critical logs a critical error message
func (l *Logger) Critical(message string, args ...interface{}) {
	l.Log(CRITICAL, message, args...)
}

// Debug logs a debug message
func (l *Logger) Debug(message string, args ...interface{}) {
	l.Log(DEBUG, message, args...)
}

// Info logs an info message
func (l *Logger) Info(message string, args ...interface{}) {
	l.Log(INFO, message, args...)
}

// Global logger instance
var AppLogger *Logger

// InitLogger initializes the global logger
func InitLogger(environment string) error {
	var err error
	AppLogger, err = NewLogger("app.log", environment)
	return err
}

// CloseLogger closes the global logger
func CloseLogger() error {
	if AppLogger != nil {
		return AppLogger.Close()
	}
	return nil
}
