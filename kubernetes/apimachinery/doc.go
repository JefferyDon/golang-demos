/*
Package apimachinery contains lots of advanced usage of client-go SDK, including:

1. Using "k8s.io/apimachinery/pkg/util/yaml" to convert YAML or JSON configurations to
runtime.RawExtension and use this variable to achieve functions which commands below can do:

	kubectl apply -f ${FILENAME}
	kubectl delete -f ${FILENAME}
*/
package apimachinery
