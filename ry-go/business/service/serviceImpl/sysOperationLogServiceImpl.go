package serviceImpl

import (
	"errors"
	"ry-go/business/dao"
	"ry-go/business/dao/daoImpl"
	"ry-go/business/domain"
	"ry-go/common/request"

	"github.com/labstack/echo/v4"
)

// SysOperationLogServiceImpl 操作日志服务实现
type SysOperationLogServiceImpl struct {
	SysOperationLogDao dao.SysOperationLogDao
}

// NewSysOperationLogServiceImpl 操作日志服务实现创建
func NewSysOperationLogServiceImpl(s *daoImpl.SysOperationLogDaoImpl) *SysOperationLogServiceImpl {
	return &SysOperationLogServiceImpl{
		SysOperationLogDao: s,
	}
}

func (s *SysOperationLogServiceImpl) Insert(e echo.Context, dept *domain.SysOperationLog) (*domain.SysOperationLog, error) {
	return s.SysOperationLogDao.Insert(e.Request().Context(), dept)
}

func (s *SysOperationLogServiceImpl) Update(e echo.Context, dept *domain.SysOperationLog) (int64, error) {
	return s.SysOperationLogDao.Update(e.Request().Context(), dept)
}

func (s *SysOperationLogServiceImpl) BatchInsert(e echo.Context, list []*domain.SysOperationLog) ([]int64, error) {
	return s.SysOperationLogDao.BatchInsert(e.Request().Context(), list)
}

func (s *SysOperationLogServiceImpl) SelectById(e echo.Context, id int64) (*domain.SysOperationLog, error) {
	if id == 0 {
		return nil, errors.New("操作日志详情参数验证失败")
	}
	return s.SysOperationLogDao.SelectById(e.Request().Context(), id)
}

func (s *SysOperationLogServiceImpl) SelectPage(e echo.Context, param *request.SysNoticePageParam) ([]*domain.SysOperationLog, int64, int64, error) {
	return s.SysOperationLogDao.SelectPage(e.Request().Context(), param)
}

func (s *SysOperationLogServiceImpl) BatchDelete(e echo.Context, ids []any) (int64, error) {
	if len(ids) <= 0 {
		return 0, errors.New("操作日志批量新增参数验证失败")
	}
	return s.SysOperationLogDao.BatchDelete(e.Request().Context(), ids)
}
