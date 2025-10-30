package handlers

import (
	"codelab/backend/internal/domain"
	"codelab/backend/internal/middlewares"
	"codelab/backend/internal/usecases/user"
	"codelab/backend/pkg/exception"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	CreateUseCase *user.CreateUserUseCase
	LoginUseCase *user.LoginUserUseCase
}

func NewUserHandler(create *user.CreateUserUseCase, login *user.LoginUserUseCase) *UserHandler {
	return &UserHandler{create, login}
}

func (h *UserHandler) Register(router *gin.Engine, auth *middlewares.AuthMiddleware) {
	group := router.Group("/api/users")
	group.Use(auth.Private())
	{
		group.POST("/new", h.Create)
		group.POST("/login", h.Login)
	}
}

// Create cria um novo usuário
// @Summary Cria um usuário
// @Description Cria um novo usuário com os dados fornecidos
// @Tags User
// @Accept json
// @Produce json
// @Param payload body domain.CreateUserDTO true "Dados do usuário a ser criado"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /api/users/new [post]
func (h *UserHandler) Create(c *gin.Context) {
	var payload domain.CreateUserDTO
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, exception.New(err.Error()))
		return
	}

	if err := h.CreateUseCase.Execute(&payload); err != nil {
		c.JSON(http.StatusInternalServerError, exception.New(err.Error()))
		return
	}

	c.Status(http.StatusNoContent)
}

// Login faz login de um usuário
// @Summary Login do usuário
// @Description Faz login do usuário
// @Tags User
// @Accept json
// @Produce json
// @Param payload body domain.LoginUserDTO true "Dados do usuário a ser criado"
// @Success 200 "Ok"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /api/users/new [post]
func (h *UserHandler) Login(c *gin.Context) {
	var payload domain.LoginUserDTO
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, exception.New(err.Error()))
		return
	}

	token, err := h.LoginUseCase.Execute(payload.Email, payload.Password); 
	if err != nil {
		c.JSON(http.StatusInternalServerError, exception.New(err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}