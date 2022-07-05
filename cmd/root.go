package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fs714/go-proj-boot/cmd/version"
	"github.com/fs714/go-proj-boot/global"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgPath string

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
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgPath, "config", "c", "", "config file (default is $HOME/conf/go-proj-boot.yml)")
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))

	rootCmd.AddCommand(version.StartCmd)
}

func initConfig() {
	if cfgPath != "" {
		viper.SetConfigFile(cfgPath)
		viper.SetConfigType("yaml")
	} else {
		viper.SetConfigName("go-proj-boot")
		viper.SetConfigType("yaml")

		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			fmt.Printf("failed to get current path with err: %s\n", err.Error())
		}

		viper.AddConfigPath(dir + "/conf")
		viper.AddConfigPath(".")
		viper.AddConfigPath("./conf")
		viper.AddConfigPath("/etc/go-proj-boot")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Printf("using config file: %s\n", viper.ConfigFileUsed())
	} else {
		fmt.Printf("failed to read configuration file with err: %s\n", err.Error())
	}
}
