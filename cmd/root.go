package cmd

import (
	"os"

	"github.com/fs714/go-proj-boot/cmd/version"
	"github.com/fs714/go-proj-boot/global"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "go-proj-boot",
	Version: global.Version,
	Short:   "A golang reference project",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(version.StartCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-proj-boot.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
