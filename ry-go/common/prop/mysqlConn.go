package prop

// MysqlConfig mysql配置
type MysqlConfig struct {
	Mysql MysqlConfigItem `mapstructure:"mysql"` // mysql
}

// MysqlConfigItem mysql配置项
type MysqlConfigItem struct {
	Url      string `mapstructure:"url"`      // 连接URL
	Port     int    `mapstructure:"port"`     // 端口
	Username string `mapstructure:"username"` // 用户名
	Password string `mapstructure:"password"` // 密码
	Database string `mapstructure:"database"` // 数据库名称
}
