package client

import (
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// ConfigByServiceAccount will generate *rest.Config by using files allocated
// by service account.
func ConfigByServiceAccount() (*rest.Config, error) {
	return adjustKubeConfig(rest.InClusterConfig())
}

// ConfigByKubeConfigPath will generate *rest.Config by reading content from
// file specified using parameter path.
func ConfigByKubeConfigPath(path string) (*rest.Config, error) {
	return adjustKubeConfig(clientcmd.BuildConfigFromFlags("", path))
}

// ConfigByKubeConfigContent will generate *rest.Config by kubeConfig content
// passed by parameter content.
func ConfigByKubeConfigContent(content string) (*rest.Config, error) {
	return adjustKubeConfig(clientcmd.RESTConfigFromKubeConfig([]byte(content)))
}

func adjustKubeConfig(cfg *rest.Config, err error) (*rest.Config, error) {
	if err != nil {
		return nil, err
	}
	cfg.Burst = 100
	cfg.QPS = 100
	return cfg, nil
}
