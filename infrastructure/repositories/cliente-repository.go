package infrastructure_repositories

import (
	"context"
	"tests/domain/entities"
	"tests/infrastructure/database"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx"
)

type ClienteRepository struct {
	db *pgx.Conn
}

func NewClienteRepository() *ClienteRepository {
	db, err := database.NewConnection()
	if err != nil {
		panic(err)
	}

	return &ClienteRepository{
		db: db,
	}
}

func (r *ClienteRepository) Criar(ctx context.Context, cliente *entities.Cliente) error {
	sqlBuilder := sq.Insert("clientes").Columns("nome", "email", "tipo").Values(cliente.Nome, cliente.Email, cliente.Tipo).PlaceholderFormat(sq.Dollar)
	query, args, _ := sqlBuilder.ToSql()

	tx, err := r.db.Begin()
	defer tx.Rollback()

	_, err = tx.Exec(query, args...)
	if err != nil {
		return err
	}

	tx.Commit()
	return nil
}
