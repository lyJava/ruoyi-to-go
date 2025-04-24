package service

import (
	"ry-go/business/domain"

	"github.com/labstack/echo/v4"
)

type SysDeptLeaderService interface {
	Save(e echo.Context, leader *domain.DeptLeader) (*domain.DeptLeader, error)
	Update(e echo.Context, leader *domain.DeptLeader) (*domain.DeptLeader, error)
	// SelectById 通过ID查询
	//
	// 参数
	//   - ctx: echo.Context上下文
	//   - id: 主键ID
	//
	// 返回
	//
	//   - *domain.DeptLeader: 部门负责人信息，未找到为nil
	//   - error 返回错误
	//
	SelectById(e echo.Context, deptLeaderId int64) (*domain.DeptLeader, error)
	DeleteById(e echo.Context, id int64) (int64, error)
	BatchDelete(e echo.Context, ids []any) (int64, error)
}
