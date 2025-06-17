package service

import (
	"ry-go/business/domain"
	"ry-go/common/request"

	"github.com/labstack/echo/v4"
)

type SysDeptService interface {
	// Save 保存
	Save(e echo.Context, dept *domain.Dept) (*domain.Dept, error)
	// Update 修改
	Update(e echo.Context, dept *domain.Dept) (*domain.Dept, error)
	SelectById(e echo.Context, id int64) (*domain.Dept, error)
	// CheckIsExist 检查是否存在
	CheckIsExist(e echo.Context, deptId int64) (bool, error)
	BatchSave(e echo.Context, list []*domain.Dept) ([]int64, error)
	BatchSaveByList(e echo.Context, list []*request.DeptInsert) ([]int64, error)
	SelectDeptList(e echo.Context, param *request.DeptPageParam) ([]*domain.Dept, error)
	SelectPage(e echo.Context, param *request.DeptPageParam) ([]*domain.Dept, int64, int64, error)
	BatchDelete(e echo.Context, ids []any) (int64, error)
	SelectListByDept(e echo.Context, dept *domain.Dept) ([]*domain.Dept, error)
	BuildDeptTree(list []*domain.Dept) ([]*domain.Dept, error)
	BuildDeptTreeSelect(list []*domain.Dept) []*domain.TreeSelect
	SelectListByRoleId(e echo.Context, roleId int64) ([]int64, error)
	SelectAll(e echo.Context) ([]*domain.Dept, error)
}
