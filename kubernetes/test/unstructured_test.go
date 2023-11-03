package test

import (
	"fmt"
	"github.com/JefferyDon/golang-demos/kubernetes/apimachinery"
	"github.com/JefferyDon/golang-demos/kubernetes/client"
	"github.com/JefferyDon/golang-demos/kubernetes/restmapper"
	"io"
	"os"
	"testing"
)

const (
	testFileName = "unstructured-test.yaml"
)

// TestApplyFiles will use apimachinery.ConvertFileToUnstructuredObjects,
// apimachinery.ApplyUnstructuredObjects and restmapper.SearchingGroupKindVersions
// to apply a YAML file.
func TestApplyFiles(t *testing.T) {
	dyClient, err := client.NewDynamicClientWithKubeConfigFilePath("kube-configs/kube.config")
	if err != nil {
		t.Fatal(err)
	}

	k8sClient, err := client.NewClientWithKubeConfigFilePath("kube-configs/kube.config")
	if err != nil {
		t.Fatal(err)
	}

	targetFile := fmt.Sprintf("test-yamls/%s", testFileName)
	fd, err := os.Open(targetFile)
	if err != nil {
		t.Fatal(err)
	}

	content, err := io.ReadAll(fd)
	if err != nil {
		t.Fatal(err)
	}

	unstructuredObjects, err := apimachinery.ConvertFileToUnstructuredObjects(string(content))
	if err != nil {
		t.Fatal(err)
	}

	for _, obj := range unstructuredObjects {
		mapping, err := restmapper.SearchingGroupKindVersions(k8sClient, obj.GroupVersionKind().GroupKind(), obj.GroupVersionKind().Version)
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("Apply unstructured data to kubernetes cluster. GroupKindVersion: %+v, Name: %s, Namespace: %s. ", obj.GroupVersionKind(), obj.GetName(), obj.GetNamespace())

		if err = apimachinery.ApplyUnstructuredObject(dyClient, mapping.Resource, obj); err != nil {
			t.Fatal(err)
		}
	}
}

// TestDeleteFiles will use apimachinery.ConvertFileToUnstructuredObjects,
// apimachinery.ApplyUnstructuredObjects and restmapper.SearchingGroupKindVersions
// to delete a YAML file.
func TestDeleteFiles(t *testing.T) {
	dyClient, err := client.NewDynamicClientWithKubeConfigFilePath("kube-configs/kube.config")
	if err != nil {
		t.Fatal(err)
	}

	k8sClient, err := client.NewClientWithKubeConfigFilePath("kube-configs/kube.config")
	if err != nil {
		t.Fatal(err)
	}

	targetFile := fmt.Sprintf("test-yamls/%s", testFileName)
	fd, err := os.Open(targetFile)
	if err != nil {
		t.Fatal(err)
	}

	content, err := io.ReadAll(fd)
	if err != nil {
		t.Fatal(err)
	}

	unstructuredObjects, err := apimachinery.ConvertFileToUnstructuredObjects(string(content))
	if err != nil {
		t.Fatal(err)
	}

	for _, obj := range unstructuredObjects {
		mapping, err := restmapper.SearchingGroupKindVersions(k8sClient, obj.GroupVersionKind().GroupKind(), obj.GroupVersionKind().Version)
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("Delete unstructured data to kubernetes cluster. GroupKindVersion: %+v, Name: %s, Namespace: %s. ", obj.GroupVersionKind(), obj.GetName(), obj.GetNamespace())

		if err = apimachinery.DeleteUnstructuredObject(dyClient, mapping.Resource, obj); err != nil {
			t.Fatal(err)
		}
	}
}
