package test

import (
	"fmt"
	"github.com/JefferyDon/golang-demos/kubernetes/client"
	"github.com/JefferyDon/golang-demos/kubernetes/remotecommand"
	"io"
	"os"
	"testing"
)

const (
	testPodName          = "shaman-0"
	testPodNamespace     = "xhzy-backend"
	testPodContainerName = "shaman"
)

func pre(t *testing.T) remotecommand.Executor {
	config, err := client.ConfigByKubeConfigPath("kube-configs/kube.config")
	if err != nil {
		t.Fatal(err)
	}
	executor, err := remotecommand.NewRemoteCommandExecutor(config)
	if err != nil {
		t.Fatal(err)
	}
	return executor
}

func TestGetFile(t *testing.T) {
	executor := pre(t)

	testFiles := []string{
		"/home/admin/logs/webapp/shaman/auth.log",
		"/home/admin/file_not_exist",
	}

	for _, file := range testFiles {
		t.Run(file, func(t *testing.T) {
			result, err := executor.GetFile(remotecommand.NewFileOptions(
				testPodName, testPodNamespace, testPodContainerName, file))
			if err != nil {
				t.Fatal(err)
			}
			fmt.Println(func(r string) string {
				if r == "" {
					return "文件不存在"
				}
				return r
			}(result))
		})
	}
}

func TestDownloadFile(t *testing.T) {
	executor := pre(t)

	type testObject struct {
		name     string
		fileName string
		writer   io.Writer
	}

	w, err := os.OpenFile("testFile", os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		t.Fatal(err)
	}

	testObjects := []testObject{
		{
			name:     "fileExistAndToFile",
			fileName: "/home/admin/test",
			writer:   w,
		},
		{
			name:     "fileExistAndToStdOut",
			fileName: "/home/admin/test",
			writer:   os.Stdout,
		},
		{
			name:     "fileNotExistAndToStdOut",
			fileName: "/home/admin/test.notExist",
			writer:   os.Stdout,
		},
	}

	for _, to := range testObjects {
		t.Run(to.name, func(t *testing.T) {
			if err := executor.DownloadFile(remotecommand.NewFileOptions(
				testPodName, testPodNamespace, testPodContainerName, to.fileName), to.writer); err != nil {
				t.Fatal(err)
			}
		})
	}
}
