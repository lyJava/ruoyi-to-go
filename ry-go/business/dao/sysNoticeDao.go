package dao

import (
	"context"
	"ry-go/business/domain"
	"ry-go/common/request"
)

// SysNoticeDao 系统公告Dao接口
type SysNoticeDao interface {
	Insert(ctx context.Context, dept *domain.SysNotice) (*domain.SysNotice, error)
	Update(ctx context.Context, dept *domain.SysNotice) (int64, error)
	BatchInsert(ctx context.Context, list []*domain.SysNotice) ([]int64, error)
	SelectById(ctx context.Context, id int64) (*domain.SysNotice, error)
	SelectPage(ctx context.Context, param *request.SysNoticePageParam) ([]*domain.SysNotice, int64, int64, error)
	BatchDelete(ctx context.Context, ids []any) (int64, error)
	SelectAll(ctx context.Context) ([]*domain.SysNotice, error)
}
