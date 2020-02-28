package pkg

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func NewKubernetsClientset() (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/markushartmann/repo/minicloud/kubeconfig")

	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfigOrDie(config), nil

}
