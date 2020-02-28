package cmd

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/aidun/minicloud/pkg"
	execute "github.com/alexellis/go-execute/pkg/v1"
	"github.com/spf13/cobra"
)

const MinicloudNamespace = "minicloud"

func MakeInit() *cobra.Command {

	var command = &cobra.Command{
		Use:          "init",
		Short:        "Installs all needed components to the cluster",
		Example:      "minikube init",
		SilenceUsage: false,
	}

	command.Run = func(cmd *cobra.Command, args []string) {
		log.Println("create minicloud namespace")
		if err := pkg.CreateNameSpace(MinicloudNamespace); err != nil {
			log.Fatal(err)
		}

		log.Println("check if helm 3 is present")
		if !checkHelmVersion() {
			log.Fatal("no helm 3 was found in path")
		}

		log.Println("add fluxcd repo to the list of repos")
		if err := addFluxcdHelmRepo(); err != nil {
			log.Fatal(err)
		}

		log.Println("install helm-operator for arm")
		if err := installHelmOperator(); err != nil {
			log.Fatal(err)
		}
	}
	return command
}

func checkHelmVersion() bool {
	helmVersion := execute.ExecTask{
		Command: "helm",
		Args:    []string{"version", "--short"},
		Shell:   true,
	}

	res, err := helmVersion.Execute()
	if err != nil {
		log.Println(err)
		return false
	}

	if strings.HasPrefix(res.Stdout, "v3") {
		return true
	}

	return false
}

func addFluxcdHelmRepo() error {
	helmRepo := execute.ExecTask{
		Command: "helm",
		Args:    []string{"repo", "add", "minicloud-fluxcd", "https://charts.fluxcd.io"},
		Shell:   true,
	}

	res, err := helmRepo.Execute()
	if err != nil {
		return err
	}

	if res.ExitCode != 0 {
		return fmt.Errorf("Error while helm repo add: %s", res.Stderr)
	}

	helmRepoUpdate := execute.ExecTask{
		Command: "helm",
		Args:    []string{"repo", "update"},
		Shell:   true,
	}

	res, err = helmRepoUpdate.Execute()
	if err != nil {
		return err
	}

	if res.ExitCode != 0 {
		return fmt.Errorf("Error while helm repo update: %s", res.Stderr)
	}

	return nil
}

func installHelmOperator() error {
	helmUpgrade := execute.ExecTask{
		Command: "helm",
		Args: []string{
			"--kubeconfig", "/Users/markushartmann/repo/minicloud/kubeconfig",
			"upgrade",
			"-i", "helm-operator",
			"minicloud-fluxcd/helm-operator",
			"--namespace", MinicloudNamespace,
			"--set", "helm.versions=v3",
			"--set", "image.repository=docker.io/onedr0p/helm-operator",
			"--set", "image.tag=latest",
		},
		Shell: true,
	}

	res, err := helmUpgrade.Execute()
	if err != nil {
		return err
	}

	if res.ExitCode != 0 {
		return fmt.Errorf("Error while helm upgade: %s", res.Stderr)
	}

	return nil
}
