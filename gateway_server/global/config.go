package global

var (
	DebugFullConfig = &DebugConfig{}
	ProxyFullConfig = &ProxyConfig{}
)

type ProxyConfig struct {
	HTTPConfig  HTTPConfig  `mapstructure:"http"`
	HTTPSConfig HTTPSConfig `mapstructure:"https"`
	TCPConfig   TCPConfig   `mapstructure:"tcp"`
	GRPCConfig  GRPCConfig  `mapstructure:"grpc"`
}

type GRPCConfig struct {
	Host string `mapstructure:"host"`
}

type GRPCServer struct {
	Host string `mapstructure:"host"`
}

type HTTPConfig struct {
	Host           string `mapstructure:"host"`
	Port           int    `mapstructure:"port"`
	ReadTimeout    int    `mapstructure:"read-timeout"`
	WriteTimeout   int    `mapstructure:"write-timeout"`
	MaxHeaderBytes int    `mapstructure:"max-header-bytes"`
}

type HTTPSConfig struct {
	Host           string `mapstructure:"host"`
	Port           int    `mapstructure:"port"`
	ReadTimeout    int    `mapstructure:"read-timeout"`
	WriteTimeout   int    `mapstructure:"write-timeout"`
	MaxHeaderBytes int    `mapstructure:"max-header-bytes"`
}

type DebugConfig struct {
	ZapConfig    ZapConfig    `mapstructure:"zap"`
	RedisConfig  RedisConfig  `mapstructure:"redis"`
	ConsulConfig ConsulConfig `mapstructure:"consul"`
	GRPCServer   GRPCServer   `mapstructure:"grpc"`
}

type TCPConfig struct {
	Host string `mapstructure:"host"`
}

type ConsulConfig struct {
	Addr string `mapstructure:"addr"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

type ZapConfig struct {
	ErrorPath string `mapstructure:"error-path"`
	OtherPath string `mapstructure:"other-path"`
	MaxSize   int    `mapstructure:"max-size"`
	MaxAge    int    `mapstructure:"max-age"`
	MaxBackup int    `mapstructure:"max-backup"`
}
