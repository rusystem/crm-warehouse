package grpc

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/rusystem/crm-warehouse/pkg/gen/proto/warehouse"
	"google.golang.org/grpc"
)

type Warehouse struct {
	ID                int64                  `gorm:"primaryKey" json:"id"` // Уникальный идентификатор склада
	Name              string                 `json:"name"`                 // Название склада
	Address           string                 `json:"address"`              // Адрес склада
	ResponsiblePerson string                 `json:"responsible_person"`   // Ответственное лицо за склад
	Phone             string                 `json:"phone"`                // Контактный телефон склада
	Email             string                 `json:"email"`                // Электронная почта для связи
	MaxCapacity       int64                  `json:"max_capacity"`         // Максимальная вместимость склада
	CurrentOccupancy  int64                  `json:"current_occupancy"`    // Текущая заполняемость склада
	OtherFields       map[string]interface{} `json:"other_fields"`         // Дополнительные пользовательские поля
	Country           string                 `json:"country"`              // Страна склада
}

type WarehouseClient struct {
	conn            *grpc.ClientConn
	warehouseClient warehouse.WarehouseServiceClient
}

func NewWarehouseClient(addr string) (*WarehouseClient, error) {
	opt := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(addr, opt...)
	if err != nil {
		return nil, err
	}

	return &WarehouseClient{
		conn:            conn,
		warehouseClient: warehouse.NewWarehouseServiceClient(conn),
	}, nil
}

func (w *WarehouseClient) Close() error {
	return w.conn.Close()
}

func (w *WarehouseClient) GetById(ctx context.Context, id int64) (Warehouse, error) {
	if id <= 0 {
		return Warehouse{}, errors.New("calls grpc: id can`t be zero")
	}

	resp, err := w.warehouseClient.GetById(ctx, &warehouse.Id{Id: id})
	if err != nil {
		return Warehouse{}, err
	}

	var otherFields map[string]interface{}
	if err = json.Unmarshal([]byte(resp.OtherFields), &otherFields); err != nil {
		return Warehouse{}, err
	}

	return Warehouse{
		ID:                resp.Id,
		Name:              resp.Name,
		Address:           resp.Address,
		ResponsiblePerson: resp.ResponsiblePerson,
		Phone:             resp.Phone,
		Email:             resp.Email,
		MaxCapacity:       resp.MaxCapacity,
		CurrentOccupancy:  resp.CurrentOccupancy,
		OtherFields:       otherFields,
		Country:           resp.Country,
	}, nil
}

func (w *WarehouseClient) Create(ctx context.Context, wh Warehouse) (int64, error) {
	otherFieldsJSON, err := json.Marshal(wh.OtherFields)
	if err != nil {
		return 0, err
	}

	resp, err := w.warehouseClient.Create(ctx, &warehouse.Warehouse{
		Id:                wh.ID,
		Name:              wh.Name,
		Address:           wh.Address,
		ResponsiblePerson: wh.ResponsiblePerson,
		Phone:             wh.Phone,
		Email:             wh.Email,
		MaxCapacity:       wh.MaxCapacity,
		CurrentOccupancy:  wh.CurrentOccupancy,
		OtherFields:       string(otherFieldsJSON),
		Country:           wh.Country,
	})
	if err != nil {
		return 0, err
	}

	return resp.Id, nil
}
