package server

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

/*
Função que cria um roteador do Gin.

Aqui, são especificadas configurações
para permitir acesso de diferentes cliente: CORS.
*/
func NewRouter() *gin.Engine {
	// Define qual roteador será usado (padrão)
	router := gin.Default()

	// Configurações do CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	return router
}

/*
Função que cria um novo servidor HTTP.

O servidor é criado com base no roteador Gin.

Usando o Fx, bilioteca que permite injeção de
dependencias, é fácil controlar start e stop do servidor.
*/
func NewServer(router *gin.Engine) *http.Server {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	return srv
}