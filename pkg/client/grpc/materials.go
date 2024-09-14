package grpc

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/rusystem/crm-warehouse/pkg/domain"
	"github.com/rusystem/crm-warehouse/pkg/gen/proto/materials"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

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
	IncomingDeliveryNumber string                 `json:"incoming_delivery_number"` // Входящий номер поставки
	OtherFields            map[string]interface{} `json:"other_fields"`             // Дополнительные пользовательские поля
	CompanyID              int64                  `json:"company_id"`               // Кабинет компании к кому привязан товар
}

type MaterialParams struct {
	Limit     int64
	Offset    int64
	CompanyId int64
}

type MaterialsClient struct {
	conn            *grpc.ClientConn
	materialsClient materials.MaterialServiceClient
}

func NewMaterialsClient(addr string) (*MaterialsClient, error) {
	opt := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(addr, opt...)
	if err != nil {
		return nil, err
	}

	return &MaterialsClient{
		conn:            conn,
		materialsClient: materials.NewMaterialServiceClient(conn),
	}, nil
}

func (mc *MaterialsClient) Close() error {
	return mc.conn.Close()
}

func (mc *MaterialsClient) CreatePlanning(ctx context.Context, material Material) (int64, error) {
	otherFieldsJSON, err := json.Marshal(material.OtherFields)
	if err != nil {
		return 0, err
	}

	resp, err := mc.materialsClient.CreatePlanning(ctx, &materials.Material{
		WarehouseId:            material.WarehouseID,
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
	})
	if err != nil {
		return 0, err
	}

	return resp.Id, nil
}

func (mc *MaterialsClient) UpdatePlanningById(ctx context.Context, material Material) error {
	otherFieldsJSON, err := json.Marshal(material.OtherFields)
	if err != nil {
		return err
	}

	_, err = mc.materialsClient.UpdatePlanning(ctx, &materials.Material{
		Id:                     material.ID,
		WarehouseId:            material.WarehouseID,
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
	})
	if err != nil {
		return err
	}

	return nil
}

func (mc *MaterialsClient) DeletePlanningById(ctx context.Context, id int64) error {
	_, err := mc.materialsClient.DeletePlanning(ctx, &materials.MaterialId{Id: id})
	return err
}

func (mc *MaterialsClient) GetPlanningById(ctx context.Context, id int64) (Material, error) {
	if id <= 0 {
		return Material{}, errors.New("materials, grpc client - invalid id")
	}

	resp, err := mc.materialsClient.GetPlanning(ctx, &materials.MaterialId{Id: id})
	if err != nil {
		if err.Error() == "rpc error: code = Unknown desc = sql: no rows in result set" {
			return Material{}, domain.ErrMaterialNotFound
		}

		return Material{}, err
	}

	var otherFields map[string]interface{}
	if err = json.Unmarshal([]byte(resp.OtherFields), &otherFields); err != nil {
		return Material{}, err
	}

	return Material{
		ID:                     resp.Id,
		WarehouseID:            resp.WarehouseId,
		ItemID:                 resp.ItemId,
		Name:                   resp.Name,
		ByInvoice:              resp.ByInvoice,
		Article:                resp.Article,
		ProductCategory:        resp.ProductCategory,
		Unit:                   resp.Unit,
		TotalQuantity:          resp.TotalQuantity,
		Volume:                 resp.Volume,
		PriceWithoutVAT:        resp.PriceWithoutVat,
		TotalWithoutVAT:        resp.TotalWithoutVat,
		SupplierID:             resp.SupplierId,
		Location:               resp.Location,
		Contract:               resp.Contract.AsTime(),
		File:                   resp.File,
		Status:                 resp.Status,
		Comments:               resp.Comments,
		Reserve:                resp.Reserve,
		ReceivedDate:           resp.ReceivedDate.AsTime(),
		LastUpdated:            resp.LastUpdated.AsTime(),
		MinStockLevel:          resp.MinStockLevel,
		ExpirationDate:         resp.ExpirationDate.AsTime(),
		ResponsiblePerson:      resp.ResponsiblePerson,
		StorageCost:            resp.StorageCost,
		WarehouseSection:       resp.WarehouseSection,
		IncomingDeliveryNumber: resp.IncomingDeliveryNumber,
		OtherFields:            otherFields,
		CompanyID:              resp.CompanyId,
	}, nil
}

func (mc *MaterialsClient) GetListPlanning(ctx context.Context, params MaterialParams) ([]Material, error) {
	var mtrls []Material

	resp, err := mc.materialsClient.GetListPlanning(ctx, &materials.MaterialParams{
		Limit:     params.Limit,
		Offset:    params.Offset,
		CompanyId: params.CompanyId,
	})
	if err != nil {
		return nil, err
	}

	for _, mtrl := range resp.Materials {
		var otherFields map[string]interface{}
		if err = json.Unmarshal([]byte(mtrl.OtherFields), &otherFields); err != nil {
			return nil, err
		}

		mtrls = append(mtrls, Material{
			ID:                     mtrl.Id,
			WarehouseID:            mtrl.WarehouseId,
			ItemID:                 mtrl.ItemId,
			Name:                   mtrl.Name,
			ByInvoice:              mtrl.ByInvoice,
			Article:                mtrl.Article,
			ProductCategory:        mtrl.ProductCategory,
			Unit:                   mtrl.Unit,
			TotalQuantity:          mtrl.TotalQuantity,
			Volume:                 mtrl.Volume,
			PriceWithoutVAT:        mtrl.PriceWithoutVat,
			TotalWithoutVAT:        mtrl.TotalWithoutVat,
			SupplierID:             mtrl.SupplierId,
			Location:               mtrl.Location,
			Contract:               mtrl.Contract.AsTime(),
			File:                   mtrl.File,
			Status:                 mtrl.Status,
			Comments:               mtrl.Comments,
			Reserve:                mtrl.Reserve,
			ReceivedDate:           mtrl.ReceivedDate.AsTime(),
			LastUpdated:            mtrl.LastUpdated.AsTime(),
			MinStockLevel:          mtrl.MinStockLevel,
			ExpirationDate:         mtrl.ExpirationDate.AsTime(),
			ResponsiblePerson:      mtrl.ResponsiblePerson,
			StorageCost:            mtrl.StorageCost,
			WarehouseSection:       mtrl.WarehouseSection,
			IncomingDeliveryNumber: mtrl.IncomingDeliveryNumber,
			OtherFields:            otherFields,
			CompanyID:              mtrl.CompanyId,
		})
	}

	return mtrls, nil
}

func (mc *MaterialsClient) MovePlanningToPurchased(ctx context.Context, id int64) (int64, int64, error) {
	resp, err := mc.materialsClient.MovePlanningToPurchased(ctx, &materials.MaterialId{Id: id})
	if err != nil {
		return 0, 0, err
	}

	return resp.Id, resp.ItemId, nil
}

func (mc *MaterialsClient) CreatePurchased(ctx context.Context, material Material) (int64, int64, error) {
	otherFieldsJSON, err := json.Marshal(material.OtherFields)
	if err != nil {
		return 0, 0, err
	}

	resp, err := mc.materialsClient.CreatePurchased(ctx, &materials.Material{
		WarehouseId:            material.WarehouseID,
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
	})
	if err != nil {
		return 0, 0, err
	}

	return resp.Id, resp.ItemId, nil
}

func (mc *MaterialsClient) UpdatePurchasedById(ctx context.Context, material Material) error {
	otherFieldsJSON, err := json.Marshal(material.OtherFields)
	if err != nil {
		return err
	}

	_, err = mc.materialsClient.UpdatePurchased(ctx, &materials.Material{
		Id:                     material.ID,
		WarehouseId:            material.WarehouseID,
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
	})
	if err != nil {
		return err
	}

	return nil
}

func (mc *MaterialsClient) DeletePurchasedById(ctx context.Context, id int64) error {
	_, err := mc.materialsClient.DeletePurchased(ctx, &materials.MaterialId{Id: id})
	return err
}

func (mc *MaterialsClient) GetPurchasedById(ctx context.Context, id int64) (Material, error) {
	if id <= 0 {
		return Material{}, errors.New("materials, grpc client - invalid id")
	}

	resp, err := mc.materialsClient.GetPurchased(ctx, &materials.MaterialId{Id: id})
	if err != nil {
		if err.Error() == "rpc error: code = Unknown desc = sql: no rows in result set" {
			return Material{}, domain.ErrMaterialNotFound
		}

		return Material{}, err
	}

	var otherFields map[string]interface{}
	if err = json.Unmarshal([]byte(resp.OtherFields), &otherFields); err != nil {
		return Material{}, err
	}

	return Material{
		ID:                     resp.Id,
		WarehouseID:            resp.WarehouseId,
		ItemID:                 resp.ItemId,
		Name:                   resp.Name,
		ByInvoice:              resp.ByInvoice,
		Article:                resp.Article,
		ProductCategory:        resp.ProductCategory,
		Unit:                   resp.Unit,
		TotalQuantity:          resp.TotalQuantity,
		Volume:                 resp.Volume,
		PriceWithoutVAT:        resp.PriceWithoutVat,
		TotalWithoutVAT:        resp.TotalWithoutVat,
		SupplierID:             resp.SupplierId,
		Location:               resp.Location,
		Contract:               resp.Contract.AsTime(),
		File:                   resp.File,
		Status:                 resp.Status,
		Comments:               resp.Comments,
		Reserve:                resp.Reserve,
		ReceivedDate:           resp.ReceivedDate.AsTime(),
		LastUpdated:            resp.LastUpdated.AsTime(),
		MinStockLevel:          resp.MinStockLevel,
		ExpirationDate:         resp.ExpirationDate.AsTime(),
		ResponsiblePerson:      resp.ResponsiblePerson,
		StorageCost:            resp.StorageCost,
		WarehouseSection:       resp.WarehouseSection,
		IncomingDeliveryNumber: resp.IncomingDeliveryNumber,
		OtherFields:            otherFields,
		CompanyID:              resp.CompanyId,
	}, nil
}

func (mc *MaterialsClient) GetListPurchased(ctx context.Context, params MaterialParams) ([]Material, error) {
	var mtrls []Material

	resp, err := mc.materialsClient.GetListPurchased(ctx, &materials.MaterialParams{
		Limit:     params.Limit,
		Offset:    params.Offset,
		CompanyId: params.CompanyId,
	})
	if err != nil {
		return nil, err
	}

	for _, mtrl := range resp.Materials {
		var otherFields map[string]interface{}
		if err = json.Unmarshal([]byte(mtrl.OtherFields), &otherFields); err != nil {
			return nil, err
		}

		mtrls = append(mtrls, Material{
			ID:                     mtrl.Id,
			WarehouseID:            mtrl.WarehouseId,
			ItemID:                 mtrl.ItemId,
			Name:                   mtrl.Name,
			ByInvoice:              mtrl.ByInvoice,
			Article:                mtrl.Article,
			ProductCategory:        mtrl.ProductCategory,
			Unit:                   mtrl.Unit,
			TotalQuantity:          mtrl.TotalQuantity,
			Volume:                 mtrl.Volume,
			PriceWithoutVAT:        mtrl.PriceWithoutVat,
			TotalWithoutVAT:        mtrl.TotalWithoutVat,
			SupplierID:             mtrl.SupplierId,
			Location:               mtrl.Location,
			Contract:               mtrl.Contract.AsTime(),
			File:                   mtrl.File,
			Status:                 mtrl.Status,
			Comments:               mtrl.Comments,
			Reserve:                mtrl.Reserve,
			ReceivedDate:           mtrl.ReceivedDate.AsTime(),
			LastUpdated:            mtrl.LastUpdated.AsTime(),
			MinStockLevel:          mtrl.MinStockLevel,
			ExpirationDate:         mtrl.ExpirationDate.AsTime(),
			ResponsiblePerson:      mtrl.ResponsiblePerson,
			StorageCost:            mtrl.StorageCost,
			WarehouseSection:       mtrl.WarehouseSection,
			IncomingDeliveryNumber: mtrl.IncomingDeliveryNumber,
			OtherFields:            otherFields,
			CompanyID:              mtrl.CompanyId,
		})
	}

	return mtrls, nil
}

func (mc *MaterialsClient) MovePurchasedToArchive(ctx context.Context, id int64) error {
	_, err := mc.materialsClient.MovePurchasedToArchive(ctx, &materials.MaterialId{Id: id})
	return err
}

func (mc *MaterialsClient) GetPlanningArchiveById(ctx context.Context, id int64) (Material, error) {
	if id <= 0 {
		return Material{}, errors.New("materials, grpc client - invalid id")
	}

	resp, err := mc.materialsClient.GetPlanningArchive(ctx, &materials.MaterialId{Id: id})
	if err != nil {
		if err.Error() == "rpc error: code = Unknown desc = sql: no rows in result set" {
			return Material{}, domain.ErrMaterialNotFound
		}

		return Material{}, err
	}

	var otherFields map[string]interface{}
	if err = json.Unmarshal([]byte(resp.OtherFields), &otherFields); err != nil {
		return Material{}, err
	}

	return Material{
		ID:                     resp.Id,
		WarehouseID:            resp.WarehouseId,
		ItemID:                 resp.ItemId,
		Name:                   resp.Name,
		ByInvoice:              resp.ByInvoice,
		Article:                resp.Article,
		ProductCategory:        resp.ProductCategory,
		Unit:                   resp.Unit,
		TotalQuantity:          resp.TotalQuantity,
		Volume:                 resp.Volume,
		PriceWithoutVAT:        resp.PriceWithoutVat,
		TotalWithoutVAT:        resp.TotalWithoutVat,
		SupplierID:             resp.SupplierId,
		Location:               resp.Location,
		Contract:               resp.Contract.AsTime(),
		File:                   resp.File,
		Status:                 resp.Status,
		Comments:               resp.Comments,
		Reserve:                resp.Reserve,
		ReceivedDate:           resp.ReceivedDate.AsTime(),
		LastUpdated:            resp.LastUpdated.AsTime(),
		MinStockLevel:          resp.MinStockLevel,
		ExpirationDate:         resp.ExpirationDate.AsTime(),
		ResponsiblePerson:      resp.ResponsiblePerson,
		StorageCost:            resp.StorageCost,
		WarehouseSection:       resp.WarehouseSection,
		IncomingDeliveryNumber: resp.IncomingDeliveryNumber,
		OtherFields:            otherFields,
		CompanyID:              resp.CompanyId,
	}, nil
}

func (mc *MaterialsClient) GetPurchasedArchiveById(ctx context.Context, id int64) (Material, error) {
	if id <= 0 {
		return Material{}, errors.New("materials, grpc client - invalid id")
	}

	resp, err := mc.materialsClient.GetPurchasedArchive(ctx, &materials.MaterialId{Id: id})
	if err != nil {
		if err.Error() == "rpc error: code = Unknown desc = sql: no rows in result set" {
			return Material{}, domain.ErrMaterialNotFound
		}

		return Material{}, err
	}

	var otherFields map[string]interface{}
	if err = json.Unmarshal([]byte(resp.OtherFields), &otherFields); err != nil {
		return Material{}, err
	}

	return Material{
		ID:                     resp.Id,
		WarehouseID:            resp.WarehouseId,
		ItemID:                 resp.ItemId,
		Name:                   resp.Name,
		ByInvoice:              resp.ByInvoice,
		Article:                resp.Article,
		ProductCategory:        resp.ProductCategory,
		Unit:                   resp.Unit,
		TotalQuantity:          resp.TotalQuantity,
		Volume:                 resp.Volume,
		PriceWithoutVAT:        resp.PriceWithoutVat,
		TotalWithoutVAT:        resp.TotalWithoutVat,
		SupplierID:             resp.SupplierId,
		Location:               resp.Location,
		Contract:               resp.Contract.AsTime(),
		File:                   resp.File,
		Status:                 resp.Status,
		Comments:               resp.Comments,
		Reserve:                resp.Reserve,
		ReceivedDate:           resp.ReceivedDate.AsTime(),
		LastUpdated:            resp.LastUpdated.AsTime(),
		MinStockLevel:          resp.MinStockLevel,
		ExpirationDate:         resp.ExpirationDate.AsTime(),
		ResponsiblePerson:      resp.ResponsiblePerson,
		StorageCost:            resp.StorageCost,
		WarehouseSection:       resp.WarehouseSection,
		IncomingDeliveryNumber: resp.IncomingDeliveryNumber,
		OtherFields:            otherFields,
		CompanyID:              resp.CompanyId,
	}, nil
}

func (mc *MaterialsClient) GetListPlanningArchive(ctx context.Context, params MaterialParams) ([]Material, error) {
	var mtrls []Material

	resp, err := mc.materialsClient.GetListPlanningArchive(ctx, &materials.MaterialParams{
		Limit:     params.Limit,
		Offset:    params.Offset,
		CompanyId: params.CompanyId,
	})
	if err != nil {
		return nil, err
	}

	for _, mtrl := range resp.Materials {
		var otherFields map[string]interface{}
		if err = json.Unmarshal([]byte(mtrl.OtherFields), &otherFields); err != nil {
			return nil, err
		}

		mtrls = append(mtrls, Material{
			ID:                     mtrl.Id,
			WarehouseID:            mtrl.WarehouseId,
			ItemID:                 mtrl.ItemId,
			Name:                   mtrl.Name,
			ByInvoice:              mtrl.ByInvoice,
			Article:                mtrl.Article,
			ProductCategory:        mtrl.ProductCategory,
			Unit:                   mtrl.Unit,
			TotalQuantity:          mtrl.TotalQuantity,
			Volume:                 mtrl.Volume,
			PriceWithoutVAT:        mtrl.PriceWithoutVat,
			TotalWithoutVAT:        mtrl.TotalWithoutVat,
			SupplierID:             mtrl.SupplierId,
			Location:               mtrl.Location,
			Contract:               mtrl.Contract.AsTime(),
			File:                   mtrl.File,
			Status:                 mtrl.Status,
			Comments:               mtrl.Comments,
			Reserve:                mtrl.Reserve,
			ReceivedDate:           mtrl.ReceivedDate.AsTime(),
			LastUpdated:            mtrl.LastUpdated.AsTime(),
			MinStockLevel:          mtrl.MinStockLevel,
			ExpirationDate:         mtrl.ExpirationDate.AsTime(),
			ResponsiblePerson:      mtrl.ResponsiblePerson,
			StorageCost:            mtrl.StorageCost,
			WarehouseSection:       mtrl.WarehouseSection,
			IncomingDeliveryNumber: mtrl.IncomingDeliveryNumber,
			OtherFields:            otherFields,
			CompanyID:              mtrl.CompanyId,
		})
	}

	return mtrls, nil
}

func (mc *MaterialsClient) GetListPurchasedArchive(ctx context.Context, params MaterialParams) ([]Material, error) {
	var mtrls []Material

	resp, err := mc.materialsClient.GetListPurchasedArchive(ctx, &materials.MaterialParams{
		Limit:     params.Limit,
		Offset:    params.Offset,
		CompanyId: params.CompanyId,
	})
	if err != nil {
		return nil, err
	}

	for _, mtrl := range resp.Materials {
		var otherFields map[string]interface{}
		if err = json.Unmarshal([]byte(mtrl.OtherFields), &otherFields); err != nil {
			return nil, err
		}

		mtrls = append(mtrls, Material{
			ID:                     mtrl.Id,
			WarehouseID:            mtrl.WarehouseId,
			ItemID:                 mtrl.ItemId,
			Name:                   mtrl.Name,
			ByInvoice:              mtrl.ByInvoice,
			Article:                mtrl.Article,
			ProductCategory:        mtrl.ProductCategory,
			Unit:                   mtrl.Unit,
			TotalQuantity:          mtrl.TotalQuantity,
			Volume:                 mtrl.Volume,
			PriceWithoutVAT:        mtrl.PriceWithoutVat,
			TotalWithoutVAT:        mtrl.TotalWithoutVat,
			SupplierID:             mtrl.SupplierId,
			Location:               mtrl.Location,
			Contract:               mtrl.Contract.AsTime(),
			File:                   mtrl.File,
			Status:                 mtrl.Status,
			Comments:               mtrl.Comments,
			Reserve:                mtrl.Reserve,
			ReceivedDate:           mtrl.ReceivedDate.AsTime(),
			LastUpdated:            mtrl.LastUpdated.AsTime(),
			MinStockLevel:          mtrl.MinStockLevel,
			ExpirationDate:         mtrl.ExpirationDate.AsTime(),
			ResponsiblePerson:      mtrl.ResponsiblePerson,
			StorageCost:            mtrl.StorageCost,
			WarehouseSection:       mtrl.WarehouseSection,
			IncomingDeliveryNumber: mtrl.IncomingDeliveryNumber,
			OtherFields:            otherFields,
			CompanyID:              mtrl.CompanyId,
		})
	}

	return mtrls, nil
}

func (mc *MaterialsClient) DeletePlanningArchiveById(ctx context.Context, id int64) error {
	_, err := mc.materialsClient.DeletePlanningArchive(ctx, &materials.MaterialId{Id: id})
	return err
}

func (mc *MaterialsClient) DeletePurchasedArchiveById(ctx context.Context, id int64) error {
	_, err := mc.materialsClient.DeletePurchasedArchive(ctx, &materials.MaterialId{Id: id})
	return err
}
