package show_config

import (
	"encoding/json"
	"fmt"

	"github.com/fs714/go-proj-boot/utils/config"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:          "config",
	Short:        "Show configuration information",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return printConfig()
	},
}

func printConfig() (err error) {
	fmt.Printf("configuration file: %s is in use\n\n", config.Viper.ConfigFileUsed())

	fmt.Println("configuration from viper:")
	conf := config.Viper.AllSettings()
	configJson, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		errors.Wrap(err, "marshal viper conf failed")
		return
	}
	fmt.Println(string(configJson))
	fmt.Println()

	fmt.Println("configuration from config structure:")
	configJson, err = json.MarshalIndent(config.Config, "", "  ")
	if err != nil {
		errors.Wrap(err, "marshal config stuct failed")
		return
	}
	fmt.Println(string(configJson))

	return
}
