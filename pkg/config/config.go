package config

import "github.com/spf13/viper"

/*
Configurações a serem utilizadas em todo o projeto.
*/

type Config struct {
	Postgres PostgresConfig
	Workers  int
	API      string
	JWT      string
}

type PostgresConfig struct {
	User     string
	Password string
	Database string
	Host     string
	Port     string
}

/*
Viper é uma biblioteca que facilita a definição
de configurações baseadas em arquivos.
No nosso caso, o arquivo base é o '.env', que armazena
senha do banco de dados, etc.

> Em um ambiente de produção, o '.env' NUNCA deve ser
vazado, nem mesmo colocado no reposiótio do github.
*/
func NewConfig() (*Config, error) {
	// Configurando para ler .env
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// Preencher a configuração com os valores lidos
	var cfg Config

	// Postgres setup
	cfg.Postgres.User = viper.GetString("POSTGRES_USER")
	cfg.Postgres.Password = viper.GetString("POSTGRES_PASSWORD")
	cfg.Postgres.Database = viper.GetString("POSTGRES_DB")
	cfg.Postgres.Host = viper.GetString("POSTGRES_HOST")
	cfg.Postgres.Port = viper.GetString("POSTGRES_PORT")

	cfg.API = viper.GetString("API_URL")

	cfg.Workers = 10

	cfg.JWT = viper.GetString("TOKEN_SECRET")

	return &cfg, nil
}
