package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/rusystem/crm-warehouse/internal/service"
	"github.com/rusystem/crm-warehouse/pkg/domain"
	"github.com/rusystem/crm-warehouse/pkg/gen/proto/materials"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MaterialsHandler struct {
	service *service.Service
}

func NewMaterialsHandler(service *service.Service) *MaterialsHandler {
	return &MaterialsHandler{
		service: service,
	}
}

func (mh *MaterialsHandler) CreatePlanning(ctx context.Context, material *materials.Material) (*materials.MaterialId, error) {
	var otherFields map[string]interface{}
	if err := json.Unmarshal([]byte(material.OtherFields), &otherFields); err != nil {
		return nil, err
	}

	id, err := mh.service.Material.CreatePlanning(ctx, domain.Material{
		WarehouseID:            material.WarehouseId,
		ItemID:                 material.ItemId,
		Name:                   material.Name,
		ByInvoice:              material.ByInvoice,
		Article:                material.Article,
		ProductCategory:        material.ProductCategory,
		Unit:                   material.Unit,
		TotalQuantity:          material.TotalQuantity,
		Volume:                 material.Volume,
		PriceWithoutVAT:        material.PriceWithoutVat,
		TotalWithoutVAT:        material.TotalWithoutVat,
		SupplierID:             material.SupplierId,
		Location:               material.Location,
		Contract:               material.Contract.AsTime(),
		File:                   material.File,
		Status:                 material.Status,
		Comments:               material.Comments,
		Reserve:                material.Reserve,
		ReceivedDate:           material.ReceivedDate.AsTime(),
		LastUpdated:            material.LastUpdated.AsTime(),
		MinStockLevel:          material.MinStockLevel,
		ExpirationDate:         material.ExpirationDate.AsTime(),
		ResponsiblePerson:      material.ResponsiblePerson,
		StorageCost:            material.StorageCost,
		WarehouseSection:       material.WarehouseSection,
		IncomingDeliveryNumber: material.IncomingDeliveryNumber,
		OtherFields:            otherFields,
		CompanyID:              material.CompanyId,
	})
	if err != nil {
		return nil, err
	}

	return &materials.MaterialId{Id: id}, nil
}

func (mh *MaterialsHandler) UpdatePlanning(ctx context.Context, material *materials.Material) (*emptypb.Empty, error) {
	var otherFields map[string]interface{}
	if err := json.Unmarshal([]byte(material.OtherFields), &otherFields); err != nil {
		return nil, err
	}

	err := mh.service.Material.UpdatePlanning(ctx, domain.Material{
		ID:                     material.Id,
		WarehouseID:            material.WarehouseId,
		ItemID:                 material.ItemId,
		Name:                   material.Name,
		ByInvoice:              material.ByInvoice,
		Article:                material.Article,
		ProductCategory:        material.ProductCategory,
		Unit:                   material.Unit,
		TotalQuantity:          material.TotalQuantity,
		Volume:                 material.Volume,
		PriceWithoutVAT:        material.PriceWithoutVat,
		TotalWithoutVAT:        material.TotalWithoutVat,
		SupplierID:             material.SupplierId,
		Location:               material.Location,
		Contract:               material.Contract.AsTime(),
		File:                   material.File,
		Status:                 material.Status,
		Comments:               material.Comments,
		Reserve:                material.Reserve,
		ReceivedDate:           material.ReceivedDate.AsTime(),
		LastUpdated:            material.LastUpdated.AsTime(),
		MinStockLevel:          material.MinStockLevel,
		ExpirationDate:         material.ExpirationDate.AsTime(),
		ResponsiblePerson:      material.ResponsiblePerson,
		StorageCost:            material.StorageCost,
		WarehouseSection:       material.WarehouseSection,
		IncomingDeliveryNumber: material.IncomingDeliveryNumber,
		OtherFields:            otherFields,
	})
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (mh *MaterialsHandler) DeletePlanning(ctx context.Context, req *materials.MaterialId) (*emptypb.Empty, error) {
	if err := mh.service.Material.DeletePlanning(ctx, req.Id); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (mh *MaterialsHandler) GetPlanning(ctx context.Context, req *materials.MaterialId) (*materials.Material, error) {
	material, err := mh.service.Material.GetPlanningById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	otherFieldsJSON, err := json.Marshal(material.OtherFields)
	if err != nil {
		return nil, err
	}

	return &materials.Material{
		Id:                     material.ID,
		WarehouseId:            material.WarehouseID,
		ItemId:                 material.ItemID,
		Name:                   material.Name,
		ByInvoice:              material.ByInvoice,
		Article:                material.Article,
		ProductCategory:        material.ProductCategory,
		Unit:                   material.Unit,
		TotalQuantity:          material.TotalQuantity,
		Volume:                 material.Volume,
		PriceWithoutVat:        material.PriceWithoutVAT,
		TotalWithoutVat:        material.TotalWithoutVAT,
		SupplierId:             material.SupplierID,
		Location:               material.Location,
		Contract:               timestamppb.New(material.Contract),
		File:                   material.File,
		Status:                 material.Status,
		Comments:               material.Comments,
		Reserve:                material.Reserve,
		ReceivedDate:           timestamppb.New(material.ReceivedDate),
		LastUpdated:            timestamppb.New(material.LastUpdated),
		MinStockLevel:          material.MinStockLevel,
		ExpirationDate:         timestamppb.New(material.ExpirationDate),
		ResponsiblePerson:      material.ResponsiblePerson,
		StorageCost:            material.StorageCost,
		WarehouseSection:       material.WarehouseSection,
		IncomingDeliveryNumber: material.IncomingDeliveryNumber,
		OtherFields:            string(otherFieldsJSON),
		CompanyId:              material.CompanyID,
	}, nil
}

func (mh *MaterialsHandler) GetListPlanning(ctx context.Context, req *materials.MaterialParams) (*materials.MaterialList, error) {
	if req.Limit <= 0 {
		return nil, errors.New("materials, grpc handler - invalid limit")
	}

	if req.Offset < 0 {
		return nil, errors.New("materials, grpc handler - invalid offset")
	}

	if req.CompanyId <= 0 {
		return nil, errors.New("materials, grpc handler - invalid company id")
	}

	mtrls, err := mh.service.Material.GetPlanningList(ctx, domain.MaterialParams{
		Limit:     req.Limit,
		Offset:    req.Offset,
		CompanyId: req.CompanyId,
	})
	if err != nil {
		return nil, err
	}

	resp := make([]*materials.Material, 0, len(mtrls))

	for _, mtrl := range mtrls {
		otherFieldsJSON, err := json.Marshal(mtrl.OtherFields)
		if err != nil {
			return nil, err
		}

		resp = append(resp, &materials.Material{
			Id:                     mtrl.ID,
			WarehouseId:            mtrl.WarehouseID,
			ItemId:                 mtrl.ItemID,
			Name:                   mtrl.Name,
			ByInvoice:              mtrl.ByInvoice,
			Article:                mtrl.Article,
			ProductCategory:        mtrl.ProductCategory,
			Unit:                   mtrl.Unit,
			TotalQuantity:          mtrl.TotalQuantity,
			Volume:                 mtrl.Volume,
			PriceWithoutVat:        mtrl.PriceWithoutVAT,
			TotalWithoutVat:        mtrl.TotalWithoutVAT,
			SupplierId:             mtrl.SupplierID,
			Location:               mtrl.Location,
			Contract:               timestamppb.New(mtrl.Contract),
			File:                   mtrl.File,
			Status:                 mtrl.Status,
			Comments:               mtrl.Comments,
			Reserve:                mtrl.Reserve,
			ReceivedDate:           timestamppb.New(mtrl.ReceivedDate),
			LastUpdated:            timestamppb.New(mtrl.LastUpdated),
			MinStockLevel:          mtrl.MinStockLevel,
			ExpirationDate:         timestamppb.New(mtrl.ExpirationDate),
			ResponsiblePerson:      mtrl.ResponsiblePerson,
			StorageCost:            mtrl.StorageCost,
			WarehouseSection:       mtrl.WarehouseSection,
			IncomingDeliveryNumber: mtrl.IncomingDeliveryNumber,
			OtherFields:            string(otherFieldsJSON),
			CompanyId:              mtrl.CompanyID,
		})
	}

	return &materials.MaterialList{
		Materials: resp,
	}, nil
}

func (mh *MaterialsHandler) MovePlanningToPurchased(ctx context.Context, req *materials.MaterialId) (*materials.MaterialId, error) {
	id, itemId, err := mh.service.Material.MovePlanningToPurchased(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &materials.MaterialId{Id: id, ItemId: itemId}, nil
}

func (mh *MaterialsHandler) CreatePurchased(ctx context.Context, material *materials.Material) (*materials.MaterialId, error) {
	var otherFields map[string]interface{}
	if err := json.Unmarshal([]byte(material.OtherFields), &otherFields); err != nil {
		return nil, err
	}

	id, itemID, err := mh.service.Material.CreatePurchased(ctx, domain.Material{
		WarehouseID:            material.WarehouseId,
		ItemID:                 material.ItemId,
		Name:                   material.Name,
		ByInvoice:              material.ByInvoice,
		Article:                material.Article,
		ProductCategory:        material.ProductCategory,
		Unit:                   material.Unit,
		TotalQuantity:          material.TotalQuantity,
		Volume:                 material.Volume,
		PriceWithoutVAT:        material.PriceWithoutVat,
		TotalWithoutVAT:        material.TotalWithoutVat,
		SupplierID:             material.SupplierId,
		Location:               material.Location,
		Contract:               material.Contract.AsTime(),
		File:                   material.File,
		Status:                 material.Status,
		Comments:               material.Comments,
		Reserve:                material.Reserve,
		ReceivedDate:           material.ReceivedDate.AsTime(),
		LastUpdated:            material.LastUpdated.AsTime(),
		MinStockLevel:          material.MinStockLevel,
		ExpirationDate:         material.ExpirationDate.AsTime(),
		ResponsiblePerson:      material.ResponsiblePerson,
		StorageCost:            material.StorageCost,
		WarehouseSection:       material.WarehouseSection,
		IncomingDeliveryNumber: material.IncomingDeliveryNumber,
		OtherFields:            otherFields,
		CompanyID:              material.CompanyId,
	})
	if err != nil {
		return nil, err
	}

	return &materials.MaterialId{Id: id, ItemId: itemID}, nil
}

func (mh *MaterialsHandler) UpdatePurchased(ctx context.Context, material *materials.Material) (*emptypb.Empty, error) {
	var otherFields map[string]interface{}
	if err := json.Unmarshal([]byte(material.OtherFields), &otherFields); err != nil {
		return nil, err
	}

	err := mh.service.Material.UpdatePurchased(ctx, domain.Material{
		ID:                     material.Id,
		WarehouseID:            material.WarehouseId,
		ItemID:                 material.ItemId,
		Name:                   material.Name,
		ByInvoice:              material.ByInvoice,
		Article:                material.Article,
		ProductCategory:        material.ProductCategory,
		Unit:                   material.Unit,
		TotalQuantity:          material.TotalQuantity,
		Volume:                 material.Volume,
		PriceWithoutVAT:        material.PriceWithoutVat,
		TotalWithoutVAT:        material.TotalWithoutVat,
		SupplierID:             material.SupplierId,
		Location:               material.Location,
		Contract:               material.Contract.AsTime(),
		File:                   material.File,
		Status:                 material.Status,
		Comments:               material.Comments,
		Reserve:                material.Reserve,
		ReceivedDate:           material.ReceivedDate.AsTime(),
		LastUpdated:            material.LastUpdated.AsTime(),
		MinStockLevel:          material.MinStockLevel,
		ExpirationDate:         material.ExpirationDate.AsTime(),
		ResponsiblePerson:      material.ResponsiblePerson,
		StorageCost:            material.StorageCost,
		WarehouseSection:       material.WarehouseSection,
		IncomingDeliveryNumber: material.IncomingDeliveryNumber,
		OtherFields:            otherFields,
	})
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (mh *MaterialsHandler) DeletePurchased(ctx context.Context, req *materials.MaterialId) (*emptypb.Empty, error) {
	if err := mh.service.Material.DeletePurchased(ctx, req.Id); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (mh *MaterialsHandler) GetPurchased(ctx context.Context, req *materials.MaterialId) (*materials.Material, error) {
	material, err := mh.service.Material.GetPurchasedById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	otherFieldsJSON, err := json.Marshal(material.OtherFields)
	if err != nil {
		return nil, err
	}

	return &materials.Material{
		Id:                     material.ID,
		WarehouseId:            material.WarehouseID,
		ItemId:                 material.ItemID,
		Name:                   material.Name,
		ByInvoice:              material.ByInvoice,
		Article:                material.Article,
		ProductCategory:        material.ProductCategory,
		Unit:                   material.Unit,
		TotalQuantity:          material.TotalQuantity,
		Volume:                 material.Volume,
		PriceWithoutVat:        material.PriceWithoutVAT,
		TotalWithoutVat:        material.TotalWithoutVAT,
		SupplierId:             material.SupplierID,
		Location:               material.Location,
		Contract:               timestamppb.New(material.Contract),
		File:                   material.File,
		Status:                 material.Status,
		Comments:               material.Comments,
		Reserve:                material.Reserve,
		ReceivedDate:           timestamppb.New(material.ReceivedDate),
		LastUpdated:            timestamppb.New(material.LastUpdated),
		MinStockLevel:          material.MinStockLevel,
		ExpirationDate:         timestamppb.New(material.ExpirationDate),
		ResponsiblePerson:      material.ResponsiblePerson,
		StorageCost:            material.StorageCost,
		WarehouseSection:       material.WarehouseSection,
		IncomingDeliveryNumber: material.IncomingDeliveryNumber,
		OtherFields:            string(otherFieldsJSON),
		CompanyId:              material.CompanyID,
	}, nil
}

func (mh *MaterialsHandler) GetListPurchased(ctx context.Context, req *materials.MaterialParams) (*materials.MaterialList, error) {
	if req.Limit <= 0 {
		return nil, errors.New("materials, grpc handler - invalid limit")
	}

	if req.Offset < 0 {
		return nil, errors.New("materials, grpc handler - invalid offset")
	}

	if req.CompanyId <= 0 {
		return nil, errors.New("materials, grpc handler - invalid company id")
	}

	mtrls, err := mh.service.Material.GetPurchasedList(ctx, domain.MaterialParams{
		Limit:     req.Limit,
		Offset:    req.Offset,
		CompanyId: req.CompanyId,
	})
	if err != nil {
		return nil, err
	}

	resp := make([]*materials.Material, 0, len(mtrls))

	for _, mtrl := range mtrls {
		otherFieldsJSON, err := json.Marshal(mtrl.OtherFields)
		if err != nil {
			return nil, err
		}

		resp = append(resp, &materials.Material{
			Id:                     mtrl.ID,
			WarehouseId:            mtrl.WarehouseID,
			ItemId:                 mtrl.ItemID,
			Name:                   mtrl.Name,
			ByInvoice:              mtrl.ByInvoice,
			Article:                mtrl.Article,
			ProductCategory:        mtrl.ProductCategory,
			Unit:                   mtrl.Unit,
			TotalQuantity:          mtrl.TotalQuantity,
			Volume:                 mtrl.Volume,
			PriceWithoutVat:        mtrl.PriceWithoutVAT,
			TotalWithoutVat:        mtrl.TotalWithoutVAT,
			SupplierId:             mtrl.SupplierID,
			Location:               mtrl.Location,
			Contract:               timestamppb.New(mtrl.Contract),
			File:                   mtrl.File,
			Status:                 mtrl.Status,
			Comments:               mtrl.Comments,
			Reserve:                mtrl.Reserve,
			ReceivedDate:           timestamppb.New(mtrl.ReceivedDate),
			LastUpdated:            timestamppb.New(mtrl.LastUpdated),
			MinStockLevel:          mtrl.MinStockLevel,
			ExpirationDate:         timestamppb.New(mtrl.ExpirationDate),
			ResponsiblePerson:      mtrl.ResponsiblePerson,
			StorageCost:            mtrl.StorageCost,
			WarehouseSection:       mtrl.WarehouseSection,
			IncomingDeliveryNumber: mtrl.IncomingDeliveryNumber,
			OtherFields:            string(otherFieldsJSON),
			CompanyId:              mtrl.CompanyID,
		})
	}

	return &materials.MaterialList{
		Materials: resp,
	}, nil
}

func (mh *MaterialsHandler) MovePurchasedToArchive(ctx context.Context, req *materials.MaterialId) (*emptypb.Empty, error) {
	if err := mh.service.Material.MovePurchasedToArchive(ctx, req.Id); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (mh *MaterialsHandler) GetPlanningArchive(ctx context.Context, req *materials.MaterialId) (*materials.Material, error) {
	material, err := mh.service.Material.GetPlanningArchiveById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	otherFieldsJSON, err := json.Marshal(material.OtherFields)
	if err != nil {
		return nil, err
	}

	return &materials.Material{
		Id:                     material.ID,
		WarehouseId:            material.WarehouseID,
		ItemId:                 material.ItemID,
		Name:                   material.Name,
		ByInvoice:              material.ByInvoice,
		Article:                material.Article,
		ProductCategory:        material.ProductCategory,
		Unit:                   material.Unit,
		TotalQuantity:          material.TotalQuantity,
		Volume:                 material.Volume,
		PriceWithoutVat:        material.PriceWithoutVAT,
		TotalWithoutVat:        material.TotalWithoutVAT,
		SupplierId:             material.SupplierID,
		Location:               material.Location,
		Contract:               timestamppb.New(material.Contract),
		File:                   material.File,
		Status:                 material.Status,
		Comments:               material.Comments,
		Reserve:                material.Reserve,
		ReceivedDate:           timestamppb.New(material.ReceivedDate),
		LastUpdated:            timestamppb.New(material.LastUpdated),
		MinStockLevel:          material.MinStockLevel,
		ExpirationDate:         timestamppb.New(material.ExpirationDate),
		ResponsiblePerson:      material.ResponsiblePerson,
		StorageCost:            material.StorageCost,
		WarehouseSection:       material.WarehouseSection,
		IncomingDeliveryNumber: material.IncomingDeliveryNumber,
		OtherFields:            string(otherFieldsJSON),
		CompanyId:              material.CompanyID,
	}, nil
}

func (mh *MaterialsHandler) GetPurchasedArchive(ctx context.Context, req *materials.MaterialId) (*materials.Material, error) {
	material, err := mh.service.Material.GetPurchasedArchiveById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	otherFieldsJSON, err := json.Marshal(material.OtherFields)
	if err != nil {
		return nil, err
	}

	return &materials.Material{
		Id:                     material.ID,
		WarehouseId:            material.WarehouseID,
		ItemId:                 material.ItemID,
		Name:                   material.Name,
		ByInvoice:              material.ByInvoice,
		Article:                material.Article,
		ProductCategory:        material.ProductCategory,
		Unit:                   material.Unit,
		TotalQuantity:          material.TotalQuantity,
		Volume:                 material.Volume,
		PriceWithoutVat:        material.PriceWithoutVAT,
		TotalWithoutVat:        material.TotalWithoutVAT,
		SupplierId:             material.SupplierID,
		Location:               material.Location,
		Contract:               timestamppb.New(material.Contract),
		File:                   material.File,
		Status:                 material.Status,
		Comments:               material.Comments,
		Reserve:                material.Reserve,
		ReceivedDate:           timestamppb.New(material.ReceivedDate),
		LastUpdated:            timestamppb.New(material.LastUpdated),
		MinStockLevel:          material.MinStockLevel,
		ExpirationDate:         timestamppb.New(material.ExpirationDate),
		ResponsiblePerson:      material.ResponsiblePerson,
		StorageCost:            material.StorageCost,
		WarehouseSection:       material.WarehouseSection,
		IncomingDeliveryNumber: material.IncomingDeliveryNumber,
		OtherFields:            string(otherFieldsJSON),
		CompanyId:              material.CompanyID,
	}, nil
}

func (mh *MaterialsHandler) GetListPlanningArchive(ctx context.Context, req *materials.MaterialParams) (*materials.MaterialList, error) {
	if req.Limit <= 0 {
		return nil, errors.New("materials, grpc handler - invalid limit")
	}

	if req.Offset < 0 {
		return nil, errors.New("materials, grpc handler - invalid offset")
	}

	if req.CompanyId <= 0 {
		return nil, errors.New("materials, grpc handler - invalid company id")
	}

	mtrls, err := mh.service.Material.GetPlanningArchiveList(ctx, domain.MaterialParams{
		Limit:     req.Limit,
		Offset:    req.Offset,
		CompanyId: req.CompanyId,
	})
	if err != nil {
		return nil, err
	}

	resp := make([]*materials.Material, 0, len(mtrls))

	for _, mtrl := range mtrls {
		otherFieldsJSON, err := json.Marshal(mtrl.OtherFields)
		if err != nil {
			return nil, err
		}

		resp = append(resp, &materials.Material{
			Id:                     mtrl.ID,
			WarehouseId:            mtrl.WarehouseID,
			ItemId:                 mtrl.ItemID,
			Name:                   mtrl.Name,
			ByInvoice:              mtrl.ByInvoice,
			Article:                mtrl.Article,
			ProductCategory:        mtrl.ProductCategory,
			Unit:                   mtrl.Unit,
			TotalQuantity:          mtrl.TotalQuantity,
			Volume:                 mtrl.Volume,
			PriceWithoutVat:        mtrl.PriceWithoutVAT,
			TotalWithoutVat:        mtrl.TotalWithoutVAT,
			SupplierId:             mtrl.SupplierID,
			Location:               mtrl.Location,
			Contract:               timestamppb.New(mtrl.Contract),
			File:                   mtrl.File,
			Status:                 mtrl.Status,
			Comments:               mtrl.Comments,
			Reserve:                mtrl.Reserve,
			ReceivedDate:           timestamppb.New(mtrl.ReceivedDate),
			LastUpdated:            timestamppb.New(mtrl.LastUpdated),
			MinStockLevel:          mtrl.MinStockLevel,
			ExpirationDate:         timestamppb.New(mtrl.ExpirationDate),
			ResponsiblePerson:      mtrl.ResponsiblePerson,
			StorageCost:            mtrl.StorageCost,
			WarehouseSection:       mtrl.WarehouseSection,
			IncomingDeliveryNumber: mtrl.IncomingDeliveryNumber,
			OtherFields:            string(otherFieldsJSON),
			CompanyId:              mtrl.CompanyID,
		})
	}

	return &materials.MaterialList{
		Materials: resp,
	}, nil
}

func (mh *MaterialsHandler) GetListPurchasedArchive(ctx context.Context, req *materials.MaterialParams) (*materials.MaterialList, error) {
	if req.Limit <= 0 {
		return nil, errors.New("materials, grpc handler - invalid limit")
	}

	if req.Offset < 0 {
		return nil, errors.New("materials, grpc handler - invalid offset")
	}

	if req.CompanyId <= 0 {
		return nil, errors.New("materials, grpc handler - invalid company id")
	}

	mtrls, err := mh.service.Material.GetPurchasedArchiveList(ctx, domain.MaterialParams{
		Limit:     req.Limit,
		Offset:    req.Offset,
		CompanyId: req.CompanyId,
	})
	if err != nil {
		return nil, err
	}

	resp := make([]*materials.Material, 0, len(mtrls))

	for _, mtrl := range mtrls {
		otherFieldsJSON, err := json.Marshal(mtrl.OtherFields)
		if err != nil {
			return nil, err
		}

		resp = append(resp, &materials.Material{
			Id:                     mtrl.ID,
			WarehouseId:            mtrl.WarehouseID,
			ItemId:                 mtrl.ItemID,
			Name:                   mtrl.Name,
			ByInvoice:              mtrl.ByInvoice,
			Article:                mtrl.Article,
			ProductCategory:        mtrl.ProductCategory,
			Unit:                   mtrl.Unit,
			TotalQuantity:          mtrl.TotalQuantity,
			Volume:                 mtrl.Volume,
			PriceWithoutVat:        mtrl.PriceWithoutVAT,
			TotalWithoutVat:        mtrl.TotalWithoutVAT,
			SupplierId:             mtrl.SupplierID,
			Location:               mtrl.Location,
			Contract:               timestamppb.New(mtrl.Contract),
			File:                   mtrl.File,
			Status:                 mtrl.Status,
			Comments:               mtrl.Comments,
			Reserve:                mtrl.Reserve,
			ReceivedDate:           timestamppb.New(mtrl.ReceivedDate),
			LastUpdated:            timestamppb.New(mtrl.LastUpdated),
			MinStockLevel:          mtrl.MinStockLevel,
			ExpirationDate:         timestamppb.New(mtrl.ExpirationDate),
			ResponsiblePerson:      mtrl.ResponsiblePerson,
			StorageCost:            mtrl.StorageCost,
			WarehouseSection:       mtrl.WarehouseSection,
			IncomingDeliveryNumber: mtrl.IncomingDeliveryNumber,
			OtherFields:            string(otherFieldsJSON),
			CompanyId:              mtrl.CompanyID,
		})
	}

	return &materials.MaterialList{
		Materials: resp,
	}, nil
}

func (mh *MaterialsHandler) DeletePlanningArchive(ctx context.Context, req *materials.MaterialId) (*emptypb.Empty, error) {
	if err := mh.service.Material.DeletePlanningArchive(ctx, req.Id); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (mh *MaterialsHandler) DeletePurchasedArchive(ctx context.Context, req *materials.MaterialId) (*emptypb.Empty, error) {
	if err := mh.service.Material.DeletePurchasedArchive(ctx, req.Id); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (mh *MaterialsHandler) SearchMaterial(ctx context.Context, req *materials.MaterialParams) (*materials.MaterialList, error) {
	if req.Limit <= 0 {
		return nil, errors.New("materials, grpc handler - invalid limit")
	}

	if req.Offset < 0 {
		return nil, errors.New("materials, grpc handler - invalid offset")
	}

	if req.CompanyId <= 0 {
		return nil, errors.New("materials, grpc handler - invalid company id")
	}

	mtrls, err := mh.service.Material.Search(ctx, domain.Param{
		Limit:     req.Limit,
		Offset:    req.Offset,
		CompanyId: req.CompanyId,
		Query:     req.Query,
	})
	if err != nil {
		return nil, err
	}

	resp := make([]*materials.Material, 0, len(mtrls))

	for _, mtrl := range mtrls {
		otherFieldsJSON, err := json.Marshal(mtrl.OtherFields)
		if err != nil {
			return nil, err
		}

		resp = append(resp, &materials.Material{
			Id:                     mtrl.ID,
			WarehouseId:            mtrl.WarehouseID,
			ItemId:                 mtrl.ItemID,
			Name:                   mtrl.Name,
			ByInvoice:              mtrl.ByInvoice,
			Article:                mtrl.Article,
			ProductCategory:        mtrl.ProductCategory,
			Unit:                   mtrl.Unit,
			TotalQuantity:          mtrl.TotalQuantity,
			Volume:                 mtrl.Volume,
			PriceWithoutVat:        mtrl.PriceWithoutVAT,
			TotalWithoutVat:        mtrl.TotalWithoutVAT,
			SupplierId:             mtrl.SupplierID,
			Location:               mtrl.Location,
			Contract:               timestamppb.New(mtrl.Contract),
			File:                   mtrl.File,
			Status:                 mtrl.Status,
			Comments:               mtrl.Comments,
			Reserve:                mtrl.Reserve,
			ReceivedDate:           timestamppb.New(mtrl.ReceivedDate),
			LastUpdated:            timestamppb.New(mtrl.LastUpdated),
			MinStockLevel:          mtrl.MinStockLevel,
			ExpirationDate:         timestamppb.New(mtrl.ExpirationDate),
			ResponsiblePerson:      mtrl.ResponsiblePerson,
			StorageCost:            mtrl.StorageCost,
			WarehouseSection:       mtrl.WarehouseSection,
			IncomingDeliveryNumber: mtrl.IncomingDeliveryNumber,
			OtherFields:            string(otherFieldsJSON),
			CompanyId:              mtrl.CompanyID,
		})
	}

	return &materials.MaterialList{
		Materials: resp,
	}, nil
}

func (mh *MaterialsHandler) CreateMaterialCategory(ctx context.Context, category *materials.MaterialCategory) (*materials.MaterialCategoryId, error) {
	id, err := mh.service.Category.Create(ctx, domain.MaterialCategory{
		Name:        category.Name,
		CompanyID:   category.CompanyId,
		Description: category.Description,
		Slug:        category.Slug,
		CreatedAt:   category.CreatedAt.AsTime(),
		UpdatedAt:   category.UpdatedAt.AsTime(),
		IsActive:    category.IsActive,
		ImgURL:      category.ImgUrl,
	})
	if err != nil {
		return nil, err
	}

	return &materials.MaterialCategoryId{Id: id}, nil
}

func (mh *MaterialsHandler) GetByIdMaterialCategory(ctx context.Context, req *materials.MaterialCategoryId) (*materials.MaterialCategory, error) {
	category, err := mh.service.Category.GetById(ctx, req.Id, req.CompanyId)
	if err != nil {
		return nil, err
	}

	return &materials.MaterialCategory{
		Id:          category.ID,
		Name:        category.Name,
		CompanyId:   category.CompanyID,
		Description: category.Description,
		Slug:        category.Slug,
		CreatedAt:   timestamppb.New(category.CreatedAt),
		UpdatedAt:   timestamppb.New(category.UpdatedAt),
		IsActive:    category.IsActive,
		ImgUrl:      category.ImgURL,
	}, nil
}

func (mh *MaterialsHandler) UpdateMaterialCategory(ctx context.Context, category *materials.MaterialCategory) (*emptypb.Empty, error) {
	if err := mh.service.Category.Update(ctx, domain.MaterialCategory{
		ID:          category.Id,
		Name:        category.Name,
		CompanyID:   category.CompanyId,
		Description: category.Description,
		Slug:        category.Slug,
		CreatedAt:   category.CreatedAt.AsTime(),
		UpdatedAt:   category.UpdatedAt.AsTime(),
		IsActive:    category.IsActive,
		ImgURL:      category.ImgUrl,
	}); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (mh *MaterialsHandler) DeleteMaterialCategory(ctx context.Context, req *materials.MaterialCategoryId) (*emptypb.Empty, error) {
	if err := mh.service.Category.Delete(ctx, req.Id, req.CompanyId); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (mh *MaterialsHandler) GetListMaterialCategory(ctx context.Context, req *materials.MaterialParams) (*materials.MaterialCategoryList, error) {
	if req.Limit <= 0 {
		return nil, errors.New("material categories, grpc handler - invalid limit")
	}

	if req.Offset < 0 {
		return nil, errors.New("material categories, grpc handler - invalid offset")
	}

	if req.CompanyId <= 0 {
		return nil, errors.New("material categories, grpc handler - invalid company id")
	}

	categories, err := mh.service.Category.List(ctx, domain.Param{
		Limit:     req.Limit,
		Offset:    req.Offset,
		CompanyId: req.CompanyId,
		Query:     req.Query,
	})
	if err != nil {
		return nil, err
	}

	resp := make([]*materials.MaterialCategory, 0, len(categories))

	for _, c := range categories {
		resp = append(resp, &materials.MaterialCategory{
			Id:          c.ID,
			Name:        c.Name,
			CompanyId:   c.CompanyID,
			Description: c.Description,
			Slug:        c.Slug,
			CreatedAt:   timestamppb.New(c.CreatedAt),
			UpdatedAt:   timestamppb.New(c.UpdatedAt),
			IsActive:    c.IsActive,
			ImgUrl:      c.ImgURL,
		})
	}

	return &materials.MaterialCategoryList{
		MaterialCategories: resp,
	}, nil
}

func (mh *MaterialsHandler) SearchMaterialCategory(ctx context.Context, req *materials.MaterialParams) (*materials.MaterialCategoryList, error) {
	if req.Limit <= 0 {
		return nil, errors.New("material categories, grpc handler - invalid limit")
	}

	if req.Offset < 0 {
		return nil, errors.New("material categories, grpc handler - invalid offset")
	}

	if req.CompanyId <= 0 {
		return nil, errors.New("material categories, grpc handler - invalid company id")
	}

	categories, err := mh.service.Category.Search(ctx, domain.Param{
		Limit:     req.Limit,
		Offset:    req.Offset,
		CompanyId: req.CompanyId,
		Query:     req.Query,
	})
	if err != nil {
		return nil, err
	}

	resp := make([]*materials.MaterialCategory, 0, len(categories))

	for _, c := range categories {
		resp = append(resp, &materials.MaterialCategory{
			Id:          c.ID,
			Name:        c.Name,
			CompanyId:   c.CompanyID,
			Description: c.Description,
			Slug:        c.Slug,
			CreatedAt:   timestamppb.New(c.CreatedAt),
			UpdatedAt:   timestamppb.New(c.UpdatedAt),
			IsActive:    c.IsActive,
			ImgUrl:      c.ImgURL,
		})
	}

	return &materials.MaterialCategoryList{
		MaterialCategories: resp,
	}, nil
}
