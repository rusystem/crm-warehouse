package service

import (
	"context"
	"github.com/rusystem/crm-warehouse/internal/repository"
	"github.com/rusystem/crm-warehouse/pkg/domain"
)

type Material interface {
	CreatePlanning(ctx context.Context, material domain.Material) (int64, error)
	UpdatePlanning(ctx context.Context, material domain.Material) error
	DeletePlanning(ctx context.Context, id int64) error
	GetPlanningById(ctx context.Context, id int64) (domain.Material, error)
	GetPlanningList(ctx context.Context, params domain.MaterialParams) ([]domain.Material, error)
	MovePlanningToPurchased(ctx context.Context, id int64) (int64, int64, error)

	CreatePurchased(ctx context.Context, material domain.Material) (int64, int64, error)
	UpdatePurchased(ctx context.Context, material domain.Material) error
	DeletePurchased(ctx context.Context, id int64) error
	GetPurchasedById(ctx context.Context, id int64) (domain.Material, error)
	GetPurchasedList(ctx context.Context, params domain.MaterialParams) ([]domain.Material, error)
	MovePurchasedToArchive(ctx context.Context, id int64) error

	GetPlanningArchiveById(ctx context.Context, id int64) (domain.Material, error)
	GetPurchasedArchiveById(ctx context.Context, id int64) (domain.Material, error)
	GetPlanningArchiveList(ctx context.Context, params domain.MaterialParams) ([]domain.Material, error)
	GetPurchasedArchiveList(ctx context.Context, params domain.MaterialParams) ([]domain.Material, error)
	DeletePlanningArchive(ctx context.Context, id int64) error
	DeletePurchasedArchive(ctx context.Context, id int64) error
}

type MaterialService struct {
	repo *repository.Repository
}

func NewMaterialService(repo *repository.Repository) *MaterialService {
	return &MaterialService{
		repo: repo,
	}
}

func (ms *MaterialService) CreatePlanning(ctx context.Context, material domain.Material) (int64, error) {
	return ms.repo.Materials.CreatePlanning(ctx, material)
}

func (ms *MaterialService) UpdatePlanning(ctx context.Context, material domain.Material) error {
	return ms.repo.Materials.UpdatePlanning(ctx, material)
}

func (ms *MaterialService) DeletePlanning(ctx context.Context, id int64) error {
	return ms.repo.Materials.DeletePlanning(ctx, id)
}

func (ms *MaterialService) GetPlanningById(ctx context.Context, id int64) (domain.Material, error) {
	return ms.repo.Materials.GetPlanningById(ctx, id)
}

func (ms *MaterialService) GetPlanningList(ctx context.Context, params domain.MaterialParams) ([]domain.Material, error) {
	return ms.repo.Materials.GetPlanningList(ctx, params)
}

func (ms *MaterialService) MovePlanningToPurchased(ctx context.Context, id int64) (int64, int64, error) {
	return ms.repo.Materials.MovePlanningToPurchased(ctx, id)
}

func (ms *MaterialService) CreatePurchased(ctx context.Context, material domain.Material) (int64, int64, error) {
	return ms.repo.Materials.CreatePurchased(ctx, material)
}

func (ms *MaterialService) UpdatePurchased(ctx context.Context, material domain.Material) error {
	return ms.repo.Materials.UpdatePurchased(ctx, material)
}

func (ms *MaterialService) DeletePurchased(ctx context.Context, id int64) error {
	return ms.repo.Materials.DeletePurchased(ctx, id)
}

func (ms *MaterialService) GetPurchasedById(ctx context.Context, id int64) (domain.Material, error) {
	return ms.repo.Materials.GetPurchasedById(ctx, id)
}

func (ms *MaterialService) GetPurchasedList(ctx context.Context, params domain.MaterialParams) ([]domain.Material, error) {
	return ms.repo.Materials.GetPurchasedList(ctx, params)
}

func (ms *MaterialService) MovePurchasedToArchive(ctx context.Context, id int64) error {
	return ms.repo.Materials.MovePurchasedToArchive(ctx, id)
}

func (ms *MaterialService) GetPlanningArchiveById(ctx context.Context, id int64) (domain.Material, error) {
	return ms.repo.Materials.GetPlanningArchiveById(ctx, id)
}

func (ms *MaterialService) GetPurchasedArchiveById(ctx context.Context, id int64) (domain.Material, error) {
	return ms.repo.Materials.GetPurchasedArchiveById(ctx, id)
}

func (ms *MaterialService) GetPlanningArchiveList(ctx context.Context, params domain.MaterialParams) ([]domain.Material, error) {
	return ms.repo.Materials.GetPlanningArchiveList(ctx, params)
}

func (ms *MaterialService) GetPurchasedArchiveList(ctx context.Context, params domain.MaterialParams) ([]domain.Material, error) {
	return ms.repo.Materials.GetPurchasedArchiveList(ctx, params)
}

func (ms *MaterialService) DeletePlanningArchive(ctx context.Context, id int64) error {
	return ms.repo.Materials.DeletePlanningArchive(ctx, id)
}

func (ms *MaterialService) DeletePurchasedArchive(ctx context.Context, id int64) error {
	return ms.repo.Materials.DeletePurchasedArchive(ctx, id)
}
