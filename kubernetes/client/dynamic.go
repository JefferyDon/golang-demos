package client

import (
	"fmt"
	"k8s.io/client-go/dynamic"
	"os"
)

func NewDynamicClientUsingServiceAccount() (dynamic.Interface, error) {
	kubeConfig, err := ConfigByServiceAccount()
	if err != nil {
		return nil, err
	}
	return dynamic.NewForConfig(kubeConfig)
}

func NewDynamicClientWithKubeConfigFilePath(kubeConfigFilePath string) (dynamic.Interface, error) {
	kubeConfig, err := ConfigByKubeConfigPath(kubeConfigFilePath)
	if err != nil {
		return nil, err
	}
	return dynamic.NewForConfig(kubeConfig)
}

func NewDynamicClientWithDefaultKubeConfigFilePath() (dynamic.Interface, error) {
	return NewDynamicClientWithKubeConfigFilePath(fmt.Sprintf("%s/.kube/config", os.Getenv("HOME")))
}

func NewDynamicWithKubeConfigString(kc string) (dynamic.Interface, error) {
	kubeConfig, err := ConfigByKubeConfigContent(kc)
	if err != nil {
		return nil, err
	}
	return dynamic.NewForConfig(kubeConfig)
}
