package memcache

import (
	"encoding/json"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/rusystem/crm-warehouse/internal/config"
	"github.com/rusystem/crm-warehouse/pkg/domain"
)

type Suppliers interface {
	GetById(id int64) (domain.Supplier, error)
	AddById(supplier domain.Supplier) error
}

type SuppliersMemcacheRepository struct {
	cfg   *config.Config
	cache *memcache.Client
}

func NewSuppliersMemcacheRepository(cfg *config.Config, cache *memcache.Client) *SuppliersMemcacheRepository {
	return &SuppliersMemcacheRepository{
		cfg:   cfg,
		cache: cache,
	}
}

func (smr *SuppliersMemcacheRepository) GetById(id int64) (domain.Supplier, error) {
	item, err := smr.cache.Get(fmt.Sprintf("supplier_%d", id))
	if err != nil {
		return domain.Supplier{}, err
	}

	var supplier domain.Supplier
	if err = json.Unmarshal(item.Value, &supplier); err != nil {
		return domain.Supplier{}, fmt.Errorf("failed to unmarshal supplier data: %v", err)
	}

	return supplier, nil
}

func (smr *SuppliersMemcacheRepository) AddById(supplier domain.Supplier) error {
	data, err := json.Marshal(supplier)
	if err != nil {
		return fmt.Errorf("failed to marshal supplier data: %v", err)
	}

	return smr.cache.Set(&memcache.Item{
		Key:   fmt.Sprintf("supplier_%d", supplier.ID),
		Value: data,
	})
}
