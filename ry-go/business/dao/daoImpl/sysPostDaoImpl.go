package daoImpl

import (
	"context"
	"errors"
	"fmt"
	"ry-go/business/domain"
	"ry-go/common/request"
	"ry-go/utils"
	"strings"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// SysPostDaoImpl 岗位信息Dao实现结构体
type SysPostDaoImpl struct {
	Gorm *gorm.DB
}

// NewSysPostDaoImpl 岗位信息Dao实现创建
func NewSysPostDaoImpl(db *gorm.DB) *SysPostDaoImpl {
	return &SysPostDaoImpl{
		Gorm: db,
	}
}

func (impl *SysPostDaoImpl) Insert(ctx context.Context, sysPost *domain.SysPost) (*domain.SysPost, error) {
	if err := impl.Gorm.WithContext(ctx).Create(&sysPost).Error; err != nil {
		return nil, errors.New("岗位信息新增失败")
	}
	// 如果这里边有使用钩子函数，那么这里需要重新查询而不是直接返回sysPost
	insertPost, err := impl.SelectById(ctx, sysPost.Id)
	if err != nil {
		return nil, errors.New("岗位信息新增失败")
	}
	return insertPost, nil
}

func (impl *SysPostDaoImpl) Update(ctx context.Context, sysPost *domain.SysPost) (int64, error) {
	result := impl.Gorm.WithContext(ctx).
		Where("id = ?", sysPost.Id).
		Omit("id").
		Save(&sysPost)
	if result.Error != nil {
		zap.L().Sugar().Error("岗位信息修改事务执行错误===", zap.Error(result.Error))
		return 0, errors.New("岗位信息修改失败")
	}

	return result.RowsAffected, nil
}

func (impl *SysPostDaoImpl) BatchInsert(ctx context.Context, list []*domain.SysPost) ([]int64, error) {
	if len(list) <= 0 {
		return nil, errors.New("岗位信息批量新增参数验证失败")
	}

	zap.L().Sugar().Infof("岗位信息批量新增BatchInsert参数:list===%v", list)

	result := impl.Gorm.WithContext(ctx).CreateInBatches(list, 50)
	if result.Error != nil {
		zap.L().Sugar().Error("岗位信息批量新增BatchInsert错误===", zap.Error(result.Error))
		return nil, errors.New("岗位信息批量新增失败")
	}

	var ids []int64
	for _, item := range list {
		ids = append(ids, item.Id)
	}

	zap.L().Sugar().Infof("岗位信息批量新增BatchInsert成功,成功条数===%d,返回ids:===%v", result.RowsAffected, ids)
	return ids, nil
}

func (impl *SysPostDaoImpl) SelectById(ctx context.Context, id int64) (*domain.SysPost, error) {
	var sysPost *domain.SysPost
	// 主键为id的时候First(&sysPost, id)默认使用主键ID查询
	if err := impl.Gorm.WithContext(ctx).Model(&domain.SysPost{}).First(&sysPost, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("岗位信息未找到")
		}
		return nil, errors.New("岗位信息查询失败")
	}
	return sysPost, nil
}

func (impl *SysPostDaoImpl) SelectPage(ctx context.Context, param *request.PostPageParam) ([]*domain.SysPost, int64, int64, error) {
	page, size := utils.SetPageDefault(param.Page, param.Size)
	param.Page = page
	param.Size = size

	// 计算分页参数
	offset := (param.Page - 1) * param.Size

	// 初始化 GORM db
	db := impl.Gorm.WithContext(ctx).Model(&domain.SysPost{})

	if param.PostName != "" {
		db.Where("post_name ILIKE ?", "%"+param.PostName+"%")
	}

	if param.PostCode != "" {
		db.Where("post_code ILIKE ?", "%"+param.PostCode+"%")
	}

	if param.PostStatus != "" {
		db.Where("post_status = ?", param.PostStatus)
	}

	// 时间范围处理
	if param.BeginTime != "" && param.EndTime != "" {
		if begin, end, err := parseTimeRange(param.BeginTime, param.EndTime); err == nil {
			db.Where("create_time BETWEEN ? AND ?", begin, end)
		}
	}

	// 查询总记录数
	var total int64
	if err := db.Count(&total).Error; err != nil {
		zap.L().Sugar().Error("获取岗位信息总记录数异常:", zap.Error(err))
		return nil, 0, 0, errors.New("获取岗位信息总记录数失败")
	}

	if param.Column == "" && param.Order == "" {
		db.Order("id ASC")
	} else {
		db.Order(fmt.Sprintf("%s %s", utils.CamelToSnakeCase(param.Column), strings.ToUpper(param.Order)))
	}

	// 查询用户列表
	var list []*domain.SysPost
	if err := db.
		Offset(int(offset)).Limit(int(param.Size)).
		Find(&list).Error; err != nil {
		return nil, 0, 0, errors.New("获取岗位信息列表失败")
	}

	return list, total, utils.CalcTotalPage(total, param.Size), nil
}

func (impl *SysPostDaoImpl) BatchDelete(ctx context.Context, ids []any) (int64, error) {
	if len(ids) <= 0 {
		return 0, errors.New("岗位信息批量删除参数验证失败")
	}
	zap.L().Sugar().Infof("岗位信息批量删除参数:ids===%v", ids)
	result := impl.Gorm.WithContext(ctx).Model(&domain.SysPost{}).Where("id IN (?)", ids).Delete(nil)
	if result.Error != nil {
		zap.L().Sugar().Error("岗位信息批量删除异常", zap.Error(result.Error))
		return 0, errors.New("批量删除失败")
	}
	zap.L().Sugar().Infof("岗位信息批量删除成功执行条数===%d", result.RowsAffected)
	return result.RowsAffected, nil
}

//goland:noinspection SqlResolve,SqlCaseVsIf,SqlNoDataSourceInspection
func (impl *SysPostDaoImpl) SelectPostIdsByUserId(ctx context.Context, userId int64) ([]int64, error) {
	var postIds []int64
	querySql := `SELECT post_id FROM sys_user_post WHERE user_id = ?`
	if err := impl.Gorm.WithContext(ctx).Raw(querySql, userId).Find(&postIds).Error; err != nil {
		return nil, errors.New("查询用户岗位失败")
	}
	return postIds, nil
}

//goland:noinspection SqlResolve,SqlCaseVsIf,SqlNoDataSourceInspection
func (impl *SysPostDaoImpl) SelectPostNameListByIds(ctx context.Context, ids []int64) ([]string, error) {
	var postNameList []string
	querySql := `SELECT post_name FROM sys_post WHERE id IN ?`
	if err := impl.Gorm.WithContext(ctx).Raw(querySql, ids).Find(&postNameList).Error; err != nil {
		return nil, errors.New("查询用户岗位名称失败")
	}
	return postNameList, nil
}

func (impl *SysPostDaoImpl) SelectAll(ctx context.Context) ([]*domain.SysPost, error) {
	var list []*domain.SysPost
	if err := impl.Gorm.WithContext(ctx).Model(&domain.SysPost{}).Find(&list).Error; err != nil {
		return nil, errors.New("查询所有岗位失败")
	}
	return list, nil
}

//goland:noinspection SqlResolve,SqlCaseVsIf,SqlNoDataSourceInspection
func (impl *SysPostDaoImpl) DeleteByUserId(ctx context.Context, userId int64) (int64, error) {
	var count int64
	querySql := `DELETE FROM sys_user_post WHERE user_id = ?`
	if err := impl.Gorm.WithContext(ctx).Raw(querySql, userId).Count(&count).Error; err != nil {
		return 0, fmt.Errorf("用户ID:%d删除绑定岗位失败, %w", userId, err)
	}
	return count, nil
}

//goland:noinspection SqlResolve,SqlCaseVsIf,SqlNoDataSourceInspection
func (impl *SysPostDaoImpl) DeleteByUserIdList(ctx context.Context, userIds []int64) (int64, error) {
	var count int64
	querySql := `DELETE FROM sys_user_post WHERE user_id IN ?`
	if err := impl.Gorm.WithContext(ctx).Raw(querySql, userIds).Count(&count).Error; err != nil {
		return 0, fmt.Errorf("用户ID:%v删除绑定岗位失败, %w", userIds, err)
	}
	return count, nil
}
