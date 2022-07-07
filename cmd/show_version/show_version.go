package show_version

import (
	"fmt"

	"github.com/fs714/go-proj-boot/utils/version"
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
	versionInfo := `Version Information:
Version:		%s
Commit:			%s
Build Time:		%s
Go Version:		%s`

	fmt.Println(fmt.Sprintf(versionInfo, version.BaseVersion, version.GitVersion, version.BuildTime, version.GoVersion))
}
