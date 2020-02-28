package cmd

import (
	log "github.com/sirupsen/logrus"

	"github.com/aidun/minicloud/pkg/apps"
	"github.com/spf13/cobra"
)

func MakeInstall() *cobra.Command {

	var command = &cobra.Command{
		Use:          "install",
		Short:        "Installs all needed apps to the cluster",
		Example:      "minikube install",
		SilenceUsage: false,
	}

	command.Run = func(cmd *cobra.Command, args []string) {
		for _, app := range apps.AppList {
			log.Infof("Updating app %s", app.Name)
			if err := app.Install(); err != nil {
				log.Warnf("Can not install app %s: %s", app.Name, err)
			}
		}
	}
	return command
}
