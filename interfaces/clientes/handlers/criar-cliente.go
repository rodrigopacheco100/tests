package clientes_handlers

import (
	"log"
	"net/http"
	useCases "tests/application/use-cases"
	infrastructure_repositories "tests/infrastructure/repositories"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// @name createClienteDTO
type createClienteDTO struct {
	Nome  string `json:"nome" validate:"required" example:"João Silva"`
	Email string `json:"email" validate:"required,email" example:"joao@email.com"`
	Tipo  string `json:"tipo" validate:"required,oneof=fisico juridico" example:"fisico"`
}

type createClienteResponse struct {
	Message string `json:"message"`
	Id      string `json:"id"`
}

// CriarClienteHandler cria um novo cliente
// @Summary      Criar um novo cliente
// @Description  Cria um novo cliente no sistema validando nome e email
// @Tags         clientes
// @Accept       json
// @Produce      json
// @Param        request body createClienteDTO true "Dados do cliente"
// @Success      200  {object}  createClienteResponse
// @Failure      400  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /clientes/ [post]
func CriarClienteHandler(ctx *gin.Context) {
	var clienteDTO createClienteDTO
	if err := ctx.ShouldBindJSON(&clienteDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(&clienteDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	clienteRepository := infrastructure_repositories.NewClienteRepository()
	clienteUseCase := useCases.NewCriarClienteUseCase(clienteRepository)

	input := useCases.CriarClienteUseCaseInput{
		Nome:  clienteDTO.Nome,
		Email: clienteDTO.Email,
		Tipo:  clienteDTO.Tipo,
	}

	output, err := clienteUseCase.Handle(ctx.Request.Context(), &input)
	if err != nil {
		log.Printf("[ERROR] 500 - POST /clientes/: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Cliente criado com sucesso!",
		"id":      output.Cliente.Id,
	})
}
