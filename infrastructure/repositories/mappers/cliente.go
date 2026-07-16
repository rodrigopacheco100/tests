package repository_mappers

import (
	"tests/domain/entities"
	infra_entities "tests/infrastructure/entities"
)

func MapClienteToDomain(cliente *infra_entities.Cliente) *entities.Cliente {
	return &entities.Cliente{
		Id:        cliente.Id,
		Nome:      cliente.Nome,
		Email:     cliente.Email,
		Tipo:      cliente.Tipo,
		CreatedAt: cliente.CreatedAt,
		UpdatedAt: cliente.UpdatedAt,
		DeletedAt: cliente.DeletedAt,
	}
}
