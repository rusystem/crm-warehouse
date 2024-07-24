package handler

import (
	"context"
	"encoding/json"
	"github.com/rusystem/crm-warehouse/internal/service"
	"github.com/rusystem/crm-warehouse/pkg/domain"
	"github.com/rusystem/crm-warehouse/pkg/gen/proto/warehouse"
)

type WarehouseHandler struct {
	service *service.Service
}

func NewWarehouseHandler(service *service.Service) *WarehouseHandler {
	return &WarehouseHandler{
		service: service,
	}
}

func (wh *WarehouseHandler) GetById(ctx context.Context, id *warehouse.Id) (*warehouse.Warehouse, error) {
	whs, err := wh.service.Warehouse.GetById(ctx, id.Id)
	if err != nil {
		return nil, err
	}

	otherFieldsJSON, err := json.Marshal(whs.OtherFields)
	if err != nil {
		return nil, err
	}

	return &warehouse.Warehouse{
		Id:                whs.ID,
		Name:              whs.Name,
		Address:           whs.Address,
		ResponsiblePerson: whs.ResponsiblePerson,
		Phone:             whs.Phone,
		Email:             whs.Email,
		MaxCapacity:       whs.MaxCapacity,
		CurrentOccupancy:  whs.CurrentOccupancy,
		OtherFields:       string(otherFieldsJSON),
		Country:           whs.Country,
	}, nil
}

func (wh *WarehouseHandler) Create(ctx context.Context, whs *warehouse.Warehouse) (*warehouse.Id, error) {
	var otherFields map[string]interface{}
	if err := json.Unmarshal([]byte(whs.OtherFields), &otherFields); err != nil {
		return nil, err
	}

	id, err := wh.service.Warehouse.Create(ctx, domain.Warehouse{
		ID:                whs.Id,
		Name:              whs.Name,
		Address:           whs.Address,
		ResponsiblePerson: whs.ResponsiblePerson,
		Phone:             whs.Phone,
		Email:             whs.Email,
		MaxCapacity:       whs.MaxCapacity,
		CurrentOccupancy:  whs.CurrentOccupancy,
		OtherFields:       otherFields,
		Country:           whs.Country,
	})
	if err != nil {
		return nil, err
	}

	return &warehouse.Id{Id: id}, nil
}
