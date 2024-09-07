package domain

// Warehouse представляет данные о складе
type Warehouse struct {
	ID                int64                  `json:"id"`                 // Уникальный идентификатор склада
	Name              string                 `json:"name"`               // Название склада
	Address           string                 `json:"address"`            // Адрес склада
	ResponsiblePerson string                 `json:"responsible_person"` // Ответственное лицо за склад
	Phone             string                 `json:"phone"`              // Контактный телефон склада
	Email             string                 `json:"email"`              // Электронная почта для связи
	MaxCapacity       int64                  `json:"max_capacity"`       // Максимальная вместимость склада
	CurrentOccupancy  int64                  `json:"current_occupancy"`  // Текущая заполняемость склада
	OtherFields       map[string]interface{} `json:"other_fields"`       // Дополнительные пользовательские поля
	Country           string                 `json:"country"`            // Страна склада
	CompanyID         int64                  `json:"company_id"`         // ID компании
}
