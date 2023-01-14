package config

import (
	"common/cst"
	"common/datasize"
	"common/etcd"
	"common/logs"
	"common/registry"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	ConfFilePath = "../../conf/meta-server.yaml"
)

type Config struct {
	Port        string          `yaml:"port" env:"PORT" env-default:"8090"`
	RpcPort     string          `yaml:"rpc-port" env:"RPC_PORT" env-default:"4090"`
	LogLevel    logs.Level      `yaml:"log-level" env:"LOG_LEVEL"`
	DataDir     string          `yaml:"data-dir" env:"DATA_DIR"`
	Cluster     ClusterConfig   `yaml:"cluster" env-prefix:"CLUSTER"`
	Registry    registry.Config `yaml:"registry" env-prefix:"REGISTRY"`
	Etcd        etcd.Config     `yaml:"etcd" env-prefix:"ETCD"`
	HashSlot    HashSlotConfig  `yaml:"hash-slot" env-prefix:"HASH_SLOT"`
	Cache       CacheConfig     `yaml:"cache" env-prefix:"CACHE"`
	filePath    string          `yaml:"-" env:"-"`
	persistLock sync.Locker     `yaml:"-" env:"-"`
}

func (c *Config) initialize(filePath string) {
	if c.DataDir == "" {
		c.DataDir = os.TempDir()
	}
	c.filePath, _ = filepath.Abs(filePath)
	c.persistLock = &sync.Mutex{}
	c.Cluster.LogLevel = string(c.LogLevel)
	c.Cluster.StoreDir = filepath.Join(c.DataDir, c.Registry.ServerID+"_raft")
	c.Cluster.Port = c.RpcPort
	if c.Cluster.Enable {
		c.Cluster.ID = c.Registry.ServerID
		c.HashSlot.StoreID = c.Cluster.GroupID
	} else {
		c.HashSlot.StoreID = c.Registry.ServerID
	}
}

func (c *Config) Persist() error {
	c.persistLock.Lock()
	defer c.persistLock.Unlock()
	fi, err := os.OpenFile(c.filePath, cst.OS.WriteFlag, cst.OS.ModeUser)
	if err != nil {
		return fmt.Errorf("write data to config '%s': %w", c.filePath, err)
	}
	defer fi.Close()
	enc := yaml.NewEncoder(fi)
	defer enc.Close()
	enc.SetIndent(2)
	err = enc.Encode(c)
	if err != nil {
		return fmt.Errorf("marshal config to yaml: %w", err)
	}
	return nil
}

type CacheConfig struct {
	TTL           time.Duration     `yaml:"ttl" env:"TTL" env-default:"20m"`
	CleanInterval time.Duration     `yaml:"clean-interval" env:"CLEAN_INTERVAL" env-default:"10m"`
	MaxSize       datasize.DataSize `yaml:"max-size" env:"MAX_SIZE" env-default:"128MB"`
}

type HashSlotConfig struct {
	StoreID        string        `yaml:"-" env:"-"` //StoreID could be Cluster.GroupID or Registry.ServerId
	Slots          []string      `yaml:"slots" env-separator:"," env-default:"0-16384"`
	PrepareTimeout time.Duration `yaml:"prepare-timeout" env-default:"10s"`
}

type ClusterConfig struct {
	Enable           bool          `yaml:"enable" env:"ENABLE" env-default:"false"`
	Bootstrap        bool          `yaml:"bootstrap" env:"BOOTSTRAP" env-default:"false"`
	GroupID          string        `yaml:"group-id" env:"GROUP_ID" env-default:"raft"`
	ElectionTimeout  time.Duration `yaml:"election-timeout" env:"ELECTION_TIMEOUT" env-default:"900ms"`
	HeartbeatTimeout time.Duration `yaml:"heartbeat-timeout" env:"HEARTBEAT_TIMEOUT" env-default:"800ms"`
	ID               string        `yaml:"-" env:"-"` //ID equals to Registry.ServerId
	Port             string        `yaml:"-" env:"-"` //Port equals to Config.RpcPort
	LogLevel         string        `yaml:"-" env:"-"`
	StoreDir         string        `yaml:"-" env:"-"`
	Nodes            []string      `yaml:"nodes" env:"NODES" env-separator:","`
}

func ReadConfig() *Config {
	var conf Config
	if err := cleanenv.ReadConfig(ConfFilePath, &conf); err != nil {
		panic(err)
	}
	logs.Std().Infof("read config from %s", ConfFilePath)
	conf.initialize(ConfFilePath)
	return &conf
}

func ReadConfigFrom(path string) *Config {
	var conf Config
	if err := cleanenv.ReadConfig(path, &conf); err != nil {
		if os.IsNotExist(err) {
			return ReadConfig()
		}
		panic(err)
	}
	logs.Std().Infof("read config from %s", path)
	conf.initialize(path)
	return &conf
}
