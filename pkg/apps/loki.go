package apps

import (
	log "github.com/sirupsen/logrus"

	"k8s.io/helm/pkg/chartutil"
)

func CreateLokiApp() *App {

	valuesRaw := []byte(`
grafana:
  enabled: true
  service:
    type: LoadBalancer
  sidecar:
    datasources:
      enabled: false
    dasboards:
      enabled: false`)

	values, err := chartutil.ReadValues(valuesRaw)

	if err != nil {
		log.Fatal(err)
		log.Fatal("loki values can´t parsed")
	}

	return &App{
		Namespace:       "loki",
		Name:            "loki-stack",
		ChartVersion:    "0.32.0",
		ChartName:       "loki-stack",
		ChartRepository: "https://grafana.github.io/loki/charts",
		Values:          values.AsMap(),
	}
}
