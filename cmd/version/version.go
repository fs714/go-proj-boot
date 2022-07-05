package version

import (
	"fmt"

	"github.com/fs714/go-proj-boot/global"
	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Run: func(cmd *cobra.Command, args []string) {
		printVersion()
	},
}

func printVersion() {
	version := `Version Information:
Version:		%s
Commit:			%s
Build Time:		%s
Go Version:		%s`

	fmt.Println(fmt.Sprintf(version, global.BaseVersion, global.GitVersion, global.BuildTime, global.GoVersion))
}
