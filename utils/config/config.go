package config

import "github.com/spf13/viper"

var (
	Viper  *viper.Viper
	Config config
)

type common struct {
	RunMode   string `mapstructure:"run_mode"`
	Profiling bool   `mapstructure:"profiling"`
}

type logging struct {
	Level string `mapstructure:"level"`
}

type httpServer struct {
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	ReadTimeout  int    `mapstructure:"read_timeout"`
	WriteTimeout int    `mapstructure:"write_timeout"`
}

type database struct {
	Host              string `mapstructure:"host"`
	Port              string `mapstructure:"port"`
	User              string `mapstructure:"user"`
	Pass              string `mapstructure:"pass"`
	Name              string `mapstructure:"name"`
	MaxOpenConnection int    `mapstructure:"max_open_connection"`
	MaxIdleConnection int    `mapstructure:"max_idle_connection"`
	MaxLifeTime       int    `mapstructure:"max_life_time"`
}

type jwt struct {
	Secret  string `mapstructure:"secret"`
	Timeout int    `mapstructure:"timeout"`
}

type config struct {
	Common     common     `mapstructure:"common"`
	Logging    logging    `mapstructure:"logging"`
	HttpServer httpServer `mapstructure:"http_server"`
	Database   database   `mapstructure:"database"`
	Jwt        jwt        `mapstructure:"jwt"`
}
