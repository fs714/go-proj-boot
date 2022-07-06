package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fs714/go-proj-boot/cmd/config"
	"github.com/fs714/go-proj-boot/cmd/version"
	"github.com/fs714/go-proj-boot/global"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgPath string
var logLevel string

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
	global.Viper = viper.New()
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgPath, "config", "c", "", "config file path")

	rootCmd.PersistentFlags().StringVarP(&logLevel, "log-level", "", "info", "Set logging level, could be info or debug")
	global.Viper.BindPFlag("logging.level", rootCmd.PersistentFlags().Lookup("log-level"))
	global.Viper.BindEnv("logging.level", "LOGGING_LEVEL")

	rootCmd.AddCommand(version.StartCmd)
	rootCmd.AddCommand(config.StartCmd)
}

func initConfig() {
	if cfgPath != "" {
		global.Viper.SetConfigFile(cfgPath)
		global.Viper.SetConfigType("yaml")
	} else {
		global.Viper.SetConfigName("go-proj-boot")
		global.Viper.SetConfigType("yaml")

		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			fmt.Printf("failed to get current path with err: %s\n", err.Error())
		}

		global.Viper.AddConfigPath(dir + "/conf")
		global.Viper.AddConfigPath(".")
		global.Viper.AddConfigPath("./conf")
		global.Viper.AddConfigPath("/etc/go-proj-boot")
	}

	err := global.Viper.ReadInConfig()
	if err != nil {
		fmt.Printf("failed to read configuration file with err: %s\n", err.Error())
		os.Exit(1)
	}

	err = global.Viper.Unmarshal(&global.Config)
	if err != nil {
		fmt.Printf("failed to unmarshal configuration to structure with err: %s\n", err.Error())
		os.Exit(1)
	}
}
