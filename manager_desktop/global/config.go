package global

var (
	DebugFullConfig = &DebugConfig{}
	ProxyFullConfig = &ProxyConfig{}
)

type ProxyConfig struct {
	HTTPConfig  HTTPConfig  `mapstructure:"http"`
	HTTPSConfig HTTPSConfig `mapstructure:"https"`
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
	ZapConfig     ZapConfig     `mapstructure:"zap"`
	MysqlConfig   MysqlConfig   `mapstructure:"mysql"`
	ServerConfig  ServerConfig  `mapstructure:"server"`
	ClusterConfig ClusterConfig `mapstructure:"cluster"`
	RedisConfig   RedisConfig   `mapstructure:"redis"`
	GRPCServer    GRPCServer    `mapstructure:"grpc"`
}

type GRPCServer struct {
	Host string `mapstructure:"host"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

type ClusterConfig struct {
	Host          string `mapstructure:"host"`
	Port          int    `mapstructure:"port"`
	SSLPort       int    `mapstructure:"ssl-port"`
	WebSocketPort int    `mapstructure:"websocket-port"`
}

type MysqlConfig struct {
	Dsn             string `mapstructure:"dsn"`
	MaxOpenConn     int    `mapstructure:"max-open-conn"`
	MaxIdleConn     int    `mapstructure:"max-idle-conn"`
	MaxConnLifeTime int    `mapstructure:"max-conn-life-time"`
}

type ZapConfig struct {
	ErrorPath string `mapstructure:"error-path"`
	OtherPath string `mapstructure:"other-path"`
	MaxSize   int    `mapstructure:"max-size"`
	MaxAge    int    `mapstructure:"max-age"`
	MaxBackup int    `mapstructure:"max-backup"`
}

type ServerConfig struct {
	Addr string `mapstructure:"addr"`
}
