package controller

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"os"
	"ry-go/business/domain"
	"ry-go/business/service"
	"ry-go/business/service/serviceImpl"
	"ry-go/common/request"
	"ry-go/common/response"
	"ry-go/utils"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

// UserController 系统用户控制器
type UserController struct {
	userService    service.SysUserService
	captchaService service.CaptchaService
	deptService    service.SysDeptService
	roleService    service.SysRoleService
	menuService    service.SysMenuService
	tokenService   service.TokenService
}

// UserControllerInit 系统用户控制器初始化
func UserControllerInit(
	s *serviceImpl.SysUserServiceImpl,
	c *serviceImpl.CaptchaServiceImpl,
	d *serviceImpl.SysDeptServiceImpl,
	r *serviceImpl.SysRoleServiceImpl,
	m *serviceImpl.SysMenuServiceImpl,
	t *serviceImpl.TokenServiceImpl,
) *UserController {
	return &UserController{
		userService:    s,
		captchaService: c,
		deptService:    d,
		roleService:    r,
		menuService:    m,
		tokenService:   t,
	}
}

var validatorMap = map[string]govalidator.CustomTypeValidator{
	"customerSexValidator":    domain.CustomerSexValidator,
	"customerStatusValidator": domain.CustomerStatusValidator,
}

// SaveHandler 系统用户新增
func (controller *UserController) SaveHandler(c echo.Context) error {
	var sysUser *domain.User
	if err := c.Bind(&sysUser); err != nil {
		response.NewRespCodeMsg(c, http.StatusInternalServerError, "用户新增参数解析错误")
		return err
	}

	if sysUser.DeptId == 0 {
		sysUser.DeptId = 9999999
	}
	if err := UseCustomValidator(sysUser, validatorMap, c); err != nil {
		return err
	}

	user, err := controller.userService.Insert(c, sysUser)
	if err != nil {
		response.NewRespCodeErr(c, http.StatusInternalServerError, err)
		return err
	}
	response.NewResponse(c, http.StatusOK, "用户新增成功", user)
	return nil
}

// UpdateHandler 系统用户修改
func (controller *UserController) UpdateHandler(c echo.Context) error {
	var updateUser *domain.UserUpdate
	if err := c.Bind(&updateUser); err != nil {
		response.NewRespCodeMsg(c, http.StatusInternalServerError, "用户修改参数解析错误")
		return err
	}

	if err := UseCustomValidator(updateUser, validatorMap, c); err != nil {
		return err
	}

	loginUser, _ := controller.tokenService.GetLoginUser(c)

	result, err := controller.userService.Update(c, UpdateUserConvertToUser(updateUser, loginUser))
	if result < 1 {
		response.NewRespCodeErr(c, http.StatusInternalServerError, err)
		return err
	}
	response.NewRespCodeMsg(c, http.StatusOK, "用户修改成功")
	return nil
}

func UpdateUserConvertToUser(updateUser *domain.UserUpdate, loginUser *domain.LoginUser) *domain.User {
	user := &domain.User{
		UpdateTime: domain.LocalDateTimeNow(""),
	}
	if loginUser != nil {
		user.UpdateBy = loginUser.Username
	}
	user.Id = updateUser.Id
	user.DeptId = updateUser.DeptId
	user.Username = updateUser.Username
	user.UserStatus = updateUser.UserStatus
	user.Email = updateUser.Email
	user.Avatar = updateUser.Avatar
	user.Nickname = updateUser.Nickname
	user.PhoneNo = updateUser.PhoneNo
	user.Remarks = updateUser.Remarks
	user.Sex = updateUser.Sex
	user.RoleIds = updateUser.RoleIds
	user.PostIds = updateUser.PostIds

	return user
}

// BatchInsertHandler 系统用户批量新增
func (controller *UserController) BatchInsertHandler(c echo.Context) error {
	var list []*domain.User
	if err := c.Bind(&list); err != nil {
		response.NewRespCodeMsg(c, http.StatusInternalServerError, "用户批量新增参数解析错误")
		return err
	}

	/*var errList []string
	for _, user := range list {
		if err := RegisterCustomValidator(user, validatorMap); err != nil {
			errList = append(errList, err.Error())
		}
	}

	if len(errList) > 0 {
		response.NewResponse(c, http.StatusInternalServerError, "校验失败", errList)
		return errors.New("校验失败")
	}*/

	if err := UseCustomValidators(list, validatorMap, c, true); err != nil {
		return err
	}

	ids, err := controller.userService.BatchInert(c, list)
	if len(ids) <= 0 {
		response.NewRespCodeMsg(c, http.StatusInternalServerError, "用户新增失败")
		return err
	}
	if err != nil {
		response.NewRespCodeErr(c, http.StatusInternalServerError, err)
		return err
	}

	response.NewResponse(c, http.StatusOK, "用户新增成功", len(ids))
	return nil
}

// DetailHandler 系统用户详情查询
func (controller *UserController) DetailHandler(c echo.Context) error {
	user, err := controller.userService.SelectById(c, cast.ToInt64(c.Param("id")))
	if err != nil {
		response.NewRespCodeErr(c, http.StatusInternalServerError, err)
		return err
	}

	response.NewResponse(c, http.StatusOK, "系统用户详情查询成功", user)
	return nil
}

// DeleteHandler 系统用户删除
func (controller *UserController) DeleteHandler(c echo.Context) error {
	ids, err := GlobalDeleteHandler(c)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	if len(ids) > 20 {
		response.NewRespCodeMsg(c, http.StatusInternalServerError, "单次批量不能超过20条")
		return errors.New("单次批量不能超过20条")
	}

	count, err := controller.userService.BatchDelete(c, ids)
	if err != nil {
		response.NewRespCodeErr(c, http.StatusInternalServerError, err)
		return err
	}
	response.NewRespCodeMsg(c, http.StatusOK, fmt.Sprintf("成功删除条数：%d", count))
	return nil
}

// BatchDeleteHandler 系统用户批量删除
func (controller *UserController) BatchDeleteHandler(c echo.Context) error {
	var ids []any
	if err := c.Bind(&ids); err != nil {
		response.NewRespCodeMsg(c, http.StatusInternalServerError, "用户批量删除参数解析错误")
		return err
	}

	count, err := controller.userService.BatchDelete(c, ids)
	if err != nil {
		response.NewRespCodeErr(c, http.StatusInternalServerError, err)
		return err
	}

	response.NewResponse(c, http.StatusOK, "用户删除成功", count)
	return nil
}

// SelectPageHandler 系统用户分页查询
func (controller *UserController) SelectPageHandler(c echo.Context) error {
	param := new(request.UserPageParam)
	if err := c.Bind(param); err != nil {
		response.NewRespCodeMsg(c, http.StatusInternalServerError, "分页参数解析失败")
		return err
	}

	list, total, page, err := controller.userService.SelectPage(c, param)
	if err != nil {
		response.NewRespCodeErr(c, http.StatusInternalServerError, err)
		return err
	}

	pageData := response.NewPageData(list, total, page)

	response.NewResponse(c, http.StatusOK, "系统用户分页查询成功", pageData)
	return nil
}

// ExportExcelHandler 导出excel处理器
func (controller *UserController) ExportExcelHandler(e echo.Context) error {
	size := cast.ToInt64(e.QueryParam("size"))
	if size <= 0 {
		response.NewRespCodeMsg(e, http.StatusBadRequest, "导出条数参数无效")
		return nil
	}

	// 查询总条数
	totalCount, err := controller.userService.GetTotalCount(e)
	if err != nil {
		zap.L().Sugar().Errorf("查询总条数失败: %v", err)
		response.NewRespCodeMsg(e, http.StatusInternalServerError, "导出excel失败")
		return nil
	}

	if totalCount == 0 {
		response.NewRespCodeMsg(e, http.StatusOK, "当前无数据可导出")
		return nil
	}

	// size大于总条数，则size为总条数
	if size > totalCount {
		size = totalCount
	}

	var list []*domain.UserForExcel

	// 定义每批次查询的大小
	batchSize := int64(2000)
	// 计算查询次数
	totalBatches := int(math.Ceil(float64(size) / float64(batchSize)))
	zap.L().Sugar().Infof("计算查询次数===%d", totalBatches)
	// 分批次查询
	for batch := 0; batch < totalBatches; batch++ {
		limit := batchSize
		offset := int64(batch) * batchSize
		if offset+batchSize > size {
			limit = size - offset // 最后一次查询可能不足 batchSize
		}

		zap.L().Sugar().Infof("计算查询limit===%d,offset==%d", limit, offset)

		batchList, err := controller.userService.SelectList(e, limit, offset)
		if err != nil {
			response.NewResponse(e, http.StatusInternalServerError, err.Error(), nil)
			return err
		}

		zap.L().Sugar().Infof("查询的batchList长度:%d", len(batchList))

		// 合并结果
		list = append(list, batchList...)
	}

	if len(list) == 0 {
		response.NewRespCodeMsg(e, http.StatusInternalServerError, "当前无数据可导出")
		return nil
	}

	currentPath, err := os.Getwd()
	if err != nil {
		zap.L().Sugar().Errorf("获取当前目录错误===%v", err)
		response.NewRespCodeMsg(e, http.StatusInternalServerError, "导出excel失败")
		return err
	}
	zap.L().Sugar().Infof("获取当前目录:%s", currentPath)

	// 设置表头
	headers := []string{
		"主键ID",
		"部门ID",
		"用户账号",
		"用户昵称",
		"用户类型",
		"用户邮箱",
		"手机号码",
		"用户性别",
		"头像地址",
		"帐号状态",
		"创建人",
		"创建时间",
		"修改人",
		"修改时间",
		"备注",
		"是否删除",
	}
	excelFilePath, err := utils.WriteDataToExcel("/excel/"+utils.NormalUUID()+".xlsx", headers, list)
	if err != nil {
		response.NewRespCodeMsg(e, http.StatusInternalServerError, "导出excel失败")
		return err
	}
	return utils.DownloadFile(excelFilePath, e)
}

// CaptchaHandler 验证码
func (controller *UserController) CaptchaHandler(e echo.Context) error {
	return utils.GenerateCaptcha(e)
}

// VerifyCaptchaHandler 验证验证码
func (controller *UserController) VerifyCaptchaHandler(e echo.Context) error {
	return utils.VerifyCaptcha(e)
}

// CaptchaRedisHandler 验证码
func (controller *UserController) CaptchaRedisHandler(e echo.Context) error {
	return controller.captchaService.GenerateCaptcha(e)
}

// VerifyCaptchaRedisHandler 验证验证码
func (controller *UserController) VerifyCaptchaRedisHandler(e echo.Context) error {
	return controller.captchaService.VerifyCaptcha(e)
}

// DownloadExcelBufferHandler 导出Excel，使用缓冲流
func (controller *UserController) DownloadExcelBufferHandler(c echo.Context) error {
	data, err := controller.userService.SelectAll(c)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}
	headers := []string{"ID", "部门ID", "昵称", "用户名", "用户类型", "邮箱", "手机号", "性别", "头像地址", "密码", "启用状态", "创建人", "创建时间", "修改人", "修改时间", "删除状态", "备注信息", "部门名称"}
	return DownloadExcelBuffer(c, "用户信息_"+time.Now().Format("20060102150405")+".xlsx", "用户信息", headers, data, true)
}