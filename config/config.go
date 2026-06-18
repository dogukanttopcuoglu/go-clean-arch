package config

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	App    AppConfig
	HTTP   HTTPConfig
	Logger LoggerConfig
}

type AppConfig struct {
	Environment string
}

type HTTPConfig struct {
	Port string
}

type LoggerConfig struct {
	Level       string
	ServiceName string
}

func Load() (*Config, error) {
	v := viper.New()
	v.SetDefault("app.environment", "development")
	v.SetDefault("http.port", "8080")
	v.SetDefault("logger.level", "debug")
	v.SetDefault("logger.service_name", "clean-lab-api")

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		var notFound viper.ConfigFileNotFoundError
		if !errors.As(err, &notFound) {
			return nil, fmt.Errorf("read config file: %w", err)
		}
	}

	v.SetEnvPrefix("CLEAN_LAB")
	v.SetEnvKeyReplacer(strings.NewReplacer(
		".", "_",
		"-", "_",
	))

	v.AutomaticEnv()

	cfg := &Config{
		App: AppConfig{
			Environment: v.GetString("app.environment"),
		},
		HTTP: HTTPConfig{
			Port: v.GetString("http.port"),
		},
		Logger: LoggerConfig{
			Level:       v.GetString("logger.level"),
			ServiceName: v.GetString("logger.service_name"),
		},
	}

	return cfg, nil
}
