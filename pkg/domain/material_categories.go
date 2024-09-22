package domain

import (
	"time"
)

type MaterialCategory struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	CompanyID   int64     `json:"company_id"`
	Description string    `json:"description"`
	Slug        string    `json:"slug"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	IsActive    bool      `json:"is_active"`
	ImgURL      string    `json:"img_url"`
}
