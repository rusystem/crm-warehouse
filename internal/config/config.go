package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
	"time"
)

const (
	CONFIG_DIR       = "configs"
	CONFIG_PROD_FILE = "prod"
	CONFIG_DEV_FILE  = "dev"
)

type Config struct {
	Postgres Postgres
	Nats     Nats
	IsProd   bool

	Grpc struct {
		Port int64 `mapstructure:"port"`
	} `mapstructure:"grpc"`

	Cache struct {
		Ttl int64 `mapstructure:"ttl"`
	} `mapstructure:"cache"`

	Ctx struct {
		Ttl time.Duration `mapstructure:"ttl"`
	} `mapstructure:"ctx"`
}

type Postgres struct {
	Host     string
	Port     string
	DBName   string
	User     string
	Password string
	SSLMode  string
}

type Nats struct {
	Address        string
	TotalWait      time.Duration `mapstructure:"total_wait"`
	ReconnectDelay time.Duration `mapstructure:"reconnect_delay"`
	Timeout        time.Duration `mapstructure:"timeout"`
}

type Telegram struct {
	BotToken string `vault:"telegram_bot_token"`
	ChatId   string `vault:"telegram_chat_id"`
}

func New(isProd bool) (*Config, error) {
	cfg := new(Config)

	viper.AddConfigPath(CONFIG_DIR)
	viper.SetConfigName(CONFIG_DEV_FILE)

	if isProd {
		viper.SetConfigName(CONFIG_PROD_FILE)
		cfg.IsProd = true
	}

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	if err := envconfig.Process("postgres", &cfg.Postgres); err != nil {
		return nil, err
	}

	if err := envconfig.Process("nats", &cfg.Nats); err != nil {
		return nil, err
	}

	return cfg, nil
}
