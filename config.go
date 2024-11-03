package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	query    string
	dsn      string
	duration time.Duration
}

func (c *Config) Duration() time.Duration {
	return c.duration
}

func (c *Config) DSN() string {
	return c.dsn
}

func (c *Config) Query() string {
	return c.query
}

func NewConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	query := os.Getenv("QUERY")
	if query == "" {
		log.Fatal("QUERY environment variable is empty")
	}

	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Fatal("DSN environment variable is empty")
	}

	durationStr := os.Getenv("DURATION")
	if durationStr == "" {
		log.Fatal("DURATION environment variable is empty")
	}

	duration, err := strconv.Atoi(durationStr)
	if err != nil {
		log.Fatal("DURATION environment variable is invalid", err)
	}

	return &Config{
		query:    query,
		dsn:      dsn,
		duration: time.Millisecond * time.Duration(duration),
	}
}
