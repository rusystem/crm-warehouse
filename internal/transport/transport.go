package transport

import (
	"github.com/rusystem/crm-warehouse/internal/service"
	"github.com/rusystem/crm-warehouse/internal/transport/handler"
)

type Handler struct {
	Warehouse *handler.WarehouseHandler
	Supplier  *handler.SupplierHandler
	Materials *handler.MaterialsHandler
}

func New(service *service.Service) *Handler {
	return &Handler{
		Warehouse: handler.NewWarehouseHandler(service),
		Supplier:  handler.NewSupplierHandler(service),
		Materials: handler.NewMaterialsHandler(service),
	}
}
