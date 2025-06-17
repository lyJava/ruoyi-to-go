package serviceImpl

import (
	"ry-go/business/dao"
	"ry-go/business/dao/daoImpl"
	"ry-go/business/domain"
	"ry-go/common/request"

	"github.com/labstack/echo/v4"
)

// SysDeptServiceImpl 系统部门服务结构体
type SysDeptServiceImpl struct {
	DeptDao dao.SysDeptDao
}

// NewSysDeptServiceImpl 系统部门服务实现创建
func NewSysDeptServiceImpl(pd *daoImpl.SysDeptDaoImpl) *SysDeptServiceImpl {
	return &SysDeptServiceImpl{
		DeptDao: pd,
	}
}

// Save 新增
func (dp *SysDeptServiceImpl) Save(e echo.Context, dept *domain.Dept) (*domain.Dept, error) {
	/*tx, err := dp.Db.Begin()
	if err != nil {
		panic(err)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else {
			err = tx.Commit()
		}
	}()*/
	deptReturn, err := dp.DeptDao.Insert(e.Request().Context(), dept)
	if err != nil {
		return nil, err // 返回错误，defer 中会处理回滚
	}
	return deptReturn, nil
}

// Update 更新
func (dp *SysDeptServiceImpl) Update(c echo.Context, dept *domain.Dept) (*domain.Dept, error) {
	/*tx, err := dp.Db.Begin()
	if err != nil {
		panic(err)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else {
			err = tx.Commit()
		}
	}()*/

	deptReturn, err := dp.DeptDao.Update(c.Request().Context(), dept)
	if err != nil {
		return nil, err
	}
	return deptReturn, nil
}

func (dp *SysDeptServiceImpl) SelectById(e echo.Context, id int64) (*domain.Dept, error) {
	return dp.DeptDao.SelectById(e.Request().Context(), id)
}

// CheckIsExist 验证部门是否存在
func (dp *SysDeptServiceImpl) CheckIsExist(e echo.Context, deptId int64) (bool, error) {
	return dp.DeptDao.CheckDept(e.Request().Context(), deptId)
}

func (dp *SysDeptServiceImpl) BatchSave(e echo.Context, list []*domain.Dept) ([]int64, error) {
	return dp.DeptDao.BatchInsert(e.Request().Context(), list)
}

func (dp *SysDeptServiceImpl) BatchSaveByList(e echo.Context, list []*request.DeptInsert) ([]int64, error) {
	return dp.DeptDao.BatchInsertByList(e.Request().Context(), list)
}

func (dp *SysDeptServiceImpl) SelectDeptList(e echo.Context, param *request.DeptPageParam) ([]*domain.Dept, error) {
	return dp.DeptDao.SelectDeptList(e.Request().Context(), param)
}

func (dp *SysDeptServiceImpl) SelectPage(e echo.Context, param *request.DeptPageParam) ([]*domain.Dept, int64, int64, error) {
	return dp.DeptDao.SelectPage(e.Request().Context(), param)
}

func (dp *SysDeptServiceImpl) BatchDelete(e echo.Context, ids []any) (int64, error) {
	return dp.DeptDao.BatchDelete(e.Request().Context(), ids)
}

func (dp *SysDeptServiceImpl) SelectListByDept(e echo.Context, dept *domain.Dept) ([]*domain.Dept, error) {
	return dp.DeptDao.SelectListByDept(e.Request().Context(), dept)
}

func (dp *SysDeptServiceImpl) BuildDeptTree(list []*domain.Dept) ([]*domain.Dept, error) {
	return dp.DeptDao.BuildDeptTree(list)
}

func (dp *SysDeptServiceImpl) BuildDeptTreeSelect(list []*domain.Dept) []*domain.TreeSelect {
	return dp.DeptDao.BuildDeptTreeSelect(list)
}

func (dp *SysDeptServiceImpl) SelectListByRoleId(e echo.Context, roleId int64) ([]int64, error) {
	return dp.DeptDao.SelectListByRoleId(e.Request().Context(), roleId)
}

func (dp *SysDeptServiceImpl) SelectAll(e echo.Context) ([]*domain.Dept, error) {
	return dp.DeptDao.SelectAll(e.Request().Context())
}
