package grpc

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/rusystem/crm-warehouse/pkg/domain"
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
	CompanyId         int64                  `json:"company_id"`           // Уникальный идентификатор компании
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

	resp, err := w.warehouseClient.GetById(ctx, &warehouse.WarehouseId{Id: id})
	if err != nil {
		if err.Error() == "rpc error: code = Unknown desc = sql: no rows in result set" {
			return Warehouse{}, domain.ErrWarehouseNotFound
		}

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
		CompanyId:         resp.CompanyId,
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
		CompanyId:         wh.CompanyId,
	})
	if err != nil {
		return 0, err
	}

	return resp.Id, nil
}

func (w *WarehouseClient) Update(ctx context.Context, wh Warehouse) error {
	otherFieldsJSON, err := json.Marshal(wh.OtherFields)
	if err != nil {
		return err
	}

	_, err = w.warehouseClient.Update(ctx, &warehouse.Warehouse{
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
		return err
	}

	return nil
}

func (w *WarehouseClient) Delete(ctx context.Context, id int64) error {
	_, err := w.warehouseClient.Delete(ctx, &warehouse.WarehouseId{Id: id})
	return err
}

func (w *WarehouseClient) GetList(ctx context.Context, companyId int64) ([]Warehouse, error) {
	var warehouses []Warehouse
	resp, err := w.warehouseClient.GetList(ctx, &warehouse.WarehouseCompanyId{Id: companyId})
	if err != nil {
		return warehouses, err
	}

	for _, wh := range resp.Warehouses {
		var otherFields map[string]interface{}
		if err = json.Unmarshal([]byte(wh.OtherFields), &otherFields); err != nil {
			return warehouses, err
		}

		warehouses = append(warehouses, Warehouse{
			ID:                wh.Id,
			Name:              wh.Name,
			Address:           wh.Address,
			ResponsiblePerson: wh.ResponsiblePerson,
			Phone:             wh.Phone,
			Email:             wh.Email,
			MaxCapacity:       wh.MaxCapacity,
			CurrentOccupancy:  wh.CurrentOccupancy,
			OtherFields:       otherFields,
			Country:           wh.Country,
			CompanyId:         wh.CompanyId,
		})
	}

	return warehouses, nil
}

func (w *WarehouseClient) GetResponsiblePerson(ctx context.Context, companyId int64) ([]domain.User, error) {
	resp, err := w.warehouseClient.GetResponsibleUsers(ctx, &warehouse.WarehouseCompanyId{Id: companyId})
	if err != nil {
		return nil, err
	}

	var users []domain.User
	for _, v := range resp.Users {
		var lastLogin sql.NullTime

		if !v.LastLogin.AsTime().IsZero() {
			lastLogin = sql.NullTime{
				Time:  v.LastLogin.AsTime(),
				Valid: true,
			}
		}

		users = append(users, domain.User{
			ID:                       v.Id,
			CompanyID:                v.CompanyId,
			Username:                 v.Username,
			Name:                     v.Name,
			Email:                    v.Email,
			Phone:                    v.Phone,
			PasswordHash:             v.PasswordHash,
			CreatedAt:                v.CreatedAt.AsTime(),
			UpdatedAt:                v.UpdatedAt.AsTime(),
			LastLogin:                lastLogin,
			IsActive:                 v.IsActive,
			Role:                     v.Role,
			Language:                 v.Language,
			Country:                  v.Country,
			IsApproved:               v.IsApproved,
			IsSendSystemNotification: v.IsSendSystemNotification,
			Sections:                 v.Sections,
			Position:                 v.Position,
		})
	}

	return users, nil
}
