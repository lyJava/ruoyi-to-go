package serviceImpl

import (
	"encoding/json"
	"errors"
	"ry-go/business/domain"
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
