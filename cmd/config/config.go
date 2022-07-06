package config

import (
	"encoding/json"
	"fmt"

	"github.com/fs714/go-proj-boot/global"
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
	fmt.Printf("configuration file: %s is in use\n\n", global.Viper.ConfigFileUsed())

	fmt.Println("configuration from viper:")
	config := global.Viper.AllSettings()
	configJson, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Printf("failed to marshal viper config with err: %s\n", err.Error())
		return
	}
	fmt.Println(string(configJson))
	fmt.Println()

	fmt.Println("configuration from global config structure:")
	configJson, err = json.MarshalIndent(global.Config, "", "  ")
	if err != nil {
		fmt.Printf("failed to marshal global config with err: %s\n", err.Error())
		return
	}
	fmt.Println(string(configJson))

	return
}
