package serviceImpl

import (
	"errors"
	"fmt"
	"ry-go/business/domain"
	"ry-go/configuation"
	"ry-go/utils"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type TokenServiceImpl struct {
	redisClient *redis.Client
}

func NewTokenServiceImpl(r *redis.Client) *TokenServiceImpl {
	return &TokenServiceImpl{
		redisClient: r,
	}
}

func (s *TokenServiceImpl) CreateToken(e echo.Context, loginUser *domain.LoginUser) (string, error) {
	// 创建token
	claims := domain.MyClaims{
		Uuid:     loginUser.Uuid,
		Username: loginUser.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "ry-go-server",
			Subject:   loginUser.Uuid,
			Audience:  jwt.ClaimStrings{"test", "api-service"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(configuation.JwtConfig.Item.Duration)), // 过期时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                           // 可选：生效时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                           // 签发时间
			ID:        utils.ShortUUID(),
		},
	}

	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := tokenClaim.SignedString([]byte(configuation.JwtConfig.Item.Secret))
	if err != nil {
		return "", errors.New("用户凭证信息返回错误，请稍后重试")
	}

	cacheKey := configuation.CacheConfig.Cache.Login.Prefix + loginUser.Uuid
	// 缓存登陆用户信息，有效期为30分钟
	_, err = utils.Set(e.Request().Context(), s.redisClient,
		cacheKey,
		loginUser,
		configuation.CacheConfig.Cache.Login.Expiration*3)
	if err != nil {
		return "", errors.New("用户凭证信息返回错误，请稍后重试")
	}

	return tokenStr, nil
}

func (s *TokenServiceImpl) CacheToken(e echo.Context, uuid, token string) error {
	_ = configuation.CacheConfig.Cache.Auth.Prefix + uuid

	return nil
}

func (s *TokenServiceImpl) GetLoginUser(e echo.Context) (*domain.LoginUser, error) {
	var loginUser *domain.LoginUser

	token, err := s.GetToken(e, configuation.JwtConfig.Item.RespHeaderKey, configuation.JwtConfig.Item.AuthPrefix)
	if err != nil {
		return nil, err
	}

	if token == "" {
		return nil, errors.New("未找到登录凭证")
	}
	myClaims, err := utils.ParseToken(token, configuation.JwtConfig.Item.Secret)
	if err != nil {
		return nil, err
	}
	// 获取用户缓存信息
	userCacheKey := configuation.CacheConfig.Cache.Login.Prefix + myClaims.Uuid

	loginUser, err = utils.GetTyped(e.Request().Context(), s.redisClient, userCacheKey, loginUser)
	if err != nil {
		return nil, fmt.Errorf("解析用户信息失败: %w", err)
	}

	// 验证用户uuid一致性
	if loginUser.Uuid != myClaims.Uuid {

		return nil, errors.New("用户信息不匹配")
	}

	return loginUser, nil
}

func (s *TokenServiceImpl) SetLoginUser(e echo.Context, loginUser *domain.LoginUser) error {
	// 设置登陆时间与过期时间
	loginUser.LoginTime = time.Now().Unix()
	loginUser.ExpireTime = time.Now().Add(configuation.CacheConfig.Cache.Login.Expiration).Unix()

	loginUserCacheKey := configuation.CacheConfig.Cache.Login.Prefix + loginUser.Uuid
	_, err := utils.Set(e.Request().Context(), s.redisClient,
		loginUserCacheKey,
		loginUser,
		configuation.CacheConfig.Cache.Login.Expiration)
	if err != nil {
		return err
	}

	return nil
}

func (s *TokenServiceImpl) DeleteLoginUser(e echo.Context) error {
	loginUser, err := s.GetLoginUser(e)
	if err != nil {
		return err
	}
	// 删除缓存的用户信息
	_, err = utils.Del(e.Request().Context(), s.redisClient, configuation.CacheConfig.Cache.Login.Prefix+loginUser.Uuid)
	if err != nil {
		return err
	}

	// 删除缓存token
	_, err = utils.Del(e.Request().Context(), s.redisClient,
		configuation.CacheConfig.Cache.Auth.Prefix+loginUser.Uuid)
	if err != nil {
		return err
	}
	return nil
}

func (s *TokenServiceImpl) GetToken(e echo.Context, tokenKey, authPrefix string) (string, error) {
	return utils.GetJwtTokenFromEcho(e, tokenKey, authPrefix)
}
