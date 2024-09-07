package service

import (
	"context"
	"github.com/rusystem/crm-warehouse/internal/repository"
	"github.com/rusystem/crm-warehouse/pkg/domain"
)

type Supplier interface {
	Create(ctx context.Context, supplier domain.Supplier) (int64, error)
	GetById(ctx context.Context, id int64) (domain.Supplier, error)
	Update(ctx context.Context, supplier domain.Supplier) error
	Delete(ctx context.Context, id int64) error
	GetListByCompanyId(ctx context.Context, id int64) ([]domain.Supplier, error)
}

type SupplierService struct {
	repo *repository.Repository
}

func NewSupplierService(repo *repository.Repository) *SupplierService {
	return &SupplierService{
		repo: repo,
	}
}

func (ss *SupplierService) Create(ctx context.Context, supplier domain.Supplier) (int64, error) {
	return ss.repo.Suppliers.Create(ctx, supplier)
}

func (ss *SupplierService) GetById(ctx context.Context, id int64) (domain.Supplier, error) {
	return ss.repo.Suppliers.GetById(ctx, id)
}

func (ss *SupplierService) Update(ctx context.Context, supplier domain.Supplier) error {
	return ss.repo.Suppliers.Update(ctx, supplier)
}

func (ss *SupplierService) Delete(ctx context.Context, id int64) error {
	return ss.repo.Suppliers.Delete(ctx, id)
}

func (ss *SupplierService) GetListByCompanyId(ctx context.Context, id int64) ([]domain.Supplier, error) {
	return ss.repo.Suppliers.GetListByCompanyId(ctx, id)
}
