package serviceImpl

import (
	"errors"
	"ry-go/business/dao"
	"ry-go/business/dao/daoImpl"
	"ry-go/business/domain"
	"ry-go/common/request"

	"github.com/labstack/echo/v4"
)

// SysNoticeServiceImpl 系统公告服务结构体
type SysNoticeServiceImpl struct {
	NoticeDao dao.SysNoticeDao
}

// NewSysNoticeServiceImpl 系统公告服务实现创建
func NewSysNoticeServiceImpl(dp *daoImpl.SysNoticeDaoImpl) *SysNoticeServiceImpl {
	return &SysNoticeServiceImpl{
		NoticeDao: dp,
	}
}

func (dp *SysNoticeServiceImpl) Insert(e echo.Context, dept *domain.SysNotice) (*domain.SysNotice, error) {
	return dp.NoticeDao.Insert(e.Request().Context(), dept)
}

func (dp *SysNoticeServiceImpl) Update(e echo.Context, dept *domain.SysNotice) (int64, error) {
	return dp.NoticeDao.Update(e.Request().Context(), dept)
}

func (dp *SysNoticeServiceImpl) BatchSave(e echo.Context, list []*domain.SysNotice) ([]int64, error) {
	if len(list) <= 0 {
		return nil, errors.New("系统公告批量新增参数验证失败")
	}
	return dp.NoticeDao.BatchInsert(e.Request().Context(), list)
}

func (dp *SysNoticeServiceImpl) SelectById(e echo.Context, id int64) (*domain.SysNotice, error) {
	if id == 0 {
		return nil, errors.New("查询参数验证失败")
	}
	return dp.NoticeDao.SelectById(e.Request().Context(), id)
}

func (dp *SysNoticeServiceImpl) SelectPage(e echo.Context, param *request.SysNoticePageParam) ([]*domain.SysNotice, int64, int64, error) {
	return dp.NoticeDao.SelectPage(e.Request().Context(), param)
}

func (dp *SysNoticeServiceImpl) BatchDelete(e echo.Context, ids []any) (int64, error) {
	if len(ids) <= 0 {
		return 0, errors.New("系统公告批量删除参数验证失败")
	}
	return dp.NoticeDao.BatchDelete(e.Request().Context(), ids)
}
