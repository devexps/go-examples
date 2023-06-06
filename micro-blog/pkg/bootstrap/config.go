package bootstrap

import (
	"github.com/devexps/go-examples/micro-blog/api/gen/go/common/conf"
	"github.com/devexps/go-micro/v2/config"
	"github.com/devexps/go-micro/v2/config/file"
	"github.com/devexps/go-micro/v2/log"
	"os"
)

type ConfigType string

const (
	ConfigTypeLocalFile ConfigType = "file"
	ConfigTypeConsul    ConfigType = "consul"
	ConfigTypeEtcd      ConfigType = "etcd"

	remoteConfigSourceConfigFile = "remote.yaml"
)

// LoadRemoteConfigSourceConfigs loads local config from remote config source
func LoadRemoteConfigSourceConfigs(configPath string) (*conf.RemoteConfig, error) {
	configPath = configPath + "/" + remoteConfigSourceConfigFile
	if !pathExists(configPath) {
		return nil, nil
	}
	cfg := config.New(
		config.WithSource(
			NewFileConfigSource(configPath),
		),
	)
	defer func(cfg config.Config) {
		err := cfg.Close()
		if err != nil {
			panic(err)
		}
	}(cfg)

	var err error

	if err = cfg.Load(); err != nil {
		return nil, err
	}

	var rc conf.Bootstrap
	if err = cfg.Scan(&rc); err != nil {
		return nil, err
	}

	return rc.Config, err
}

// NewConfigProvider creates a configuration
func NewConfigProvider(configPath string) config.Config {
	rc, err := LoadRemoteConfigSourceConfigs(configPath)
	if err != nil {
		log.Error("LoadRemoteConfigSourceConfigs: ", err.Error())
	}
	if rc != nil {
		return config.New(
			config.WithSource(
				NewFileConfigSource(configPath),
				NewRemoteConfigSource(rc),
			),
		)
	} else {
		return config.New(
			config.WithSource(
				NewFileConfigSource(configPath),
			),
		)
	}
}

// NewFileConfigSource creates a local file configuration source
func NewFileConfigSource(filePath string) config.Source {
	return file.NewSource(filePath)
}

// NewRemoteConfigSource creates a remote configuration source
func NewRemoteConfigSource(c *conf.RemoteConfig) config.Source {
	switch ConfigType(c.Type) {
	default:
		fallthrough
	case ConfigTypeLocalFile:
		return nil
	case ConfigTypeConsul:
		return nil // TODO
	case ConfigTypeEtcd:
		return nil // TODO
	}
}

// LoadBootstrapConfig loader bootstrap configuration
func LoadBootstrapConfig(configPath string) *conf.Bootstrap {
	cfg := NewConfigProvider(configPath)
	if err := cfg.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := cfg.Scan(&bc); err != nil {
		panic(err)
	}

	if bc.Server == nil {
		bc.Server = &conf.Server{}
		_ = cfg.Scan(&bc.Server)
	}

	if bc.Data == nil {
		bc.Data = &conf.Data{}
		_ = cfg.Scan(&bc.Data)
	}

	if bc.Auth == nil {
		bc.Auth = &conf.Auth{}
		_ = cfg.Scan(&bc.Auth)
	}

	if bc.Trace == nil {
		bc.Trace = &conf.Tracer{}
		_ = cfg.Scan(&bc.Trace)
	}

	if bc.Logger == nil {
		bc.Logger = &conf.Logger{}
		_ = cfg.Scan(&bc.Logger)
	}

	if bc.Registry == nil {
		bc.Registry = &conf.Registry{}
		_ = cfg.Scan(&bc.Registry)
	}

	return &bc
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
