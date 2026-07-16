package useCases

import (
	"context"
	"tests/application/repositories"
	"tests/domain/entities"
)

type CriarClienteUseCase struct {
	repository repositories.IClienteRepository
}

func NewCriarClienteUseCase(repository repositories.IClienteRepository) *CriarClienteUseCase {
	return &CriarClienteUseCase{
		repository: repository,
	}
}

type CriarClienteUseCaseInput struct {
	Nome  string
	Email string
	Tipo  entities.ClienteTipo
}

type CriarClienteUseCaseOutput struct {
	Cliente *entities.Cliente
}

func (c *CriarClienteUseCase) Handle(ctx context.Context, input *CriarClienteUseCaseInput) (CriarClienteUseCaseOutput, error) {
	cliente := entities.NewCliente(input.Nome, input.Email, input.Tipo)

	err := c.repository.Criar(ctx, cliente)
	if err != nil {
		return CriarClienteUseCaseOutput{}, err
	}

	return CriarClienteUseCaseOutput{
		Cliente: cliente,
	}, nil
}
