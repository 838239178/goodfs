package config

import (
	"common/etcd"
	"common/logs"
	"common/registry"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	ConfFilePath = "../conf/meta-server.yaml"
)

type Config struct {
	Port     string          `yaml:"port" env:"PORT" env-default:"4091"`
	LogLevel logs.Level      `yaml:"log-level" env:"LOG_LEVEL"`
	DataDir  string          `ymal:"data-dir" env:"DATA_DIR" env-default:"/tmp/goodfs"`
	Cluster  ClusterConfig   `yaml:"cluster" env-prefix:"CLUSTER"`
	Registry registry.Config `yaml:"registry" env-prefix:"REGISTRY"`
	Etcd     etcd.Config     `yaml:"etcd" env-prefix:"ETCD"`
}

type ClusterConfig struct {
	Bootstrap       bool          `yaml:"bootstrap" env:"BOOTSTRAP" env-default:"false"`
	ElectionTimeout time.Duration `yaml:"election-timeout" env:"ELECTION_TIMEOUT" env-default:"300ms"`
	LogLevel        string        `yaml:"log-level" env:"LOG_LEVEL" env-default:"INFO"`
	StoreDir        string        `yaml:"store-dir" env:"STORE_DIR" env-default:"/tmp/goodfs_metaserver"`
	Nodes           []string      `yaml:"nodes" env:"NODES" env-required:"true" env-separator:","`
}

func ReadConfig() Config {
	return ReadConfigFrom(ConfFilePath)
}

func ReadConfigFrom(path string) Config {
	var conf Config
	if err := cleanenv.ReadConfig(path, &conf); err != nil {
		panic(err)
	}
	return conf
}
