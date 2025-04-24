package prop

import "time"

// JwtConfig jwt配置
type JwtConfig struct {
	Item struct {
		Secret        string        `mapstructure:"secret"`          // 加密字符串
		Duration      time.Duration `mapstructure:"duration"`        // 有效时间
		RespHeaderKey string        `mapstructure:"resp-header-key"` // 响应头键
		AuthPrefix    string        `mapstructure:"auth-prefix"`     // 请求凭证开头
	} `mapstructure:"jwt"`
}
