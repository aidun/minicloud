module github.com/aidun/minicloud

go 1.14

require (
	github.com/alexellis/go-execute v0.0.0-20191207085904-961405ea7544
	github.com/fluxcd/helm-operator v1.0.0-rc9
	github.com/googleapis/gnostic v0.4.1 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.6
	k8s.io/api v0.17.3
	k8s.io/apimachinery v0.17.3
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/helm v2.16.1+incompatible
	k8s.io/utils v0.0.0-20200124190032-861946025e34 // indirect
)

replace github.com/fluxcd/helm-operator/pkg/install => github.com/fluxcd/helm-operator/pkg/install v0.0.0-20200221120503-95c993a51505

replace github.com/docker/docker => github.com/docker/docker v1.4.2-0.20200226173334-8a05747fb6bf

replace helm.sh/helm/v3 => helm.sh/helm/v3 v3.1.1

replace (
	github.com/fluxcd/flux => github.com/fluxcd/flux v1.18.0
	github.com/fluxcd/flux/pkg/install => github.com/fluxcd/flux/pkg/install v0.0.0-20200206191601-8b676b003ab0
)

replace k8s.io/client-go => k8s.io/client-go v0.17.2
