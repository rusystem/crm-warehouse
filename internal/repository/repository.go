package repository

import (
	"database/sql"
	"github.com/rusystem/crm-warehouse/internal/config"
)

type Repository struct {
	Suppliers *SuppliersRepository
	Warehouse *WarehouseRepository
	Materials *MaterialsRepository
	Category  *MaterialCategoriesRepository
}

func New(cfg *config.Config, postgres *sql.DB) *Repository {
	return &Repository{
		Suppliers: NewSuppliersRepository(cfg, postgres),
		Warehouse: NewWarehouseRepository(cfg, postgres),
		Materials: NewMaterialsRepository(cfg, postgres),
		Category:  NewMaterialCategoriesRepository(cfg, postgres),
	}
}
