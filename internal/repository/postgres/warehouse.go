package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rusystem/crm-warehouse/pkg/domain"
)

type Warehouse interface {
	Create(ctx context.Context, warehouse domain.Warehouse) (int64, error)
	GetById(ctx context.Context, id int64) (domain.Warehouse, error)
}

type WarehousePostgresRepository struct {
	db *sql.DB
}

func NewWarehousePostgresRepository(db *sql.DB) *WarehousePostgresRepository {
	return &WarehousePostgresRepository{db: db}
}

func (wpr *WarehousePostgresRepository) Create(ctx context.Context, warehouse domain.Warehouse) (int64, error) {
	otherFieldsJSON, err := json.Marshal(warehouse.OtherFields)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal other_fields to JSON: %v", err)
	}

	query := fmt.Sprintf(`
    INSERT INTO %s (
        name, address, responsible_person, phone, email,
        max_capacity, current_occupancy, other_fields, country
    ) VALUES (
        $1, $2, $3, $4, $5,
        $6, $7, $8, $9
    ) RETURNING id
    `, domain.TableWarehouse)

	var id int64
	if err = wpr.db.QueryRow(query,
		warehouse.Name, warehouse.Address, warehouse.ResponsiblePerson, warehouse.Phone, warehouse.Email,
		warehouse.MaxCapacity, warehouse.CurrentOccupancy, otherFieldsJSON, warehouse.Country,
	).Scan(&id); err != nil {
		return 0, fmt.Errorf("failed to insert warehouse: %v", err)
	}

	return id, nil
}

func (wpr *WarehousePostgresRepository) GetById(ctx context.Context, id int64) (domain.Warehouse, error) {
	query := fmt.Sprintf(`
    SELECT
        id, name, address, responsible_person, phone, email,
        max_capacity, current_occupancy, other_fields, country
    FROM %s
    WHERE id = $1;
    `, domain.TableWarehouse)

	var warehouse domain.Warehouse
	var otherFieldsJSON []byte

	row := wpr.db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&warehouse.ID, &warehouse.Name, &warehouse.Address, &warehouse.ResponsiblePerson, &warehouse.Phone, &warehouse.Email,
		&warehouse.MaxCapacity, &warehouse.CurrentOccupancy, &otherFieldsJSON, &warehouse.Country,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Warehouse{}, fmt.Errorf("warehouse with ID %d not found", id)
		}

		return domain.Warehouse{}, fmt.Errorf("failed to get warehouse by ID: %v", err)
	}

	if err = json.Unmarshal(otherFieldsJSON, &warehouse.OtherFields); err != nil {
		return domain.Warehouse{}, fmt.Errorf("failed to unmarshal other_fields JSON: %v", err)
	}

	return warehouse, nil
}
