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

type Suppliers interface {
	Create(ctx context.Context, supplier domain.Supplier) (int64, error)
	GetById(ctx context.Context, id int64) (domain.Supplier, error)
}

type SuppliersRepository struct {
	cfg   *config.Config
	cache memcache.Suppliers
	psql  postgres.Suppliers
}

func NewSuppliersRepository(cfg *config.Config, cache *cache.Client, db *sql.DB) *SuppliersRepository {
	return &SuppliersRepository{
		cfg:   cfg,
		cache: memcache.NewSuppliersMemcacheRepository(cfg, cache),
		psql:  postgres.NewSuppliersPostgresRepository(db),
	}
}

func (sr *SuppliersRepository) Create(ctx context.Context, supplier domain.Supplier) (int64, error) {
	return sr.psql.Create(ctx, supplier)
}

func (sr *SuppliersRepository) GetById(ctx context.Context, id int64) (domain.Supplier, error) {
	if id == 0 {
		return domain.Supplier{}, domain.ErrEmptyId
	}

	supplier, err := sr.cache.GetById(id)
	if err == nil {
		return supplier, err
	}

	if !errors.Is(err, cache.ErrCacheMiss) {
		return domain.Supplier{}, err
	}

	supplier, err = sr.psql.GetById(ctx, id)
	if err == nil {
		if err = sr.cache.AddById(supplier); err != nil {
			return domain.Supplier{}, nil
		}
	}

	return supplier, nil
}
