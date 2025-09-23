package main

import (
	"codelab/backend/internal/usecases/neighborhood"
	"codelab/backend/internal/usecases/weather"
	"codelab/backend/pkg/config"
	"codelab/backend/pkg/database"
	"codelab/backend/pkg/migrations"
	"codelab/backend/pkg/server"
	"net/http"

	"go.uber.org/fx"
)


/*
Função principal.

Executa um aplicativo usando Fx para realizar a 
injeção de dependências.
Desse modo, struct e ponteiros que são
necessários como parâmetros de função serão
fornecidos automaticamente.

Claro, o Fx deve conhecer os métodos de inicialização
para esses tipos de dados.

Por exemplo, Fx fará a injeção de um *gorm.DB se
alguma função registrada produza como resultado
*gorm.DB.
*/
func main() {
	app := fx.New(
		config.Module,
		database.Module,
		neighborhood.Module,
		weather.Module,
		server.Module,

		// Database migration
		fx.Invoke(migrations.ApplyMigrations),

		// Server setup
		fx.Invoke(func(*http.Server) {}),
	)

	app.Run()
}