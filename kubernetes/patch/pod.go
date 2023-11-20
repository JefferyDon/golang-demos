package patch

import (
	"context"
	"encoding/json"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
	"k8s.io/client-go/kubernetes"
)

// PodStatusPatch shows how to update pod status
// by using patch method. Actually, this is the same as PatchPodStatus
// function provided by package "k8s.io/kubernetes/pkg/util/pod".
// It means you can simply update pod status by using PatchPodStatus functions, like:
//
//	import (
//	  	statusutil "k8s.io/kubernetes/pkg/util/pod"
//	)
//
//	 func main(){
//	  	newPod, patchBytes, unchanged, err := statusutil.PatchPodStatus(m.kubeClient, pod.Namespace, pod.Name, pod.UID, pod.Status, mergedStatus)
//		klog.V(3).InfoS("Patch status for pod", "pod", klog.KObj(pod), "patch", string(patchBytes))
//	 }
func PodStatusPatch(client kubernetes.Interface, name, namespace string, oldStatus, newStatus v1.PodStatus) error {
	patchBytes, err := preparePatchBytesForPodStatus(namespace, name, oldStatus, newStatus)
	if err != nil {
		return err
	}

	_, err = client.CoreV1().Pods(namespace).Patch(context.Background(), name, types.StrategicMergePatchType,
		patchBytes, metaV1.PatchOptions{}, "status")
	if err != nil {
		return fmt.Errorf("failed to patch status %q for pod %q/%q: %v", patchBytes, namespace, name, err)
	}
	return nil
}

func preparePatchBytesForPodStatus(namespace, name string, oldPodStatus, newPodStatus v1.PodStatus) ([]byte, error) {
	oldData, err := json.Marshal(v1.Pod{
		Status: oldPodStatus,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to Marshal oldData for pod %q/%q: %v", namespace, name, err)
	}

	newData, err := json.Marshal(v1.Pod{
		Status: newPodStatus,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to Marshal newData for pod %q/%q: %v", namespace, name, err)
	}

	patchBytes, err := strategicpatch.CreateTwoWayMergePatch(oldData, newData, v1.Pod{})
	if err != nil {
		return nil, fmt.Errorf("failed to CreateTwoWayMergePatch for pod %q/%q: %v", namespace, name, err)
	}
	return patchBytes, nil
}
