package daoImpl

import (
	"context"
	"errors"
	"fmt"
	"ry-go/business/dao"
	"ry-go/business/domain"
	"ry-go/common/request"
	"ry-go/utils"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// SysDeptDaoImpl 系统部门Dao结构
type SysDeptDaoImpl struct {
	Gorm    *gorm.DB
	RoleDao dao.SysRoleDao
}

// NewSysDeptDaoImpl 系统部门Dao创建
func NewSysDeptDaoImpl(db *gorm.DB, r *SysRoleDaoImpl) *SysDeptDaoImpl {
	return &SysDeptDaoImpl{
		Gorm:    db,
		RoleDao: r,
	}
}

//goland:noinspection SqlResolve,SqlCaseVsIf,SqlNoDataSourceInspection,SqlDialectInspection
func (impl *SysDeptDaoImpl) Insert(ctx context.Context, dept *domain.Dept) (*domain.Dept, error) {
	exist, err := impl.CheckDeptByName(ctx, dept.DeptName)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("%s已存在，请勿重复添加", dept.DeptName)
	}

	dept.DelFlag = "0"
	// 如果主键字段默认不是id，调用create需要使用Omit("主键ID")忽略
	result := impl.Gorm.WithContext(ctx).
		//Model(&domain.Dept{}).
		//Omit("dept_id").
		Create(&dept)
	if result.Error != nil {
		zap.L().Sugar().Errorf("部门新增异常===%v", result.Error)
		return nil, errors.New("部门新增失败")
	}
	zap.L().Sugar().Infof("部门新增返回主键ID===%v", dept.Id)
	deptNew, err := impl.SelectDeptById(ctx, dept.Id)
	if err != nil {
		return nil, errors.New("部门新增失败")
	}
	return deptNew, nil
}

//goland:noinspection SqlResolve,SqlCaseVsIf,SqlNoDataSourceInspection,SqlDialectInspection
func (impl *SysDeptDaoImpl) Update(ctx context.Context, dept *domain.Dept) (*domain.Dept, error) {
	exist, err := impl.CheckDeptByIdAndName(ctx, dept.Id, dept.DeptName)
	if err != nil {
		zap.L().Sugar().Errorw("Failed to check department by id and name",
			"部门ID===", dept.Id,
			"部门名称===", dept.DeptName,
			"error===", err)
		return nil, errors.New("部门名称验证失败")
	}

	checkDept, _ := impl.CheckDept(ctx, dept.Id)
	zap.L().Sugar().Infof("当前部门启用状态===%+v", checkDept)

	if exist {
		errExist := fmt.Errorf("部门名称:%s 已存在", dept.DeptName)
		zap.L().Sugar().Error(errExist)
		return nil, errExist
	}

	result := impl.Gorm.WithContext(ctx).
		Where("id = ?", dept.Id).
		//Omit("id").
		Save(&dept)
	if result.Error != nil {
		zap.L().Sugar().Errorf("部门修改事务执行错误===%+v", result.Error)
		return nil, errors.New("部门修改失败")
	}
	rowsAffected := result.RowsAffected

	var deptNew *domain.Dept
	if rowsAffected == 1 {
		deptNew, err = impl.SelectDeptById(ctx, dept.Id)
		if err != nil {
			zap.L().Sugar().Errorf("当前部门==%d不存在", dept.Id)
			return nil, errors.New("当前部门不存在")
		}
	}
	return deptNew, nil
}

func (impl *SysDeptDaoImpl) SelectById(ctx context.Context, id int64) (*domain.Dept, error) {
	var dept *domain.Dept
	if err := impl.Gorm.WithContext(ctx).First(&dept, id).Error; err != nil {
		return nil, err
	}
	return dept, nil
}

//goland:noinspection SqlResolve,SqlCaseVsIf,SqlNoDataSourceInspection,SqlDialectInspection
func (impl *SysDeptDaoImpl) CheckDept(ctx context.Context, deptId int64) (bool, error) {
	var dept domain.Dept

	result := impl.Gorm.WithContext(ctx).Model(&domain.Dept{}).
		Where("del_flag", "0").First(&dept, deptId)
	if result.Error == gorm.ErrRecordNotFound {
		zap.L().Sugar().Errorf("当前部门ID:%d,名称:%s查询不存在 ===%+v", deptId, dept.DeptName, result.Error)
		return false, errors.New("当前部门不存在")
	}

	if dept.DeptStatus == "1" {
		zap.L().Sugar().Errorf("当前部门ID:%d,名称:%s暂未启用", deptId, dept.DeptName)
		return false, errors.New("当前部门未启用")
	}
	return true, nil
}

//goland:noinspection SqlResolve,SqlCaseVsIf,SqlNoDataSourceInspection,SqlDialectInspection
func (impl *SysDeptDaoImpl) CheckDeptByName(ctx context.Context, deptName string) (bool, error) {
	var count int64
	if err := impl.Gorm.WithContext(ctx).Model(&domain.Dept{}).
		Where("dept_name", deptName).
		Count(&count).Error; err != nil {
		zap.L().Sugar().Error("通过名称查询部门条数异常===%+v", err)
		return false, errors.New("通过名称查询部门失败")
	}
	return count > 0, nil
}

//goland:noinspection SqlResolve,SqlCaseVsIf,SqlNoDataSourceInspection,SqlDialectInspection
func (impl *SysDeptDaoImpl) SelectDeptByName(ctx context.Context, deptName string) (*domain.Dept, error) {
	var dept domain.Dept
	result := impl.Gorm.WithContext(ctx).Model(&domain.Dept{}).Where("dept_name", deptName).First(&dept)
	if result.Error != nil {
		zap.L().Sugar().Error("通过名称查询部门异常===%+v", result.Error)
		return nil, errors.New("通过名称查询部门失败")
	}
	return &dept, nil
}

//goland:noinspection SqlResolve,SqlCaseVsIf,SqlNoDataSourceInspection,SqlDialectInspection
func (impl *SysDeptDaoImpl) SelectDeptById(ctx context.Context, id int64) (*domain.Dept, error) {
	var dept domain.Dept
	// 主键设置为id，可以直接使用如下方式查询
	result := impl.Gorm.WithContext(ctx).First(&dept, id)
	if result.Error != nil {
		zap.L().Sugar().Error("通过ID查询部门异常===%+v", result.Error)
		return nil, errors.New("通过ID查询部门失败")
	}
	return &dept, nil
}

func (impl *SysDeptDaoImpl) BatchInsert(ctx context.Context, list []*domain.Dept) ([]int64, error) {
	if len(list) <= 0 {
		return nil, errors.New("部门批量新增参数验证失败")
	}

	result := impl.Gorm.WithContext(ctx).CreateInBatches(list, 50)
	if result.Error != nil {
		zap.L().Sugar().Errorf("动态批量新增BatchInsert错误===%+v", result.Error)
		return nil, errors.New("部门批量新增失败")
	}

	var ids []int64

	for _, dept := range list {
		ids = append(ids, dept.Id)
	}

	zap.L().Sugar().Infof("部门批量新增BatchInsert成功条数===%d,ids===%+v", result.RowsAffected, ids)

	return ids, nil
}

func (impl *SysDeptDaoImpl) BatchInsertByList(ctx context.Context, list []*request.DeptInsert) ([]int64, error) {
	var ids []int64

	if len(list) < 0 {
		return nil, errors.New("导入的数据条数必须大于0")
	}

	if err := impl.Gorm.WithContext(ctx).Table("sys_dept").CreateInBatches(&list, 1000).Error; err != nil {
		zap.L().Sugar().Errorf("动态批量新增BatchInsertByList错误===%+v", err)
		return nil, errors.New("部门批量新增失败")
	}

	for _, item := range list {
		ids = append(ids, item.Id)
	}

	return ids, nil
}

func (impl *SysDeptDaoImpl) SelectPage(ctx context.Context, param *request.DeptPageParam) ([]*domain.Dept, int64, int64, error) {
	page, size := utils.SetPageDefault(param.Page, param.Size)
	param.Page = page
	param.Size = size

	// 计算分页参数
	offset := (param.Page - 1) * param.Size
	// 初始化 GORM 查询
	db := impl.Gorm.WithContext(ctx).Model(&domain.Dept{})
	// 动态构建查询条件
	if param.DeptName != "" {
		db.Where(`dept_name ILIKE ?`, "%"+param.DeptName+"%")
	}
	if param.Remarks != "" {
		db.Where(`remarks ILIKE ?`, "%"+param.Remarks+"%")
	}
	if param.Status != "" {
		db.Where("dept_status = ?", param.Status)
	}
	if param.UpdateBy != "" {
		db.Where(`update_by ILIKE ?`, "%"+param.UpdateBy+"%")
	}

	// 查询总记录数
	var total int64
	if err := db.Count(&total).Error; err != nil {
		zap.L().Sugar().Errorf("获取总记录数异常: %+v", err)
		return nil, 0, 0, errors.New("获取总记录数失败")
	}

	// 查询用户列表
	var list []*domain.Dept
	if err := db.
		Offset(int(offset)).Limit(int(param.Size)).
		Select("id", "parent_id", "ancestors", "dept_name", "order_num", "dept_status", "create_by",
			"create_time", "update_by", "update_time", "remarks", "del_flag", "leader", "phone", "email").
		Find(&list).Error; err != nil {
		zap.L().Error("获取部门列表异常",
			zap.Error(err),
			zap.String("function", "SelectPage"),
			zap.Any("param", param),
		)
		return nil, 0, 0, errors.New("获取用户列表失败")
	}

	if len(list) > 0 {
		for _, dept := range list {
			if dept != nil && dept.ParentId != 0 {
				parentInfo, err := impl.SelectById(ctx, dept.ParentId)
				if err != nil {
					continue
				}
				dept.ParentName = parentInfo.DeptName
			}
		}
	}

	return list, total, utils.CalcTotalPage(total, param.Size), nil
}

func (impl *SysDeptDaoImpl) SelectDeptList(ctx context.Context, param *request.DeptPageParam) ([]*domain.Dept, error) {
	var list []*domain.Dept

	db := impl.Gorm.WithContext(ctx).Model(&domain.Dept{}).Where("del_flag = ?", "0")
	// 动态构建查询条件
	if param.DeptName != "" {
		db.Where(`dept_name ILIKE ?`, "%"+param.DeptName+"%")
	}
	if param.Remarks != "" {
		db.Where(`remarks ILIKE ?`, "%"+param.Remarks+"%")
	}
	if param.Status != "" {
		db.Where("dept_status = ?", param.Status)
	}

	if param.UpdateBy != "" {
		db.Where(`update_by ILIKE ?`, "%"+param.UpdateBy+"%")
	}

	if err := db.Order("parent_id, order_num").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (impl *SysDeptDaoImpl) BatchDelete(ctx context.Context, ids []any) (int64, error) {
	if len(ids) <= 0 {
		return 0, errors.New("部门批量删除参数验证失败")
	}
	result := impl.Gorm.WithContext(ctx).Model(&domain.Dept{}).Where("id IN (?)", ids).Delete(nil)
	if result.Error != nil {
		zap.L().Sugar().Errorf("部门批量删除异常: %+v", result.Error)
		return 0, errors.New("批量删除失败")
	}
	count := result.RowsAffected
	zap.L().Sugar().Infof("部门批量删除成功条数===%d,删除参数ids===%+v", count, ids)
	return count, nil
}

//goland:noinspection SqlResolve,SqlCaseVsIf,SqlNoDataSourceInspection,SqlDialectInspection
func (impl *SysDeptDaoImpl) CheckDeptByIdAndName(ctx context.Context, deptId int64, deptName string) (bool, error) {
	var count int64
	if err := impl.Gorm.WithContext(ctx).Model(&domain.Dept{}).
		Where("dept_name = ? AND id != ?", deptName, deptId).
		Count(&count).Error; err != nil {
		zap.L().Sugar().Errorf("通过ID与名称查询查询部门条数错误===%+v", err)
		return false, errors.New("通过ID与名称查询查询部门失败")
	}
	return count > 0, nil
}

func (impl *SysDeptDaoImpl) SelectListByDept(ctx context.Context, dept *domain.Dept) ([]*domain.Dept, error) {
	var deptList []*domain.Dept

	db := impl.Gorm.WithContext(ctx).Model(&domain.Dept{})
	if dept == nil {
		db.Where("del_flag = ?", "0")
	} else {
		if dept.DeptStatus != "" {
			db.Where("dept_status = ?", dept.DeptStatus)
		}

		if dept.DeptName != "" {
			db.Where("dept_name ILIKE CONCAT('%' ||?|| '%')", dept.DeptName)
		}
	}

	if err := db.Find(&deptList).Error; err != nil {
		return nil, errors.New("通过部门查询部门失败")
	}
	return deptList, nil
}

func (impl *SysDeptDaoImpl) BuildDeptTree(list []*domain.Dept) ([]*domain.Dept, error) {
	var deptList []*domain.Dept
	var deptIdList []any
	for _, dept := range list {
		deptIdList = append(deptIdList, dept.Id)
	}

	for _, dept := range list {
		if utils.Contains(deptIdList, dept.ParentId) {
			resourceFN(list, dept)
			deptList = append(deptList, dept)
		}
	}

	if len(deptIdList) == 0 {
		deptList = list
	}

	return deptList, nil
}

func (impl *SysDeptDaoImpl) BuildDeptTreeSelect(list []*domain.Dept) []*domain.TreeSelect {

	var allList []*domain.Dept

	if err := impl.Gorm.WithContext(context.Background()).Where("dept_status = ?", "0").Find(&allList).Error; err != nil {
		zap.L().Sugar().Info("获取父节点错误===")
	}

	return BuildDeptTree(allList)
}

func BuildDeptTree(depts []*domain.Dept) []*domain.TreeSelect {
	// 第一步：建立ID到部门的映射
	deptMap := make(map[int64]*domain.Dept)
	for _, d := range depts {
		deptMap[d.Id] = d
	}

	// 第二步：构建树形结构
	var roots []*domain.Dept
	for _, d := range depts {
		parentId := d.ParentId
		if parentId == 0 {
			roots = append(roots, d)
			continue
		}

		// 将当前节点挂载到父节点下
		if parent, exists := deptMap[parentId]; exists {
			parent.Children = append(parent.Children, d)
		}
	}

	var treeSelectList []*domain.TreeSelect
	for _, dept := range roots {
		treeSelect := domain.NewTreeSelectByDept(dept)
		treeSelectList = append(treeSelectList, treeSelect)
	}

	return treeSelectList
}

//goland:noinspection SqlResolve,SqlCaseVsIf,SqlNoDataSourceInspection,SqlDialectInspection
func (impl *SysDeptDaoImpl) SelectListByRoleId(ctx context.Context, roleId int64) ([]int64, error) {
	roleInfo, err := impl.RoleDao.SelectById(ctx, roleId)
	if err != nil {
		return nil, errors.New("获取用户角色失败")
	}

	if roleInfo == nil {
		return nil, errors.New("角色信息不存在")
	}

	var deptIds []int64

	var (
		querySql string
		args     = []interface{}{roleInfo.Id} // 默认第一个参数是 role_id
	)

	switch roleInfo.DeptCheckStrictly {
	case "0":
		querySql = `SELECT
						d.id
					FROM
						sys_dept d
						LEFT JOIN sys_role_dept rd ON d.id = rd.dept_id
					WHERE
						rd.role_id = ?`
	case "1":

		querySql = `SELECT
							d.id
						FROM
							sys_dept d
							LEFT JOIN sys_role_dept rd ON d.id = rd.dept_id
						WHERE
							rd.role_id = ?
							AND d.id NOT IN (
								SELECT
									d.parent_id
								FROM
									sys_dept d
									INNER JOIN sys_role_dept rd ON d.id = rd.dept_id
									AND rd.role_id = ?
							)`
		args = append(args, roleInfo.Id)
	}
	if err = impl.Gorm.WithContext(ctx).Raw(querySql, args...).Find(&deptIds).Error; err != nil {
		return nil, errors.New("通过角色查询部门失败")
	}

	return deptIds, nil
}

func resourceFN(list []*domain.Dept, dept *domain.Dept) {
	childList := GetChildList(list, dept)
	dept.Children = childList
	for _, tChild := range childList {
		if HasChild(list, tChild) {
			resourceFN(list, tChild)
		}
	}
}

func GetChildList(list []*domain.Dept, dept *domain.Dept) []*domain.Dept {
	var tList []*domain.Dept
	for _, tc := range list {
		if tc.ParentId != 0 && tc.ParentId == dept.Id {
			tList = append(tList, tc)
		}
	}
	return tList
}

func HasChild(list []*domain.Dept, dept *domain.Dept) bool {
	return len(GetChildList(list, dept)) > 0
}
