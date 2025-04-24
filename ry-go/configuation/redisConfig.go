package configuation

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	redisOnce     sync.Once
	redisInstance *redis.Client
)

// GlobalStartTime 全局启动时间
var GlobalStartTime time.Time

// GlobalRedisClient 全局redis客户端
var GlobalRedisClient *redis.Client

func InitRedisClient() (*redis.Client, error) {
	var redisInitError error
	redisOnce.Do(func() {

		opts := &redis.Options{
			Addr:     RedisConfig.Redis.Address,
			Password: RedisConfig.Redis.Password,
			DB:       RedisConfig.Redis.Db,
			// 补充推荐配置项
			DialTimeout:  10 * time.Second,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			PoolSize:     20,
		}

		client := redis.NewClient(opts)
		if client == nil {
			redisInitError = errors.New("redis客户端初始化失败")
			return
		}

		// 测试连接是否有效（使用Ping）
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		if err := client.Ping(ctx).Err(); err != nil {
			redisInitError = fmt.Errorf("ping redis connection failed: %w", err)
			return
		}

		redisInstance = client
	})

	if redisInitError != nil {
		return nil, redisInitError
	}

	return redisInstance, redisInitError
}
