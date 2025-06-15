package serviceImpl

import (
	"context"
	"errors"
	"ry-go/business/dao"
	"ry-go/business/dao/daoImpl"
	"ry-go/business/domain"
	"ry-go/common/request"

	"github.com/labstack/echo/v4"
)

// SysLoginRecordServiceImpl 登录记录服务实现
type SysLoginRecordServiceImpl struct {
	SysLoginRecordDao dao.SysLoginRecordDao
}

// NewSysLoginRecordServiceImpl 登录记录服务实现创建
func NewSysLoginRecordServiceImpl(s *daoImpl.SysLoginRecordDaoImpl) *SysLoginRecordServiceImpl {
	return &SysLoginRecordServiceImpl{
		SysLoginRecordDao: s,
	}
}

func (s *SysLoginRecordServiceImpl) Insert(ctx context.Context, dept *domain.SysLoginRecord) (*domain.SysLoginRecord, error) {
	return s.SysLoginRecordDao.Insert(ctx, dept)
}

func (s *SysLoginRecordServiceImpl) Update(e echo.Context, dept *domain.SysLoginRecord) (int64, error) {
	return s.SysLoginRecordDao.Update(e.Request().Context(), dept)
}

func (s *SysLoginRecordServiceImpl) BatchInsert(e echo.Context, list []*domain.SysLoginRecord) ([]int64, error) {
	return s.SysLoginRecordDao.BatchInsert(e.Request().Context(), list)
}

func (s *SysLoginRecordServiceImpl) SelectById(e echo.Context, id int64) (*domain.SysLoginRecord, error) {
	if id == 0 {
		return nil, errors.New("登录记录详情参数验证失败")
	}
	return s.SysLoginRecordDao.SelectById(e.Request().Context(), id)
}

func (s *SysLoginRecordServiceImpl) SelectPage(e echo.Context, param *request.SysNoticePageParam) ([]*domain.SysLoginRecord, int64, int64, error) {
	return s.SysLoginRecordDao.SelectPage(e.Request().Context(), param)
}

func (s *SysLoginRecordServiceImpl) BatchDelete(e echo.Context, ids []any) (int64, error) {
	if len(ids) <= 0 {
		return 0, errors.New("登录记录批量新增参数验证失败")
	}
	return s.SysLoginRecordDao.BatchDelete(e.Request().Context(), ids)
}

func (s *SysLoginRecordServiceImpl) SelectLastRecode(e echo.Context, username string) (*domain.SysLoginRecord, error) {
	return s.SysLoginRecordDao.SelectLastRecode(e.Request().Context(), username)
}


func (s *SysLoginRecordServiceImpl)SelectAll(e echo.Context) ([]*domain.SysLoginRecord, error) {
	return s.SysLoginRecordDao.SelectAll(e.Request().Context())
}