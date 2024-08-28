package database

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/rusystem/crm-warehouse/internal/config"
)

func NewMemcache(cfg *config.Config) (*memcache.Client, error) {
	// init memory cache
	mc := memcache.New(fmt.Sprintf("%s:%d", cfg.Memcache.Host, cfg.Memcache.Port))

	// ping memcache
	if err := mc.Ping(); err != nil {
		return nil, err
	}

	return mc, nil
}
