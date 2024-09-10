package postgres

import (
	"context"
	"database/sql"
	"github.com/rusystem/crm-warehouse/pkg/domain"
)

type Materials interface {
	CreatePlanning(ctx context.Context, material domain.Material) (int64, error)
	UpdatePlanning(ctx context.Context)
}

type MaterialsPostgresRepository struct {
	psql *sql.DB
}

func NewMaterialsPostgresRepository(psql *sql.DB) *MaterialsPostgresRepository {
	return &MaterialsPostgresRepository{
		psql: psql,
	}
}
