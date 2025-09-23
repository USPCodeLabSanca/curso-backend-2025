package database

import (
	"codelab/backend/pkg/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*
Cria uma conexão com o banco de dados via GORM.
Recebe as variáveis da conexão via struct config.
*/
func NewDatabaseConnection(config *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Postgres.Host,
		config.Postgres.Port,
		config.Postgres.User,
		config.Postgres.Password,
		config.Postgres.Database,
		"disable",
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}