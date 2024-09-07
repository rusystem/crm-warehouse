package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/rusystem/crm-warehouse/internal/service"
	"github.com/rusystem/crm-warehouse/pkg/domain"
	"github.com/rusystem/crm-warehouse/pkg/gen/proto/supplier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (sh *SupplierHandler) GetById(ctx context.Context, id *supplier.SupplierId) (*supplier.Supplier, error) {
	spl, err := sh.service.Supplier.GetById(ctx, id.Id)
	if err != nil {
		if errors.Is(err, domain.ErrSupplierNotFound) {
			return nil, status.Errorf(codes.NotFound, "supplier with ID %d not found", id.Id)
		}

		return nil, status.Errorf(codes.Internal, "internal server error - %v", err)
	}

	otherFieldsJSON, err := json.Marshal(spl.OtherFields)
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
		OtherFields:       string(otherFieldsJSON),
		CompanyId:         spl.CompanyID,
	}, nil
}

func (sh *SupplierHandler) Create(ctx context.Context, spl *supplier.Supplier) (*supplier.SupplierId, error) {
	var otherFields map[string]interface{}
	if err := json.Unmarshal([]byte(spl.OtherFields), &otherFields); err != nil {
		return nil, err
	}

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
		OtherFields:       otherFields,
		CompanyID:         spl.CompanyId,
	})
	if err != nil {
		return nil, err
	}

	return &supplier.SupplierId{Id: id}, nil
}

func (sh *SupplierHandler) Update(ctx context.Context, spl *supplier.Supplier) (*emptypb.Empty, error) {
	var otherFields map[string]interface{}
	if err := json.Unmarshal([]byte(spl.OtherFields), &otherFields); err != nil {
		return nil, err
	}

	if err := sh.service.Supplier.Update(ctx, domain.Supplier{
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
		OtherFields:       otherFields,
		CompanyID:         spl.CompanyId,
	}); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (sh *SupplierHandler) Delete(ctx context.Context, req *supplier.SupplierId) (*emptypb.Empty, error) {
	if err := sh.service.Supplier.Delete(ctx, req.Id); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (sh *SupplierHandler) GetList(ctx context.Context, id *supplier.SupplierCompanyId) (*supplier.SupplierList, error) {
	suppliers, err := sh.service.Supplier.GetListByCompanyId(ctx, id.Id)
	if err != nil {
		return nil, err
	}

	var resp []*supplier.Supplier
	for _, s := range suppliers {
		otherFieldsJSON, err := json.Marshal(s.OtherFields)
		if err != nil {
			return nil, err
		}

		resp = append(resp, &supplier.Supplier{
			Id:                s.ID,
			Name:              s.Name,
			LegalAddress:      s.LegalAddress,
			ActualAddress:     s.ActualAddress,
			WarehouseAddress:  s.WarehouseAddress,
			ContactPerson:     s.ContactPerson,
			Phone:             s.Phone,
			Email:             s.Email,
			Website:           s.Website,
			ContractNumber:    s.ContractNumber,
			ProductCategories: s.ProductCategories,
			PurchaseAmount:    s.PurchaseAmount,
			Balance:           s.Balance,
			ProductTypes:      s.ProductTypes,
			Comments:          s.Comments,
			Files:             s.Files,
			Country:           s.Country,
			Region:            s.Region,
			TaxId:             s.TaxID,
			BankDetails:       s.BankDetails,
			RegistrationDate:  timestamppb.New(s.RegistrationDate),
			PaymentTerms:      s.PaymentTerms,
			IsActive:          s.IsActive,
			OtherFields:       string(otherFieldsJSON),
			CompanyId:         s.CompanyID,
		})
	}

	return &supplier.SupplierList{Suppliers: resp}, nil
}
