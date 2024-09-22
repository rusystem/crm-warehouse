package service

import (
	"context"
	"github.com/rusystem/crm-warehouse/internal/repository"
	"github.com/rusystem/crm-warehouse/pkg/domain"
)

type Category interface {
	Create(ctx context.Context, category domain.MaterialCategory) (int64, error)
	GetById(ctx context.Context, id, companyId int64) (domain.MaterialCategory, error)
	Update(ctx context.Context, category domain.MaterialCategory) error
	Delete(ctx context.Context, id, companyId int64) error
	List(ctx context.Context, param domain.Param) ([]domain.MaterialCategory, error)
	Search(ctx context.Context, param domain.Param) ([]domain.MaterialCategory, error)
}

type MaterialCategoryService struct {
	repo *repository.Repository
}

func NewMaterialCategoryService(repo *repository.Repository) *MaterialCategoryService {
	return &MaterialCategoryService{
		repo: repo,
	}
}

func (mc *MaterialCategoryService) Create(ctx context.Context, category domain.MaterialCategory) (int64, error) {
	return mc.repo.Category.Create(ctx, category)
}

func (mc *MaterialCategoryService) GetById(ctx context.Context, id, companyId int64) (domain.MaterialCategory, error) {
	return mc.repo.Category.GetById(ctx, id, companyId)
}

func (mc *MaterialCategoryService) Update(ctx context.Context, category domain.MaterialCategory) error {
	return mc.repo.Category.Update(ctx, category)
}

func (mc *MaterialCategoryService) Delete(ctx context.Context, id, companyId int64) error {
	return mc.repo.Category.Delete(ctx, id, companyId)
}

func (mc *MaterialCategoryService) List(ctx context.Context, param domain.Param) ([]domain.MaterialCategory, error) {
	return mc.repo.Category.List(ctx, param)
}

func (mc *MaterialCategoryService) Search(ctx context.Context, param domain.Param) ([]domain.MaterialCategory, error) {
	return mc.repo.Category.Search(ctx, param)
}
