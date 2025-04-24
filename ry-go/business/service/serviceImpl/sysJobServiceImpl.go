package serviceImpl

import (
	"errors"
	"ry-go/business/dao"
	"ry-go/business/dao/daoImpl"
	"ry-go/business/domain"
	"ry-go/common/request"

	"github.com/labstack/echo/v4"
)

// SysJobServiceImpl 系统任务服务实现
type SysJobServiceImpl struct {
	SysJobDao dao.SysJobDao
}

// NewSysJobServiceImpl 系统任务服务实现创建
func NewSysJobServiceImpl(s *daoImpl.SysJobDaoImpl) *SysJobServiceImpl {
	return &SysJobServiceImpl{
		SysJobDao: s,
	}
}

func (s *SysJobServiceImpl) Insert(e echo.Context, dept *domain.SysJob) (*domain.SysJob, error) {
	return s.SysJobDao.Insert(e.Request().Context(), dept)
}

func (s *SysJobServiceImpl) Update(e echo.Context, dept *domain.SysJob) (int64, error) {
	return s.SysJobDao.Update(e.Request().Context(), dept)
}

func (s *SysJobServiceImpl) BatchInsert(e echo.Context, list []*domain.SysJob) ([]int64, error) {
	return s.SysJobDao.BatchInsert(e.Request().Context(), list)
}

func (s *SysJobServiceImpl) SelectById(e echo.Context, id int64) (*domain.SysJob, error) {
	if id == 0 {
		return nil, errors.New("系统任务详情参数验证失败")
	}
	return s.SysJobDao.SelectById(e.Request().Context(), id)
}

func (s *SysJobServiceImpl) SelectPage(e echo.Context, param *request.SysNoticePageParam) ([]*domain.SysJob, int64, int64, error) {
	return s.SysJobDao.SelectPage(e.Request().Context(), param)
}

func (s *SysJobServiceImpl) BatchDelete(e echo.Context, ids []any) (int64, error) {
	if len(ids) <= 0 {
		return 0, errors.New("系统任务批量新增参数验证失败")
	}
	return s.SysJobDao.BatchDelete(e.Request().Context(), ids)
}

func (s *SysJobServiceImpl) UpdateJobStatus(e echo.Context, id int64, status string) (int64, error) {
	return s.SysJobDao.UpdateJobStatus(e.Request().Context(), id, status)
}

func (s *SysJobServiceImpl) UpdateJobStatusBatch(e echo.Context, list []*request.JobChangeParam) (int64, error) {
	if len(list) <= 0 {
		return 0, errors.New("系统任务批量修改状态参数验证失败")
	}
	return s.SysJobDao.UpdateJobStatusBatch(e.Request().Context(), list)
}
