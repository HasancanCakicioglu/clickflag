#!/bin/bash

# Log Monitoring Script
# Usage: ./scripts/log-monitor.sh [check|clean|status]

LOG_DIR="./logs"
MAX_SIZE_MB=100  # Maximum total log directory size in MB

case "$1" in
    "check")
        echo "=== Log Directory Status ==="
        echo "Directory: $LOG_DIR"
        
        if [ ! -d "$LOG_DIR" ]; then
            echo "Log directory does not exist!"
            exit 1
        fi
        
        # Check total size
        total_size=$(du -sm "$LOG_DIR" | cut -f1)
        echo "Total size: ${total_size}MB"
        
        # Check file count
        file_count=$(find "$LOG_DIR" -name "*.log*" | wc -l)
        echo "Log files: $file_count"
        
        # List files with sizes
        echo ""
        echo "=== Log Files ==="
        ls -lah "$LOG_DIR"/*.log* 2>/dev/null || echo "No log files found"
        
        # Check if size exceeds limit
        if [ "$total_size" -gt "$MAX_SIZE_MB" ]; then
            echo ""
            echo "⚠️  WARNING: Log directory size (${total_size}MB) exceeds limit (${MAX_SIZE_MB}MB)"
            echo "Run './scripts/log-monitor.sh clean' to clean old logs"
        fi
        ;;
        
    "clean")
        echo "=== Cleaning Old Logs ==="
        
        if [ ! -d "$LOG_DIR" ]; then
            echo "Log directory does not exist!"
            exit 1
        fi
        
        # Remove files older than 7 days
        echo "Removing log files older than 7 days..."
        find "$LOG_DIR" -name "*.log.*" -mtime +7 -delete
        
        # Keep only the 10 most recent files
        echo "Keeping only the 10 most recent log files..."
        ls -t "$LOG_DIR"/*.log.* 2>/dev/null | tail -n +11 | xargs -r rm
        
        echo "Cleanup completed!"
        ;;
        
    "status")
        echo "=== Log System Status ==="
        
        # Check if application is running
        if pgrep -f "clickflag-backend" > /dev/null; then
            echo "✅ Application is running"
        else
            echo "❌ Application is not running"
        fi
        
        # Check log directory
        if [ -d "$LOG_DIR" ]; then
            echo "✅ Log directory exists"
            total_size=$(du -sm "$LOG_DIR" 2>/dev/null | cut -f1 || echo "0")
            echo "📊 Log directory size: ${total_size}MB"
        else
            echo "❌ Log directory does not exist"
        fi
        
        # Check Docker logs
        if docker ps | grep -q "clickflag-backend"; then
            echo "✅ Docker container is running"
            docker_logs_size=$(docker logs clickflag-backend 2>/dev/null | wc -c)
            docker_logs_size_mb=$((docker_logs_size / 1024 / 1024))
            echo "📊 Docker logs size: ${docker_logs_size_mb}MB"
        else
            echo "❌ Docker container is not running"
        fi
        ;;
        
    *)
        echo "Usage: $0 {check|clean|status}"
        echo ""
        echo "Commands:"
        echo "  check   - Check log directory status and sizes"
        echo "  clean   - Clean old log files"
        echo "  status  - Show overall log system status"
        exit 1
        ;;
esac
