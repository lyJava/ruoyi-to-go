package prop

import "time"

// CacheConfig 缓存配置
type CacheConfig struct {
	// RSA密钥配置
	Cache struct {
		Captcha KeyItem `mapstructure:"captcha"` // 验证码
		Login   KeyItem `mapstructure:"login"`   // 用户登陆
		Auth    struct {
			Prefix     string        `mapstructure:"prefix"`
			Expiration time.Duration `mapstructure:"expiration"`
		} `mapstructure:"auth"`
	} `mapstructure:"cache"`
}

type KeyItem struct {
	Prefix     string        `mapstructure:"prefix"`     // 用户登陆
	Expiration time.Duration `mapstructure:"expiration"` // 设置有效时间
}
