package clientes_handlers

import (
	"log"
	"net/http"
	"strconv"
	useCases "tests/application/use-cases"
	"tests/domain/shared"
	infrastructure_repositories "tests/infrastructure/repositories"
	"time"

	"github.com/gin-gonic/gin"
)

// @name clienteDTO
type clienteDTO struct {
	Id        string     `json:"id"`
	Nome      string     `json:"nome"`
	Email     string     `json:"email"`
	Tipo      string     `json:"tipo"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

// @name listarClientesResponse
type listarClientesResponse struct {
	Items      []clienteDTO `json:"items"`
	TotalPages int          `json:"totalPages"`
	Total      int          `json:"total"`
}

// ListarClientesHandler lista todos os clientes paginados
// @Summary      Listar clientes
// @Description  Retorna uma lista paginada de clientes cadastrados
// @Tags         clientes
// @Accept       json
// @Produce      json
// @Param        page query int false "Página atual"
// @Param        limit query int false "Limite de registros por página"
// @Success      200  {object}  listarClientesResponse
// @Failure      500  {object}  map[string]interface{}
// @Router       /clientes [get]
func ListarClientesHandler(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 || limit > 100 {
		limit = 10
	}

	clienteRepository := infrastructure_repositories.NewClienteRepository()
	clienteUseCase := useCases.NewListarClientesUseCase(clienteRepository)

	input := useCases.ListarClientesUseCaseInput{
		Pagination: &shared.PaginationRequest{
			Page:  page,
			Limit: limit,
		},
	}

	output, err := clienteUseCase.Handle(ctx.Request.Context(), &input)
	if err != nil {
		log.Printf("[ERROR] 500 - GET /clientes: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	clientes := make([]clienteDTO, len(output.Clientes.Items))
	for i, cliente := range output.Clientes.Items {
		clientes[i] = clienteDTO{
			Id:        cliente.Id,
			Nome:      cliente.Nome,
			Email:     cliente.Email,
			Tipo:      cliente.Tipo,
			CreatedAt: cliente.CreatedAt,
			UpdatedAt: cliente.UpdatedAt,
			DeletedAt: cliente.DeletedAt,
		}
	}

	ctx.JSON(http.StatusOK, listarClientesResponse{
		Items:      clientes,
		TotalPages: output.Clientes.TotalPages,
		Total:      output.Clientes.Total,
	})
}
