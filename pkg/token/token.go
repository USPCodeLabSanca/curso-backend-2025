package token

import (
	"codelab/backend/pkg/config"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

/*
Structs para configuração da autenticação.

O token doer é uma interface que pode se adaptar a
diferentes tipos de tokens. Aqui será usado o JWT.

A struct JWTClaims indica quais dados serão
armazenados no token.
*/
type TokenDoer interface {
	Encrypt(data map[string]any, expiration time.Time) (string, error)
	Decrypt(token string) (bool, map[string]any, error)
}

type JWTClaims struct {
	Data map[string]any       `json:"data"`
	jwt.RegisteredClaims
}

type doer struct {
	key []byte
}

func NewJWTDoer(conf *config.Config) TokenDoer {
	return &doer{
		key: []byte(conf.JWT),
	}
}

/*
Função que recebe um pacote de dados e um tempo
de expiração e retorna um token JWT válido.
*/
func (d doer) Encrypt(data map[string]any, expiration time.Time) (string, error) {
	claims := &JWTClaims{
		Data: data,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(d.key)
}

/*
Função que recebe um token JWT e tenta validá-lo.
Em caso de sucesso, os dados contidos no token são retornados.
*/
func (d doer) Decrypt(token string) (bool, map[string]any, error) {
	claims := &JWTClaims{}

	result, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (any, error) {
		return d.key, nil
	})

	if result == nil || !result.Valid {
        return false, nil, errors.New("invalid token")
    }

	if err != nil {
		return false, nil, err
	}

	recovered, ok := result.Claims.(*JWTClaims)
	if !ok {
		return false, nil, errors.New("failed to parse token")
	}

	return true, recovered.Data, nil
}
