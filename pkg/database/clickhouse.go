package database

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/rusystem/crm-warehouse/internal/config"
	"net"
	"time"
)

type Clickhouse struct {
	cfg  *config.Config
	conn clickhouse.Conn
}

func NewClickhouse(cfg *config.Config) *Clickhouse {
	return &Clickhouse{cfg: cfg}
}

func (c *Clickhouse) Init() (clickhouse.Conn, error) {
	dialCount := 0

	var err error
	c.conn, err = clickhouse.Open(&clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%d", c.cfg.Clickhouse.Host, c.cfg.Clickhouse.Port)},
		Auth: clickhouse.Auth{
			Database: c.cfg.Clickhouse.Database,
			Username: c.cfg.Clickhouse.Username,
			Password: c.cfg.Clickhouse.Password,
		},
		DialContext: func(ctx context.Context, addr string) (net.Conn, error) {
			dialCount++
			var d net.Dialer
			return d.DialContext(ctx, "tcp", addr)
		},
		Debug: false, //if you want to debug change Debug field onto "true"
		Debugf: func(format string, v ...interface{}) {
			fmt.Printf(format, v)
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 300,
		},
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		DialTimeout:      time.Duration(10) * time.Second,
		MaxOpenConns:     300,
		MaxIdleConns:     150,
		ConnMaxLifetime:  time.Duration(10) * time.Minute,
		ConnOpenStrategy: clickhouse.ConnOpenInOrder,
	})
	if err != nil {
		return nil, err
	}
	if err = c.conn.Ping(context.Background()); err != nil {
		return nil, err
	}

	return c.conn, err
}
