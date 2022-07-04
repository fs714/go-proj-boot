/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/fs714/go-proj-boot/global"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Run: func(cmd *cobra.Command, args []string) {
		version := `Version Information:
Version:		%s
Commit:			%s
Build Time:		%s
Go Version:		%s`

		fmt.Println(fmt.Sprintf(version, global.BaseVersion, global.GitVersion, global.BuildTime, global.GoVersion))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
