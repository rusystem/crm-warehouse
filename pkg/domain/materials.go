package domain

import "time"

// Material представляет структуру товара
type Material struct {
	ID                     int64                  `json:"id"`                       // Уникальный идентификатор записи
	WarehouseID            int64                  `json:"warehouse_id"`             // Id склада
	ItemID                 int64                  `json:"item_id"`                  // Идентификатор товара
	Name                   string                 `json:"name"`                     // Наименование товара
	ByInvoice              string                 `json:"by_invoice"`               // Накладная на товар
	Article                string                 `json:"article"`                  // Артикул товара
	ProductCategory        string                 `json:"product_category"`         // Категория товара
	Unit                   string                 `json:"unit"`                     // Единица измерения
	TotalQuantity          int64                  `json:"total_quantity"`           // Общее количество товара
	Volume                 int64                  `json:"volume"`                   // Объем товара
	PriceWithoutVAT        float64                `json:"price_without_vat"`        // Цена без НДС
	TotalWithoutVAT        float64                `json:"total_without_vat"`        // Общая стоимость без НДС
	SupplierID             int64                  `json:"supplier_id"`              // Поставщик товара
	Location               string                 `json:"location"`                 // Локация на складе
	Contract               time.Time              `json:"contract"`                 // Дата договора
	File                   string                 `json:"file"`                     // Файл, связанный с товаром
	Status                 string                 `json:"status"`                   // Статус товара
	Comments               string                 `json:"comments"`                 // Комментарии
	Reserve                string                 `json:"reserve"`                  // Резерв товара
	ReceivedDate           time.Time              `json:"received_date"`            // Дата поступления товара
	LastUpdated            time.Time              `json:"last_updated"`             // Дата последнего обновления информации о товаре
	MinStockLevel          int64                  `json:"min_stock_level"`          // Минимальный уровень запаса
	ExpirationDate         time.Time              `json:"expiration_date"`          // Срок годности товара
	ResponsiblePerson      string                 `json:"responsible_person"`       // Ответственное лицо за товар
	StorageCost            float64                `json:"storage_cost"`             // Стоимость хранения товара
	WarehouseSection       string                 `json:"warehouse_section"`        // Секция склада, где хранится товар
	Barcode                string                 `json:"barcode"`                  // Штрих-код товара
	IncomingDeliveryNumber string                 `json:"incoming_delivery_number"` // Входящий номер поставки
	OtherFields            map[string]interface{} `json:"other_fields"`             // Дополнительные пользовательские поля
	CompanyID              int64                  `json:"company_id"`               // Кабинет компании к кому привязан товар
}
