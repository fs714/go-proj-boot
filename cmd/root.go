package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fs714/go-proj-boot/cmd/show_config"
	"github.com/fs714/go-proj-boot/cmd/show_version"
	"github.com/fs714/go-proj-boot/utils/config"
	"github.com/fs714/go-proj-boot/utils/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgPath string
var logLevel string

var rootCmd = &cobra.Command{
	Use:     "go-proj-boot",
	Version: version.Version,
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
	config.Viper = viper.New()
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgPath, "config", "c", "", "config file path")

	rootCmd.PersistentFlags().StringVarP(&logLevel, "log-level", "", "info", "Set logging level, could be info or debug")
	config.Viper.BindPFlag("logging.level", rootCmd.PersistentFlags().Lookup("log-level"))
	config.Viper.BindEnv("logging.level", "LOGGING_LEVEL")

	rootCmd.AddCommand(show_version.StartCmd)
	rootCmd.AddCommand(show_config.StartCmd)
}

func initConfig() {
	if cfgPath != "" {
		config.Viper.SetConfigFile(cfgPath)
		config.Viper.SetConfigType("yaml")
	} else {
		config.Viper.SetConfigName("go-proj-boot")
		config.Viper.SetConfigType("yaml")

		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			fmt.Printf("failed to get current path with err: %s\n", err.Error())
		}

		config.Viper.AddConfigPath(dir + "/conf")
		config.Viper.AddConfigPath(".")
		config.Viper.AddConfigPath("./conf")
		config.Viper.AddConfigPath("/etc/go-proj-boot")
	}

	err := config.Viper.ReadInConfig()
	if err != nil {
		fmt.Printf("failed to read configuration file with err: %s\n", err.Error())
		os.Exit(1)
	}

	err = config.Viper.Unmarshal(&config.Config)
	if err != nil {
		fmt.Printf("failed to unmarshal configuration to structure with err: %s\n", err.Error())
		os.Exit(1)
	}
}
