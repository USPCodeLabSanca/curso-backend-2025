package handlers

import (
	"codelab/backend/internal/domain"
	"codelab/backend/internal/middlewares"
	"codelab/backend/internal/usecases/neighborhood"
	"codelab/backend/pkg/exception"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
Handlers: rotas HTTP referentes aos bairros.

Cada rota está associada à um caso de uso de bairro.
As requisições podem ou não ter conteúdo (corpo / parâmetros).
As respostas podem ou não ter conteúdo.

Códigos HTTP e métodos adequados são usados para bom
funcionamento e padronização.
*/

type NeighborhoodHandler struct {
	CreateUseCase      *neighborhood.CreateNeighborhoodUseCase
	UpdateUseCase      *neighborhood.UpdateNeighborhoodUseCase
	DeleteUseCase      *neighborhood.DeleteNeighborhoodUseCase
	FindByIdUseCase    *neighborhood.FindNeighborhoodByIdUseCase
	FindAllUseCase     *neighborhood.FindAllNeighborhoodsUseCase
	FindActiveUseCase  *neighborhood.FindActiveNeighborhoodsUseCase
	FindWeatherUseCase *neighborhood.FindNeighborhoodWeatherUseCase
}

func NewNeighborhoodHandler(
	create *neighborhood.CreateNeighborhoodUseCase,
	update *neighborhood.UpdateNeighborhoodUseCase,
	delete *neighborhood.DeleteNeighborhoodUseCase,
	findById *neighborhood.FindNeighborhoodByIdUseCase,
	findAll *neighborhood.FindAllNeighborhoodsUseCase,
	findActive *neighborhood.FindActiveNeighborhoodsUseCase,
	findWeather *neighborhood.FindNeighborhoodWeatherUseCase,
) *NeighborhoodHandler {
	return &NeighborhoodHandler{
		CreateUseCase:     create,
		UpdateUseCase:     update,
		DeleteUseCase:     delete,
		FindByIdUseCase:   findById,
		FindAllUseCase:    findAll,
		FindActiveUseCase: findActive,
		FindWeatherUseCase: findWeather,
	}
}

func (h *NeighborhoodHandler) Register(router *gin.Engine, auth *middlewares.AuthMiddleware) {
	group := router.Group("/api/neighborhoods")
	group.Use(auth.Private())
	{
		group.POST("/new", h.Create)
		group.GET("/", h.FindAll)
		group.GET("/active", h.FindActive)
		group.GET("/:id", h.FindById)
		group.PUT("/", h.Update)
		group.DELETE("/:id", h.Delete)
		group.GET("/:id/weather/:start/:end", h.FindNeighborhoodWeather)
	}
}

// Create cria um novo bairro
// @Summary Cria um bairro
// @Description Cria um novo bairro com os dados fornecidos
// @Tags Neighborhood
// @Accept json
// @Produce json
// @Param payload body domain.CreateNeighborhoodDTO true "Dados do bairro a ser criado"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /api/neighborhoods/new [post]
func (h *NeighborhoodHandler) Create(c *gin.Context) {
	var payload domain.CreateNeighborhoodDTO
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

// Update atualiza um bairro existente
// @Summary Atualiza um bairro
// @Description Atualiza os dados de um bairro existente
// @Tags Neighborhood
// @Accept json
// @Produce json
// @Param neighborhood body domain.Neighborhood true "Dados do bairro"
// @Success 204 "No Content"
// @Failure 400 {object} exception.Exception
// @Failure 500 {object} exception.Exception
// @Router /api/neighborhoods [put]
func (h *NeighborhoodHandler) Update(c *gin.Context) {
	var neighborhood domain.Neighborhood
	if err := c.ShouldBindJSON(&neighborhood); err != nil {
		c.JSON(http.StatusBadRequest, exception.New(err.Error()))
		return
	}

	if err := h.UpdateUseCase.Execute(&neighborhood); err != nil {
		c.JSON(http.StatusInternalServerError, exception.New(err.Error()))
		return
	}

	c.Status(http.StatusNoContent)
}

// Delete remove um bairro pelo ID
// @Summary Deleta um bairro
// @Description Remove um bairro existente pelo ID
// @Tags Neighborhood
// @Param id path int true "ID do bairro"
// @Success 204 "No Content"
// @Failure 400 {object} exception.Exception
// @Failure 500 {object} exception.Exception
// @Router /api/neighborhoods/{id} [delete]
func (h *NeighborhoodHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, exception.New("Invalid ID"))
		return
	}

	if err := h.DeleteUseCase.Execute(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, exception.New(err.Error()))
		return
	}

	c.Status(http.StatusNoContent)
}

// FindById retorna um bairro pelo ID
// @Summary Busca bairro por ID
// @Description Retorna os dados de um bairro pelo seu ID
// @Tags Neighborhood
// @Param id path int true "ID do bairro"
// @Success 200 {object} domain.Neighborhood
// @Failure 400 {object} exception.Exception
// @Failure 404 {object} map[string]string "Not Found"
// @Failure 500 {object} exception.Exception
// @Router /api/neighborhoods/{id} [get]
func (h *NeighborhoodHandler) FindById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, exception.New("Invalid ID"))
		return
	}

	neighborhood, err := h.FindByIdUseCase.Execute(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, exception.New(err.Error()))
		return
	}
	if neighborhood == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, neighborhood)
}

// FindAll retorna todos os bairros
// @Summary Lista todos os bairros
// @Description Retorna todos os bairros cadastrados
// @Tags Neighborhood
// @Success 200 {array} domain.Neighborhood
// @Failure 500 {object} exception.Exception
// @Router /api/neighborhoods [get]
func (h *NeighborhoodHandler) FindAll(c *gin.Context) {
	neighborhoods, err := h.FindAllUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, exception.New(err.Error()))
		return
	}

	c.JSON(http.StatusOK, neighborhoods)
}

// FindActive retorna todos os bairros ativos
// @Summary Lista bairros ativos
// @Description Retorna apenas os bairros ativos
// @Tags Neighborhood
// @Success 200 {array} domain.Neighborhood
// @Failure 500 {object} exception.Exception
// @Router /api/neighborhoods/active [get]
func (h *NeighborhoodHandler) FindActive(c *gin.Context) {
	neighborhoods, err := h.FindActiveUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, exception.New(err.Error()))
		return
	}

	c.JSON(http.StatusOK, neighborhoods)
}

// FindNeighborhoodWeather retorna o clima de um bairro entre datas
// @Summary Clima de um bairro
// @Description Retorna os registros de clima de um bairro entre duas datas
// @Tags Neighborhood
// @Param id path int true "ID do bairro"
// @Param start path string true "Data inicial (YYYY-MM-DD)"
// @Param end path string true "Data final (YYYY-MM-DD)"
// @Success 200 {object} domain.Neighborhood "Bairro com dados de clima"
// @Failure 400 {object} exception.Exception
// @Failure 500 {object} exception.Exception
// @Router /api/neighborhoods/{id}/weather/{start}/{end} [get]
func (h *NeighborhoodHandler) FindNeighborhoodWeather(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, exception.New("Invalid ID"))
		return
	}

	// Pegar datas da URL
	start := c.Param("start")
	end := c.Param("end")

	neighborhood, err := h.FindWeatherUseCase.Execute(uint(id), start, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, exception.New(err.Error()))
		return
	}

	c.JSON(200, neighborhood)
}