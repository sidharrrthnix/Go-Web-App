package main

import (
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	models "web-app.com/simple/sql"
)

func main() {
	cfg := models.DefaultPostgresConfig()
	// Override to use postgres database initially (for creating simple_test)
	cfg.Database = "postgres"

	db, err := models.Open(cfg)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer db.Close()

	fmt.Println("Connected to postgres database")

	// Check if 'simple_test' database exists (for testing only)
	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = 'simple_test')").Scan(&exists)
	if err != nil {
		log.Fatalf("Failed to check database: %v", err)
	}

	if !exists {
		fmt.Println("Creating 'simple_test' database...")
		_, err = db.Exec("CREATE DATABASE simple_test")
		if err != nil {
			log.Fatalf("Failed to create database: %v", err)
		}
		fmt.Println("Database 'simple_test' created!")
	} else {
		fmt.Println("Database 'simple_test' already exists")
	}

	db.Close()

	// Small delay to ensure database is ready
	time.Sleep(1 * time.Second)

	// Now connect to 'simple_test' database with retry
	cfg.Database = "simple_test"

	// Create users table (matching sql/user.go schema)
	_, err = db.Exec(`DROP TABLE IF EXISTS users CASCADE`)
	if err != nil {
		log.Fatalf("Failed to drop users table: %v", err)
	}

	_, err = db.Exec(`
	CREATE TABLE users (
		id SERIAL PRIMARY KEY,
		email TEXT UNIQUE NOT NULL,
		password_hash TEXT NOT NULL
	)
`)
	if err != nil {
		log.Fatalf("Failed to create users table: %v", err)
	}
	fmt.Println("Users table created")

	// Create orders table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS orders (
			id SERIAL PRIMARY KEY,
			user_id INT NOT NULL,
			amount INT NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		log.Fatalf("Failed to create orders table: %v", err)
	}
	fmt.Println("Orders table created")

	fmt.Println("✅ All done! Database and tables are ready.")
}
