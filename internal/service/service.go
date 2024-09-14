package service

import (
	"github.com/nats-io/nats.go"
	"github.com/rusystem/crm-warehouse/internal/repository"
)

type Service struct {
	Supplier  Supplier
	Warehouse Warehouse
	Material  Material
}

func New(repo *repository.Repository, nc *nats.Conn) *Service {
	return &Service{
		Supplier:  NewSupplierService(repo),
		Warehouse: NewWarehouseService(repo),
		Material:  NewMaterialService(repo),
	}
}
