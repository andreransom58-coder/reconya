package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type DatabaseType string

const (
	SQLite DatabaseType = "sqlite"
)

type Config struct {
	JwtKey       []byte
	Port         string
	DatabaseType DatabaseType
	SQLitePath   string
	Username     string
	Password     string
	DatabaseName string
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	databaseName := os.Getenv("DATABASE_NAME")
	if databaseName == "" {
		return nil, fmt.Errorf("DATABASE_NAME environment variable is not set")
	}

	username := os.Getenv("LOGIN_USERNAME")
	password := os.Getenv("LOGIN_PASSWORD")
	if username == "" || password == "" {
		return nil, fmt.Errorf("LOGIN_USERNAME or LOGIN_PASSWORD environment variables are not set")
	}

	jwtSecret := os.Getenv("JWT_SECRET_KEY")
	if jwtSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET_KEY environment variable is not set")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3008"
	}

	dbType := string(SQLite)

	config := &Config{
		JwtKey:       []byte(jwtSecret),
		Port:         port,
		DatabaseType: DatabaseType(dbType),
		Username:     username,
		Password:     password,
		DatabaseName: databaseName,
	}

	sqlitePath := os.Getenv("SQLITE_PATH")
	if sqlitePath == "" {
		sqlitePath = filepath.Join("data", fmt.Sprintf("%s.db", databaseName))
	}
	config.SQLitePath = sqlitePath

	return config, nil
}
