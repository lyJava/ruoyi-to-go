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

// SysNoticeDaoImpl 系系统公告Dao接口实现
type SysNoticeDaoImpl struct {
	Gorm *gorm.DB
}

// NewSysNoticeDaoImpl 系统公告Dao接口实现创建
func NewSysNoticeDaoImpl(db *gorm.DB) *SysNoticeDaoImpl {
	return &SysNoticeDaoImpl{
		Gorm: db,
	}
}

func (impl *SysNoticeDaoImpl) Insert(ctx context.Context, notice *domain.SysNotice) (*domain.SysNotice, error) {
	// 如果主键字段默认不是id，调用create需要使用Omit("主键ID")忽略
	if err := impl.Gorm.WithContext(ctx).Create(&notice).Error; err != nil {
		zap.L().Sugar().Errorf("系统公告新增异常===%+v", err)
		return nil, errors.New("系统公告新增失败")
	}
	// 如果这里边有使用钩子函数，那么这里需要重新查询而不是直接返回notice
	//zap.L().Sugar().Infof("系统公告新增返回===%v", notice)
	sysNotice, err := impl.SelectById(ctx, notice.Id)
	if err != nil {
		return nil, errors.New("系统公告新增失败")
	}
	return sysNotice, nil
}

func (impl *SysNoticeDaoImpl) Update(ctx context.Context, notice *domain.SysNotice) (int64, error) {
	result := impl.Gorm.WithContext(ctx).
		Where("id = ?", notice.Id).
		Omit("id").
		Save(&notice)
	if result.Error != nil {
		zap.L().Sugar().Errorf("系统公告修改事务执行错误===%+v", result.Error)
		return 0, errors.New("系统公告修改失败")
	}

	return result.RowsAffected, nil
}

func (impl *SysNoticeDaoImpl) BatchInsert(ctx context.Context, list []*domain.SysNotice) ([]int64, error) {
	zap.L().Sugar().Infof("系统公告批量新增BatchInsert参数:list===%v", list)

	result := impl.Gorm.WithContext(ctx).CreateInBatches(list, 50)
	if result.Error != nil {
		zap.L().Sugar().Errorf("系统公告批量新增BatchInsert错误===%+v", result.Error)
		return nil, errors.New("系统公告批量新增失败")
	}

	var ids []int64
	for _, dept := range list {
		ids = append(ids, dept.Id)
	}

	zap.L().Sugar().Infof("系统公告批量新增(BatchInsert)成功,成功条数===%d,返回ids:===%v", result.RowsAffected, ids)
	return ids, nil
}

func (impl *SysNoticeDaoImpl) SelectById(ctx context.Context, id int64) (*domain.SysNotice, error) {
	var newNotice *domain.SysNotice
	// 主键为id的时候First(&newNotice, id)默认使用主键ID查询
	if err := impl.Gorm.WithContext(ctx).Model(&domain.SysNotice{}).First(&newNotice, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("系统公告未找到")
		}
		zap.L().Sugar().Errorf("查询公告失败===%+v", err)
		return nil, errors.New("系统公告查询失败")
	}
	return newNotice, nil
}

func (impl *SysNoticeDaoImpl) SelectPage(ctx context.Context, param *request.SysNoticePageParam) ([]*domain.SysNotice, int64, int64, error) {

	page, size := utils.SetPageDefault(param.Page, param.Size)
	param.Page = page
	param.Size = size

	// 计算分页参数
	offset := (param.Page - 1) * param.Size
	// 初始化 GORM 查询
	db := impl.Gorm.WithContext(ctx).Model(&domain.SysNotice{})
	// 动态构建查询条件
	if param.Title != "" {
		db.Where("notice_title ILIKE ?", "%"+param.Title+"%")
	}
	if param.Type != "" {
		db.Where("notice_type ILIKE ?", "%"+param.Type+"%")
	}
	if param.CreateBy != "" {
		db.Where("create_by ILIKE ?", "%"+param.CreateBy+"%")
	}

	if param.Status != "" {
		db.Where("notice_status = ?", param.Status)
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
		zap.L().Sugar().Errorf("获取系统公告总记录数异常:%+v", err)
		return nil, 0, 0, errors.New("获取系统公告总记录数失败")
	}

	PageParamOrder(db, param.Column, param.Order)

	// 查询用户列表
	var list []*domain.SysNotice
	if err := db.
		Offset(int(offset)).Limit(int(param.Size)).
		//Select("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks").
		Find(&list).Error; err != nil {
		zap.L().Error("获取系统公告列表异常",
			zap.Error(err),
			zap.String("function", "SelectPage"),
			zap.Any("param", param),
		)
		return nil, 0, 0, errors.New("获取系统公告列表失败")
	}

	return list, total, utils.CalcTotalPage(total, param.Size), nil
}

func (impl *SysNoticeDaoImpl) BatchDelete(ctx context.Context, ids []any) (int64, error) {
	zap.L().Sugar().Infof("系统公告批量删除参数:ids===%v", ids)
	result := impl.Gorm.WithContext(ctx).Model(&domain.SysNotice{}).
		//Where("id IN (?)", ids).
		//Delete(nil)
		Delete(&domain.SysNotice{}, ids)
	if result.Error != nil {
		zap.L().Sugar().Errorf("系统公告批量删除异常:%+v", result.Error)
		return 0, errors.New("批量删除失败")
	}
	zap.L().Sugar().Infof("系统公告批量删除成功执行条数===%d", result.RowsAffected)
	return result.RowsAffected, nil
}

// PageParamOrder 分页排序
//
// 参数
//
//	-db: gorm的db
//	-column:列名
//	-order:排序(desc/asc)
func PageParamOrder(db *gorm.DB, column, order string) {
	if column == "" || order == "" {
		db.Order("id DESC")
	} else {
		db.Order(fmt.Sprintf("%s %s", utils.CamelToSnakeCase(column), strings.ToUpper(order)))
	}
}
