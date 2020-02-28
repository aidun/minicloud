package main

import (
	"fmt"
	"os"

	"github.com/aidun/minicloud/cmd"
	"github.com/spf13/cobra"
)

func main() {

	var rootCmd = &cobra.Command{
		Use: "minicloud",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("test")
		},
	}

	rootCmd.AddCommand(cmd.MakeVersion())
	rootCmd.AddCommand(cmd.MakeInit())
	rootCmd.AddCommand(cmd.MakeInstall())
	rootCmd.AddCommand(cmd.MakeUpdate())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
