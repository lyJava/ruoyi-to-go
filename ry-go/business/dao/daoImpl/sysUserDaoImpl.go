package daoImpl

import (
	"context"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"ry-go/business/dao"
	"ry-go/business/domain"
	"ry-go/common/request"
	"ry-go/configuation"
	"ry-go/utils"
	"strings"

	"github.com/spf13/cast"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// SysUserDaoImpl 系统用户Dao结构
type SysUserDaoImpl struct {
	Gorm        *gorm.DB
	DeptDao     dao.SysDeptDao
	PostDao     dao.SysPostDao
	RoleDao     dao.SysRoleDao
	UserRoleDao dao.SysUserRoleDao
	UserPostDao dao.SysUserPostDao
}

// NewSysUserDaoImpl 系统用户Dao创建
func NewSysUserDaoImpl(db *gorm.DB,
	dept *SysDeptDaoImpl,
	post *SysPostDaoImpl,
	role *SysRoleDaoImpl,
	ur *SysUserRoleDaoImpl,
	up *SysUserPostDaoImpl,
) *SysUserDaoImpl {
	return &SysUserDaoImpl{
		Gorm:        db,
		DeptDao:     dept,
		PostDao:     post,
		RoleDao:     role,
		UserRoleDao: ur,
		UserPostDao: up,
	}
}

func hashUserPassword(user *domain.User) {
	hashPassword, _ := utils.HashPassword(user.UserPwd, bcrypt.DefaultCost)
	user.UserPwd = hashPassword
}

//goland:noinspection SqlResolve,SqlCaseVsIf,SqlNoDataSourceInspection,SqlDialectInspection
func (userDao *SysUserDaoImpl) Insert(ctx context.Context, user *domain.User) (*domain.User, error) {
	hashUserPassword(user)

	if err := userDao.Gorm.WithContext(ctx).Create(user).Error; err != nil {
		zap.L().Sugar().Errorf("系统用户新增错误===%+v", err)
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	if user.Id != 0 {
		id := user.Id
		roleIds := user.RoleIds
		if len(roleIds) > 0 {
			var userRoleList []*domain.SysUserRole
			for _, roleId := range roleIds {
				userRoleList = append(userRoleList, &domain.SysUserRole{
					UserId:     id,
					RoleId:     roleId,
					CreateTime: domain.LocalDateTimeNow(""),
				})
			}

			insertRoleCount, err := userDao.UserRoleDao.BatchInsert(ctx, userRoleList)
			if err != nil {
				return nil, errors.New("新增用户失败")
			}
			zap.L().Sugar().Infof("新增用户角色返回条数===%d", insertRoleCount)
		}

		postIds := user.PostIds
		if len(postIds) > 0 {
			var userPostList []*domain.SysUserPost
			for _, postId := range postIds {
				userPostList = append(userPostList, &domain.SysUserPost{
					UserId:     id,
					PostId:     postId,
					CreateTime: domain.LocalDateTimeNow(""),
				})
			}

			// 新增用户岗位
			insertPostCount, err := userDao.UserPostDao.BatchInsert(ctx, userPostList)
			if err != nil {
				zap.L().Sugar().Errorf("删除用户岗位失败===%+v", err)
				return nil, errors.New("新增用户失败")
			}
			zap.L().Sugar().Infof("新增用户岗位返回条数===%d", insertPostCount)
		}
	}

	return user, nil
}

//goland:noinspection SqlResolve,SqlCaseVsIf,SqlNoDataSourceInspection,SqlDialectInspection
func (userDao *SysUserDaoImpl) Update(ctx context.Context, user *domain.User) (int64, error) {
	hashUserPassword(user)
	result := userDao.Gorm.WithContext(ctx).Save(user)
	if result.Error != nil {
		zap.L().Sugar().Errorf("系统用户更新失败===%+v", result.Error)
		return 0, errors.New("系统用户更新失败")
	}

	rowsAffected := result.RowsAffected
	if rowsAffected == 0 {
		return 0, errors.New("系统用户更新失败")
	}

	roleIds := user.RoleIds
	if len(roleIds) > 0 {
		var userRoleList []*domain.SysUserRole
		for _, roleId := range roleIds {
			userRoleList = append(userRoleList, &domain.SysUserRole{
				UserId:     user.Id,
				RoleId:     roleId,
				CreateTime: domain.LocalDateTimeNow(""),
			})
		}
		// 先删除用户角色
		delRoleCount, err := userDao.UserRoleDao.DeleteByUserId(ctx, user.Id)
		if err != nil {
			zap.L().Sugar().Errorf("删除用户角色失败===%+v", err)
			return 0, errors.New("删除用户角色失败")
		}
		zap.L().Sugar().Infof("删除用户角色返回条数===%d", delRoleCount)

		insertRoleCount, err := userDao.UserRoleDao.BatchInsert(ctx, userRoleList)
		if err != nil {
			zap.L().Sugar().Errorf("新增用户角色失败===%+v", err)
			return 0, errors.New("新增用户角色失败")
		}
		zap.L().Sugar().Infof("新增用户角色返回条数===%d", insertRoleCount)

	} else {
		delRoleCount, err := userDao.UserRoleDao.DeleteByUserId(ctx, user.Id)
		if err != nil {
			zap.L().Sugar().Errorf("删除用户角色失败===%+v", err)
			return 0, errors.New("删除用户角色失败")
		}
		zap.L().Sugar().Infof("删除用户角色返回条数===%d", delRoleCount)
	}

	postIds := user.PostIds
	if len(postIds) > 0 {
		var userPostList []*domain.SysUserPost
		for _, postId := range postIds {
			userPostList = append(userPostList, &domain.SysUserPost{
				UserId:     user.Id,
				PostId:     postId,
				CreateTime: domain.LocalDateTimeNow(""),
			})
		}

		// 先删除用户岗位
		_, err := userDao.UserPostDao.DeleteByUserId(ctx, user.Id)
		if err != nil {
			return 0, errors.New("删除用户岗位失败")
		}
		// 新增用户岗位
		_, err = userDao.UserPostDao.BatchInsert(ctx, userPostList)
		if err != nil {
			return 0, errors.New("删除用户岗位失败")
		}

	} else {
		_, err := userDao.PostDao.DeleteByUserId(ctx, user.Id)
		if err != nil {
			return 0, errors.New("删除用户岗位失败")
		}
	}

	return rowsAffected, nil
}

func (userDao *SysUserDaoImpl) SelectById(ctx context.Context, id int64) (*domain.User, error) {
	var userInfo *domain.User
	if err := userDao.Gorm.WithContext(ctx).
		Model(&domain.User{}).
		Select("*, mask_password(user_pwd) AS user_pwd").
		First(&userInfo, id).Error; err != nil {
		zap.L().Sugar().Errorf("系统用户详情查询失败===%+v", err)
		return nil, errors.New("系统用户详情查询失败")
	}

	if userInfo != nil {
		deptInfo, err := userDao.DeptDao.SelectById(ctx, userInfo.DeptId)
		if err != nil {
			return nil, err
		}
		if deptInfo != nil {
			userInfo.DeptName = deptInfo.DeptName
		}

		roleIds, err := userDao.RoleDao.SelectIdsByUserId(ctx, userInfo.Id)
		if err != nil {
			return nil, err
		}
		userInfo.RoleIds = roleIds

		sysRoles, err := userDao.RoleDao.SelectAll(ctx)
		if err != nil {
			return nil, err
		}
		if domain.IsAdmin(userInfo.Id) {
			userInfo.Roles = sysRoles
		} else {
			allRoleFilterAdmin, _ := userDao.RoleDao.SelectAllFilterAdmin(sysRoles, true)
			userInfo.Roles = allRoleFilterAdmin
		}

		postIds, err := userDao.PostDao.SelectPostIdsByUserId(ctx, userInfo.Id)
		if err != nil {
			return nil, err
		}
		userInfo.PostIds = postIds

		postList, err := userDao.PostDao.SelectAll(ctx)
		if err != nil {
			return nil, err
		}
		userInfo.Posts = postList
	}

	return userInfo, nil
}

func (userDao *SysUserDaoImpl) SelectByUsername(ctx context.Context, username string) (*domain.User, error) {
	var user *domain.User
	if err := userDao.Gorm.WithContext(ctx).
		Select("*, mask_phone(phone_no) AS phone_no, mask_password(user_pwd) AS user_pwd").
		Where("username = ?", username).
		First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
		}
		return nil, errors.New("查询用户错误，请稍后再试")
	}
	return user, nil
}

func (userDao *SysUserDaoImpl) DeleteById(ctx context.Context, id int64) (int64, error) {
	result := userDao.Gorm.WithContext(ctx).Delete(&domain.User{}, id)
	if result.Error != nil {
		zap.L().Sugar().Errorf("系统用户删除失败===%+v", result.Error)
		return 0, errors.New("系统用户删除失败")
	}

	userDelCount := result.RowsAffected
	if userDelCount > 0 {
		_, err := userDao.UserRoleDao.DeleteByUserId(ctx, id)
		if err != nil {
			return 0, err
		}
		_, err = userDao.PostDao.DeleteByUserId(ctx, id)
		if err != nil {
			return 0, err
		}
	}

	return userDelCount, nil
}

func (userDao *SysUserDaoImpl) BatchInsert(ctx context.Context, list []*domain.User) ([]int64, error) {
	for _, user := range list {
		user.CreateTime = domain.LocalDateTimeNow("")
		hashUserPassword(user)
	}

	result := userDao.Gorm.WithContext(ctx).Create(list)
	if result.Error != nil {
		zap.L().Sugar().Errorf("系统用户批量新增BatchInsert错误===%+v", result.Error)
		return nil, errors.New("系统用户批量新增失败")
	}

	var ids []int64
	for _, item := range list {
		ids = append(ids, item.Id)
	}

	zap.L().Sugar().Infof("系统用户批量新增BatchInsert成功,成功条数===%d,返回ids:===%v", result.RowsAffected, ids)
	return ids, nil
}

func (userDao *SysUserDaoImpl) BatchDelete(ctx context.Context, ids []any) (int64, error) {
	result := userDao.Gorm.WithContext(ctx).
		Delete(&domain.User{}, ids)
	if result.Error != nil {
		zap.L().Sugar().Error("系统用户批量删除错误===", zap.Error(result.Error))
		return 0, errors.New("系统用户批量删除失败")
	}

	userDelCount := result.RowsAffected
	if userDelCount > 0 {
		_, err := userDao.UserRoleDao.DeleteByUserIdList(ctx, utils.ToInt64Slice(ids))
		if err != nil {
			zap.L().Sugar().Error("删除绑定角色失败===", zap.Error(result.Error))
		}
		_, err = userDao.PostDao.DeleteByUserIdList(ctx, utils.ToInt64Slice(ids))
		if err != nil {
			zap.L().Sugar().Error("删除绑定岗位失败===", zap.Error(result.Error))
		}
	}

	return userDelCount, nil
}

func (userDao *SysUserDaoImpl) SelectPage(ctx context.Context, param *request.UserPageParam) ([]*domain.User, int64, int64, error) {
	page, size := utils.SetPageDefault(param.Page, param.Size)
	param.Page = page
	param.Size = size

	// 计算分页参数
	offset := (param.Page - 1) * param.Size

	// 初始化 GORM 查询
	db := userDao.Gorm.WithContext(ctx).Model(&domain.User{})
	// 动态构建查询条件
	if param.Username != "" {
		db.Where(`username ILIKE ?`, "%"+param.Username+"%")
	}
	if param.Nickname != "" {
		db.Where(`nickname ILIKE ?`, "%"+param.Nickname+"%")
	}
	if param.PhoneNo != "" {
		db.Where(`phone_no ILIKE ?`, "%"+param.PhoneNo+"%")
	}
	if param.Sex != "" {
		db.Where("sex = ?", param.Sex)
	}
	if param.UserStatus != "" {
		db.Where("user_status = ?", param.UserStatus)
	}
	if param.DeptId != "" {
		db.Where("dept_id = ?", cast.ToInt64(param.DeptId))
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
		zap.L().Sugar().Error("系统用户总记录数异常:", zap.Error(err))
		return nil, 0, 0, errors.New("系统用户总记录数失败")
	}

	db.Offset(int(offset)).
		Select("*, mask_phone(phone_no) AS phone_no, mask_password(user_pwd) AS user_pwd").
		Limit(int(param.Size))
	if param.Column == "" {
		db.Order("id ASC")
	} else {
		db.Order(fmt.Sprintf("%s %s", param.Column, strings.ToUpper(param.Order)))
	}

	// 查询用户列表
	var list []*domain.User
	if err := db.Find(&list).Error; err != nil {
		zap.L().Error("系统用户列表异常",
			zap.Error(err),
			zap.String("function", "SelectPage"),
			zap.Any("param", param),
		)
		return nil, 0, 0, errors.New("系统用户列表失败")
	}

	if len(list) > 0 {

		for _, user := range list {
			user.IsAdmin = domain.IsAdmin(user.Id)
			deptInfo, err := userDao.DeptDao.SelectById(ctx, user.DeptId)
			if err != nil {
				zap.L().Sugar().Error("查询用户部门名称失败")
			}
			if deptInfo != nil {
				user.DeptName = deptInfo.DeptName
			}

			postIds, err := userDao.PostDao.SelectPostIdsByUserId(ctx, user.Id)
			if err != nil {
				zap.L().Sugar().Error("查询用户岗位ID切片失败")
			}
			postNameList, err := userDao.PostDao.SelectPostNameListByIds(ctx, postIds)
			if err != nil {
				zap.L().Sugar().Error("查询用户岗位ID切片失败")
			}
			user.PostNameArray = postNameList

			nameListByUserId, err := userDao.RoleDao.SelectRoleNameListByUserId(ctx, user.Id)
			if err != nil {
				zap.L().Sugar().Error("查询用户角色名称切片失败")
			}
			user.RoleNameArray = nameListByUserId

		}

	}

	return list, total, utils.CalcTotalPage(total, param.Size), nil
}

//goland:noinspection SqlResolve,SqlCaseVsIf,SqlNoDataSourceInspection
func (userDao *SysUserDaoImpl) SelectList(ctx context.Context, limit, offset int64) ([]*domain.UserForExcel, error) {
	result := userDao.Gorm.WithContext(ctx).Raw(
		`SELECT
				id,
				dept_id,
				username,
				nickname,
				CASE
					WHEN user_type = '00' THEN '系统用户'
					WHEN user_type = '01' THEN '普通用户'
					WHEN user_type = '02' THEN '测试用户'
					ELSE '其他用户'
				END AS user_type,
				CASE
					WHEN email IS NULL THEN ''
					ELSE email
				END AS email,
				CASE
					WHEN phone_no IS NULL THEN ''
					ELSE phone_no
				END AS phone_no,
				CASE
					WHEN sex = '0' THEN '女'
					WHEN sex = '0' THEN '男'
					ELSE '其他'
				END AS sex,
				avatar,
				CASE
					WHEN user_pwd IS NULL THEN ''
					ELSE user_pwd
				END AS user_pwd,
				CASE
					WHEN user_status = '0' THEN '正常'
					ELSE '停用'
				END AS user_status,
				CASE
					WHEN create_by IS NULL THEN ''
					ELSE create_by
				END AS create_by,
				CASE
					WHEN create_time IS NOT NULL THEN to_char(create_time, 'YYYY-MM-DD HH24:MI:SS')
					ELSE ''
				END AS create_time,
				CASE
					WHEN update_by IS NULL THEN ''
					ELSE update_by
				END AS update_by,
				CASE
					WHEN update_time IS NOT NULL THEN to_char(update_time, 'YYYY-MM-DD HH24:MI:SS')
					ELSE ''
				END AS update_time,
				CASE
					WHEN remarks IS NULL THEN ''
					ELSE remarks
				END AS remarks,
				CASE
					WHEN del_flag = '0' THEN '正常'
					ELSE '删除'
				END AS del_flag
			FROM
				sys_user
			ORDER BY id DESC
			LIMIT $1 OFFSET $2`, limit, offset)

	if result.Error != nil {
		zap.L().Sugar().Errorf("用户分页查询行数错误===%+v", zap.Error(result.Error))
		return nil, errors.New("查询数据失败")
	}

	rows, err := result.Rows()
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
	}(rows)

	var list []*domain.UserForExcel

	for rows.Next() {
		user := &domain.UserForExcel{}
		if err = rows.Scan(
			&user.Id,
			&user.DeptId,
			&user.Username,
			&user.Nickname,
			&user.UserType,
			&user.Email,
			&user.PhoneNo,
			&user.Sex,
			&user.Avatar,
			&user.UserPwd,
			&user.UserStatus,
			&user.CreateBy,
			&user.CreateTime,
			&user.UpdateBy,
			&user.UpdateTime,
			&user.Remarks,
			&user.DelFlag,
		); err != nil {
			zap.L().Sugar().Errorf("用户分页row转换===%+v", err)
			return nil, err
		}
		list = append(list, user)
	}

	return list, nil
}

//goland:noinspection SqlResolve,SqlCaseVsIf,SqlNoDataSourceInspection
func (userDao *SysUserDaoImpl) GetTotalCount(ctx context.Context) (int64, error) {
	var total int64
	result := userDao.Gorm.WithContext(ctx).Raw("SELECT COUNT(*) FROM sys_user").Scan(&total)
	if result.Error != nil {
		zap.L().Sugar().Errorf("查询总条数失败: %+v", result.Error)
		return 0, result.Error
	}
	return total, nil
}

func (userDao *SysUserDaoImpl) Login(ctx context.Context, user *domain.UserLogin) (*domain.User, error) {

	username := user.Username
	var loginUser *domain.User
	if err := userDao.Gorm.WithContext(ctx).Where("username = ?", username).First(&loginUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("登录错误，请稍后再试")
		}
		return nil, errors.New("系统错误，请稍后再试")
	}

	encodePwdByte, err := base64.StdEncoding.DecodeString(user.UserPwd)
	if err != nil {
		return nil, errors.New("密码格式无效")
	}

	// 通过密钥文件解密
	decryptPwdByte, err := utils.RsaDecrypt([]byte(configuation.RsaConfig.Rsa.Key.Private), encodePwdByte)

	if err != nil {
		return nil, errors.New("系统错误，请稍后再试")
	}

	decryptPassword := string(decryptPwdByte)

	// 验证密码hash是否通过
	isSame := utils.ComparePassword(loginUser.UserPwd, decryptPassword)

	if !isSame {
		return nil, errors.New("用户名与密码不匹配")
	}

	loginUser.UserPwd = ""
	return loginUser, nil
}
