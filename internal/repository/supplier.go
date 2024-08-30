package repository

import (
	"context"
	"database/sql"
	"github.com/rusystem/crm-warehouse/internal/config"
	"github.com/rusystem/crm-warehouse/internal/repository/postgres"
	"github.com/rusystem/crm-warehouse/pkg/domain"
)

type Suppliers interface {
	Create(ctx context.Context, supplier domain.Supplier) (int64, error)
	GetById(ctx context.Context, id int64) (domain.Supplier, error)
}

type SuppliersRepository struct {
	cfg  *config.Config
	psql postgres.Suppliers
}

func NewSuppliersRepository(cfg *config.Config, db *sql.DB) *SuppliersRepository {
	return &SuppliersRepository{
		cfg:  cfg,
		psql: postgres.NewSuppliersPostgresRepository(db),
	}
}

func (sr *SuppliersRepository) Create(ctx context.Context, supplier domain.Supplier) (int64, error) {
	return sr.psql.Create(ctx, supplier)
}

func (sr *SuppliersRepository) GetById(ctx context.Context, id int64) (domain.Supplier, error) {
	return sr.psql.GetById(ctx, id)
}
