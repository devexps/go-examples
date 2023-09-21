package bootstrap

import (
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	"github.com/devexps/go-examples/k8s/api/gen/go/common/conf"

	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-micro/v2/registry"

	// consul
	consulMicro "github.com/devexps/go-micro/registry/consul/v2"
	consulClient "github.com/hashicorp/consul/api"

	// etcd
	etcdMicro "github.com/devexps/go-micro/registry/etcd/v2"
	etcdClient "go.etcd.io/etcd/client/v3"

	// k8s
	k8sMicro "github.com/devexps/go-micro/registry/k8s/v2"
)

type RegistryType string

const (
	RegistryTypeConsul RegistryType = "consul"
	RegistryTypeEtcd   RegistryType = "etcd"
	RegistryTypeK8s    RegistryType = "k8s"
)

// NewRegistry creates new registry
func NewRegistry(cfg *conf.Registry) registry.Registrar {
	if cfg == nil {
		return nil
	}
	switch RegistryType(cfg.Type) {
	case RegistryTypeConsul:
		return NewConsulRegistry(cfg)
	case RegistryTypeEtcd:
		return NewEtcdRegistry(cfg)
	case RegistryTypeK8s:
		return NewK8sRegistry(cfg)
	}
	return nil
}

// NewDiscovery creates a discovery client
func NewDiscovery(cfg *conf.Registry) registry.Discovery {
	if cfg == nil {
		return nil
	}
	switch RegistryType(cfg.Type) {
	case RegistryTypeConsul:
		return NewConsulRegistry(cfg)
	case RegistryTypeEtcd:
		return NewEtcdRegistry(cfg)
	case RegistryTypeK8s:
		return NewK8sRegistry(cfg)
	}
	return nil
}

// NewConsulRegistry creates a registry discovery client - Consul
func NewConsulRegistry(c *conf.Registry) *consulMicro.Registry {
	cfg := consulClient.DefaultConfig()
	cfg.Address = c.Consul.Address
	cfg.Scheme = c.Consul.Scheme

	var cli *consulClient.Client
	var err error
	if cli, err = consulClient.NewClient(cfg); err != nil {
		log.Fatal(err)
	}
	return consulMicro.New(cli, consulMicro.WithHealthCheck(c.Consul.HealthCheck))
}

// NewEtcdRegistry creates a registry discovery client - Etcd
func NewEtcdRegistry(c *conf.Registry) *etcdMicro.Registry {
	cfg := etcdClient.Config{
		Endpoints: c.Etcd.Endpoints,
	}
	var err error
	var cli *etcdClient.Client
	if cli, err = etcdClient.New(cfg); err != nil {
		log.Fatal(err)
	}
	return etcdMicro.New(cli)
}

// NewK8sRegistry creates a registry discovery client - Kubernetes
func NewK8sRegistry(c *conf.Registry) *k8sMicro.Registry {
	clientSet, _, err := getK8sClientSet(c.K8S.GetMasterUrl())
	if err != nil {
		log.Fatal(err)
	}
	r := k8sMicro.NewRegistry(clientSet)
	r.Start()
	return r
}

func getK8sClientSet(masterUrl string) (*kubernetes.Clientset, string, error) {
	kubeConfig := ""
	restConfig, err := rest.InClusterConfig()
	if err != nil {
		kubeConfig = filepath.Join(homedir.HomeDir(), ".kube", "config")
		restConfig, err = clientcmd.BuildConfigFromFlags(masterUrl, kubeConfig)
		if err != nil {
			return nil, kubeConfig, err
		}
	}
	clientSet, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, kubeConfig, err
	}
	return clientSet, kubeConfig, nil
}
