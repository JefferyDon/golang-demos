package client

import (
	"fmt"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"os"
)

func NewCoreV1ClientUsingServiceAccount() (*v1.CoreV1Client, error) {
	kubeConfig, err := ConfigByServiceAccount()
	if err != nil {
		return nil, err
	}
	return v1.NewForConfig(kubeConfig)
}

func NewCoreV1ClientWithKubeConfigFilePath(kubeConfigFilePath string) (*v1.CoreV1Client, error) {
	kubeConfig, err := ConfigByKubeConfigPath(kubeConfigFilePath)
	if err != nil {
		return nil, err
	}
	return v1.NewForConfig(kubeConfig)
}

func NewCoreV1ClientWithDefaultKubeConfigFilePath() (*v1.CoreV1Client, error) {
	return NewCoreV1ClientWithKubeConfigFilePath(fmt.Sprintf("%s/.kube/config", os.Getenv("HOME")))
}

func NewCoreV1ClientWithKubeConfigString(kc string) (*v1.CoreV1Client, error) {
	kubeConfig, err := ConfigByKubeConfigContent(kc)
	if err != nil {
		return nil, err
	}
	return v1.NewForConfig(kubeConfig)
}
