package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	cmd_config "github.com/fs714/go-proj-boot/cmd/config"
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

	cmd_server.InitStartCmd()

	rootCmd.AddCommand(cmd_version.StartCmd)
	rootCmd.AddCommand(cmd_config.StartCmd)
	rootCmd.AddCommand(cmd_server.StartCmd)
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
			log.ParseLevel(config.Config.Logging.Level), true)
		log.ResetCurrentLog(logger)
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

		logger := log.NewTeeWithRotate(tops, true)
		log.ResetCurrentLog(logger)
	}
}
