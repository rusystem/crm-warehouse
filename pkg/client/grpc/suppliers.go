package grpc

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/rusystem/crm-warehouse/pkg/gen/proto/supplier"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Supplier struct {
	ID                int64                  `gorm:"primaryKey" json:"id"` // Уникальный идентификатор поставщика
	Name              string                 `json:"name"`                 // Наименование поставщика
	LegalAddress      string                 `json:"legal_address"`        // Юридический адрес поставщика
	ActualAddress     string                 `json:"actual_address"`       // Фактический адрес поставщика
	WarehouseAddress  string                 `json:"warehouse_address"`    // Адрес склада поставщика
	ContactPerson     string                 `json:"contact_person"`       // Контактное лицо у поставщика
	Phone             string                 `json:"phone"`                // Телефон поставщика
	Email             string                 `json:"email"`                // Электронная почта поставщика
	Website           string                 `json:"website"`              // Сайт поставщика
	ContractNumber    string                 `json:"contract_number"`      // Номер и дата договора с поставщиком
	ProductCategories string                 `json:"product_categories"`   // Категории товаров, поставляемых поставщиком
	PurchaseAmount    float64                `json:"purchase_amount"`      // Общая сумма закупок у поставщика
	Balance           float64                `json:"balance"`              // Баланс по поставщику
	ProductTypes      int64                  `json:"product_types"`        // Количество типов товаров от поставщика
	Comments          string                 `json:"comments"`             // Комментарии
	Files             string                 `json:"files"`                // Ссылки на файлы или документы
	Country           string                 `json:"country"`              // Страна поставщика
	Region            string                 `json:"region"`               // Регион или штат поставщика
	TaxID             string                 `json:"tax_id"`               // Идентификационный номер налогоплательщика (ИНН)
	BankDetails       string                 `json:"bank_details"`         // Банковские реквизиты поставщика
	RegistrationDate  time.Time              `json:"registration_date"`    // Дата регистрации поставщика
	PaymentTerms      string                 `json:"payment_terms"`        // Условия оплаты по контракту
	IsActive          bool                   `json:"is_active"`            // Статус активности поставщика (активен/неактивен)
	OtherFields       map[string]interface{} `json:"other_fields"`         // Дополнительные пользовательские поля
}

type SuppliersClient struct {
	conn           *grpc.ClientConn
	supplierClient supplier.SupplierServiceClient
}

func NewSuppliersClient(addr string) (*SuppliersClient, error) {
	opt := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(addr, opt...)
	if err != nil {
		return nil, err
	}

	return &SuppliersClient{
		conn:           conn,
		supplierClient: supplier.NewSupplierServiceClient(conn),
	}, nil
}

func (s *SuppliersClient) Close() error {
	return s.conn.Close()
}

func (s *SuppliersClient) GetById(ctx context.Context, id int64) (Supplier, error) {
	if id <= 0 {
		return Supplier{}, errors.New("calls grpc: id can`t be zero")
	}

	resp, err := s.supplierClient.GetById(ctx, &supplier.Id{Id: id})
	if err != nil {
		return Supplier{}, err
	}

	var otherFields map[string]interface{}
	if err = json.Unmarshal([]byte(resp.OtherFields), &otherFields); err != nil {
		return Supplier{}, err
	}

	return Supplier{
		ID:                resp.Id,
		Name:              resp.Name,
		LegalAddress:      resp.LegalAddress,
		ActualAddress:     resp.ActualAddress,
		WarehouseAddress:  resp.WarehouseAddress,
		ContactPerson:     resp.ContactPerson,
		Phone:             resp.Phone,
		Email:             resp.Email,
		Website:           resp.Website,
		ContractNumber:    resp.ContractNumber,
		ProductCategories: resp.ProductCategories,
		PurchaseAmount:    resp.PurchaseAmount,
		Balance:           resp.Balance,
		ProductTypes:      resp.ProductTypes,
		Comments:          resp.Comments,
		Files:             resp.Files,
		Country:           resp.Country,
		Region:            resp.Region,
		TaxID:             resp.TaxId,
		BankDetails:       resp.BankDetails,
		RegistrationDate:  resp.RegistrationDate.AsTime(),
		PaymentTerms:      resp.PaymentTerms,
		IsActive:          resp.IsActive,
		OtherFields:       otherFields,
	}, nil
}

func (s *SuppliersClient) Create(ctx context.Context, spl Supplier) (int64, error) {
	otherFieldsJSON, err := json.Marshal(spl.OtherFields)
	if err != nil {
		return 0, err
	}

	resp, err := s.supplierClient.Create(ctx, &supplier.Supplier{
		Id:                spl.ID,
		Name:              spl.Name,
		LegalAddress:      spl.LegalAddress,
		ActualAddress:     spl.ActualAddress,
		WarehouseAddress:  spl.WarehouseAddress,
		ContactPerson:     spl.ContactPerson,
		Phone:             spl.Phone,
		Email:             spl.Email,
		Website:           spl.Website,
		ContractNumber:    spl.ContractNumber,
		ProductCategories: spl.ProductCategories,
		PurchaseAmount:    spl.PurchaseAmount,
		Balance:           spl.Balance,
		ProductTypes:      spl.ProductTypes,
		Comments:          spl.Comments,
		Files:             spl.Files,
		Country:           spl.Country,
		Region:            spl.Region,
		TaxId:             spl.TaxID,
		BankDetails:       spl.BankDetails,
		RegistrationDate:  timestamppb.New(spl.RegistrationDate),
		PaymentTerms:      spl.PaymentTerms,
		IsActive:          spl.IsActive,
		OtherFields:       string(otherFieldsJSON),
	})
	if err != nil {
		return 0, err
	}

	return resp.Id, nil
}
