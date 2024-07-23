package memcache

import (
	"encoding/json"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/rusystem/crm-warehouse/internal/config"
	"github.com/rusystem/crm-warehouse/pkg/domain"
)

type Warehouse interface {
	GetById(id int64) (domain.Warehouse, error)
	AddById(warehouse domain.Warehouse) error
}

type WarehouseMemcacheRepository struct {
	cfg   *config.Config
	cache *memcache.Client
}

func NewWarehouseMemcacheRepository(cfg *config.Config, cache *memcache.Client) *WarehouseMemcacheRepository {
	return &WarehouseMemcacheRepository{
		cfg:   cfg,
		cache: cache,
	}
}

func (wmr *WarehouseMemcacheRepository) GetById(id int64) (domain.Warehouse, error) {
	item, err := wmr.cache.Get(fmt.Sprintf("warehouse_%d", id))
	if err != nil {
		return domain.Warehouse{}, err
	}

	var warehouse domain.Warehouse
	if err = json.Unmarshal(item.Value, &warehouse); err != nil {
		return domain.Warehouse{}, fmt.Errorf("failed to unmarshal warehouse data: %v", err)
	}

	return warehouse, nil
}

func (wmr *WarehouseMemcacheRepository) AddById(warehouse domain.Warehouse) error {
	data, err := json.Marshal(warehouse)
	if err != nil {
		return fmt.Errorf("failed to marshal warehouse data: %v", err)
	}

	return wmr.cache.Set(&memcache.Item{
		Key:   fmt.Sprintf("warehouse_%d", warehouse.ID),
		Value: data,
	})
}
