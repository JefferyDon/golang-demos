package test

import (
	"context"
	"github.com/JefferyDon/golang-demos/kubernetes/client"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"testing"
)

const (
	targetDeploymentNamespace = "kube-system"
	targetDeploymentName      = "coredns"
)

func doSomething(client kubernetes.Interface) error {
	deploy, err := client.AppsV1().
		Deployments(targetDeploymentNamespace).
		Get(context.Background(), targetDeploymentName, v1.GetOptions{})
	if err != nil {
		return err
	}
	println(deploy.Name, deploy.Namespace)
	return nil
}

func TestNewClientWithKubeConfigFilePath(t *testing.T) {
	c, err := client.NewClientWithKubeConfigFilePath("kube-configs/kube.config")
	if err != nil {
		t.Fatal(err)
	}

	if err = doSomething(c); err != nil {
		t.Fatal(err)
	}
}
