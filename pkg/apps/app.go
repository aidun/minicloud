package apps

import (
	"fmt"

	"github.com/aidun/minicloud/pkg"
	helmreleasev1 "github.com/fluxcd/helm-operator/pkg/apis/helm.fluxcd.io/v1"
	helm3 "github.com/fluxcd/helm-operator/pkg/client/clientset/versioned"
	"github.com/fluxcd/helm-operator/pkg/helm"
	"k8s.io/client-go/tools/clientcmd"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type App struct {
	Namespace       string
	ChartRepository string
	ChartName       string
	ChartVersion    string
	Name            string
	Values          helm.Values
}

func (app *App) Install() error {

	config, _ := clientcmd.BuildConfigFromFlags("", "/Users/markushartmann/repo/minicloud/kubeconfig")
	clientset := helm3.NewForConfigOrDie(config)

	if err := pkg.CreateNameSpace(app.Namespace); err != nil {
		return fmt.Errorf("Error creating the namespace %s", app.Namespace)
	}

	_, err := clientset.HelmV1().HelmReleases("minicloud").Get(app.Name, metav1.GetOptions{})

	if err != nil && err.Error() == fmt.Sprintf("helmreleases.helm.fluxcd.io \"%s\" not found", app.Name) {
		_, err := clientset.HelmV1().HelmReleases("minicloud").Create(app.GetHelmRelease())
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("%s already installed. Try minicloud update instead.", app.Name)
	}

	return nil
}

func (app *App) Update() error {
	config, _ := clientcmd.BuildConfigFromFlags("", "/Users/markushartmann/repo/minicloud/kubeconfig")
	clientset := helm3.NewForConfigOrDie(config)

	actualVersion, err := clientset.HelmV1().HelmReleases("minicloud").Get(app.Name, metav1.GetOptions{})

	if err != nil && err.Error() == fmt.Sprintf("HelmRelease \"%s\" not found", app.Name) {
		return fmt.Errorf("%s not installed to the cluster. Run minicloud install first.", app.Name)
	}

	helmRelease := app.GetHelmRelease()
	helmRelease.SetResourceVersion(actualVersion.GetResourceVersion())
	_, err = clientset.HelmV1().HelmReleases("minicloud").Update(helmRelease)
	if err != nil {
		return err
	}

	newVersion, err := clientset.HelmV1().HelmReleases("minicloud").Get(app.Name, metav1.GetOptions{})

	if actualVersion.GetResourceVersion() == newVersion.GetResourceVersion() {
		return fmt.Errorf("No update needed")
	}

	return nil
}

func (app *App) GetHelmRelease() *helmreleasev1.HelmRelease {
	helmReleaseSpec := helmreleasev1.HelmReleaseSpec{
		HelmVersion:     "v3",
		ReleaseName:     app.Name,
		TargetNamespace: app.Namespace,
		SkipCRDs:        false,
		ChartSource: helmreleasev1.ChartSource{
			nil,
			&helmreleasev1.RepoChartSource{
				RepoURL: app.ChartRepository,
				Name:    app.ChartName,
				Version: app.ChartVersion,
			},
		},
		HelmValues: helmreleasev1.HelmValues{
			Values: app.Values,
		},
	}

	return &helmreleasev1.HelmRelease{
		ObjectMeta: metav1.ObjectMeta{
			Name: app.Name,
		},
		Spec: helmReleaseSpec,
	}
}
