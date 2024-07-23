package handler

import (
	"context"
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

	return &warehouse.Warehouse{
		Id:                whs.ID,
		Name:              whs.Name,
		Address:           whs.Address,
		ResponsiblePerson: whs.ResponsiblePerson,
		Phone:             whs.Phone,
		Email:             whs.Email,
		MaxCapacity:       whs.MaxCapacity,
		CurrentOccupancy:  whs.CurrentOccupancy,
		OtherFields:       whs.OtherFields,
		Country:           whs.Country,
	}, nil
}

func (wh *WarehouseHandler) Create(ctx context.Context, whs *warehouse.Warehouse) (*warehouse.Id, error) {
	id, err := wh.service.Warehouse.Create(ctx, domain.Warehouse{
		ID:                whs.Id,
		Name:              whs.Name,
		Address:           whs.Address,
		ResponsiblePerson: whs.ResponsiblePerson,
		Phone:             whs.Phone,
		Email:             whs.Email,
		MaxCapacity:       whs.MaxCapacity,
		CurrentOccupancy:  whs.CurrentOccupancy,
		OtherFields:       whs.OtherFields,
		Country:           whs.Country,
	})
	if err != nil {
		return nil, err
	}

	return &warehouse.Id{Id: id}, nil
}
