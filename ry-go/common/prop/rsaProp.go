package prop

// RsaConfig RSA配置
type RsaConfig struct {
	// RSA密钥配置
	Rsa struct {
		Key RsaKeyItem `mapstructure:"key"` // 密钥类型
	} `mapstructure:"rsa"`
}

type RsaKeyItem struct {
	Public  string `mapstructure:"public"`  // 公钥
	Private string `mapstructure:"private"` // 私钥
}
