package restmapper

import (
	"fmt"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/restmapper"
)

func SearchingGroupKindVersions(client kubernetes.Interface, gk schema.GroupKind, versions ...string) (*meta.RESTMapping, error) {
	gr, err := restmapper.GetAPIGroupResources(client.Discovery())
	if err != nil {
		return nil, err
	}
	mapper := restmapper.NewDiscoveryRESTMapper(gr)
	mapping, err := mapper.RESTMapping(gk, versions...)
	if err != nil {
		return nil, err
	}
	if mapping == nil {
		return nil, fmt.Errorf("GroupKindVersion not found! GroupKind: %s, Versions: %+v. ", gk.String(), versions)
	}
	return mapping, nil
}
