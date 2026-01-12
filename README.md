# Simple Web App

A Go web application with user authentication, built using Chi router and PostgreSQL.

## Tech Stack

- Go 1.25+
- PostgreSQL
- Chi Router
- Gorilla CSRF Protection
- Air (live reload)

## Prerequisites

- Docker and Docker Compose
- Go 1.25 or higher

## Setup

1. Start the database:
```bash
docker-compose up -d
```

2. Run the application:
```bash
go run main.go
```

Or use Air for live reload:
```bash
air
```

The application will be available at `http://localhost:8080`

## Services

- **App**: `http://localhost:8080`
- **Adminer** (Database UI): `http://localhost:3333`
- **PostgreSQL**: `localhost:5434`

## Database

- **User**: postgres
- **Password**: postgres
- **Database**: simple
- **Port**: 5434

## Features

- User registration and authentication
- CSRF protection
- Static pages (Home, Contact, FAQ)
- Template rendering with Tailwind CSS

## Project Structure

```
.
├── cmd/              # Experimental and utility code
├── controllers/      # HTTP handlers
├── model/            # Database models and SQL
├── routes/           # Route definitions
├── sql/              # Database connection and queries
├── templates/        # HTML templates
├── views/            # Template parsing logic
└── main.go           # Application entry point
```
