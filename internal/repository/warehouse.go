package repository

import (
	"context"
	"database/sql"
	"errors"
	cache "github.com/bradfitz/gomemcache/memcache"
	"github.com/rusystem/crm-warehouse/internal/config"
	memcache "github.com/rusystem/crm-warehouse/internal/repository/memcache"
	"github.com/rusystem/crm-warehouse/internal/repository/postgres"
	"github.com/rusystem/crm-warehouse/pkg/domain"
)

type Warehouse interface {
	Create(ctx context.Context, warehouse domain.Warehouse) (int64, error)
	GetById(ctx context.Context, id int64) (domain.Warehouse, error)
}

type WarehouseRepository struct {
	cfg   *config.Config
	cache memcache.Warehouse
	psql  postgres.Warehouse
}

func NewWarehouseRepository(cfg *config.Config, cache *cache.Client, psql *sql.DB) *WarehouseRepository {
	return &WarehouseRepository{
		cfg:   cfg,
		cache: memcache.NewWarehouseMemcacheRepository(cfg, cache),
		psql:  postgres.NewWarehousePostgresRepository(psql),
	}
}

func (wr *WarehouseRepository) Create(ctx context.Context, warehouse domain.Warehouse) (int64, error) {
	return wr.psql.Create(ctx, warehouse)
}

func (wr *WarehouseRepository) GetById(ctx context.Context, id int64) (domain.Warehouse, error) {
	if id == 0 {
		return domain.Warehouse{}, domain.ErrEmptyId
	}

	warehouse, err := wr.cache.GetById(id)
	if err == nil {
		return warehouse, nil
	}

	if !errors.Is(err, cache.ErrCacheMiss) {
		return domain.Warehouse{}, err
	}

	warehouse, err = wr.psql.GetById(ctx, id)
	if err == nil {
		if err = wr.cache.AddById(warehouse); err != nil {
			return domain.Warehouse{}, err
		}
	}

	return warehouse, err
}
