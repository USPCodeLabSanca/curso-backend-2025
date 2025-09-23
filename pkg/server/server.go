package server

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

/*
Função que cria um roteador do Gin.

Aqui, são especificadas configurações
para permitir acesso de diferentes cliente: CORS.
*/
func NewRouter() *gin.Engine {
	// Default router
	router := gin.Default()

	// CORS config
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST"},
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
func NewServer(lc fx.Lifecycle, router *gin.Engine) *http.Server {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Add server to lifecycle
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}

			// Start server using goroutine
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return srv
}