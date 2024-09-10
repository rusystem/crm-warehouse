package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/rusystem/crm-warehouse/internal/service"
	"github.com/rusystem/crm-warehouse/pkg/domain"
	"github.com/rusystem/crm-warehouse/pkg/gen/proto/warehouse"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type WarehouseHandler struct {
	service *service.Service
}

func NewWarehouseHandler(service *service.Service) *WarehouseHandler {
	return &WarehouseHandler{
		service: service,
	}
}

func (wh *WarehouseHandler) GetById(ctx context.Context, id *warehouse.WarehouseId) (*warehouse.Warehouse, error) {
	whs, err := wh.service.Warehouse.GetById(ctx, id.Id)
	if err != nil {
		if errors.Is(err, domain.ErrWarehouseNotFound) {
			return nil, status.Errorf(codes.NotFound, "warehouse with ID %d not found", id.Id)
		}

		return nil, status.Errorf(codes.Internal, "internal server error - %v", err)
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
		CompanyId:         whs.CompanyID,
	}, nil
}

func (wh *WarehouseHandler) Create(ctx context.Context, whs *warehouse.Warehouse) (*warehouse.WarehouseId, error) {
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
		CompanyID:         whs.CompanyId,
	})
	if err != nil {
		return nil, err
	}

	return &warehouse.WarehouseId{Id: id}, nil
}

func (wh *WarehouseHandler) Update(ctx context.Context, whs *warehouse.Warehouse) (*emptypb.Empty, error) {
	var otherFields map[string]interface{}
	if err := json.Unmarshal([]byte(whs.OtherFields), &otherFields); err != nil {
		return nil, err
	}

	if err := wh.service.Warehouse.Update(ctx, domain.Warehouse{
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
	}); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (wh *WarehouseHandler) Delete(ctx context.Context, req *warehouse.WarehouseId) (*emptypb.Empty, error) {
	if err := wh.service.Warehouse.Delete(ctx, req.Id); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (wh *WarehouseHandler) GetList(ctx context.Context, req *warehouse.WarehouseCompanyId) (*warehouse.WarehouseList, error) {
	warehouses, err := wh.service.Warehouse.GetListByCompanyId(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	var resp []*warehouse.Warehouse
	for _, w := range warehouses {
		otherFieldsJSON, err := json.Marshal(w.OtherFields)
		if err != nil {
			return nil, err
		}

		resp = append(resp, &warehouse.Warehouse{
			Id:                w.ID,
			Name:              w.Name,
			Address:           w.Address,
			ResponsiblePerson: w.ResponsiblePerson,
			Phone:             w.Phone,
			Email:             w.Email,
			MaxCapacity:       w.MaxCapacity,
			CurrentOccupancy:  w.CurrentOccupancy,
			OtherFields:       string(otherFieldsJSON),
			Country:           w.Country,
			CompanyId:         w.CompanyID,
		})
	}

	return &warehouse.WarehouseList{Warehouses: resp}, nil
}
