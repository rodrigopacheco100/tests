package useCases

import (
	"context"
	"tests/application/repositories"
	"tests/domain/entities"
	"tests/domain/shared"
)

type ListarClientesUseCase struct {
	repository repositories.IClienteRepository
}

func NewListarClientesUseCase(repository repositories.IClienteRepository) *ListarClientesUseCase {
	return &ListarClientesUseCase{
		repository: repository,
	}
}

type ListarClientesUseCaseInput struct {
	Pagination *shared.PaginationRequest
}

type ListarClientesUseCaseOutput struct {
	Clientes *shared.PaginationResult[entities.Cliente]
}

func (c *ListarClientesUseCase) Handle(ctx context.Context, input *ListarClientesUseCaseInput) (ListarClientesUseCaseOutput, error) {
	clientes, err := c.repository.Listar(ctx, input.Pagination)
	if err != nil {
		return ListarClientesUseCaseOutput{}, err
	}

	return ListarClientesUseCaseOutput{
		Clientes: clientes,
	}, nil
}
