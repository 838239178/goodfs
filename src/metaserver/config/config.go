package config

import (
	"common/datasize"
	"common/etcd"
	"common/logs"
	"common/registry"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	ConfFilePath = "../../conf/admin-server.yaml"
)

type Config struct {
	Port     string          `yaml:"port" env:"PORT" env-default:"8090"`
	RpcPort  string          `yaml:"rpc-port" env:"RPC_PORT" env-default:"4090"`
	LogLevel logs.Level      `yaml:"log-level" env:"LOG_LEVEL"`
	DataDir  string          `yaml:"data-dir" env:"DATA_DIR" env-default:"/tmp/goodfs"`
	Cluster  ClusterConfig   `yaml:"cluster" env-prefix:"CLUSTER"`
	Registry registry.Config `yaml:"registry" env-prefix:"REGISTRY"`
	Etcd     etcd.Config     `yaml:"etcd" env-prefix:"ETCD"`
	HashSlot HashSlotConfig  `yaml:"hash-slot" env-prefix:"HASH_SLOT"`
	Cache    CacheConfig     `yaml:"cache" env-prefix:"CACHE"`
}

func (c *Config) initialize() {
	c.Cluster.Port = c.RpcPort
	if c.Cluster.Enable {
		c.Cluster.ID = c.Registry.ServerID
		c.HashSlot.StoreID = c.Cluster.GroupID
	} else {
		c.HashSlot.StoreID = c.Registry.ServerID
	}
}

type CacheConfig struct {
	TTL           time.Duration     `yaml:"ttl" env:"TTL" env-default:"1h"`
	CleanInterval time.Duration     `yaml:"clean-interval" env:"CLEAN_INTERVAL" env-default:"1h"`
	MaxSize       datasize.DataSize `yaml:"max-size" env:"MAX_SIZE" env-default:"128MB"`
}

type HashSlotConfig struct {
	StoreID        string        `yaml:"-" env:"-"`
	Slots          []string      `yaml:"slots" env-separator:"," env-default:"0-16384"`
	PrepareTimeout time.Duration `yaml:"prepare-timeout" env-default:"10s"`
}

type ClusterConfig struct {
	Enable           bool          `yaml:"enable" env:"ENABLE" env-default:"false"`
	Bootstrap        bool          `yaml:"bootstrap" env:"BOOTSTRAP" env-default:"false"`
	ID               string        `yaml:"-" env:"-"`
	GroupID          string        `yaml:"group-id" env:"GROUP_ID" env-required:"true"`
	Port             string        `yaml:"-" env:"-"`
	LogLevel         string        `yaml:"log-level" env:"LOG_LEVEL" env-default:"INFO"`
	StoreDir         string        `yaml:"store-dir" env:"STORE_DIR" env-default:"/tmp/goodfs_metaserver"`
	ElectionTimeout  time.Duration `yaml:"election-timeout" env:"ELECTION_TIMEOUT" env-default:"900ms"`
	HeartbeatTimeout time.Duration `yaml:"heartbeat-timeout" env:"HEARTBEAT_TIMEOUT" env-default:"800ms"`
	Nodes            []string      `yaml:"nodes" env:"NODES" env-separator:","`
}

func ReadConfig() Config {
	var conf Config
	if err := cleanenv.ReadConfig(ConfFilePath, &conf); err != nil {
		panic(err)
	}
	logs.Std().Infof("read config from %s", ConfFilePath)
	conf.initialize()
	return conf
}

func ReadConfigFrom(path string) Config {
	var conf Config
	if err := cleanenv.ReadConfig(path, &conf); err != nil {
		if os.IsNotExist(err) {
			return ReadConfig()
		}
		panic(err)
	}
	logs.Std().Infof("read config from %s", path)
	conf.initialize()
	return conf
}