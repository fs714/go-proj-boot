package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fs714/go-proj-boot/cmd/show_config"
	"github.com/fs714/go-proj-boot/cmd/show_version"
	"github.com/fs714/go-proj-boot/cmd/start_server"
	"github.com/fs714/go-proj-boot/pkg/utils/config"
	"github.com/fs714/go-proj-boot/pkg/utils/log"
	"github.com/fs714/go-proj-boot/pkg/utils/version"
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

	rootCmd.PersistentFlags().StringVarP(&logLevel, "log-level", "", config.DefaultConfig.Logging.Level,
		"Set logging level, could be info or debug")
	config.Viper.BindPFlag("logging.level", rootCmd.PersistentFlags().Lookup("log-level"))
	config.Viper.BindEnv("logging.level", "LOGGING_LEVEL")

	rootCmd.AddCommand(show_version.StartCmd)
	rootCmd.AddCommand(show_config.StartCmd)
	rootCmd.AddCommand(start_server.StartCmd)
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

	initLog()
}

func initLog() {
	if config.Config.Logging.File == "" {
		logger := log.New(os.Stderr, log.ParseFormat(config.Config.Logging.Format),
			log.ParseLevel(config.Config.Logging.Level), log.WithCaller(true))
		log.ResetDefault(logger)
	} else {
		var tops = []log.TeeWithRotateOption{
			{
				Filename:   config.Config.Logging.File,
				MaxSize:    config.Config.Logging.MaxSize,
				MaxAge:     config.Config.Logging.MaxAge,
				MaxBackups: config.Config.Logging.MaxBackups,
				Compress:   config.Config.Logging.Compress,
				Lef: func(lvl log.Level) bool {
					return lvl >= log.ParseLevel(config.Config.Logging.Level)
				},
				F: log.ParseFormat(config.Config.Logging.Format),
			},
		}

		logger := log.NewTeeWithRotate(tops, log.WithCaller(true))
		log.ResetDefault(logger)
	}
}
