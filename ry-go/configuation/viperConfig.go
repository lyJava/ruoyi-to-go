package configuation

import (
	"fmt"
	"go.uber.org/zap"
	"ry-go/common/prop"
	"ry-go/utils"
	"sync"

	"github.com/spf13/viper"
)

// ReadConfig 读取配置
//
// 参数
//
//	-path：配置路径
//	-name：配置文件名(不包括后缀)
//	-configType：配置文件路径
//
// 返回
//
//	-viper：*viper.Viper
func ReadConfig(path, name, configType string) *viper.Viper {
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType(configType)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("viper config read error===%w", err))
	}
	return viper.GetViper()
}

var (
	loadOnce       sync.Once             // 只执行一次
	ViperConfig    *viper.Viper          // 全局Viper配置
	ServerConfig   prop.ServerConfig     // 全局服务配置
	PostgresConfig prop.PostgresqlConfig // 全局postgres配置
	RedisConfig    prop.RedisConfig      // 全局redis配置
	RsaConfig      prop.RsaConfig        // 全局RSA密钥配置
	CacheConfig    prop.CacheConfig      // 全局缓存键等配置
	JwtConfig      prop.JwtConfig        // 全局JWT配置
)

// InitGlobalConfig 初始化全局配置
func InitGlobalConfig() error {
	var loadErr error
	loadOnce.Do(func() {
		ViperConfig = ReadConfig("resources", "application", "yml")
		if err := ViperConfig.Unmarshal(&ServerConfig); err != nil {
			loadErr = fmt.Errorf("failed to unmarshal server config: %w", err)
			return
		}

		if err := ViperConfig.Unmarshal(&PostgresConfig); err != nil {
			loadErr = fmt.Errorf("failed to unmarshal postgres config: %w", err)
			return
		}

		if err := ViperConfig.Unmarshal(&RedisConfig); err != nil {
			loadErr = fmt.Errorf("failed to unmarshal redis config: %w", err)
			return
		}

		if err := ViperConfig.Unmarshal(&RsaConfig); err != nil {
			loadErr = fmt.Errorf("failed to unmarshal rsa config: %w", err)
			return
		}

		if err := ViperConfig.Unmarshal(&CacheConfig); err != nil {
			loadErr = fmt.Errorf("failed to unmarshal cache config: %w", err)
			return
		}

		if err := ViperConfig.Unmarshal(&JwtConfig); err != nil {
			loadErr = fmt.Errorf("failed to unmarshal jwt config: %w", err)
			return
		}
	})

	zap.L().Sugar().Infof("全局viper配置路径===%s", ViperConfig.ConfigFileUsed())
	zap.L().Sugar().Infof("echo配置信息===\n%s", utils.ToJsonFormat(ServerConfig))
	zap.L().Sugar().Infof("postgresql配置信息===\n%s", utils.ToJsonFormat(PostgresConfig))
	zap.L().Sugar().Infof("redis配置信息===\n%s", utils.ToJsonFormat(RedisConfig))
	zap.L().Sugar().Infof("cache配置信息===\n%s", utils.ToJsonFormat(CacheConfig))
	zap.L().Sugar().Infof("jwt配置信息===\n%s", utils.ToJsonFormat(JwtConfig))

	return loadErr
}
