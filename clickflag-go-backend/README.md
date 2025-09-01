# Country Code API

A high-performance country code API built with Go Fiber framework. Stores country codes in in-memory cache and writes to database via background processor.

## Features

- **Thread-safe in-memory cache**: Country codes and values stored in memory
- **Background processing**: Writes pending updates to database every 5 seconds
- **SQLite3 database**: Lightweight and fast database support
- **RESTful API**: GET and POST endpoints
- **Health check**: System status monitoring
- **Graceful shutdown**: Safe application termination
- **CORS support**: Cross-origin request handling
- **Rate limiting**: 10 requests per second per IP address
- **Logging**: Detailed logging system

## Supported Country Codes

The API accepts the following country codes (ISO 3166-1 alpha-2 standard):
- `TR` - Turkey
- `US` - United States
- `UK` - United Kingdom
- `DE` - Germany
- `FR` - France
- `TW` - Taiwan
- `JP` - Japan
- `KR` - South Korea
- `CN` - China
- `IN` - India
- `BR` - Brazil
- `RU` - Russia
- `CA` - Canada
- `AU` - Australia
- `MX` - Mexico
- and more ...

## Installation

### Requirements

- Go 1.25.0 or higher
- SQLite3
- Docker (optional)

### Package Installation

```bash
go mod tidy
```

### Environment Variables

Create environment file:

```bash
# Copy example file
cp env.example .env

# Edit .env file
nano .env
```

**Example .env file:**
```env
# Server Configuration
PORT=8080
HOST=0.0.0.0

# Database Configuration
DATABASE_PATH=./data/countries.db

# Logging
LOG_LEVEL=info
ENVIRONMENT=production

# Docker Configuration
DOCKER_PORT=8080
DOCKER_CONTAINER_NAME=clickflag-backend
```

## Running

### Local Development

```bash
go run cmd/server/main.go
```

Or build and run:

```bash
go build -o bin/server cmd/server/main.go
./bin/server
```

### Docker Deployment

#### Production Build

```bash
# Build Docker image
docker build -t clickflag-backend .

# Run container
docker run -p 8080:8080 -v $(pwd)/data:/app/data clickflag-backend
```

#### Docker Compose

```bash
# Production
docker-compose up -d
```

#### Makefile Commands

```bash
# Help
make help

# Docker build
make docker-build

# Production run
make docker-run

# Stop containers
make docker-stop

# View logs
make docker-logs

# Log management
make logs-check    # Check log directory status
make logs-clean    # Clean old log files
make logs-status   # Show log system status
```

## API Endpoints

### 1. Health Check
```
GET /health
```

**Response:**
```json
{
  "status": "healthy",
  "timestamp": "2024-01-01T12:00:00Z",
  "cache": {
    "last_refresh": "2024-01-01T12:00:00Z",
    "has_pending": false
  }
}
```

### 2. Get All Countries
```
GET /api/v1/countries
```

**Response:**
```json
{
  "success": true,
  "message": "Countries retrieved successfully",
  "data": {
    "TR": 5,
    "US": 12,
    "UK": 8,
    "DE": 15
  }
}
```

### 3. Add Country Code
```
POST /api/v1/countries
Content-Type: application/json

{
  "country_code": "TR"
}
```

**Response:**
```json
{
  "success": true,
  "message": "Country code added to pending updates successfully",
  "data": {
    "country_code": "TR",
    "status": "pending"
  }
}
```

## Project Structure

```
clickflag-go-backend/
├── cmd/
│   └── server/
│       └── main.go          # Main application file
├── config/
│   └── config.go            # Configuration management
├── database/
│   └── database.go          # Database operations
├── handlers/
│   └── country.go           # HTTP handlers
├── middleware/
│   └── middleware.go        # Middleware functions
├── models/
│   └── country.go           # Data models
├── processor/
│   └── background.go        # Background processor
├── cache/
│   └── cache.go             # In-memory cache system
├── migrations/
│   └── 001_create_countries_table.sql  # Database migration
├── go.mod                   # Go module file
└── README.md               # This file
```

## Production Log Management

### Log Levels
- **DEBUG**: Detailed debugging information (development only)
- **INFO**: General application information
- **WARNING**: Warning messages
- **ERROR**: Error messages
- **CRITICAL**: Critical errors that require immediate attention

### Log Rotation & Retention
- **Development**: 
  - 5MB file size limit
  - Maximum 5 log files
  - Keep logs for 30 days
- **Production**: 
  - 10MB file size limit
  - Maximum 10 log files
  - Keep logs for 7 days
- **Docker Logs**: 
  - 10MB file size limit
  - Maximum 3 log files
- **Backup naming**: `app.log.2024-01-01-15-04-05` format

### Production Log Configuration
```env
ENVIRONMENT=production
LOG_LEVEL=info
```

### Log File Locations
- **Development**: `./logs/app.log`
- **Production**: `/app/logs/app.log` (Docker container)

### Monitoring Logs
```bash
# View current logs
tail -f logs/app.log

# View rotated logs
ls -la logs/

# Search for errors
grep "ERROR" logs/app.log

# Search for critical errors
grep "CRITICAL" logs/app.log

# View Docker logs
docker logs clickflag-backend

# View Docker logs with follow
docker logs -f clickflag-backend

# Check log file sizes
du -h logs/*

# Check total log directory size
du -sh logs/
```

## Technical Details

### Cache System
- Thread-safe in-memory cache
- Separate map for pending updates
- Safe writing with atomic operations

### Background Processor
- Runs every 5 seconds using cron job (UTC synchronized)
- Writes pending updates to database
- Automatically refreshes cache
- Uses cron expression: `*/5 * * * * *`

### Database
- Uses SQLite3
- Validates country codes with CHECK constraint
- Fast queries with indexes

### Middleware
- CORS support
- Request logging
- Error recovery
- Request timeout (3 seconds)

## Testing

### cURL Testing

```bash
# Health check
curl http://localhost:8080/health

# Get all countries
curl http://localhost:8080/api/v1/countries

# Add country code
curl -X POST http://localhost:8080/api/v1/countries \
  -H "Content-Type: application/json" \
  -d '{"country_code": "TR"}'
```

## Development

### Adding New Features
1. Add new function in relevant package
2. Define endpoint in handler
3. Test it
4. Update README

### Log Levels
- `info`: General information
- `error`: Error messages
- `debug`: Detailed logs for development

## License

This project is licensed under the MIT License.
