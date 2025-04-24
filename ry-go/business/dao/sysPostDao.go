package dao

import (
	"context"
	"ry-go/business/domain"
	"ry-go/common/request"
)

// SysPostDao 岗位信息Dao接口
type SysPostDao interface {
	Insert(ctx context.Context, sysPost *domain.SysPost) (*domain.SysPost, error)
	Update(ctx context.Context, sysPost *domain.SysPost) (int64, error)
	BatchInsert(ctx context.Context, list []*domain.SysPost) ([]int64, error)
	SelectById(ctx context.Context, id int64) (*domain.SysPost, error)
	SelectPage(ctx context.Context, param *request.PostPageParam) ([]*domain.SysPost, int64, int64, error)
	BatchDelete(ctx context.Context, ids []any) (int64, error)
	SelectPostIdsByUserId(ctx context.Context, userId int64) ([]int64, error)
	SelectPostNameListByIds(ctx context.Context, ids []int64) ([]string, error)
	SelectAll(ctx context.Context) ([]*domain.SysPost, error)
	DeleteByUserId(ctx context.Context, userId int64) (int64, error)
	DeleteByUserIdList(ctx context.Context, userIds []int64) (int64, error)
}
