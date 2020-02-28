package cmd

import (
	log "github.com/sirupsen/logrus"

	"github.com/aidun/minicloud/pkg/apps"
	"github.com/spf13/cobra"
)

func MakeUpdate() *cobra.Command {

	var command = &cobra.Command{
		Use:          "update",
		Short:        "Updates all apps of the minicloud cluster",
		Example:      "minikube update",
		SilenceUsage: false,
	}

	command.Run = func(cmd *cobra.Command, args []string) {
		for _, app := range apps.AppList {
			log.Infof("Installing app %s", app.Name)
			if err := app.Update(); err != nil {
				log.Warnf("Can not update app %s: %s", app.Name, err)
			}
		}
	}
	return command
}
