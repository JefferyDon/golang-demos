package apimachinery

import (
	"bytes"
	"fmt"
	"io"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	serializerYaml "k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	"k8s.io/apimachinery/pkg/util/yaml"
)

// ConvertFileToUnstructuredObjects will convert content in YAML or JSON file to an array of
// unstructured.Unstructured.
//
// Note: You can use "---" as block separator in YAML file and this function
// will distract all blocks and save them as an array mentioned before.
func ConvertFileToUnstructuredObjects(content string) ([]*unstructured.Unstructured, error) {
	var (
		result      []*unstructured.Unstructured
		fileDecoder = yaml.NewYAMLOrJSONDecoder(bytes.NewBufferString(content), 4096)
	)

	for {
		var rawObj runtime.RawExtension
		err := fileDecoder.Decode(&rawObj)
		// break loop once reaches the end of file.
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("Failed to decode file content to runtime.RawExtension, err: %s. ", err)
		}
		obj, _, err := serializerYaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme).
			Decode(rawObj.Raw, nil, nil)
		if err != nil {
			return nil, fmt.Errorf("Failed to decode runtime.RawExtension to unstructured object, err: %s. ", err)
		}
		unstructuredMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
		if err != nil {
			return nil, fmt.Errorf("Failed to convert unstructured object to unstructured map, err: %s. ", err)
		}
		result = append(result, &unstructured.Unstructured{
			Object: unstructuredMap,
		})
	}
	return result, nil
}
