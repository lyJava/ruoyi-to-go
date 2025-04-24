package service

import (
	"ry-go/business/domain"
	"ry-go/common/request"

	"github.com/labstack/echo/v4"
)

// SysPostService 岗位信息服务接口
type SysPostService interface {
	Insert(e echo.Context, sysPost *domain.SysPost) (*domain.SysPost, error)
	Update(e echo.Context, sysPost *domain.SysPost) (int64, error)
	BatchInsert(e echo.Context, list []*domain.SysPost) ([]int64, error)
	SelectById(e echo.Context, id int64) (*domain.SysPost, error)
	SelectPage(e echo.Context, param *request.PostPageParam) ([]*domain.SysPost, int64, int64, error)
	BatchDelete(e echo.Context, ids []any) (int64, error)
	SelectAll(e echo.Context) ([]*domain.SysPost, error)
}
