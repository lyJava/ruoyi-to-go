package daoImpl

import (
	"context"
	"errors"
	"fmt"
	"ry-go/business/domain"
	"ry-go/common/request"
	"ry-go/utils"
	"strings"

	"github.com/go-redis/redis/v8"

	"gorm.io/gorm"
)

// SysDictDataDaoImpl 字典数据Dao实现结构体
type SysDictDataDaoImpl struct {
	Gorm        *gorm.DB
	RedisClient *redis.Client
}

// NewSysDictDataDaoImpl 字典数据Dao实现创建
func NewSysDictDataDaoImpl(db *gorm.DB, client *redis.Client) *SysDictDataDaoImpl {
	return &SysDictDataDaoImpl{
		Gorm:        db,
		RedisClient: client,
	}
}

func (impl *SysDictDataDaoImpl) Insert(ctx context.Context, sysDictData *domain.SysDictData) (*domain.SysDictData, error) {
	if err := impl.Gorm.WithContext(ctx).Create(&sysDictData).Error; err != nil {
		return nil, errors.New("字典数据新增失败")
	}
	// 如果这里边有使用钩子函数，那么这里需要重新查询而不是直接返回sysDictData
	sysDictData, err := impl.SelectById(ctx, sysDictData.Id)
	if err != nil {
		return nil, errors.New("字典数据新增失败")
	}
	return sysDictData, nil
}

func (impl *SysDictDataDaoImpl) Update(ctx context.Context, sysDictData *domain.SysDictData) (int64, error) {
	// 更新操作不需要设置Model
	result := impl.Gorm.WithContext(ctx).
		Where("id = ?", sysDictData.Id).
		Omit("id").
		Save(&sysDictData)
	if result.Error != nil {
		return 0, errors.New("字典数据修改失败")
	}

	return result.RowsAffected, nil
}

func (impl *SysDictDataDaoImpl) BatchInsert(ctx context.Context, list []*domain.SysDictData) ([]int64, error) {
	if len(list) <= 0 {
		return nil, errors.New("字典数据批量新增参数验证失败")
	}

	result := impl.Gorm.WithContext(ctx).CreateInBatches(list, 50)
	if result.Error != nil {
		return nil, errors.New("字典数据批量新增失败")
	}

	var ids []int64
	for _, item := range list {
		ids = append(ids, item.Id)
	}
	return ids, nil
}

func (impl *SysDictDataDaoImpl) SelectById(ctx context.Context, id int64) (*domain.SysDictData, error) {
	var sysDictData *domain.SysDictData
	// 主键为id的时候First(&sysDictData, id)默认使用主键ID查询
	if err := impl.Gorm.WithContext(ctx).Model(&domain.SysDictData{}).First(&sysDictData, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("字典数据未找到")
		}
		return nil, errors.New("字典数据查询失败")
	}
	return sysDictData, nil
}

func (impl *SysDictDataDaoImpl) SelectPage(ctx context.Context, param *request.DictDataPageParam) ([]*domain.SysDictData, int64, int64, error) {
	page, size := utils.SetPageDefault(param.Page, param.Size)
	param.Page = page
	param.Size = size

	// 计算分页参数
	offset := (param.Page - 1) * param.Size

	// 初始化 GORM 查询
	db := impl.Gorm.WithContext(ctx).Model(&domain.SysDictData{})

	if param.DictType != "" {
		db.Where("dict_type ILIKE ?", "%"+param.DictType+"%")
	}

	if param.DictLabel != "" {
		db.Where("dict_label ILIKE ?", "%"+param.DictLabel+"%")
	}

	if param.DictStatus != "" {
		db.Where("dict_status = ?", param.DictStatus)
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
		return nil, 0, 0, errors.New("获取字典数据总记录数失败")
	}

	if param.Column == "" && param.Order == "" {
		db.Order("id ASC")
	} else {
		db.Order(fmt.Sprintf("%s %s", utils.CamelToSnakeCase(param.Column), strings.ToUpper(param.Order)))
	}

	// 查询用户列表
	var list []*domain.SysDictData
	if err := db.
		Offset(int(offset)).Limit(int(param.Size)).
		Find(&list).Error; err != nil {
		return nil, 0, 0, errors.New("获取字典数据列表失败")
	}

	return list, total, utils.CalcTotalPage(total, param.Size), nil
}

func (impl *SysDictDataDaoImpl) BatchDelete(ctx context.Context, ids []any) (int64, error) {
	if len(ids) <= 0 {
		return 0, errors.New("字典数据批量删除参数验证失败")
	}
	result := impl.Gorm.WithContext(ctx).Model(&domain.SysDictData{}).Where("id IN (?)", ids).Delete(nil)
	if result.Error != nil {
		return 0, errors.New("批量删除失败")
	}
	return result.RowsAffected, nil
}

func (impl *SysDictDataDaoImpl) SelectDictDataByType(ctx context.Context, dictType string) ([]*domain.SysDictData, error) {
	var list []*domain.SysDictData
	if err := impl.Gorm.WithContext(ctx).
		Model(&domain.SysDictData{}).
		Select("id", "dict_sort", "dict_label", "dict_value", "dict_type").
		Where("dict_type = ? AND dict_status = ?", dictType, "0").
		Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
