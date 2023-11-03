// Package remotecommand provides some usage for package "k8s.io/client-go/tools/remotecommand",
// for more information, please see: https://pkg.go.dev/k8s.io/client-go/tools/remotecommand
package remotecommand

import (
	"archive/tar"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
	"net/url"
	"os"
)

type File struct {
	Name string
	Kind string
	Size string
}

type Executor interface {
	GetFile(options fileOptions) (string, error)
	DownloadFile(options fileOptions, writers ...io.Writer) error
}

type remoteCommandExecutor struct {
	kubeConfig *rest.Config
	client     *v1.CoreV1Client
}

func NewRemoteCommandExecutor(kubeConfig *rest.Config) (Executor, error) {
	c, err := v1.NewForConfig(kubeConfig)
	if err != nil {
		return nil, err
	}
	return &remoteCommandExecutor{
		kubeConfig: kubeConfig,
		client:     c,
	}, nil
}

// GetFile will return file name if file exists, and return an empty string if not.
func (r *remoteCommandExecutor) GetFile(options fileOptions) (string, error) {
	cmd := []string{"/bin/sh", "-c", fmt.Sprintf("stat %s -c %%n:%%F:%%s", options.fileAbsPath)}
	req := r.client.RESTClient().Get().
		Name(options.podName).
		Namespace(options.namespace).
		Resource("pods").
		SubResource("exec").
		VersionedParams(&corev1.PodExecOptions{
			Container: options.containerName,
			Command:   cmd,
			Stdin:     true,
			Stderr:    true,
			Stdout:    true,
			TTY:       false,
		}, scheme.ParameterCodec)
	reader, err := r.execute(req.URL())
	if err != nil {
		return "", err
	}

	var (
		data    = make([]byte, 1)
		allData []byte
	)

	for {
		_, err = reader.Read(data)
		if err != nil {
			break
		}
		allData = append(allData, data...)
	}
	return string(allData), nil
}

// DownloadFile will download file from container and copy file content
// with writers in parameter. You can use your writers to operate file content
// once DownloadFile has been successfully done.
func (r *remoteCommandExecutor) DownloadFile(options fileOptions, writers ...io.Writer) error {
	fileStat, err := r.GetFile(options)
	if err != nil {
		return err
	}

	if fileStat == "" {
		return fmt.Errorf("File Not Exist. ")
	}

	cmd := []string{"tar", "-cf", "-", options.fileAbsPath}
	req := r.client.RESTClient().Get().
		Namespace(options.namespace).
		Name(options.podName).
		Resource("pods").
		SubResource("exec").
		VersionedParams(&corev1.PodExecOptions{
			Container: options.containerName,
			Command:   cmd,
			Stdin:     true,
			Stdout:    true,
			Stderr:    true,
			TTY:       false,
		}, scheme.ParameterCodec)

	reader, err := r.execute(req.URL())
	if err != nil {
		return err
	}

	tarReader := tar.NewReader(reader)
	if _, err = tarReader.Next(); err != nil {
		return err
	}

	for _, w := range writers {
		if _, err = io.Copy(w, tarReader); err != nil {
			return err
		}
	}

	return nil
}

func (r *remoteCommandExecutor) execute(req *url.URL) (io.Reader, error) {
	executor, err := remotecommand.NewSPDYExecutor(r.kubeConfig, "POST", req)
	if err != nil {
		return nil, err
	}
	reader, writer := io.Pipe()

	go func() {
		defer writer.Close()
		err = executor.Stream(remotecommand.StreamOptions{
			Stdin:             os.Stdin,
			Stdout:            writer,
			Stderr:            os.Stderr,
			Tty:               false,
			TerminalSizeQueue: nil,
		})
		if err != nil {
			logrus.WithError(err).WithField("req", req.String()).Error("Failed to get SPDY stream!")
		}
	}()

	return reader, nil
}
