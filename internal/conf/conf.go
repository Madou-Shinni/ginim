package conf

var Conf = new(ProfileInfo)

type ProfileInfo struct {
	*App         `mapstructure:"app"`
	*MysqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
	*JwtConfig   `mapstructure:"jwt"`
	Github       `mapstructure:"github"`
}

// 系统配置
type App struct {
	Env        string `mapstructure:"env"`
	MachineID  int64  `mapstructure:"machineID"`
	ServerPort int    `mapstructure:"server-port"`
}

// mysql配置
type MysqlConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DBName       string `mapstructure:"dbname"`
	MaxIdleConns int    `mapstructure:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns"`
}

// redis配置
type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

// jwt配置
type JwtConfig struct {
	AccessExpire  int64  `mapstructure:"access-expire"`
	RefreshExpire int64  `mapstructure:"refresh-expire"`
	Issuer        string `mapstructure:"issuer"`
	Secret        string `mapstructure:"secret"`
}

// Github github oauth配置
type Github struct {
	ClientId     string `mapstructure:"client-id"`
	ClientSecret string `mapstructure:"client-secret"`
	RedirectUrl  string `mapstructure:"redirect-url"`
}
