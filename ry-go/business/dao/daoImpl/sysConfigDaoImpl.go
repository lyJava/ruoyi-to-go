package daoImpl

import (
	"context"
	"errors"
	"fmt"
	"log"
	"ry-go/business/domain"
	"ry-go/common/request"
	"ry-go/utils"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// SysConfigDaoImpl 系统配置Dao实现结构体
type SysConfigDaoImpl struct {
	Gorm        *gorm.DB
	RedisClient *redis.Client
}

// NewSysConfigDaoImpl 系统配置Dao实现创建
func NewSysConfigDaoImpl(db *gorm.DB, client *redis.Client) *SysConfigDaoImpl {
	return &SysConfigDaoImpl{
		Gorm:        db,
		RedisClient: client,
	}
}

func (impl *SysConfigDaoImpl) Insert(ctx context.Context, sysConfig *domain.SysConfig) (*domain.SysConfig, error) {
	if err := impl.Gorm.WithContext(ctx).Create(&sysConfig).Error; err != nil {
		zap.L().Sugar().Error("系统配置新增异常===", zap.Error(err))
		return nil, errors.New("系统配置新增失败")
	}
	// 如果这里边有使用钩子函数，那么这里需要重新查询而不是直接返回sysConfig
	sysConfig, err := impl.SelectById(ctx, sysConfig.ConfigId)
	if err != nil {
		return nil, errors.New("系统配置新增失败")
	}
	return sysConfig, nil
}

func (impl *SysConfigDaoImpl) Update(ctx context.Context, sysConfig *domain.SysConfig) (int64, error) {
	result := impl.Gorm.WithContext(ctx).
		Where("config_id = ?", sysConfig.ConfigId).
		Omit("config_id").
		Save(&sysConfig)
	if result.Error != nil {
		zap.L().Sugar().Error("系统配置修改事务执行错误===", zap.Error(result.Error))
		return 0, errors.New("系统配置修改失败")
	}

	return result.RowsAffected, nil
}

func (impl *SysConfigDaoImpl) BatchInsert(ctx context.Context, list []*domain.SysConfig) ([]int64, error) {
	if len(list) <= 0 {
		return nil, errors.New("系统配置批量新增参数验证失败")
	}

	zap.L().Sugar().Infof("系统配置批量新增BatchInsert参数:list===%v", list)

	result := impl.Gorm.WithContext(ctx).CreateInBatches(list, 50)
	if result.Error != nil {
		zap.L().Sugar().Error("系统配置批量新增BatchInsert错误===", zap.Error(result.Error))
		return nil, errors.New("系统配置批量新增失败")
	}

	var ids []int64
	for _, item := range list {
		ids = append(ids, item.ConfigId)
	}

	zap.L().Sugar().Infof("系统配置批量新增BatchInsert成功,成功条数===%d,返回ids:===%v", result.RowsAffected, ids)
	return ids, nil
}

func (impl *SysConfigDaoImpl) SelectById(ctx context.Context, id int64) (*domain.SysConfig, error) {
	var sysConfig *domain.SysConfig
	if err := impl.Gorm.WithContext(ctx).Model(&domain.SysConfig{}).Where("config_id = ?", id).First(&sysConfig).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("系统配置未找到")
		}
		zap.L().Error("查询SysConfig失败", zap.Error(err))
		return nil, errors.New("系统配置查询失败")
	}
	return sysConfig, nil
}

func (impl *SysConfigDaoImpl) SelectPage(ctx context.Context, param *request.ConfigPageParam) ([]*domain.SysConfig, int64, int64, error) {
	page, size := utils.SetPageDefault(param.Page, param.Size)
	param.Page = page
	param.Size = size

	// 计算分页参数
	offset := (param.Page - 1) * param.Size
	// 初始化 GORM 查询
	db := impl.Gorm.WithContext(ctx).Model(&domain.SysConfig{})

	if param.ConfigName != "" {
		db.Where("config_name ILIKE ?", "%"+param.ConfigName+"%")
	}

	if param.ConfigKey != "" {
		db.Where("config_key ILIKE ?", "%"+param.ConfigKey+"%")
	}

	if param.ConfigType != "" {
		db.Where("config_type = ?", param.ConfigType)
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
		zap.L().Sugar().Error("获取系统配置总记录数异常:", zap.Error(err))
		return nil, 0, 0, errors.New("获取系统配置总记录数失败")
	}

	if param.Column == "" && param.Order == "" {
		db.Order("config_id ASC")
	} else {
		db.Order(fmt.Sprintf("%s %s", utils.CamelToSnakeCase(param.Column), strings.ToUpper(param.Order)))
	}

	// 查询用户列表
	var list []*domain.SysConfig
	if err := db.
		Offset(int(offset)).Limit(int(param.Size)).
		Find(&list).Error; err != nil {
		zap.L().Error("获取系统配置列表异常",
			zap.Error(err),
			zap.String("function", "SelectPage"),
			zap.Any("param", param),
		)
		return nil, 0, 0, errors.New("获取系统配置列表失败")
	}

	return list, total, utils.CalcTotalPage(total, param.Size), nil
}

func (impl *SysConfigDaoImpl) BatchDelete(ctx context.Context, ids []any) (int64, error) {
	if len(ids) <= 0 {
		return 0, errors.New("系统配置批量删除参数验证失败")
	}
	zap.L().Sugar().Infof("系统配置批量删除参数:ids===%v", ids)
	result := impl.Gorm.WithContext(ctx).Model(&domain.SysConfig{}).Where("config_id IN (?)", ids).Delete(nil)
	if result.Error != nil {
		zap.L().Sugar().Error("系统配置批量删除异常", zap.Error(result.Error))
		return 0, errors.New("批量删除失败")
	}
	zap.L().Sugar().Infof("系统配置批量删除成功执行条数===%d", result.RowsAffected)
	return result.RowsAffected, nil
}

func (impl *SysConfigDaoImpl) SelectConfigByKey(ctx context.Context, key string) (*domain.SysConfig, error) {
	var obj *domain.SysConfig
	if err := impl.Gorm.WithContext(ctx).
		Model(&domain.SysConfig{}).
		Where("config_key = ?", key).
		Find(&obj).Error; err != nil {
		return nil, err
	}
	return obj, nil
}

func (impl *SysConfigDaoImpl) SelectConfigList(ctx context.Context) ([]*domain.SysConfig, error) {
	var list []*domain.SysConfig
	if err := impl.Gorm.WithContext(ctx).Model(&domain.SysConfig{}).Select("config_key", "config_value").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (impl *SysConfigDaoImpl) RefreshConfigCache(ctx context.Context, list []*domain.SysConfig) error {
	if len(list) < 0 {
		return errors.New("数据为空，无法执行缓存刷新")
	}
	cacheKeyPrefix := "sys_config:"

	for _, config := range list {
		if config == nil {
			continue
		}
		cacheKey := cacheKeyPrefix + config.ConfigKey
		result, err := utils.Set(ctx, impl.RedisClient, cacheKey, config.ConfigValue, 1*time.Hour)
		if err != nil {
			return err
		}
		log.Printf("系统配置key==%s 缓存刷新结果===%s", cacheKey, result)
	}

	return nil
}

func (impl *SysConfigDaoImpl) ClearConfigCache(ctx context.Context, list []*domain.SysConfig) error {
	if len(list) < 0 {
		return errors.New("数据为空，无法执行缓存清理")
	}
	cacheKeyPrefix := "sys_config:"
	for _, config := range list {
		if config == nil {
			continue
		}
		cacheKey := cacheKeyPrefix + config.ConfigKey
		count, err := utils.Del(ctx, impl.RedisClient, cacheKey)
		if err != nil {
			return err
		}
		log.Printf("系统配置key==%s 缓存删除结果===%t", cacheKey, count > 0)
	}
	return nil
}

func (impl *SysConfigDaoImpl) SelectAll(ctx context.Context) ([]*domain.SysConfig, error) {
	var (
		list    []*domain.SysConfig
		allData []*domain.SysConfig
		count   int64
		page    = 1
		size    = 10000
	)
	db := impl.Gorm.WithContext(ctx).Model(&domain.SysConfig{})
	if err := db.Count(&count).Error; err != nil {
		zap.L().Sugar().Errorf("查询系统配置总条数错误===%v", err)
		return nil, errors.New("查询系统配置总条数失败")
	}
	zap.L().Sugar().Infof("获取系统配置分页===%d,当前分页批数===%d", count, size)

	// 分页大小大于或等于总条数，则直接一次性查询
	if int64(size) >= count {
		if err := db.Order("config_id").Find(&allData).Error; err != nil {
			zap.L().Sugar().Errorf("查询系统配置所有数据错误===%v", err)
			return nil, errors.New("查询系统配置所有数据失败")
		}
		zap.L().Sugar().Infof("分页大小大于或等于总条数一次性查询所有数据的条数===%d", len(allData))
		return allData, nil
	}

	for {
		list = list[:0]
		if err := db.Order("config_id").Limit(size).Offset((page - 1) * size).Find(&list).Error; err != nil {
			zap.L().Sugar().Errorf("查询系统配置分页错误===%v", err)
			return nil, errors.New("查询系统配置分页失败")
		}
		zap.L().Sugar().Infof("第%d次查询系统配置分页数据条数===%d", page, len(list))

		if len(list) == 0 {
			zap.L().Sugar().Infof("第%d次查询系统配置分页数据条数为0退出循环", page)
			break
		}
		allData = append(allData, list...)
		page++
	}
	return allData, nil
}