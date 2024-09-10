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
	Update(ctx context.Context, warehouse domain.Warehouse) error
	Delete(ctx context.Context, id int64) error
	GetListByCompanyId(ctx context.Context, id int64) ([]domain.Warehouse, error)
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
        max_capacity, current_occupancy, other_fields, country, company_id
    ) VALUES (
        $1, $2, $3, $4, $5,
        $6, $7, $8, $9, $10
    ) RETURNING id
    `, domain.TableWarehouse)

	var id int64
	if err = wpr.db.QueryRowContext(ctx, query,
		warehouse.Name, warehouse.Address, warehouse.ResponsiblePerson, warehouse.Phone, warehouse.Email,
		warehouse.MaxCapacity, warehouse.CurrentOccupancy, otherFieldsJSON, warehouse.Country, warehouse.CompanyID,
	).Scan(&id); err != nil {
		return 0, fmt.Errorf("failed to insert warehouse: %v", err)
	}

	return id, nil
}

func (wpr *WarehousePostgresRepository) GetById(ctx context.Context, id int64) (domain.Warehouse, error) {
	query := fmt.Sprintf(`
    SELECT
        id, name, address, responsible_person, phone, email,
        max_capacity, current_occupancy, other_fields, country, company_id
    FROM %s
    WHERE id = $1;
    `, domain.TableWarehouse)

	var warehouse domain.Warehouse
	var otherFieldsJSON []byte

	row := wpr.db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&warehouse.ID, &warehouse.Name, &warehouse.Address, &warehouse.ResponsiblePerson, &warehouse.Phone, &warehouse.Email,
		&warehouse.MaxCapacity, &warehouse.CurrentOccupancy, &otherFieldsJSON, &warehouse.Country, &warehouse.CompanyID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Warehouse{}, domain.ErrWarehouseNotFound
		}

		return domain.Warehouse{}, fmt.Errorf("failed to get warehouse by ID: %v", err)
	}

	if err = json.Unmarshal(otherFieldsJSON, &warehouse.OtherFields); err != nil {
		return domain.Warehouse{}, fmt.Errorf("failed to unmarshal other_fields JSON: %v", err)
	}

	return warehouse, nil
}

func (wpr *WarehousePostgresRepository) Update(ctx context.Context, warehouse domain.Warehouse) error {
	otherFieldsJSON, err := json.Marshal(warehouse.OtherFields)
	if err != nil {
		return fmt.Errorf("failed to marshal other_fields to JSON: %v", err)
	}

	query := fmt.Sprintf(`
	UPDATE %s
	SET
		name = $1, address = $2, responsible_person = $3, phone = $4, email = $5,
		max_capacity = $6, current_occupancy = $7, other_fields = $8, country = $9
	WHERE id = $10
	`, domain.TableWarehouse)

	_, err = wpr.db.ExecContext(ctx, query,
		warehouse.Name, warehouse.Address, warehouse.ResponsiblePerson, warehouse.Phone, warehouse.Email,
		warehouse.MaxCapacity, warehouse.CurrentOccupancy, otherFieldsJSON, warehouse.Country,
		warehouse.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update warehouse: %v", err)
	}

	return nil
}

func (wpr *WarehousePostgresRepository) Delete(ctx context.Context, id int64) error {
	_, err := wpr.db.ExecContext(ctx, fmt.Sprintf("DELETE FROM %s WHERE id = $1", domain.TableWarehouse), id)
	return err
}

func (wpr *WarehousePostgresRepository) GetListByCompanyId(ctx context.Context, id int64) ([]domain.Warehouse, error) {
	query := fmt.Sprintf(`
	SELECT
		id, name, address, responsible_person, phone, email,
		max_capacity, current_occupancy, other_fields, country, company_id
	FROM %s
	WHERE company_id = $1;
	`, domain.TableWarehouse)

	rows, err := wpr.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get warehouses by company ID: %v", err)
	}
	defer func(rows *sql.Rows) {
		if err = rows.Close(); err != nil {
			return
		}
	}(rows)

	var warehouses []domain.Warehouse
	for rows.Next() {
		var warehouse domain.Warehouse
		var otherFieldsJSON []byte
		if err = rows.Scan(
			&warehouse.ID, &warehouse.Name, &warehouse.Address, &warehouse.ResponsiblePerson, &warehouse.Phone, &warehouse.Email,
			&warehouse.MaxCapacity, &warehouse.CurrentOccupancy, &otherFieldsJSON, &warehouse.Country, &warehouse.CompanyID,
		); err != nil {
			return nil, fmt.Errorf("failed to scan warehouse: %v", err)
		}

		if err = json.Unmarshal(otherFieldsJSON, &warehouse.OtherFields); err != nil {
			return nil, fmt.Errorf("failed to unmarshal other_fields JSON: %v", err)
		}

		warehouses = append(warehouses, warehouse)
	}

	return warehouses, nil
}
