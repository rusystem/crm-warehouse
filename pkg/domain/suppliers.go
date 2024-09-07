package domain

import "time"

type Supplier struct {
	ID                int64                  `json:"id"`                 // Уникальный идентификатор поставщика
	Name              string                 `json:"name"`               // Наименование поставщика
	LegalAddress      string                 `json:"legal_address"`      // Юридический адрес поставщика
	ActualAddress     string                 `json:"actual_address"`     // Фактический адрес поставщика
	WarehouseAddress  string                 `json:"warehouse_address"`  // Адрес склада поставщика
	ContactPerson     string                 `json:"contact_person"`     // Контактное лицо у поставщика
	Phone             string                 `json:"phone"`              // Телефон поставщика
	Email             string                 `json:"email"`              // Электронная почта поставщика
	Website           string                 `json:"website"`            // Сайт поставщика
	ContractNumber    string                 `json:"contract_number"`    // Номер и дата договора с поставщиком
	ProductCategories string                 `json:"product_categories"` // Категории товаров, поставляемых поставщиком
	PurchaseAmount    float64                `json:"purchase_amount"`    // Общая сумма закупок у поставщика
	Balance           float64                `json:"balance"`            // Баланс по поставщику
	ProductTypes      int64                  `json:"product_types"`      // Количество типов товаров от поставщика
	Comments          string                 `json:"comments"`           // Комментарии
	Files             string                 `json:"files"`              // Ссылки на файлы или документы
	Country           string                 `json:"country"`            // Страна поставщика
	Region            string                 `json:"region"`             // Регион или штат поставщика
	TaxID             string                 `json:"tax_id"`             // Идентификационный номер налогоплательщика (ИНН)
	BankDetails       string                 `json:"bank_details"`       // Банковские реквизиты поставщика
	RegistrationDate  time.Time              `json:"registration_date"`  // Дата регистрации поставщика
	PaymentTerms      string                 `json:"payment_terms"`      // Условия оплаты по контракту
	IsActive          bool                   `json:"is_active"`          // Статус активности поставщика (активен/неактивен)
	OtherFields       map[string]interface{} `json:"other_fields"`       // Дополнительные пользовательские поля
	CompanyID         int64                  `json:"company_id"`         // ID компании
}
