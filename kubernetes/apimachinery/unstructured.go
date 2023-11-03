package apimachinery

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/client-go/dynamic"
)

// ApplyUnstructuredObject is used to create or update configurations in kubernetes
//
// As for parameter dynamic.Interface, you can get it by client.NewDynamicClient*;
//
// As for parameter schema.GroupVersionResource, you can get it by restmapper.SearchingGroupKindVersions;
func ApplyUnstructuredObject(client dynamic.Interface, gksr schema.GroupVersionResource, obj *unstructured.Unstructured) error {

	_, err := client.Resource(gksr).Namespace(obj.GetNamespace()).Get(context.Background(), obj.GetName(), metav1.GetOptions{})
	if err != nil && !errors.IsNotFound(err) {
		return err
	}

	// if err is nil, it means the target has already been deployed to kubernetes cluster,
	// so we update it!
	if err == nil {
		if err = updateUnstructuredObject(client, gksr, obj); err != nil {
			return err
		}
		return nil
	}

	// create unstructured object!
	if err = createUnstructuredObject(client, gksr, obj); err != nil {
		return err
	}

	return nil
}

func createUnstructuredObject(client dynamic.Interface, gksr schema.GroupVersionResource, obj *unstructured.Unstructured) error {
	_, err := client.Resource(gksr).Namespace(obj.GetNamespace()).Create(context.Background(), obj, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}

func updateUnstructuredObject(client dynamic.Interface, gksr schema.GroupVersionResource, obj *unstructured.Unstructured) error {
	body, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = client.Resource(gksr).Namespace(obj.GetNamespace()).Patch(context.Background(), obj.GetName(), types.MergePatchType, body, metav1.PatchOptions{})
	if err != nil {
		return err
	}
	return nil
}

// DeleteUnstructuredObject is used to delete configurations in kubernetes
//
// As for parameter dynamic.Interface, you can get it by client.NewDynamicClient*;
//
// As for parameter schema.GroupVersionResource, you can get it by restmapper.SearchingGroupKindVersions;
func DeleteUnstructuredObject(client dynamic.Interface, gksr schema.GroupVersionResource, obj *unstructured.Unstructured) error {

	_, err := client.Resource(gksr).Namespace(obj.GetNamespace()).Get(context.Background(), obj.GetName(), metav1.GetOptions{})
	if err != nil && !errors.IsNotFound(err) {
		return err
	}

	// if err is nil, it means the target exists in kubernetes cluster,
	// so we can delete it!
	if err == nil {
		if err = deleteUnstructuredObject(client, gksr, obj); err != nil {
			return err
		}
		return nil
	}

	return fmt.Errorf("Resource not exist in kubernetes cluster! GroupKindVersion: %+v, Name: %s, Namespace: %s. ",
		gksr, obj.GetName(), obj.GetNamespace())
}

func deleteUnstructuredObject(client dynamic.Interface, gksr schema.GroupVersionResource, obj *unstructured.Unstructured) error {
	return client.Resource(gksr).Namespace(obj.GetNamespace()).Delete(context.Background(), obj.GetName(), metav1.DeleteOptions{})
}
