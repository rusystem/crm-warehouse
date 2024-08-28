package mq

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/rusystem/crm-warehouse/internal/config"
	"github.com/rusystem/crm-warehouse/pkg/logger"
)

func NewNats(cfg *config.Config) (*nats.Conn, error) {
	opts := nats.GetDefaultOptions()
	opts.Name = "NATS Sample Subscriber"
	opts.Timeout = cfg.Nats.Timeout              // 300 секунд
	opts.ReconnectWait = cfg.Nats.ReconnectDelay // 2 секунды
	opts.AllowReconnect = true
	opts.MaxReconnect = -1
	opts.Url = fmt.Sprintf("%s", cfg.Nats.Address)

	opts.DisconnectedErrCB = func(nc *nats.Conn, err error) {
		logger.Warn(fmt.Sprintf("NATS disconnected, will attempt reconnects for %v, err: %v",
			cfg.Nats.ReconnectDelay.Minutes(), err))
	}
	opts.ReconnectedCB = func(nc *nats.Conn) {
		logger.Warn(fmt.Sprintf("NATS reconnected, url: %v", nc.ConnectedUrl()))
	}
	opts.ClosedCB = func(nc *nats.Conn) {
		logger.Warn(fmt.Sprintf("NATS closed, last err: %v", nc.LastError()))
	}

	nc, err := opts.Connect()
	if err != nil {
		return nil, err
	}

	return nc, err
}
