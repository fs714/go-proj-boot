package config

import "github.com/spf13/viper"

var (
	Viper         *viper.Viper
	DefaultConfig Configuration = GetDefaultConfig()
	Config        Configuration = DefaultConfig
)

func GetDefaultConfig() Configuration {
	return Configuration{
		Common: Common{
			RunMode:   "release",
			Profiling: false,
		},
		Logging: Logging{
			File:       "",
			Level:      "info",
			Format:     "console",
			MaxSize:    10,
			MaxAge:     16,
			MaxBackups: 16,
			Compress:   true,
		},
		HttpServer: HttpServer{
			Host:         "0.0.0.0",
			Port:         "9500",
			ReadTimeout:  60,
			WriteTimeout: 60,
		},
		Database: Database{
			User: "mikasa",
			Pass: "mikasa",
			Name: "titan",
			Master: DBNodeGroup{
				MaxOpenConnection: 2,
				MaxIdleConnection: 1,
				MaxLifeTime:       21600,
				Nodes: []DBNode{
					{
						Host: "127.0.0.1",
						Port: "5432",
					},
				},
			},
			Slave: DBNodeGroup{
				MaxOpenConnection: 2,
				MaxIdleConnection: 1,
				MaxLifeTime:       21600,
				Nodes: []DBNode{
					{
						Host: "127.0.0.1",
						Port: "5432",
					},
				},
			},
		},
		Jwt: Jwt{
			Secret:         "348eee0c-4c3a-4ddc-a054-a563d8f9396b",
			ExpiresTime:    3600,
			BufferTime:     86400,
			CookieName:     "token",
			CookiePath:     "/",
			CookieDomain:   "",
			SecurityCookie: false,
		},
	}
}

type Common struct {
	RunMode   string `mapstructure:"run_mode"`
	Profiling bool   `mapstructure:"profiling"`
}

type Logging struct {
	File       string `mapstructure:"file"`
	Level      string `mapstructure:"level"`
	Format     string `mapstructure:"format"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
	Compress   bool   `mapstructure:"compress"`
}

type HttpServer struct {
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	ReadTimeout  int    `mapstructure:"read_timeout"`
	WriteTimeout int    `mapstructure:"write_timeout"`
}

type DBNode struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type DBNodeGroup struct {
	MaxOpenConnection int      `mapstructure:"max_open_connection"`
	MaxIdleConnection int      `mapstructure:"max_idle_connection"`
	MaxLifeTime       int      `mapstructure:"max_life_time"`
	Nodes             []DBNode `mapstructure:"nodes"`
}

type Database struct {
	User   string      `mapstructure:"user"`
	Pass   string      `mapstructure:"pass"`
	Name   string      `mapstructure:"name"`
	Master DBNodeGroup `mapstructure:"master"`
	Slave  DBNodeGroup `mapstructure:"slave"`
}

type Jwt struct {
	Secret         string `mapstructure:"secret"`
	ExpiresTime    int    `mapstructure:"expires_time"`
	BufferTime     int    `mapstructure:"buffer_time"`
	CookieName     string `mapstructure:"cookie_name"`
	CookiePath     string `mapstructure:"cookie_path"`
	CookieDomain   string `mapstructure:"cookie_domain"`
	SecurityCookie bool   `mapstructure:"security_cookie"`
}

type Configuration struct {
	Common     Common     `mapstructure:"common"`
	Logging    Logging    `mapstructure:"logging"`
	HttpServer HttpServer `mapstructure:"http_server"`
	Database   Database   `mapstructure:"database"`
	Jwt        Jwt        `mapstructure:"jwt"`
}
