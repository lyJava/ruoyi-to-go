package prop

// ServerConfig 服务配置结构
type ServerConfig struct {
	Server ServerConfigItem `mapstructure:"server"`
}

// ServerConfigItem 服务结构配置项
type ServerConfigItem struct {
	Port          int    `mapstructure:"port"`            // 端口
	Path          string `mapstructure:"path"`            // 前缀
	Version       string `mapstructure:"version"`         // 版本
	BodyLimitSize int64  `mapstructure:"body-limit-size"` // 请求体大小(字节)
	MaxFileSize   int64  `mapstructure:"max-file-size"`   // 上传总文件大小(字节)
}
