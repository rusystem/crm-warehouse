package repository

import (
	"context"
	"database/sql"
	"github.com/rusystem/crm-warehouse/internal/config"
	"github.com/rusystem/crm-warehouse/internal/repository/postgres"
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

type MaterialCategoriesRepository struct {
	cfg  *config.Config
	psql postgres.Category
}

func NewMaterialCategoriesRepository(cfg *config.Config, db *sql.DB) *MaterialCategoriesRepository {
	return &MaterialCategoriesRepository{
		cfg:  cfg,
		psql: postgres.NewMaterialCategoriesPostgresRepository(db),
	}
}

func (mc *MaterialCategoriesRepository) Create(ctx context.Context, category domain.MaterialCategory) (int64, error) {
	return mc.psql.Create(ctx, category)
}

func (mc *MaterialCategoriesRepository) GetById(ctx context.Context, id, companyId int64) (domain.MaterialCategory, error) {
	return mc.psql.GetById(ctx, id, companyId)
}

func (mc *MaterialCategoriesRepository) Update(ctx context.Context, category domain.MaterialCategory) error {
	return mc.psql.Update(ctx, category)
}

func (mc *MaterialCategoriesRepository) Delete(ctx context.Context, id, companyId int64) error {
	return mc.psql.Delete(ctx, id, companyId)
}

func (mc *MaterialCategoriesRepository) List(ctx context.Context, param domain.Param) ([]domain.MaterialCategory, error) {
	return mc.psql.List(ctx, param)
}

func (mc *MaterialCategoriesRepository) Search(ctx context.Context, param domain.Param) ([]domain.MaterialCategory, error) {
	return mc.psql.Search(ctx, param)
}
