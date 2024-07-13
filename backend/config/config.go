package config

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type (
	Database struct {
		Host     string `mapstructure:"host" validate:"required"`
		Port     int    `mapstructure:"port" validate:"required"`
		User     string `mapstructure:"user" validate:"required"`
		Password string `mapstructure:"password" validate:"required"`
		DBName   string `mapstructure:"dbname" validate:"required"`
		SSLMode  string `mapstructure:"sslmode" validate:"required"`
		Schema   string `mapstructure:"schema" validate:"required"`
	}

	Server struct {
		Port         int           `mapstructure:"port" validate:"required"`
		AllowOrigins []string      `mapstructure:"allowOrigins" validate:"required"`
		Timeout      time.Duration `mapstructure:"timeout" validate:"required"`
		BodyLimit    string        `mapstructure:"bodyLimit" validate:"required"`
	}

	Config struct {
		Database *Database `mapstructure:"database" validate:"required"`
		Server   *Server   `mapstructure:"server" validate:"required"`
	}
)

func GettingConfig() *Config {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	configInstance := &Config{}

	configInstance.Database = &Database{
		Host:     os.Getenv("DB_HOST"),
		Port:     getIntEnv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		Schema:   os.Getenv("DB_SCHEMA"),
	}

	// Server configuration
	configInstance.Server = &Server{
		Port:         getIntEnv("SERVER_PORT"),
		AllowOrigins: strings.Split(os.Getenv("SERVER_ALLOW_ORIGINS"), ","),
		Timeout:      getDurationEnv("SERVER_TIMEOUT"),
		BodyLimit:    os.Getenv("SERVER_BODY_LIMIT"),
	}

	validate := validator.New()
	if err := validate.Struct(configInstance); err != nil {
		panic(err)
	}

	return configInstance
}

// getIntEnv parses an integer environment variable or panics if it is invalid or missing.
func getIntEnv(key string) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		panic(err)
	}
	return value
}

// getDurationEnv parses a duration environment variable or panics if it is invalid or missing.
func getDurationEnv(key string) time.Duration {
	value, err := time.ParseDuration(os.Getenv(key))
	if err != nil {
		panic(err)
	}
	return value
}
