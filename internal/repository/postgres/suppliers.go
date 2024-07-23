package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rusystem/crm-warehouse/pkg/domain"
)

type Suppliers interface {
	Create(ctx context.Context, supplier domain.Supplier) (int64, error)
	GetById(ctx context.Context, id int64) (domain.Supplier, error)
}

type SuppliersPostgresRepository struct {
	psql *sql.DB
}

func NewSuppliersPostgresRepository(psql *sql.DB) *SuppliersPostgresRepository {
	return &SuppliersPostgresRepository{psql: psql}
}

func (sr *SuppliersPostgresRepository) Create(ctx context.Context, supplier domain.Supplier) (int64, error) {
	otherFieldsJSON, err := json.Marshal(supplier.OtherFields)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal other_fields to JSON: %v", err)
	}

	query := fmt.Sprintf(`
		INSERT INTO %s (name, legal_address, actual_address, warehouse_address, contact_person, phone, email, 
		                       website, contract_number, product_categories, purchase_amount, balance, product_types, 
		                       comments, files, country, region, tax_id, bank_details, registration_date, payment_terms, 
		                       is_active, other_fields) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23) RETURNING id`,
		domain.TableSupplier)

	var id int64
	if err = sr.psql.QueryRowContext(ctx, query, supplier.Name, supplier.LegalAddress, supplier.ActualAddress,
		supplier.WarehouseAddress, supplier.ContactPerson, supplier.Phone, supplier.Email, supplier.Website,
		supplier.ContractNumber, supplier.ProductCategories, supplier.PurchaseAmount, supplier.Balance, supplier.ProductTypes,
		supplier.Comments, supplier.Files, supplier.Country, supplier.Region, supplier.TaxID, supplier.BankDetails,
		supplier.RegistrationDate, supplier.PaymentTerms, supplier.IsActive, otherFieldsJSON,
	).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (sr *SuppliersPostgresRepository) GetById(ctx context.Context, id int64) (domain.Supplier, error) {
	query := fmt.Sprintf(`
    SELECT
        id, name, legal_address, actual_address, warehouse_address,
        contact_person, phone, email, website, contract_number,
        product_categories, purchase_amount, balance, product_types,
        comments, files, country, region, tax_id, bank_details,
        registration_date, payment_terms, is_active, other_fields
    FROM %s
    WHERE id = $1;
    `, domain.TableSupplier)

	var supplier domain.Supplier
	var otherFieldsJSON []byte

	// Выполнение запроса и сканирование результата в объект Supplier
	row := sr.psql.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&supplier.ID, &supplier.Name, &supplier.LegalAddress, &supplier.ActualAddress,
		&supplier.WarehouseAddress, &supplier.ContactPerson, &supplier.Phone, &supplier.Email,
		&supplier.Website, &supplier.ContractNumber, &supplier.ProductCategories, &supplier.PurchaseAmount,
		&supplier.Balance, &supplier.ProductTypes, &supplier.Comments, &supplier.Files,
		&supplier.Country, &supplier.Region, &supplier.TaxID, &supplier.BankDetails,
		&supplier.RegistrationDate, &supplier.PaymentTerms, &supplier.IsActive, &otherFieldsJSON,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Supplier{}, fmt.Errorf("supplier with ID %d not found", id)
		}

		return domain.Supplier{}, err
	}

	if err = json.Unmarshal(otherFieldsJSON, &supplier.OtherFields); err != nil {
		return domain.Supplier{}, fmt.Errorf("failed to unmarshal other_fields JSON: %v", err)
	}

	return supplier, nil
}
