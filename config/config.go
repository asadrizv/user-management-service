package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/spf13/viper"
)

// Config represents the application configuration
type Config struct {
	Server struct {
		Port int `yaml:"port" env:"SERVER_PORT"`
	} `yaml:"server"`
	Database struct {
		ConnectionString string `yaml:"connection_string" env:"DATABASE_CONNECTION_STRING"`
	} `yaml:"database"`
	Logging struct {
		Level string `yaml:"level" env:"LOG_LEVEL"`
	} `yaml:"logging"`
	RabbitMQURL string `yaml:"rabbitMQURL" env:"AMQP_URL"`
}

// Get loads and returns the configuration object
func Get() (*Config, error) {
	// Load the environment variables
	viper.AutomaticEnv()

	// Unmarshal the config into a struct
	config := &Config{}
	err := viper.Unmarshal(config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Parse the env tags
	err = env.Parse(config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse env: %w", err)
	}

	return config, nil
}
