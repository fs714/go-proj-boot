package cmd

import (
	"fmt"
	"os"

	cmd_config "github.com/fs714/go-proj-boot/cmd/config"
	cmd_migrate "github.com/fs714/go-proj-boot/cmd/migrate"
	cmd_server "github.com/fs714/go-proj-boot/cmd/server"
	cmd_version "github.com/fs714/go-proj-boot/cmd/version"
	"github.com/fs714/go-proj-boot/pkg/utils/config"
	"github.com/fs714/go-proj-boot/pkg/utils/log"
	"github.com/fs714/go-proj-boot/pkg/utils/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgPath string
var logFile string
var logLevel string
var logFormat string

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

	log.Sync()
}

func init() {
	rootCmd.PersistentFlags().SortFlags = false
	rootCmd.Flags().SortFlags = false

	config.Viper = viper.New()
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgPath, "config", "c", "", "config file path")

	rootCmd.PersistentFlags().StringVarP(&logFile, "log-file", "", config.DefaultConfig.Logging.File,
		"Set logging file, stderr will be used if file is empty string")
	config.Viper.BindPFlag("logging.file", rootCmd.PersistentFlags().Lookup("log-file"))
	config.Viper.BindEnv("logging.file", "LOGGING_FILE")

	rootCmd.PersistentFlags().StringVarP(&logLevel, "log-level", "", config.DefaultConfig.Logging.Level,
		"Set logging level, could be info or debug")
	config.Viper.BindPFlag("logging.level", rootCmd.PersistentFlags().Lookup("log-level"))
	config.Viper.BindEnv("logging.level", "LOGGING_LEVEL")

	rootCmd.PersistentFlags().StringVarP(&logFormat, "log-format", "", config.DefaultConfig.Logging.Format,
		"Set logging format, could be console or json")
	config.Viper.BindPFlag("logging.format", rootCmd.PersistentFlags().Lookup("log-format"))
	config.Viper.BindEnv("logging.format", "LOGGING_FORMAT")

	cmd_version.InitStartCmd()
	cmd_server.InitStartCmd()
	cmd_migrate.InitStartCmd()

	rootCmd.AddCommand(cmd_version.StartCmd)
	rootCmd.AddCommand(cmd_config.StartCmd)
	rootCmd.AddCommand(cmd_server.StartCmd)
	rootCmd.AddCommand(cmd_migrate.StartCmd)
}

func initConfig() {
	if cfgPath != "" {
		config.Viper.SetConfigFile(cfgPath)
		config.Viper.SetConfigType("yaml")
	} else {
		config.Viper.SetConfigName("go-proj-boot")
		config.Viper.SetConfigType("yaml")

		dir, err := os.Getwd()
		if err != nil {
			fmt.Printf("failed to get current dir with err: %s\n", err.Error())
		}

		config.Viper.AddConfigPath("/etc/go-proj-boot")
		config.Viper.AddConfigPath(dir + "/conf")
	}

	err := config.Viper.ReadInConfig()
	if err != nil {
		fmt.Printf("failed to read configuration file, path: %s, err: %s\n",
			config.Viper.ConfigFileUsed(), err.Error())
		os.Exit(1)
	}

	err = config.Viper.Unmarshal(&config.Config)
	if err != nil {
		fmt.Printf("failed to unmarshal configuration to structure, path: %s, err: %s\n",
			config.Viper.ConfigFileUsed(), err.Error())
		os.Exit(1)
	}
}
