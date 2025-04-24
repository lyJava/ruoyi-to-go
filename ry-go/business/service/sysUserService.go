package service

import (
	"ry-go/business/domain"
	"ry-go/common/request"

	"github.com/labstack/echo/v4"
)

type SysUserService interface {
	Insert(e echo.Context, user *domain.User) (*domain.User, error)
	Update(e echo.Context, user *domain.User) (int64, error)
	SelectById(c echo.Context, id int64) (*domain.User, error)
	SelectByUsername(c echo.Context, username string) (*domain.User, error)
	DeleteById(c echo.Context, id int64) (int64, error)
	BatchInert(e echo.Context, list []*domain.User) ([]int64, error)
	BatchDelete(e echo.Context, ids []any) (int64, error)
	SelectPage(e echo.Context, param *request.UserPageParam) ([]*domain.User, int64, int64, error)
	SelectList(e echo.Context, limit, offset int64) ([]*domain.UserForExcel, error)
	GetTotalCount(e echo.Context) (int64, error)
	Login(e echo.Context, user *domain.UserLogin) (*domain.User, error)
}
