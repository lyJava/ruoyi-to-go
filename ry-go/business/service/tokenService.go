package service

import (
	"ry-go/business/domain"

	"github.com/labstack/echo/v4"
)

type TokenService interface {
	CreateToken(e echo.Context, loginUser *domain.LoginUser) (string, error)
	CacheToken(e echo.Context, uuid, token string) error
	GetLoginUser(e echo.Context) (*domain.LoginUser, error)
	SetLoginUser(e echo.Context, loginUser *domain.LoginUser) error
	DeleteLoginUser(e echo.Context) error
	GetToken(e echo.Context, tokenKey, authPrefix string) (string, error)
}
