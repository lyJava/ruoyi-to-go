package daoImpl

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"log"
	"ry-go/business/dao"
	"ry-go/business/domain"
	"ry-go/common/request"
	"ry-go/utils"
	"strings"
	"time"

	"github.com/lib/pq"

	"github.com/go-redis/redis/v8"
	_ "gorm.io/driver/postgres"
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
		log.Printf("字典数据修改事务执行错误===%+v", result.Error)
		return 0, errors.New("字典数据修改失败")
	}

	return result.RowsAffected, nil
}

func (impl *SysDictTypeDaoImpl) BatchInsert(ctx context.Context, list []*domain.SysDictType) ([]int64, error) {
	if len(list) <= 0 {
		return nil, errors.New("字典数据批量新增参数验证失败")
	}

	log.Printf("字典数据批量新增BatchInsert参数:list===%v", list)

	result := impl.Gorm.WithContext(ctx).CreateInBatches(list, 50)
	if result.Error != nil {
		log.Printf("字典数据批量新增BatchInsert错误===%+v", result.Error)
		return nil, errors.New("字典数据批量新增失败")
	}

	var ids []int64
	for _, item := range list {
		ids = append(ids, item.DictId)
	}

	log.Printf("字典数据批量新增BatchInsert成功,成功条数===%d,返回ids:===%v", result.RowsAffected, ids)
	return ids, nil
}

func (impl *SysDictTypeDaoImpl) SelectById(ctx context.Context, id int64) (*domain.SysDictType, error) {
	var sysDictType *domain.SysDictType
	// 主键为id的时候First(&sysDictType, id)默认使用主键ID查询
	if err := impl.Gorm.WithContext(ctx).Model(&domain.SysDictType{}).Where("dict_id = ?", id).First(&sysDictType).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("字典数据未找到")
		}
		log.Printf("查询SysDictType错误===%+v", err)
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
		log.Printf("获取字典数据列表异常==%+v", err)
		return nil, 0, 0, errors.New("获取字典数据列表失败")
	}

	return list, total, utils.CalcTotalPage(total, param.Size), nil
}

func (impl *SysDictTypeDaoImpl) BatchDelete(ctx context.Context, ids []any) (int64, error) {
	if len(ids) <= 0 {
		return 0, errors.New("字典数据批量删除参数验证失败")
	}
	log.Printf("字典数据批量删除参数:ids===%v", ids)
	result := impl.Gorm.WithContext(ctx).Model(&domain.SysDictType{}).Where("dict_id IN (?)", ids).Delete(nil)
	if result.Error != nil {
		log.Printf("字典数据批量删除异常==%+v", result.Error)
		return 0, errors.New("批量删除失败")
	}
	log.Printf("字典数据批量删除成功执行条数===%d", result.RowsAffected)
	return result.RowsAffected, nil
}

func (impl *SysDictTypeDaoImpl) SelectAll(ctx context.Context) ([]*domain.SysDictType, error) {
	var (
		list    []*domain.SysDictType
		allData []*domain.SysDictType
		count   int64
		page    = 1
		size    = 2000
	)
	db := impl.Gorm.WithContext(ctx).Model(&domain.SysDictType{})
	if err := db.Count(&count).Error; err != nil {
		return nil, fmt.Errorf("获取字典类型总条数失败")
	}

	// 分页大小大于或等于总条数，则直接一次性查询
	if int64(size) >= count {
		if err := db.Find(&allData).Error; err != nil {
			return nil, errors.New("获取字典类型分页失败")
		}
		return allData, nil
	}

	for {
		list = list[:0]
		if err := db.Limit(size).Offset((page - 1) * size).Find(&list).Error; err != nil {
			return nil, errors.New("获取字典类型分页失败")
		}

		if len(list) == 0 {
			break
		}
		allData = append(allData, list...)
		page++
	}
	return allData, nil
}

func (impl *SysDictTypeDaoImpl) ClearCache(ctx context.Context) error {
	all, err := selectIdAdType(ctx, impl)
	if err != nil {
		return err
	}

	if len(all) > 0 {
		for _, dictType := range all {
			if dictType == nil {
				continue
			}

			if dictType.DictType == "" {
				continue
			}
			cacheKey := fmt.Sprintf("sys_dict:%s", dictType.DictType)
			count, delErr := utils.Del(ctx, impl.RedisClient, cacheKey)
			if delErr != nil {
				log.Printf("failed to delete cache data key==%+v,err==%+v", cacheKey, delErr)
			}
			log.Printf("字典缓存key===%s 删除数量==%d", cacheKey, count)
		}
	}
	return nil
}

func (impl *SysDictTypeDaoImpl) RefreshCache(ctx context.Context) error {
	all, err := selectIdAdType(ctx, impl)
	if err != nil {
		return err
	}

	if len(all) > 0 {
		for _, dictType := range all {
			if dictType == nil {
				continue
			}

			dictTypeStr := dictType.DictType

			dataList, dataErr := impl.DataDao.SelectDictDataByType(ctx, dictTypeStr)
			if dataErr != nil {
				log.Printf("failed to load dict data err==%+v", dataErr)
			}

			if len(dataList) == 0 {
				log.Printf("cache type = %s data list length is 0, skip it", dictTypeStr)
				continue
			}

			result, serErr := utils.Set(ctx, impl.RedisClient, fmt.Sprintf("sys_dict:%s", dictTypeStr), dataList, time.Hour)
			if serErr != nil {
				log.Printf("cache type = %s set error %+v", dictType.DictType, serErr)
				continue
			}
			log.Printf("cache type = %s set result %s", dictType.DictType, result)
		}
	}
	return nil
}

func selectIdAdType(ctx context.Context, impl *SysDictTypeDaoImpl) ([]*domain.SysDictType, error) {
	var all []*domain.SysDictType
	if err := impl.Gorm.WithContext(ctx).Model(&domain.SysDictType{}).Select("dict_id", "dict_type").Find(&all).Error; err != nil {
		return nil, err
	}
	return all, nil
}
