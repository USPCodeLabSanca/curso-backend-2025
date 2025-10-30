package middlewares

import (
	"codelab/backend/pkg/exception"
	"codelab/backend/pkg/token"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	service token.TokenDoer
}

func NewAuthMiddleware(service token.TokenDoer) *AuthMiddleware {
	return &AuthMiddleware{service}
}

/*
Middleware que indica uma rota pública.
Aqui, nenhum acesso é necessário.
*/
func (mid *AuthMiddleware) Public() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

/*
Middleware que indica uma rota privada.
Aqui, um token de acesso válido é necessário.
*/
func (mid *AuthMiddleware) Private() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := mid.verify(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, exception.New(err.Error()))
			return
		}

		c.Set("user", data["user"])
		c.Next()
	}
}

/*
Função que valida o token recebido na requisição.

O token está no cabeçalho, no formato "Bearer meu_token_aqui".
Então, é necessário fazer a leitura correta dos dados,
e, caso o token seja válido, retorná-lo como um map de atributos.
*/
func (mid *AuthMiddleware) verify(c *gin.Context) (map[string]any, error) {
	header := c.GetHeader("Authorization")
	if header == "" || !strings.HasPrefix(header, "Bearer ") {
		return nil, errors.New("token not found")
	}

	token := strings.TrimPrefix(header, "Bearer ")

	valid, data, err := mid.service.Decrypt(token)
	if err != nil || !valid {
		return nil, errors.New("invalid token")
	}

	return data, nil
}