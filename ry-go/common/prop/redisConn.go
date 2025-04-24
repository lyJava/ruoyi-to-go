package prop

// RedisConfig redis配置
type RedisConfig struct {
	Redis RedisConfigItem `mapstructure:"redis"`
}

// RedisConfigItem redis配置项
type RedisConfigItem struct {
	Address  string `mapstructure:"address"`  // 连接地址
	Password string `mapstructure:"password"` // 密码
	Db       int    `mapstructure:"db"`       // db
}
