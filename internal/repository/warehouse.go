package repository

import (
	"context"
	"database/sql"
	"github.com/rusystem/crm-warehouse/internal/config"
	"github.com/rusystem/crm-warehouse/internal/repository/postgres"
	"github.com/rusystem/crm-warehouse/pkg/domain"
)

type Warehouse interface {
	Create(ctx context.Context, warehouse domain.Warehouse) (int64, error)
	GetById(ctx context.Context, id int64) (domain.Warehouse, error)
	Update(ctx context.Context, warehouse domain.Warehouse) error
	Delete(ctx context.Context, id int64) error
	GetListByCompanyId(ctx context.Context, id int64) ([]domain.Warehouse, error)
}

type WarehouseRepository struct {
	cfg  *config.Config
	psql postgres.Warehouse
}

func NewWarehouseRepository(cfg *config.Config, psql *sql.DB) *WarehouseRepository {
	return &WarehouseRepository{
		cfg:  cfg,
		psql: postgres.NewWarehousePostgresRepository(psql),
	}
}

func (wr *WarehouseRepository) Create(ctx context.Context, warehouse domain.Warehouse) (int64, error) {
	return wr.psql.Create(ctx, warehouse)
}

func (wr *WarehouseRepository) GetById(ctx context.Context, id int64) (domain.Warehouse, error) {
	return wr.psql.GetById(ctx, id)
}

func (wr *WarehouseRepository) Update(ctx context.Context, warehouse domain.Warehouse) error {
	return wr.psql.Update(ctx, warehouse)
}

func (wr *WarehouseRepository) Delete(ctx context.Context, id int64) error {
	return wr.psql.Delete(ctx, id)
}

func (wr *WarehouseRepository) GetListByCompanyId(ctx context.Context, id int64) ([]domain.Warehouse, error) {
	return wr.psql.GetListByCompanyId(ctx, id)
}
