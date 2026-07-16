package repositories

import (
	"context"
	"tests/domain/entities"
)

type IClienteRepository interface {
	Criar(ctx context.Context, cliente *entities.Cliente) error
}
