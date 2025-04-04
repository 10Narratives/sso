package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

// Config represents the top-level configuration structure for the application.
type Config struct {
	Env      string        `yaml:"env" env-default:"local"`       // Environment (e.g., "local", "dev", "prod")
	Storage  StorageConfig `yaml:"storage"`                       // Storage-related configuration
	TokenTTL time.Duration `yaml:"token_ttl" env-required:"true"` // Token time-to-live duration
	GRPC     GRPCConfig    `yaml:"grpc"`                          // gRPC server configuration
}

// StorageConfig defines the storage-specific configuration.
type StorageConfig struct {
	Driver string `yaml:"driver" env-required:"true"`
	DSN    string `yaml:"dsn" env-required:"true"`
}

// GRPCConfig defines the gRPC server-specific configuration.
type GRPCConfig struct {
	Port    int           `yaml:"port" env-default:"4000"`
	Timeout time.Duration `yaml:"timeout" env-default:"10s"`
}

// MustLoad loads the application configuration using a path which is gotten by flag or .env file.
func MustLoad() *Config {
	configPath := fetchConfigPath()
	return MustLoadFromPath(configPath)
}

// MustLoadFromPath loads the application configuration from the specified file path.
// It panics if the file is missing, unreadable, or contains invalid configuration.
func MustLoadFromPath(configPath string) *Config {
	if configPath == "" {
		panic("cannot read configuration: " + configPath + " is empty.")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("cannot read configuration: file does not exists")
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read configuration by " + configPath + ": error is occurred " + err.Error())
	}

	return &cfg
}

// fetchConfigPath retrieves the configuration file path, either from a command-line flag or an environment variable.
// If neither is provided, it attempts to load the path from a .env file.
func fetchConfigPath() string {
	var path string

	flag.StringVar(&path, "config", "", "path to configuration file")
	flag.Parse()

	if path == "" {
		err := godotenv.Load()
		if err != nil {
			panic("flag is not set and .env cannot be loaded")
		}
		path = os.Getenv("CONFIG_PATH")
	}

	return path
}
