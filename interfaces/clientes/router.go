package clientes_router

import (
	clientes_handlers "tests/interfaces/clientes/handlers"

	"github.com/gin-gonic/gin"
)

func NewClienteRouter(router *gin.RouterGroup) {
	router.POST("/", clientes_handlers.CriarClienteHandler)
}
