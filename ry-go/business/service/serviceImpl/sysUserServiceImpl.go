package serviceImpl

import (
	"errors"
	"ry-go/business/dao"
	"ry-go/business/dao/daoImpl"
	"ry-go/business/domain"
	"ry-go/common/request"

	"github.com/labstack/echo/v4"
)

// SysUserServiceImpl 系统用户服务结构体
type SysUserServiceImpl struct {
	UserDao dao.SysUserDao
}

// NewSysUserServiceImpl 系统用户服务实现创建
func NewSysUserServiceImpl(ud *daoImpl.SysUserDaoImpl) *SysUserServiceImpl {
	return &SysUserServiceImpl{
		UserDao: ud,
	}
}

func (us *SysUserServiceImpl) Insert(e echo.Context, user *domain.User) (*domain.User, error) {
	return us.UserDao.Insert(e.Request().Context(), user)
}

func (us *SysUserServiceImpl) Update(e echo.Context, user *domain.User) (int64, error) {
	return us.UserDao.Update(e.Request().Context(), user)
}

func (us *SysUserServiceImpl) SelectById(c echo.Context, id int64) (*domain.User, error) {
	if id == 0 {
		return nil, errors.New("用户ID不能为空")
	}
	return us.UserDao.SelectById(c.Request().Context(), id)
}
func (us *SysUserServiceImpl) SelectByUsername(c echo.Context, username string) (*domain.User, error) {
	if username == "" {
		return nil, errors.New("用户名不能为空")
	}
	return us.UserDao.SelectByUsername(c.Request().Context(), username)
}

func (us *SysUserServiceImpl) DeleteById(c echo.Context, id int64) (int64, error) {
	if id == 0 {
		return 0, errors.New("删除参数验证失败")
	}
	return us.UserDao.DeleteById(c.Request().Context(), id)
}

func (us *SysUserServiceImpl) BatchInert(e echo.Context, list []*domain.User) ([]int64, error) {
	if len(list) <= 0 {
		return nil, errors.New("系统用户批量新增参数验证失败")
	}
	return us.UserDao.BatchInsert(e.Request().Context(), list)
}

func (us *SysUserServiceImpl) BatchDelete(e echo.Context, ids []any) (int64, error) {
	if len(ids) <= 0 {
		return 0, errors.New("系统用户批量删除参数验证失败")
	}
	return us.UserDao.BatchDelete(e.Request().Context(), ids)
}

func (us *SysUserServiceImpl) SelectPage(e echo.Context, param *request.UserPageParam) ([]*domain.User, int64, int64, error) {
	return us.UserDao.SelectPage(e.Request().Context(), param)
}

func (us *SysUserServiceImpl) SelectList(e echo.Context, limit, offset int64) ([]*domain.UserForExcel, error) {
	return us.UserDao.SelectList(e.Request().Context(), limit, offset)
}

func (us *SysUserServiceImpl) GetTotalCount(e echo.Context) (int64, error) {
	return us.UserDao.GetTotalCount(e.Request().Context())
}

func (us *SysUserServiceImpl) Login(e echo.Context, user *domain.UserLogin) (*domain.User, error) {
	return us.UserDao.Login(e.Request().Context(), user)
}
