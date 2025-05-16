package serviceImpl

import (
	"encoding/json"
	"errors"
	"log"
	"ry-go/business/domain"
	"ry-go/common/model"
	"ry-go/utils"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
)

type SysMonitorServiceImpl struct {
	RedisClient *redis.Client
}

func NewSysMonitorServiceImpl(c *redis.Client) *SysMonitorServiceImpl {
	return &SysMonitorServiceImpl{
		RedisClient: c,
	}
}

func (impl *SysMonitorServiceImpl) OnlineList(e echo.Context, keyPrefix string) ([]*domain.LoginUser, error) {

	allLoginValues, err := utils.GetValuesByPrefix(e.Request().Context(), impl.RedisClient, keyPrefix)
	if err != nil {
		return nil, err
	}

	if len(allLoginValues) == 0 {
		return nil, nil
	}

	var loginUserList []*domain.LoginUser

	for _, item := range allLoginValues {
		var loginUser *domain.LoginUser
		if err = json.Unmarshal([]byte(item), &loginUser); err == nil {
			loginUserList = append(loginUserList, loginUser)
		}
	}
	return loginUserList, nil
}

func (impl *SysMonitorServiceImpl) OnlineForceLogout(e echo.Context, cacheKey string) (bool, error) {
	count, err := utils.Del(e.Request().Context(), impl.RedisClient, cacheKey)
	if err != nil {
		return false, errors.New("强退操作失败，请稍后重试")
	}
	if count == 0 {
		return false, errors.New("登陆信息已过期或已下线")
	}
	return true, nil
}

func (impl *SysMonitorServiceImpl) CacheList(e echo.Context) (map[string]any, error) {

	return nil, nil
}

func (impl *SysMonitorServiceImpl) CacheListByKey(e echo.Context, cacheKey string) ([]string, error) {
	keyList, err := utils.ScanKeysByPrefix(e.Request().Context(), impl.RedisClient, cacheKey)
	if err != nil {
		return nil, errors.New("获取缓存键名失败，请稍后重试")
	}
	return keyList, nil
}

func (impl *SysMonitorServiceImpl) CacheDetail(e echo.Context, prefix, suffix string) (*model.Cache, error) {
	cache := &model.Cache{
		Name: prefix,
		Key:  suffix,
	}
	result, err := utils.GetAny(e.Request().Context(), impl.RedisClient, prefix+":"+suffix)
	if err != nil {
		return nil, errors.New("获取缓存详情，请稍后重试")
	}
	cache.Value = utils.ToJsonFormat(result)
	return cache, err
}

func (impl *SysMonitorServiceImpl) CacheClear(e echo.Context, prefix string) error {
	count, err := utils.DelByKeyList(e.Request().Context(), impl.RedisClient, prefix)
	if err != nil {
		return errors.New("删除缓存失败，请稍后重试")
	}
	log.Printf("成功删除缓存数量===%d", count)
	return nil
}
