package daoImpl

import (
	"context"
	"errors"
	"fmt"
	"ry-go/business/dao"
	"ry-go/business/domain"
	"ry-go/common/request"
	"ry-go/utils"
	"strings"
	"time"

	"github.com/lib/pq"

	_ "gorm.io/driver/postgres"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// SysDictTypeDaoImpl 字典类型Dao实现结构体
type SysDictTypeDaoImpl struct {
	Gorm        *gorm.DB
	RedisClient *redis.Client
	DataDao     dao.SysDictDataDao
}

// NewSysDictTypeDaoImpl 字典类型Dao实现创建
func NewSysDictTypeDaoImpl(db *gorm.DB, rc *redis.Client, dd *SysDictDataDaoImpl) *SysDictTypeDaoImpl {
	return &SysDictTypeDaoImpl{
		Gorm:        db,
		RedisClient: rc,
		DataDao:     dd,
	}
}

// 定义错误解包函数
func unwrapPQError(err error) *pq.Error {
	for {
		if pqErr, ok := err.(*pq.Error); ok {
			return pqErr
		}
		if unwrapped := errors.Unwrap(err); unwrapped != nil {
			err = unwrapped
		} else {
			return nil
		}
	}
}

func (impl *SysDictTypeDaoImpl) Insert(ctx context.Context, sysDictType *domain.SysDictType) (*domain.SysDictType, error) {
	if err := impl.Gorm.WithContext(ctx).Create(&sysDictType).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, errors.New("不允许重复的值")
		}

		if pqErr := unwrapPQError(err); pqErr != nil && pqErr.Code == "23505" {
			return nil, errors.New("不允许重复的值")
		}
		return nil, errors.New("字典类型新增失败")
	}
	// 如果这里边有使用钩子函数，那么这里需要重新查询而不是直接返回sysDictType
	sysDictType, err := impl.SelectById(ctx, sysDictType.DictId)
	if err != nil {
		return nil, errors.New("字典类型新增失败")
	}
	return sysDictType, nil
}

func (impl *SysDictTypeDaoImpl) Update(ctx context.Context, sysDictType *domain.SysDictType) (int64, error) {
	result := impl.Gorm.WithContext(ctx).
		Where("dict_id = ?", sysDictType.DictId).
		Omit("dict_id").
		Save(&sysDictType)
	if result.Error != nil {
		zap.L().Sugar().Error("字典数据修改事务执行错误===", zap.Error(result.Error))
		return 0, errors.New("字典数据修改失败")
	}

	return result.RowsAffected, nil
}

func (impl *SysDictTypeDaoImpl) BatchInsert(ctx context.Context, list []*domain.SysDictType) ([]int64, error) {
	if len(list) <= 0 {
		return nil, errors.New("字典数据批量新增参数验证失败")
	}

	zap.L().Sugar().Infof("字典数据批量新增BatchInsert参数:list===%v", list)

	result := impl.Gorm.WithContext(ctx).CreateInBatches(list, 50)
	if result.Error != nil {
		zap.L().Sugar().Error("字典数据批量新增BatchInsert错误===", zap.Error(result.Error))
		return nil, errors.New("字典数据批量新增失败")
	}

	var ids []int64
	for _, item := range list {
		ids = append(ids, item.DictId)
	}

	zap.L().Sugar().Infof("字典数据批量新增BatchInsert成功,成功条数===%d,返回ids:===%v", result.RowsAffected, ids)
	return ids, nil
}

func (impl *SysDictTypeDaoImpl) SelectById(ctx context.Context, id int64) (*domain.SysDictType, error) {
	var sysDictType *domain.SysDictType
	// 主键为id的时候First(&sysDictType, id)默认使用主键ID查询
	if err := impl.Gorm.WithContext(ctx).Model(&domain.SysDictType{}).Where("dict_id = ?", id).First(&sysDictType).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("字典数据未找到")
		}
		zap.L().Error("查询SysDictType失败", zap.Error(err))
		return nil, errors.New("字典数据查询失败")
	}
	return sysDictType, nil
}

func (impl *SysDictTypeDaoImpl) SelectPage(ctx context.Context, param *request.DictTypePageParam) ([]*domain.SysDictType, int64, int64, error) {
	page, size := utils.SetPageDefault(param.Page, param.Size)
	param.Page = page
	param.Size = size

	// 计算分页参数
	offset := (param.Page - 1) * param.Size
	// 初始化 GORM 查询
	db := impl.Gorm.WithContext(ctx).Model(&domain.SysDictType{})

	if param.DictName != "" {
		db.Where("dict_name ILIKE ?", "%"+param.DictName+"%")
	}

	if param.DictType != "" {
		db.Where("dict_type ILIKE ?", "%"+param.DictType+"%")
	}

	if param.DictStatus != "" {
		db.Where("status = ?", param.DictStatus)
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
		zap.L().Sugar().Error("获取字典数据总记录数异常:", zap.Error(err))
		return nil, 0, 0, errors.New("获取字典数据总记录数失败")
	}

	if param.Column == "" && param.Order == "" {
		db.Order("dict_id ASC")
	} else {
		db.Order(fmt.Sprintf("%s %s", utils.CamelToSnakeCase(param.Column), strings.ToUpper(param.Order)))
	}

	// 查询用户列表
	var list []*domain.SysDictType
	if err := db.
		Offset(int(offset)).Limit(int(param.Size)).
		Find(&list).Error; err != nil {
		zap.L().Error("获取字典数据列表异常",
			zap.Error(err),
			zap.String("function", "SelectPage"),
			zap.Any("param", param),
		)
		return nil, 0, 0, errors.New("获取字典数据列表失败")
	}

	return list, total, utils.CalcTotalPage(total, param.Size), nil
}

func (impl *SysDictTypeDaoImpl) BatchDelete(ctx context.Context, ids []any) (int64, error) {
	if len(ids) <= 0 {
		return 0, errors.New("字典数据批量删除参数验证失败")
	}
	zap.L().Sugar().Infof("字典数据批量删除参数:ids===%v", ids)
	result := impl.Gorm.WithContext(ctx).Model(&domain.SysDictType{}).Where("dict_id IN (?)", ids).Delete(nil)
	if result.Error != nil {
		zap.L().Sugar().Error("字典数据批量删除异常", zap.Error(result.Error))
		return 0, errors.New("批量删除失败")
	}
	zap.L().Sugar().Infof("字典数据批量删除成功执行条数===%d", result.RowsAffected)
	return result.RowsAffected, nil
}

func (impl *SysDictTypeDaoImpl) SelectAll(ctx context.Context) ([]*domain.SysDictType, error) {
	var list []*domain.SysDictType
	if err := impl.Gorm.WithContext(ctx).Model(&domain.SysDictType{}).Where("status = ?", cast.ToString(0)).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (impl *SysDictTypeDaoImpl) ClearCache(ctx context.Context) error {
	var all []*domain.SysDictType
	if err := impl.Gorm.WithContext(ctx).Model(&domain.SysDictType{}).Select("dict_id", "dict_type").Find(&all).Error; err != nil {
		return err
	}

	if len(all) > 0 {
		var cacheKeyList []string
		for _, dictType := range all {
			if dictType == nil {
				continue
			}

			cacheKeyList = append(cacheKeyList, fmt.Sprintf("%s%s", "sys_dict:", dictType.DictType))
		}

		if _, err := utils.Del(ctx, impl.RedisClient, cacheKeyList...); err != nil {
			zap.L().Error("failed to delete cache data",
				zap.Error(err),
				zap.Any("keyList", cacheKeyList),
			)
		}
	}
	return nil
}

func (impl *SysDictTypeDaoImpl) RefreshCache(ctx context.Context) error {
	var all []*domain.SysDictType
	if err := impl.Gorm.WithContext(ctx).Model(&domain.SysDictType{}).Select("dict_id", "dict_type").Find(&all).Error; err != nil {
		return err
	}

	if len(all) > 0 {
		for _, dictType := range all {
			if dictType == nil {
				continue
			}

			dataList, err := impl.DataDao.SelectDictDataByType(ctx, dictType.DictType)
			if err != nil {
				zap.L().Error("failed to load dict data")
			}

			if len(dataList) == 0 {
				zap.L().Sugar().Errorf("cache type = %s data list length is 0, skip it", dictType.DictType)
				continue
			}

			utils.Set(ctx, impl.RedisClient, fmt.Sprintf("%s%s", "sys_dict:", dictType.DictType), dataList, time.Hour)
		}
	}
	return nil
}
