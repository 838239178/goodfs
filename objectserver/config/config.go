package config

import (
	"common/datasize"
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	ConfFilePath = "../conf/object-server.yaml"
)

type CacheConfig struct {
	MaxSizeMB     datasize.DataSize `yaml:"max-size-mb" env:"MAX_SIZE_MB" env-default:"128"`
	TTL           time.Duration     `yaml:"ttl" env:"TTL" env-default:"1h"`
	CleanInterval time.Duration     `yaml:"clean-interval" env:"CLEAN_INTERVAL" env-default:"1h"`
	MaxItemSizeMB datasize.DataSize `yaml:"max-item-size-mb" env:"MAX_ITEM_SIZE_MB" env-default:"12"`
}

type Config struct {
	Port         string        `yaml:"port"`
	StoragePath  string        `yaml:"storage-path" env:"STROAGE_PATH" env-default:"objects"`
	TempPath     string        `yaml:"temp-path" env:"TEMP_PATH" env-default:"temp"`
	AmqpAddress  string        `yaml:"amqp-address" env:"AMQP_ADDRESS" env-required:"true"`
	BeatInterval time.Duration `yaml:"beat-interval" env:"BEAT_INTERVAL" env-default:"5s"`
	Cache        CacheConfig   `yaml:"cache" env-prefix:"CACHE"`
	Etcd         EtcdConfig    `yaml:"etcd" env-prefix:"ETCD"`
}

type EtcdConfig struct {
	Endpoint []string `yaml:"endpoint" env:"ENDPOINT" env-required:"true" env-separator:","`
	Username string   `yaml:"username" env:"USERNAME" env-required:"true"`
	Password string   `yaml:"password" env:"PASSWORD" env-required:"true"`
}

func (c *Config) LocalAddr() string {
	hn, e := os.Hostname()
	if e != nil {
		panic(e)
	}
	return fmt.Sprintf("%v:%v", hn, c.Port)
}

func ReadConfig() Config {
	var conf Config
	if err := cleanenv.ReadConfig(ConfFilePath, &conf); err != nil {
		panic(err)
	}
	return conf
}
