package serviceImpl

import (
	"errors"
	"ry-go/business/dao"
	"ry-go/business/dao/daoImpl"
	"ry-go/business/domain"
	"ry-go/common/request"

	"github.com/labstack/echo/v4"
)

// SysJobLogServiceImpl 系统任务日志服务实现
type SysJobLogServiceImpl struct {
	SysJobLogDao dao.SysJobLogDao
}

// NewSysJobLogServiceImpl 系统任务日志服务实现创建
func NewSysJobLogServiceImpl(s *daoImpl.SysJobLogDaoImpl) *SysJobLogServiceImpl {
	return &SysJobLogServiceImpl{
		SysJobLogDao: s,
	}
}

func (s *SysJobLogServiceImpl) Insert(e echo.Context, dept *domain.SysJobLog) (*domain.SysJobLog, error) {
	return s.SysJobLogDao.Insert(e.Request().Context(), dept)
}

func (s *SysJobLogServiceImpl) Update(e echo.Context, dept *domain.SysJobLog) (int64, error) {
	return s.SysJobLogDao.Update(e.Request().Context(), dept)
}

func (s *SysJobLogServiceImpl) BatchInsert(e echo.Context, list []*domain.SysJobLog) ([]int64, error) {
	return s.SysJobLogDao.BatchInsert(e.Request().Context(), list)
}

func (s *SysJobLogServiceImpl) SelectById(e echo.Context, id int64) (*domain.SysJobLog, error) {
	if id == 0 {
		return nil, errors.New("系统任务日志详情参数验证失败")
	}
	return s.SysJobLogDao.SelectById(e.Request().Context(), id)
}

func (s *SysJobLogServiceImpl) SelectPage(e echo.Context, param *request.SysNoticePageParam) ([]*domain.SysJobLog, int64, int64, error) {
	return s.SysJobLogDao.SelectPage(e.Request().Context(), param)
}

func (s *SysJobLogServiceImpl) BatchDelete(e echo.Context, ids []any) (int64, error) {
	if len(ids) <= 0 {
		return 0, errors.New("系统任务日志批量新增参数验证失败")
	}
	return s.SysJobLogDao.BatchDelete(e.Request().Context(), ids)
}
