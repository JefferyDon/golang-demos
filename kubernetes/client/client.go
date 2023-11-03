// Package client provides several ways to create a kubernetes client.
package client

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"os"
)

func NewClientUsingServiceAccount() (kubernetes.Interface, error) {
	kubeConfig, err := ConfigByServiceAccount()
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(kubeConfig)
}

func NewClientWithKubeConfigFilePath(kubeConfigFilePath string) (kubernetes.Interface, error) {
	kubeConfig, err := ConfigByKubeConfigPath(kubeConfigFilePath)
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(kubeConfig)
}

func NewClientWithDefaultKubeConfigFilePath() (kubernetes.Interface, error) {
	return NewClientWithKubeConfigFilePath(fmt.Sprintf("%s/.kube/config", os.Getenv("HOME")))
}

func NewClientWithKubeConfigString(kc string) (kubernetes.Interface, error) {
	kubeConfig, err := ConfigByKubeConfigContent(kc)
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(kubeConfig)
}
