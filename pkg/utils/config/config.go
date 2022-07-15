package config

import "github.com/spf13/viper"

var (
	Viper         *viper.Viper
	DefaultConfig config = GetDefaultConfig()
	Config        config = DefaultConfig
)

func GetDefaultConfig() config {
	return config{
		Common: common{
			RunMode:   "release",
			Profiling: false,
		},
		Logging: logging{
			File:       "",
			Level:      "info",
			Format:     "console",
			MaxSize:    10,
			MaxAge:     16,
			MaxBackups: 16,
			Compress:   true,
		},
		HttpServer: httpServer{
			Host:         "0.0.0.0",
			Port:         "9500",
			ReadTimeout:  60,
			WriteTimeout: 60,
		},
		Database: database{
			Host:              "127.0.0.1",
			Port:              "5432",
			User:              "mikasa",
			Pass:              "mikasa",
			Name:              "titan",
			MaxOpenConnection: 2,
			MaxIdleConnection: 1,
			MaxLifeTime:       21600,
		},
		Jwt: jwt{
			Secret:  "348eee0c-4c3a-4ddc-a054-a563d8f9396b",
			Timeout: 3600,
		},
	}
}

type common struct {
	RunMode   string `mapstructure:"run_mode"`
	Profiling bool   `mapstructure:"profiling"`
}

type logging struct {
	File       string `mapstructure:"file"`
	Level      string `mapstructure:"level"`
	Format     string `mapstructure:"format"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
	Compress   bool   `mapstructure:"compress"`
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
