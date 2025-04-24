package utils

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisStore 实现 base64Captcha.Store 接口
type RedisStore struct {
	client     *redis.Client
	Prefix     string        // 键前缀（如 "captcha:"）
	expiration time.Duration // 验证码过期时间
}

func NewRedisStore(client *redis.Client, prefix string, expiration time.Duration) *RedisStore {
	return &RedisStore{
		client:     client,
		Prefix:     prefix,
		expiration: expiration,
	}
}

// Set 存储验证码到 Redis
func (s *RedisStore) Set(id string, value string) error {
	key := s.Prefix + id
	if err := s.client.Set(context.Background(), key, value, s.expiration).Err(); err != nil {
		return err
	}
	return nil
}

// Get 从 Redis 获取验证码（不删除）
func (s *RedisStore) Get(id string, clear bool) string {
	key := s.Prefix + id
	val, err := s.client.Get(context.Background(), key).Result()
	if err != nil {
		return ""
	}
	// 如果需验证后删除（clear=true），这里不主动删除，由业务逻辑控制
	return val
}

// Verify 校验验证码并删除（根据 clear 参数决定）
func (s *RedisStore) Verify(id, answer string, clear bool) bool {
	key := s.Prefix + id
	stored := s.Get(id, false)
	if stored == "" || stored != answer {
		return false
	}

	if clear {
		s.client.Del(context.Background(), key).Result()
	}
	return true
}
