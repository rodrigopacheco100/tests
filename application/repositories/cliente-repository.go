package repositories

import (
	"context"
	"tests/domain/entities"
	"tests/domain/shared"
)

type IClienteRepository interface {
	Criar(ctx context.Context, cliente *entities.Cliente) error
	Listar(ctx context.Context, pagination *shared.PaginationRequest) (*shared.PaginationResult[entities.Cliente], error)
}
