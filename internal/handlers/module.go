package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Module("handlers",
	fx.Provide(NewNeighborhoodHandler),
	fx.Invoke(func (handler *NeighborhoodHandler, r *gin.Engine,) {
		handler.Register(r)
	}),
)