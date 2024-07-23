package repository

import (
	"database/sql"
	cache "github.com/bradfitz/gomemcache/memcache"
	"github.com/rusystem/crm-warehouse/internal/config"
)

type Repository struct {
	Suppliers *SuppliersRepository
	Warehouse *WarehouseRepository
}

func New(cfg *config.Config, cache *cache.Client, postgres *sql.DB) *Repository {
	return &Repository{
		Suppliers: NewSuppliersRepository(cfg, cache, postgres),
		Warehouse: NewWarehouseRepository(cfg, cache, postgres),
	}
}
