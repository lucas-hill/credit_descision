package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Port             string
	Environment      string // local, dev, prod
	DynamoDBEndpoint string // http://localhost:8000 for local
	DynamoDBTable    string
	AWSRegion        string
	LogLevel         string
}

func Load() (*Config, error) {
	cfg := &Config{
		Port:        getEnv("PORT", "8080"),
		Environment: getEnv("ENVIRONMENT", "local"),
		// DynamoDBEndpoint: getEnv("DYNAMODB_ENDPOINT", "http://localhost:8000"),
		// DynamoDBTable:    getEnv("DYNAMODB_TABLE", "CreditApplications"),
		// AWSRegion:        getEnv("AWS_REGION", "us-east-1"),
		LogLevel: getEnv("LOG_LEVEL", "info"),
	}

	if err := cfg.validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	return cfg, nil
}

func (c *Config) IsLocal() bool {
	return c.Environment == "local"
}

func (c *Config) validate() error {
	port, err := strconv.Atoi(c.Port)
	if err != nil || port < 1 || port > 65535 {
		return fmt.Errorf("invalid PORT: %s", c.Port)
	}

	validEnvs := map[string]bool{"local": true, "dev": true, "prod": true}
	if !validEnvs[c.Environment] {
		return fmt.Errorf("invalid ENVIRONMENT: %s (must be local, dev, or prod)", c.Environment)
	}

	return nil
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
