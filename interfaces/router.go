package interfaces

import (
	clientes_router "tests/interfaces/clientes"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *gin.Engine) {
	clientes_router.NewClienteRouter(app.Group("/clientes"))
}
