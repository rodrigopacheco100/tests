package infrastructure_repositories

import (
	"context"
	"math"
	"tests/domain/entities"
	"tests/domain/shared"
	"tests/infrastructure/database"
	infra_entities "tests/infrastructure/entities"
	repository_mappers "tests/infrastructure/repositories/mappers"

	sq "github.com/Masterminds/squirrel"
	pgx "github.com/jackc/pgx/v5"
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

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (r *ClienteRepository) Listar(ctx context.Context, pagination *shared.PaginationRequest) (*shared.PaginationResult[entities.Cliente], error) {
	clientesQuery, clientesArgs, _ := sq.Select("*").From("clientes").Where(sq.Eq(map[string]any{"deleted_at": nil})).PlaceholderFormat(sq.Dollar).Offset(uint64(pagination.Offset())).Limit(uint64(pagination.Limit)).ToSql()
	totalQuery, totalArgs, _ := sq.Select("COUNT(*)").From("clientes").Where(sq.Eq(map[string]any{"deleted_at": nil})).PlaceholderFormat(sq.Dollar).ToSql()

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	totalRows, err := tx.Query(ctx, totalQuery, totalArgs...)
	if err != nil {
		return nil, err
	}

	var total int
	if totalRows.Next() {
		totalRows.Scan(&total)
	}
	totalRows.Close()

	clientesRows, err := tx.Query(ctx, clientesQuery, clientesArgs...)
	if err != nil {
		return nil, err
	}
	defer clientesRows.Close()

	clientesDb, err := pgx.CollectRows(clientesRows, pgx.RowToStructByName[infra_entities.Cliente])
	if err != nil {
		return nil, err
	}

	clientes := make([]entities.Cliente, len(clientesDb))
	for i, cliente := range clientesDb {
		clientes[i] = *repository_mappers.MapClienteToDomain(&cliente)
	}

	tx.Commit(ctx)
	return &shared.PaginationResult[entities.Cliente]{
		Items:      clientes,
		TotalPages: int(math.Ceil(float64(total) / float64(pagination.Limit))),
		Total:      total,
	}, nil
}
