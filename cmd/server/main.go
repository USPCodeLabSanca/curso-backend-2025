package main

import (
	"codelab/backend/internal/handlers"
	"codelab/backend/internal/middlewares"
	"codelab/backend/internal/usecases/neighborhood"
	"codelab/backend/internal/usecases/user"
	"codelab/backend/internal/usecases/weather"
	"codelab/backend/pkg/config"
	"codelab/backend/pkg/database"
	"codelab/backend/pkg/migrations"
	"codelab/backend/pkg/server"
	"codelab/backend/pkg/token"
	"net"
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

// @title Weather API
// @version 1.0
// @description Weather API docs
// @host localhost:8080

func main() {
	// Criar configuração
	conf, err := config.NewConfig()
	if err != nil {
		panic("Falha na config")
	}

	// Conectar DB
	db, err := database.NewDatabaseConnection(conf)
	if err != nil {
		panic("Falha no banco de dados")
	}

	// Migrar
	if err := migrations.ApplyMigrations(db); err != nil {
		panic("Falha na migração")
	}

	// Criar serviço de tokens
	tokenDoer := token.NewJWTDoer(conf)

	// Criar casos de uso: usuário
	createUserUseCase := user.NewCreateUserUseCase(db)
	loginUserUseCase := user.NewLoginUserUseCase(db, tokenDoer)

	// Criar casos de uso: bairro
	createNeighborhoodUseCase := neighborhood.NewCreateNeighborhoodUseCase(db)
	updateNeighborhoodUseCase := neighborhood.NewUpdateNeighborhoodUseCase(db)
	deleteNeighborhoodUseCase := neighborhood.NewDeleteNeighborhoodUseCase(db)
	findActiveUseCase := neighborhood.NewFindActiveNeighborhoodsUseCase(db)
	findAllUseCase := neighborhood.NewFindAllNeighborhoodsUseCase(db)
	findByIdUseCase := neighborhood.NewFindNeighborhoodByIdUseCase(db)
	findWeatherUseCase := neighborhood.NewFindNeighborhoodWeatherUseCase(db)

	// Criar serviço de coleta de dados
	collectWeatherUseCase := weather.NewCollectWeatherDataUseCase(db, conf)
	collectScheduler := weather.NewWeatherScheduler(collectWeatherUseCase)

	stopChan := weather.InitWeatherScheduler(collectScheduler)
	defer close(stopChan)

	// Criar servidor
	router := server.NewRouter()
	srv := server.NewServer(router)

	// Middleware de auth
	auth := middlewares.NewAuthMiddleware(tokenDoer)

	// Conectar rotas bairro
	neighborhoodHandler := handlers.NewNeighborhoodHandler(
		createNeighborhoodUseCase,
		updateNeighborhoodUseCase,
		deleteNeighborhoodUseCase,
		findByIdUseCase,
		findAllUseCase,
		findActiveUseCase,
		findWeatherUseCase,
	)
	neighborhoodHandler.Register(router, auth)

	// Conectar rotas usuário
	userHandler := handlers.NewUserHandler(
		createUserUseCase,
		loginUserUseCase,
	)
	userHandler.Register(router, auth)

	// Conectar rotas swagger
	swaggerHandler := handlers.NewSwaggerHandler()
	swaggerHandler.Register(router)

	// Rodar
	ln, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		panic(err)
	}

	// Coloca o servidor em uma goroutine separada
	go srv.Serve(ln)
}
