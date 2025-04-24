package dao

import (
	"context"
	"ry-go/business/domain"
	"ry-go/common/request"
)

// SysUserDao 系统用户Dao
type SysUserDao interface {
	Insert(ctx context.Context, user *domain.User) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) (int64, error)
	SelectById(ctx context.Context, id int64) (*domain.User, error)
	SelectByUsername(ctx context.Context, username string) (*domain.User, error)
	DeleteById(ctx context.Context, id int64) (int64, error)
	BatchInsert(ctx context.Context, list []*domain.User) ([]int64, error)
	BatchDelete(ctx context.Context, ids []any) (int64, error)
	SelectPage(ctx context.Context, param *request.UserPageParam) ([]*domain.User, int64, int64, error)
	SelectList(ctx context.Context, limit, offset int64) ([]*domain.UserForExcel, error)
	GetTotalCount(ctx context.Context) (int64, error)
	Login(ctx context.Context, user *domain.UserLogin) (*domain.User, error)
}
