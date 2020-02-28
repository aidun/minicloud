package apps

import (
	log "github.com/sirupsen/logrus"

	"k8s.io/helm/pkg/chartutil"
)

func CreateMetallbApp() *App {

	valuesRaw := []byte(`
configInline:
  address-pools:
  - name: default
    protocol: layer2
    addresses:
    - 192.168.178.180-192.168.178.250`)
	metallbValues, err := chartutil.ReadValues(valuesRaw)

	if err != nil {
		log.Fatal(err)
		log.Fatal("metallb values can´t parsed")
	}

	return &App{
		Namespace:       "metallb",
		Name:            "metallb",
		ChartVersion:    "0.12.0",
		ChartName:       "metallb",
		ChartRepository: "https://kubernetes-charts.storage.googleapis.com",
		Values:          metallbValues.AsMap(),
	}
}
