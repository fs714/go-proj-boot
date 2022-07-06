package global

type configCommon struct {
	RunMode   string `mapstructure:"run_mode"`
	Profiling bool   `mapstructure:"profiling"`
}

type configLogging struct {
	Level string `mapstructure:"level"`
}

type configHttpServer struct {
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	ReadTimeout  int    `mapstructure:"read_timeout"`
	WriteTimeout int    `mapstructure:"write_timeout"`
}

type configDatabase struct {
	Host              string `mapstructure:"host"`
	Port              string `mapstructure:"port"`
	User              string `mapstructure:"user"`
	Pass              string `mapstructure:"pass"`
	Name              string `mapstructure:"name"`
	MaxOpenConnection int    `mapstructure:"max_open_connection"`
	MaxIdleConnection int    `mapstructure:"max_idle_connection"`
	MaxLifeTime       int    `mapstructure:"max_life_time"`
}

type configJwt struct {
	Secret  string `mapstructure:"secret"`
	Timeout int    `mapstructure:"timeout"`
}

type config struct {
	Common     configCommon     `mapstructure:"common"`
	Logging    configLogging    `mapstructure:"logging"`
	HttpServer configHttpServer `mapstructure:"http_server"`
	Database   configDatabase   `mapstructure:"database"`
	Jwt        configJwt        `mapstructure:"jwt"`
}
