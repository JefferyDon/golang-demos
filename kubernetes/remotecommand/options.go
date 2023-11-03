package remotecommand

import "strings"

type fileOptions struct {
	podName       string
	namespace     string
	containerName string
	fileAbsPath   string
}

var (
	escapeCharacters = map[string]string{
		" ": "\\ ",
		"(": "\\(",
		")": "\\)",
	}
)

// NewFileOptions will create a fileOptions structure,
// and replace all escape characters in fileAbsPath parameter.
func NewFileOptions(podName, namespace, containerName, fileAbsPath string) fileOptions {

	for s, d := range escapeCharacters {
		fileAbsPath = strings.Replace(fileAbsPath, s, d, -1)
	}
	return fileOptions{
		podName:       podName,
		namespace:     namespace,
		containerName: containerName,
		fileAbsPath:   fileAbsPath,
	}
}
