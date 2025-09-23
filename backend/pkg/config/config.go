package config

import "github.com/spf13/viper"

type Config struct {
	Postgres PostgresConfig
	Workers int
	API string
}

type PostgresConfig struct {
	User     string
	Password string
	Database string
	Host     string
	Port     string
}

func NewConfig() (*Config, error) {
	viper.SetConfigFile(".env.example")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	// Postgres setup
	cfg.Postgres.User = viper.GetString("POSTGRES_USER")
	cfg.Postgres.Password = viper.GetString("POSTGRES_PASSWORD")
	cfg.Postgres.Database = viper.GetString("POSTGRES_DB")
	cfg.Postgres.Host = viper.GetString("POSTGRES_HOST")
	cfg.Postgres.Port = viper.GetString("POSTGRES_PORT")

	cfg.API = viper.GetString("API_URL")

	return &cfg, nil
}