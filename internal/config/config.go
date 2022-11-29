package config

import (
	"os"

	"github.com/spf13/viper"
)

const (
	defaultPostgresPort = "5432"
	defaultPostgresHost = "localhost"
)

type (
	Config struct {
		DB DatabaseConfig
	}

	DatabaseConfig struct {
		Postgres PostgresConfig
	}

	PostgresConfig struct {
		User     string `mapstructure:"user"`
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Name     string `mapstructure:"dbName"`
		Password string `mapstructure:"password"`
		SSLMode  string `mapstructure:"sslMode"`
	}
)

func Init(cfgDir string) (*Config, error) {
	setupDefaultValues()

	if err := parseConfigFile(cfgDir); err != nil {
		return nil, err
	}
	var cfg Config
	if err := unmarshall(&cfg); err != nil {
		return nil, err
	}

	parseEnvFile(&cfg)
	return &cfg, nil
}

func parseConfigFile(folder string) error {
	viper.AddConfigPath(folder)
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func unmarshall(cfg *Config) error {
	return viper.UnmarshalKey("database", &cfg.DB)
}

func parseEnvFile(cfg *Config) {
	cfg.DB.Postgres.Password = os.Getenv("DB_POSTGRES_PASSWORD")
}

func setupDefaultValues() {
	viper.SetDefault("database.postgres.host", defaultPostgresHost)
	viper.SetDefault("database.postgres.port", defaultPostgresPort)
}
