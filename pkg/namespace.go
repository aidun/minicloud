package pkg

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateNameSpace(name string) error {

	clientset, err := NewKubernetsClientset()

	if err != nil {
		log.Fatal(err)
	}

	namespaceSpec := &apiv1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}

	_, err = clientset.CoreV1().Namespaces().Get(name, metav1.GetOptions{})

	if err != nil && err.Error() == fmt.Sprintf("namespaces \"%s\" not found", name) {
		_, err = clientset.CoreV1().Namespaces().Create(namespaceSpec)
		if err != nil {
			return err
		}
	}

	return nil
}
