package handlers

import (
	"codelab/backend/internal/domain"
	"codelab/backend/internal/usecases/neighborhood"
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
}

func NewNeighborhoodHandler(
	create *neighborhood.CreateNeighborhoodUseCase,
	update *neighborhood.UpdateNeighborhoodUseCase,
	delete *neighborhood.DeleteNeighborhoodUseCase,
	findById *neighborhood.FindNeighborhoodByIdUseCase,
	findAll *neighborhood.FindAllNeighborhoodsUseCase,
	findActive *neighborhood.FindActiveNeighborhoodsUseCase,
) *NeighborhoodHandler {
	return &NeighborhoodHandler{
		CreateUseCase:     create,
		UpdateUseCase:     update,
		DeleteUseCase:     delete,
		FindByIdUseCase:   findById,
		FindAllUseCase:    findAll,
		FindActiveUseCase: findActive,
	}
}

func (h *NeighborhoodHandler) Register(router *gin.Engine) {
	group := router.Group("/api/neighborhoods")
	{
		group.POST("/new", h.Create)
		group.GET("/", h.FindAll)
		group.GET("/active", h.FindActive)
		group.GET("/:id", h.FindById)
		group.PUT("/", h.Update)
		group.DELETE("/:id", h.Delete)
	}
}

func (h *NeighborhoodHandler) Create(c *gin.Context) {
	var payload domain.CreateNeighborhoodDTO
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.CreateUseCase.Execute(&payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *NeighborhoodHandler) Update(c *gin.Context) {
	var neighborhood domain.Neighborhood
	if err := c.ShouldBindJSON(&neighborhood); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.UpdateUseCase.Execute(&neighborhood); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *NeighborhoodHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	if err := h.DeleteUseCase.Execute(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *NeighborhoodHandler) FindById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	neighborhood, err := h.FindByIdUseCase.Execute(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if neighborhood == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, neighborhood)
}

func (h *NeighborhoodHandler) FindAll(c *gin.Context) {
	neighborhoods, err := h.FindAllUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, neighborhoods)
}

func (h *NeighborhoodHandler) FindActive(c *gin.Context) {
	neighborhoods, err := h.FindActiveUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, neighborhoods)
}