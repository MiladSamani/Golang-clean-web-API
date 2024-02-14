package config

import (
	"errors"
	"log"
	"os"

	"github.com/spf13/viper"
)

// Config struct holds the configuration parameters for the application.
type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
}

// ServerConfig struct holds the server-related configuration parameters.
type ServerConfig struct {
	Port    string // Port on which the server should listen.
	RunMode string // Mode in which the server should run (e.g., development, production).
}

// PostgresConfig struct holds the PostgreSQL database configuration parameters.
type PostgresConfig struct {
	Host     string // PostgreSQL database host.
	Port     string // PostgreSQL database port.
	User     string // PostgreSQL database username.
	Password string // PostgreSQL database password.
	DbName   string // PostgreSQL database name.
	SSLMode  bool   // Whether to use SSL mode for PostgreSQL connection.
}

// RedisConfig struct holds the Redis database configuration parameters.
type RedisConfig struct {
	Host              string // Redis database host.
	Port              string // Redis database port.
	Password          string // Redis database password.
	Db                string // Redis database index.
	MinIdleConnections int    // Minimum number of idle connections in the Redis connection pool.
	PoolSize          int    // Maximum number of connections in the Redis connection pool.
	PoolTimeout       int    // Maximum time (in seconds) to wait for a connection from the Redis connection pool.
}

// GetConfig reads and parses the configuration file based on the environment variable APP_ENV.
func GetConfig() *Config {
	cfgPath := getConfigPath(os.Getenv("APP_ENV")) // Get the configuration file path based on the environment.
	v, err := LoadConfig(cfgPath, "yml")            // Load the configuration file.
	if err != nil {
		log.Fatalf("Error in load config %v", err)
	}

	cfg, err := ParseConfig(v) // Parse the configuration file.
	if err != nil {
		log.Fatalf("Error in parse config %v", err)
	}

	return cfg
}

// ParseConfig parses the configuration from the given Viper instance.(3)
func ParseConfig(v *viper.Viper) (*Config, error) {
	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Printf("Unable to parse config: %v", err)
		return nil, err
	}
	return &cfg, nil
}

// LoadConfig loads the configuration from the specified file.(2)
func LoadConfig(filename string, fileType string) (*viper.Viper, error) {
	v := viper.New() // Create a new Viper instance.
	v.SetConfigType(fileType)
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig() // Read the configuration file.
	if err != nil {
		log.Printf("Unable to read config: %v", err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}
	return v, nil
}

// getConfigPath returns the path of the configuration file based on the environment.(1)
func getConfigPath(env string) string {
	if env == "docker" {
		return "config/config-docker"
	} else if env == "production" {
		return "config/config-production"
	} else {
		return "../config/config-development"
	}
}
