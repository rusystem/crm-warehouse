package service

import (
	"context"
	"github.com/rusystem/crm-warehouse/internal/repository"
	"github.com/rusystem/crm-warehouse/pkg/domain"
)

type Warehouse interface {
	Create(ctx context.Context, warehouse domain.Warehouse) (int64, error)
	GetById(ctx context.Context, id int64) (domain.Warehouse, error)
}

type WarehouseService struct {
	repo *repository.Repository
}

func NewWarehouseService(repo *repository.Repository) *WarehouseService {
	return &WarehouseService{
		repo: repo,
	}
}

func (ws *WarehouseService) Create(ctx context.Context, warehouse domain.Warehouse) (int64, error) {
	return ws.repo.Warehouse.Create(ctx, warehouse)
}

func (ws *WarehouseService) GetById(ctx context.Context, id int64) (domain.Warehouse, error) {
	return ws.repo.Warehouse.GetById(ctx, id)
}
