package global

import "github.com/spf13/viper"

var (
	BaseVersion = "0.0.1-dev"
	GitVersion  string
	GoVersion   string
	BuildTime   string
	Version     = BaseVersion + " build on " + BuildTime + " with Git Commit " + GitVersion
)

var (
	Viper  *viper.Viper
	Config config
)
