package daoImpl

import (
	"context"
	"errors"
	"fmt"
	"ry-go/business/domain"
	"ry-go/common/request"
	"ry-go/utils"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// SysLoginRecordDaoImpl 登录记录Dao实现结构体
type SysLoginRecordDaoImpl struct {
	Gorm *gorm.DB
}

// NewSysLoginRecordDaoImpl 登录记录Dao实现创建
func NewSysLoginRecordDaoImpl(db *gorm.DB) *SysLoginRecordDaoImpl {
	return &SysLoginRecordDaoImpl{
		Gorm: db,
	}
}

func (impl *SysLoginRecordDaoImpl) Insert(ctx context.Context, sysLoginRecord *domain.SysLoginRecord) (*domain.SysLoginRecord, error) {
	if err := impl.Gorm.WithContext(ctx).Create(&sysLoginRecord).Omit("info_id").Error; err != nil {
		zap.L().Sugar().Error("登录记录新增异常===", zap.Error(err))
		return nil, errors.New("登录记录新增失败")
	}
	return sysLoginRecord, nil
}

func (impl *SysLoginRecordDaoImpl) Update(ctx context.Context, sysLoginRecord *domain.SysLoginRecord) (int64, error) {
	result := impl.Gorm.WithContext(ctx).Model(&domain.SysLoginRecord{}).
		Where("info_id = ?", sysLoginRecord.InfoId).
		Omit("info_id").
		Save(&sysLoginRecord)
	if result.Error != nil {
		zap.L().Sugar().Error("登录记录修改事务执行错误===", zap.Error(result.Error))
		return 0, errors.New("登录记录修改失败")
	}

	return result.RowsAffected, nil
}

func (impl *SysLoginRecordDaoImpl) BatchInsert(ctx context.Context, list []*domain.SysLoginRecord) ([]int64, error) {
	if len(list) <= 0 {
		return nil, errors.New("登录记录批量新增参数验证失败")
	}

	zap.L().Sugar().Infof("登录记录批量新增BatchInsert参数:list===%v", list)

	result := impl.Gorm.WithContext(ctx).CreateInBatches(list, 50)
	if result.Error != nil {
		zap.L().Sugar().Error("登录记录批量新增BatchInsert错误===", zap.Error(result.Error))
		return nil, errors.New("登录记录批量新增失败")
	}

	var ids []int64
	for _, item := range list {
		ids = append(ids, item.InfoId)
	}

	zap.L().Sugar().Infof("登录记录批量新增BatchInsert成功,成功条数===%d,返回ids:===%v", result.RowsAffected, ids)
	return ids, nil
}

func (impl *SysLoginRecordDaoImpl) SelectById(ctx context.Context, id int64) (*domain.SysLoginRecord, error) {
	var sysLoginRecord *domain.SysLoginRecord
	// 主键为id的时候First(&sysLoginRecord, id)默认使用主键ID查询
	if err := impl.Gorm.WithContext(ctx).Model(&domain.SysLoginRecord{}).Where("info_id = ?", id).First(&sysLoginRecord).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("登录记录未找到")
		}
		zap.L().Error("查询SysLoginRecord失败", zap.Error(err))
		return nil, errors.New("登录记录查询失败")
	}
	return sysLoginRecord, nil
}

func (impl *SysLoginRecordDaoImpl) SelectPage(ctx context.Context, param *request.SysNoticePageParam) ([]*domain.SysLoginRecord, int64, int64, error) {
	page, size := utils.SetPageDefault(param.Page, param.Size)
	param.Page = page
	param.Size = size

	// 计算分页参数
	offset := (param.Page - 1) * param.Size
	// 初始化 GORM 查询
	db := impl.Gorm.WithContext(ctx).Model(&domain.SysLoginRecord{})
	// 查询总记录数
	var total int64
	if err := db.Count(&total).Error; err != nil {
		zap.L().Sugar().Error("获取登录记录总记录数异常:", zap.Error(err))
		return nil, 0, 0, errors.New("获取登录记录总记录数失败")
	}

	// 查询用户列表
	var list []*domain.SysLoginRecord
	if err := db.
		Offset(int(offset)).Limit(int(param.Size)).
		Find(&list).Error; err != nil {
		zap.L().Error("获取登录记录列表异常",
			zap.Error(err),
			zap.String("function", "SelectPage"),
			zap.Any("param", param),
		)
		return nil, 0, 0, errors.New("获取登录记录列表失败")
	}

	return list, total, utils.CalcTotalPage(total, param.Size), nil
}

func (impl *SysLoginRecordDaoImpl) BatchDelete(ctx context.Context, ids []any) (int64, error) {
	if len(ids) <= 0 {
		return 0, errors.New("登录记录批量删除参数验证失败")
	}
	zap.L().Sugar().Infof("登录记录批量删除参数:ids===%v", ids)
	result := impl.Gorm.WithContext(ctx).Model(&domain.SysLoginRecord{}).Where("info_id IN (?)", ids).Delete(nil)
	if result.Error != nil {
		zap.L().Sugar().Error("登录记录批量删除异常", zap.Error(result.Error))
		return 0, errors.New("批量删除失败")
	}
	zap.L().Sugar().Infof("登录记录批量删除成功执行条数===%d", result.RowsAffected)
	return result.RowsAffected, nil
}

func (impl *SysLoginRecordDaoImpl) SelectLastRecode(ctx context.Context, username string) (*domain.SysLoginRecord, error) {
	var record *domain.SysLoginRecord
	if err := impl.Gorm.WithContext(ctx).Where("user_name = ?", username).Order("info_id DESC").Take(&record).Error; err != nil {
		return nil, fmt.Errorf("获取用户:%s 最后登录记录失败", username)
	}
	return record, nil
}

func (impl *SysLoginRecordDaoImpl) SelectAll(ctx context.Context) ([]*domain.SysLoginRecord, error) {
	var (
		list    []*domain.SysLoginRecord
		allData []*domain.SysLoginRecord
		count   int64
		page    int = 1
		size    int = 2000
	)
	db := impl.Gorm.WithContext(ctx).Model(&domain.SysLoginRecord{})
	if err := db.Count(&count).Error; err != nil {
		return nil, fmt.Errorf("获取用户所有登录总条数失败")
	}

	if int64(size) >= count {
		if err := db.Find(&allData).Error; err != nil {
			return nil, fmt.Errorf("获取用户所有登录分页失败")
		}
		return allData, nil
	}

	for {
		list = list[:0]
		if err := db.Limit(size).Offset((page - 1) * size).Find(&list).Error; err != nil {
			return nil, fmt.Errorf("获取用户所有登录分页失败")
		}
		
		if len(list) == 0 {
			break
		}
		allData = append(allData, list...)
		page++
	}

	return allData, nil
}
