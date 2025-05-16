package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cast"

	"github.com/go-redis/redis/v8"
)

// Set 设置数据到缓存
//
// 参数:
//   - ctx: 上下文
//   - client: Redis 客户端
//   - key: 存储键
//   - value: 存储值
//   - duration: 过期时间
//
// 返回:
//   - 操作结果字符串（如 "OK"）
//   - 错误信息
//
//goland:noinspection GoUnusedExportedFunction
func Set(ctx context.Context, client *redis.Client, key string, value any, duration time.Duration) (string, error) {
	// 自动处理结构体序列化
	var storedValue interface{}
	switch v := value.(type) {
	case string, []byte, int64, float64, bool:
		storedValue = v
	default:
		// 尝试JSON序列化
		data, err := json.Marshal(v)
		if err != nil {
			return "", fmt.Errorf("序列化失败: %w", err)
		}
		storedValue = data
	}
	// 将 timeout 转换为持续时间
	return client.Set(ctx, key, storedValue, duration).Result()
}

// Get 获取缓存
//
// 参数:
//
//	-ctx: 上下文
//	-client: Redis 客户端
//	-key: 存储键
//
// 返回:
//
//	-string: 获取的value
//	-error: 错误信息
//
//goland:noinspection GoUnusedExportedFunction
func Get(ctx context.Context, client *redis.Client, key string) (string, error) {
	return client.Get(ctx, key).Result()
}

func GetTyped[T any](ctx context.Context, client *redis.Client, key string, result T) (T, error) {
	val, err := GetAny(ctx, client, key)
	if val == nil {
		return result, fmt.Errorf("缓存已过期==%w", err)
	}

	if err != nil {
		return result, err
	}

	// 尝试类型断言
	if t, ok := val.(T); ok {
		return t, nil
	}

	// 尝试JSON反序列化
	data, _ := json.Marshal(val)
	if err = json.Unmarshal(data, &result); err == nil {
		return result, nil
	}

	return result, fmt.Errorf("类型转换失败")
}

func GetAny(ctx context.Context, client *redis.Client, key string) (any, error) {
	result, err := client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, errors.New("键不存在")
		}
		return nil, fmt.Errorf("获取缓存失败: %w", err)
	}

	if result == "" {
		return nil, errors.New("获取的为空值")
	}

	// 尝试解析为JSON对象/数组
	var jsonObj any
	if err = json.Unmarshal([]byte(result), &jsonObj); err == nil {
		return jsonObj, nil
	}

	// 尝试解析为基本类型
	switch {
	// 布尔类型检测（支持 true/false/1/0）
	case strings.EqualFold(result, "true") || result == "1":
		return true, nil
	case strings.EqualFold(result, "false") || result == "0":
		return false, nil

	// 数字类型检测
	case isInteger(result):
		if intVal, err := strconv.ParseInt(result, 10, 64); err == nil {
			return intVal, nil
		}
	case isFloat(result):
		if floatVal, err := strconv.ParseFloat(result, 64); err == nil {
			return floatVal, nil
		}
	}

	// 默认返回字符串
	return result, nil

}

// MGet 通过多个key获取值
func MGet(ctx context.Context, client *redis.Client, keys ...string) ([]any, error) {
	return client.MGet(ctx, keys...).Result()
}

// GetValuesByPrefix 根据前缀获取所有键和对应的值
func GetValuesByPrefix(ctx context.Context, client *redis.Client, prefix string) ([]string, error) {
	// 1. 用 SCAN 安全查询所有键
	keys, err := ScanKeysByPrefix(ctx, client, prefix)
	if err != nil {
		return nil, err
	}

	// 2. 批量获取所有值
	values, err := client.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, err
	}

	// 3. 组合键值对
	var result []string
	for i, _ := range keys {
		if values[i] == nil {
			continue
		}
		result = append(result, cast.ToString(values[i]))
	}

	return result, nil
}

// ScanKeysByPrefix 迭代安全获取匹配前缀的键
func ScanKeysByPrefix(ctx context.Context, client *redis.Client, prefix string) ([]string, error) {
	var (
		cursor uint64
		keys   []string
		err    error
	)

	for {
		// 一次扫描一批
		var partKeys []string
		partKeys, cursor, err = client.Scan(ctx, cursor, prefix+"*", 1000).Result()
		if err != nil {
			return nil, err
		}

		keys = append(keys, partKeys...)

		// 游标为 0 表示扫描结束
		if cursor == 0 {
			break
		}
	}

	return keys, nil
}

// 辅助函数：判断是否为整数格式
func isInteger(s string) bool {
	if s == "" {
		return false
	}
	// 允许符号位
	if s[0] == '+' || s[0] == '-' {
		s = s[1:]
	}
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

// 辅助函数：判断是否为浮点数格式
func isFloat(s string) bool {
	if s == "" {
		return false
	}
	// 允许科学计数法（暂未实现完整校验）
	parts := strings.Split(s, ".")
	if len(parts) != 2 {
		return false
	}
	return isInteger(parts[0]) && isInteger(parts[1])
}

// Del 删除缓存
//
// 参数:
//
//	-ctx: 上下文
//	-client: Redis 客户端
//	-keys: 存储键,可以是多个
//
// 返回:
//
//	-int64: 删除成功个数
//	-error: 错误信息
//
//goland:noinspection GoUnusedExportedFunction
func Del(ctx context.Context, client *redis.Client, keys ...string) (int64, error) {
	return client.Del(ctx, keys...).Result()
}

func DelByKeyList(ctx context.Context, client *redis.Client, prefix string) (int64, error) {
	keyList, err := ScanKeysByPrefix(ctx, client, prefix)
	if err != nil {
		return 0, err
	}
	return client.Del(ctx, keyList...).Result()
}

// SetWithTimeout 带默认超时的 Set
//
// 参数:
//   - ctx: 上下文
//   - client: Redis 客户端
//   - key: 存储键
//   - value: 存储值
//   - duration: 过期时间
//   - timeout: 超时时间
//
// 返回:
//   - 操作结果字符串（如 "OK"）
//   - 错误信息
//
//goland:noinspection GoUnusedExportedFunction
func SetWithTimeout(ctx context.Context, client *redis.Client, key string, value any, duration time.Duration, timeout time.Duration) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	return client.Set(ctx, key, value, duration).Result()
}

// LPush RPush 使用RPush命令往队列右边加入
//
//goland:noinspection GoUnusedExportedFunction
func LPush(ctx context.Context, client *redis.Client, key string, value ...any) (int64, error) {
	return client.LPush(ctx, key, value).Result()
}

// RPop LPop 取出并移除左边第一个元素
//
//goland:noinspection GoUnusedExportedFunction
func RPop(ctx context.Context, client *redis.Client, key string) (string, error) {
	return client.RPop(ctx, key).Result()
}

// BRPop BLPop 取出并移除左边第一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
//
//goland:noinspection GoUnusedExportedFunction
func BRPop(ctx context.Context, client *redis.Client, timeout time.Duration, key string) ([]string, error) {
	return client.BRPop(ctx, timeout, key).Result()
}

// LLen 获取数据长度
//
//goland:noinspection GoUnusedExportedFunction
func LLen(ctx context.Context, client *redis.Client, key string) (int64, error) {
	return client.LLen(ctx, key).Result()
}

// LRange 获取数据列表
//
//goland:noinspection GoUnusedExportedFunction
func LRange(ctx context.Context, client *redis.Client, key string, start, end int64) ([]string, error) {
	return client.LRange(ctx, key, start, end).Result()
}

// HSet hash相关操作
// set hash 适合存储结构
//
//goland:noinspection GoUnusedExportedFunction
func HSet(ctx context.Context, client *redis.Client, hashKey, key string, value any) (int64, error) {
	return client.HSet(ctx, hashKey, key, value).Result()
}

// HGet get Hash
//
//goland:noinspection GoUnusedExportedFunction
func HGet(ctx context.Context, client *redis.Client, hashKey, key string) (string, error) {
	return client.HGet(ctx, hashKey, key).Result()
}

// HGetWithTimeout get Hash 带超时
//
//goland:noinspection GoUnusedExportedFunction
func HGetWithTimeout(ctx context.Context, client *redis.Client, hashKey, key string, timeout time.Duration) (string, error) {
	ctx2, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	return client.HGet(ctx2, hashKey, key).Result()
}

// HGetAll 获取所以hash ,返回map
//
//goland:noinspection GoUnusedExportedFunction
func HGetAll(ctx context.Context, client *redis.Client, hashKey string) (map[string]string, error) {
	return client.HGetAll(ctx, hashKey).Result()
}

// HDel 删除一个或多个哈希表字段
//
//goland:noinspection GoUnusedExportedFunction
func HDel(ctx context.Context, client *redis.Client, hashKey string, key ...string) (int64, error) {
	return client.HDel(ctx, hashKey, key...).Result()
}

// HExists 查看哈希表的指定字段是否存在
//
//goland:noinspection GoUnusedExportedFunction
func HExists(ctx context.Context, client *redis.Client, hashKey, key string) (bool, error) {
	return client.HExists(ctx, hashKey, key).Result()
}

// SAdd -----------------Set------------------------
// 添加Set
//
//goland:noinspection GoUnusedExportedFunction
func SAdd(ctx context.Context, client *redis.Client, key string, values ...any) (int64, error) {
	return client.SAdd(ctx, key, values).Result()
}

// SCard 获取集合的成员数
//
//goland:noinspection GoUnusedExportedFunction
func SCard(ctx context.Context, client *redis.Client, key string) (int64, error) {
	return client.SCard(ctx, key).Result()
}

// SMembers 获取集合的所有成员
//
//goland:noinspection GoUnusedExportedFunction
func SMembers(ctx context.Context, client *redis.Client, key string) ([]string, error) {
	return client.SMembers(ctx, key).Result()
}

// SRem 移除集合里的某个元素
//
//goland:noinspection GoUnusedExportedFunction
func SRem(ctx context.Context, client *redis.Client, key string, value any) (int64, error) {
	return client.SRem(ctx, key, value).Result()
}

// SPop 移除并返回set的一个随机元素(SET是无序的)
//
//goland:noinspection GoUnusedExportedFunction
func SPop(ctx context.Context, client *redis.Client, key string) (any, error) {
	return client.SPop(ctx, key).Result()
}

// ZAdd ------------------ZSet-------------------------
//
//goland:noinspection GoUnusedExportedFunction
func ZAdd(ctx context.Context, client *redis.Client, key string, values []*redis.Z) (int64, error) {
	return client.ZAdd(ctx, key, values...).Result()
}

// ZIncrBy 给指定的key和值加分
//
//goland:noinspection GoUnusedExportedFunction
func ZIncrBy(ctx context.Context, client *redis.Client, key string, increment float64, member string) (float64, error) {
	return client.ZIncrBy(ctx, key, increment, member).Result()
}

// ZRevRangeWithScores 取zSet里的前n名热度的数据
//
//goland:noinspection GoUnusedExportedFunction
func ZRevRangeWithScores(ctx context.Context, client *redis.Client, key string, start, end int64) ([]redis.Z, error) {
	return client.ZRevRangeWithScores(ctx, key, start, end).Result()
}

// Expire 给指定key 设置过期时间
//
//goland:noinspection GoUnusedExportedFunction
func Expire(ctx context.Context, client *redis.Client, key string, duration time.Duration) (bool, error) {
	return client.Expire(ctx, key, duration).Result()
}

// ExpireAt 给指定Key 设置过期时间，时间格式为UNIX时间
//
//goland:noinspection GoUnusedExportedFunction
func ExpireAt(ctx context.Context, client *redis.Client, key string, duration time.Time) (bool, error) {
	return client.ExpireAt(ctx, key, duration).Result()
}

// TTL 获取key的生存时间
//
//goland:noinspection GoUnusedExportedFunction
func TTL(ctx context.Context, client *redis.Client, key string) (time.Duration, error) {
	return client.TTL(ctx, key).Result()
}

// MSet 批量 Set
//
//goland:noinspection GoUnusedExportedFunction
func MSet(ctx context.Context, client *redis.Client, pairs map[string]any) error {
	pipe := client.Pipeline()
	for key, value := range pairs {
		pipe.Set(ctx, key, value, 0)
	}
	_, err := pipe.Exec(ctx)
	return err
}

// WithPipeline 使用Pipeline执行自定义命令
//
//goland:noinspection GoUnusedExportedFunction
func WithPipeline(ctx context.Context, client *redis.Client, fn func(pipe redis.Pipeliner) error) error {
	pipe := client.Pipeline()
	if err := fn(pipe); err != nil {
		return err
	}
	_, err := pipe.Exec(ctx)
	return err
}
