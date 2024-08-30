package repository

import (
	"database/sql"
	"github.com/rusystem/crm-warehouse/internal/config"
)

type Repository struct {
	Suppliers *SuppliersRepository
	Warehouse *WarehouseRepository
}

func New(cfg *config.Config, postgres *sql.DB) *Repository {
	return &Repository{
		Suppliers: NewSuppliersRepository(cfg, postgres),
		Warehouse: NewWarehouseRepository(cfg, postgres),
	}
}
