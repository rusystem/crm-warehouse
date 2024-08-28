package domain

import "errors"

var (
	ErrEmptyId           = errors.New("id can`t be zero")
	ErrWarehouseNotFound = errors.New("warehouse not found")
	ErrSupplierNotFound  = errors.New("supplier not found")
)
