package remotecommand

import (
	"fmt"
	"testing"
)

func TestNewFileOptions(t *testing.T) {
	const (
		podName       = "pod-test"
		namespace     = "test"
		containerName = "pod-test"
	)

	type fileTest struct {
		name   string
		path   string
		expect string
	}

	files := []fileTest{
		{
			name:   "Simple bracket test",
			path:   "/yum(1).tar.gz",
			expect: "/yum\\(1\\).tar.gz",
		},
		{
			name:   "Simple blank test",
			path:   "/go test",
			expect: "/go\\ test",
		},
		{
			name:   "Complex test for both bracket and blank",
			path:   "/Applications/Visual Studio Code.app/go test(1)",
			expect: "/Applications/Visual\\ Studio\\ Code.app/go\\ test\\(1\\)",
		},
	}

	for _, f := range files {
		t.Run(f.name, func(t *testing.T) {
			newPath := NewFileOptions(podName, namespace, containerName, f.path).fileAbsPath
			if newPath != f.expect {
				t.Errorf("Result is not what I've expected.\nTarget:%s\nResult:%s\nExpect:%s\n",
					f.path, newPath, f.expect)
				return
			}
			fmt.Printf("Result matches expectation!!\nTarget:%s\nResult:%s\nExpect:%s\n",
				f.path, newPath, f.expect)
		})
	}
}
