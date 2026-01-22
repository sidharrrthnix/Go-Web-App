package model

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func Open(config PostgresConfig) (*sql.DB, error) {
	var db *sql.DB
	var err error

	for i := 0; i < 5; i++ {
		db, err = sql.Open("pgx", config.String())
		if err == nil {
			err = db.Ping()
			if err == nil {
				log.Printf("Database connected: %s:%s/%s", config.Host, config.Port, config.Database)
				return db, nil
			}
		}
		log.Printf("Waiting for database... (attempt %d/5): %v", i+1, err)
		time.Sleep(2 * time.Second)
	}
	return nil, fmt.Errorf("failed to connect after retries: %w", err)
}

func DefaultPostgresConfig() PostgresConfig {
	return PostgresConfig{
		Host:     "localhost",
		Port:     "5434",
		User:     "postgres",
		Password: "postgres",
		Database: "simple", // Changed to 'simple' to match your setup
		SSLMode:  "disable",
	}
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
}
