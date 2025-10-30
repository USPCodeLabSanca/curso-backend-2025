package migrations

import (
	"codelab/backend/internal/domain"

	"gorm.io/gorm"
)

/*
Aplica as migrações (criação / atualização]
de tabelas) no banco utilizando as structs
de modelos definidas no domain.
Para cada struct (entidade) uma tabela é gerada.
*/
func ApplyMigrations(db *gorm.DB) error {
	return db.AutoMigrate(&domain.Neighborhood{}, &domain.WeatherData{}, &domain.User{})
}