package configuration

type Config struct {
	App        ConfigApp         `mapstructure:"app"`
	Server     ConfigServer      `mapstructure:"server"`
	Database   ConfigDatabase    `mapstructure:"database"`
	CorsConfig CorsConfiguration `mapstructure:"cors_config"`
	JWT        JWTConfiguration  `mapstructure:"jwt"`
}

type ConfigApp struct {
	Env      string `mapstructure:"env"`
	LogLevel string `mapstructure:"log_level"`
}

type ConfigAutomation struct {
	Periode int `mapstructure:"periode"`
}

type ConfigServer struct {
	Port int `mapstructure:"port"`
}

type ConfigDatabase struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	Port     string `mapstructure:"port"`
}

type CorsConfiguration struct {
	AllowedHeaders []string `mapstructure:"allowed_headers"`
	AllowedOrigins []string `mapstructure:"allowed_origins"`
	AllowedMethods []string `mapstructure:"allowed_methods"`
}

type JWTConfiguration struct {
	Secret string `mapstructure:"secret"`
}
