package handler

import (
	"context"
	"github.com/rusystem/crm-warehouse/internal/service"
	"github.com/rusystem/crm-warehouse/pkg/domain"
	"github.com/rusystem/crm-warehouse/pkg/gen/proto/supplier"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SupplierHandler struct {
	service *service.Service
}

func NewSupplierHandler(service *service.Service) *SupplierHandler {
	return &SupplierHandler{
		service: service,
	}
}

func (sh *SupplierHandler) GetById(ctx context.Context, id *supplier.Id) (*supplier.Supplier, error) {
	spl, err := sh.service.Supplier.GetById(ctx, id.Id)
	if err != nil {
		return nil, err
	}

	return &supplier.Supplier{
		Id:                spl.ID,
		Name:              spl.Name,
		LegalAddress:      spl.LegalAddress,
		ActualAddress:     spl.ActualAddress,
		WarehouseAddress:  spl.WarehouseAddress,
		ContactPerson:     spl.ContactPerson,
		Phone:             spl.Phone,
		Email:             spl.Email,
		Website:           spl.Website,
		ContractNumber:    spl.ContractNumber,
		ProductCategories: spl.ProductCategories,
		PurchaseAmount:    spl.PurchaseAmount,
		Balance:           spl.Balance,
		ProductTypes:      spl.ProductTypes,
		Comments:          spl.Comments,
		Files:             spl.Files,
		Country:           spl.Country,
		Region:            spl.Region,
		TaxId:             spl.TaxID,
		BankDetails:       spl.BankDetails,
		RegistrationDate:  timestamppb.New(spl.RegistrationDate),
		PaymentTerms:      spl.PaymentTerms,
		IsActive:          spl.IsActive,
		OtherFields:       spl.OtherFields,
	}, nil
}

func (sh *SupplierHandler) Create(ctx context.Context, spl *supplier.Supplier) (*supplier.Id, error) {
	id, err := sh.service.Supplier.Create(ctx, domain.Supplier{
		ID:                spl.Id,
		Name:              spl.Name,
		LegalAddress:      spl.LegalAddress,
		ActualAddress:     spl.ActualAddress,
		WarehouseAddress:  spl.WarehouseAddress,
		ContactPerson:     spl.ContactPerson,
		Phone:             spl.Phone,
		Email:             spl.Email,
		Website:           spl.Website,
		ContractNumber:    spl.ContractNumber,
		ProductCategories: spl.ProductCategories,
		PurchaseAmount:    spl.PurchaseAmount,
		Balance:           spl.Balance,
		ProductTypes:      spl.ProductTypes,
		Comments:          spl.Comments,
		Files:             spl.Files,
		Country:           spl.Country,
		Region:            spl.Region,
		TaxID:             spl.TaxId,
		BankDetails:       spl.BankDetails,
		RegistrationDate:  spl.RegistrationDate.AsTime(),
		PaymentTerms:      spl.PaymentTerms,
		IsActive:          spl.IsActive,
		OtherFields:       spl.OtherFields,
	})
	if err != nil {
		return nil, err
	}

	return &supplier.Id{Id: id}, nil
}
