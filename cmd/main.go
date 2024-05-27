package cmd

import (
	"database/sql"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/rusystem/crm-warehouse/internal/config"
	"github.com/rusystem/crm-warehouse/pkg/database"
	"github.com/rusystem/crm-warehouse/pkg/logger"
	"github.com/rusystem/crm-warehouse/pkg/mq"
	"github.com/rusystem/crm-warehouse/pkg/telegram"
	"os"
	"os/signal"
	"syscall"
)

// init logger
func init() {
	logger.ZapLoggerInit()
}

func main() {
	// init configs
	cfg, err := config.New(true)
	if err != nil {
		logger.Fatal(fmt.Sprintf("failed to initialize config, err: %v", err))
	}

	// init telegram bot
	tg, err := telegram.NewTelegram(cfg)
	if err != nil {
		logger.Fatal(fmt.Sprintf("failed to initialize telegram bot, err - %v", err))
	}

	// init memory cache
	mc := memcache.New(fmt.Sprintf("%s:%d", cfg.Memcache.Host, cfg.Memcache.Port))

	// init nats
	nc, err := mq.NewNats(cfg)
	if err != nil {
		logger.Fatal(fmt.Sprintf("an error occurred when try to connect to nats, err - %v\n", err))
	}
	defer nc.Close()

	// init postgres connection
	pc, err := database.NewPostgresConnection(database.PostgresConfig{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		Username: cfg.Postgres.User,
		Password: cfg.Postgres.Password,
		DBName:   cfg.Postgres.DBName,
		SSLMode:  cfg.Postgres.SSLMode,
	})
	defer func(pc *sql.DB) {
		if err = pc.Close(); err != nil {
			logger.Error(fmt.Sprintf("postgres: failed to close connection, err: %v", err.Error()))
		}
	}(pc)

	// init clickhouse connection
	ch := database.NewClickhouse(cfg)
	cdb, err := ch.Init()
	if err != nil {
		logger.Fatal(fmt.Sprintf("clickhouse: failed to connect: %v", err.Error()))
	}
	defer func(cdb clickhouse.Conn) {
		if err = cdb.Close(); err != nil {
			logger.Fatal(fmt.Sprintf("failed to close connection, err: %v", err.Error()))
		}
	}(cdb)

	// init dep-s
	tlRepo := tarantool_repo.New(cfg, tc)
	chRepo := clickhouse_repo.New(cfg, cdb)
	r := repository.New(chRepo, tlRepo)
	s := service.New(r, cfg, memCache)
	h := transport.New(s)

	//init and start grpc server
	grpcSrv := server.NewGrpcServer()
	go func() {
		if err := grpcSrv.Run(cfg.ServerGrpc.Port); err != nil {
			logger.Fatal(fmt.Sprintf("failed to start grpc server, err: %v", err))
		}
	}()
	defer grpcSrv.Stop()

	logger.Info("clickhouse-service started")

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	osSignal := <-quit

	logger.Info(fmt.Sprintf("program shutdown... call_type: %v", osSignal))
}
