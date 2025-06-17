package controller

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"ry-go/business/domain"
	"ry-go/business/service"
	"ry-go/business/service/serviceImpl"
	"ry-go/common/request"
	"ry-go/common/response"
	"ry-go/utils"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type DeptController struct {
	service      service.SysDeptService
	tokenService service.TokenService
}

// DeptControllerInit 系统部门控制器初始化
func DeptControllerInit(s *serviceImpl.SysDeptServiceImpl, t *serviceImpl.TokenServiceImpl) *DeptController {
	return &DeptController{
		service:      s,
		tokenService: t,
	}
}

// CheckDept 保存部门
func (controller *DeptController) CheckDept(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		response.NewRespCodeMsg(c, 500, "部门ID不能为空")
		return nil
	}
	// 这里需要验证类型
	deptId, err := strconv.Atoi(id)
	if err != nil {
		response.NewRespCodeMsg(c, 500, "部门ID参数验证失败")
		return nil
	}

	exist, err := controller.service.CheckIsExist(c, int64(deptId))
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	if exist {
		response.NewRespCodeMsg(c, 200, "部门信息验证通过")
		return nil
	}
	return nil
}

// SaveHandler 保存
func (controller *DeptController) SaveHandler(c echo.Context) error {
	var dept *domain.Dept
	if err := c.Bind(&dept); err != nil {
		zap.L().Sugar().Errorf("部门新增参数解析错误===%+v", err)
		response.NewRespCodeMsg(c, 500, "部门新增参数解析错误")
		return err
	}

	loginUser, err := controller.tokenService.GetLoginUser(c)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	if loginUser != nil {
		dept.CreateBy = loginUser.Username
	}

	zap.L().Sugar().Infof("部门新增入参===\n%s", utils.ToJsonFormat(dept))

	result, err := controller.service.Save(c, dept)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	zap.L().Sugar().Infof("部门新增返回===\n%s", utils.ToJsonFormat(result))
	response.NewResponse(c, 200, "部门信息新增成功", result)
	return nil
}

// UpdateDept 更新
func (controller *DeptController) UpdateDept(c echo.Context) error {
	var dept *domain.Dept
	if err := c.Bind(&dept); err != nil {
		zap.L().Sugar().Errorf("部门更新参数解析错误===%+v", err)
		response.NewRespCodeMsg(c, 500, "部门更新参数解析错误")
		return err
	}

	zap.L().Sugar().Infof("部门更新入参===\n%s", utils.ToJsonFormat(dept))

	if dept.Id == 0 {
		response.NewRespCodeMsg(c, 500, "部门ID不能为空")
		return errors.New("部门ID不能为空")
	}

	loginUser, err := controller.tokenService.GetLoginUser(c)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	if loginUser != nil {
		dept.UpdateBy = loginUser.Username
	}

	result, err := controller.service.Update(c, dept)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}
	zap.L().Sugar().Infof("部门更新返回===\n%s", utils.ToJsonFormat(result))
	response.NewResponse(c, 200, "部门信息更新成功", result)
	return nil
}

func (controller *DeptController) BatchSaveDept(c echo.Context) error {
	var list []*domain.Dept
	if err := c.Bind(&list); err != nil {
		zap.L().Sugar().Errorf("部门批量新增参数解析错误===%+v", err)
		response.NewRespCodeMsg(c, 500, "部门批量新增参数解析错误")
		return err
	}

	zap.L().Sugar().Infof("部门批量新增入参===\n%s", utils.ToJsonFormat(cast.ToSlice(list)))

	result, err := controller.service.BatchSave(c, list)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	zap.L().Sugar().Infof("部门新增返回===\n%s", utils.ToJsonFormat(result))
	response.NewResponse(c, 200, "部门信息批量新增成功", result)
	return nil
}

func (controller *DeptController) BatchSaveDeptList(c echo.Context) error {
	var list []*request.DeptInsert
	if err := c.Bind(&list); err != nil {
		zap.L().Sugar().Errorf("部门批量新增参数解析错误===%+v", err)
		response.NewRespCodeMsg(c, 500, "部门批量新增参数解析错误")
		return err
	}

	zap.L().Sugar().Infof("部门批量新增入参===\n%s", utils.ToJsonFormat(cast.ToSlice(list)))

	result, err := controller.service.BatchSaveByList(c, list)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}
	zap.L().Sugar().Infof("批量执行成功条数:%d", len(result))
	zap.L().Sugar().Infof("部门新增返回===\n%s", utils.ToJsonFormat(result))

	response.NewResponse(c, 200, "部门信息批量新增成功", result)
	return nil
}

func (controller *DeptController) ImportData(c echo.Context) error {
	fileMap, err := utils.HandlerEchoUpload(c, "file", time.Now().Format("20060102150405"), "")
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	excelAbsPath := fileMap["absPath"]

	// 导入成功后删除excel文件
	defer func(name string) {
		if err = os.Remove(name); err != nil {
			zap.L().Sugar().Errorf("删除===%s出错:%+v", excelAbsPath, err)
		}
		zap.L().Sugar().Infof("成功删除===%s", excelAbsPath)
	}(excelAbsPath)

	excelData, err := utils.ReadExcelFile(excelAbsPath)
	if err != nil {
		zap.L().Sugar().Errorf("读取Excel失败:%+v", err)
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	deptList, err := ConvertToDeptInsert(excelData)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	zap.L().Sugar().Infof("读取的Excel数据:%s", utils.ToJsonFormat(deptList))

	result, err := controller.service.BatchSaveByList(c, deptList)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	zap.L().Sugar().Infof("部门批量导入成功条数:%d", len(result))
	zap.L().Sugar().Infof("部门批量导入返回===\n%s", utils.ToJsonFormat(result))
	response.NewResponse(c, 200, fmt.Sprintf("部门信息导入成功:%d条", len(result)), result)
	return nil
}

// SelectListHandler 查询部门列表
func (controller *DeptController) SelectListHandler(c echo.Context) error {
	req := new(request.DeptPageParam)
	if err := c.Bind(req); err != nil {
		response.NewRespCodeMsg(c, http.StatusInternalServerError, "部门分页参数解析失败")
		return err
	}

	list, err := controller.service.SelectDeptList(c, req)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}
	response.NewResponse(c, 200, "查询成功", list)
	return nil
}

func (controller *DeptController) SelectByIdHandler(c echo.Context) error {
	deptId := c.Param("deptId")
	if deptId == "" {
		response.NewRespCodeMsg(c, http.StatusInternalServerError, "部门ID不能为空")
		return errors.New("部门ID不能为空")
	}

	deptInfo, err := controller.service.SelectById(c, cast.ToInt64(deptId))
	if err != nil {
		response.NewRespCodeErr(c, http.StatusInternalServerError, err)
		return err
	}

	response.NewResponse(c, 200, "查询成功", deptInfo)
	return nil
}

func (controller *DeptController) ExcludeDetailHandler(c echo.Context) error {
	deptId := c.Param("deptId")
	if deptId == "" {
		response.NewRespCodeMsg(c, http.StatusInternalServerError, "部门ID不能为空")
		return errors.New("部门ID不能为空")
	}
	id := cast.ToInt64(deptId)
	deptList, err := controller.service.SelectListByDept(c, &domain.Dept{})
	if err != nil {
		response.NewRespCodeErr(c, http.StatusInternalServerError, err)
		return err
	}

	var filtered []*domain.Dept
	if len(deptList) > 0 {
		for _, dept := range deptList {
			currentId := dept.Id
			if currentId == id {
				continue
			}

			ancestorList := strings.Split(dept.Ancestors, ",")
			if contains(ancestorList, cast.ToString(deptId)) {
				continue
			}
			filtered = append(filtered, dept)
		}
	}

	response.NewResponse(c, 200, "查询成功", filtered)
	return nil
}

// contains 检查字符串切片是否包含目标
func contains(slice []string, target string) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}

// SelectPageHandler 分页处理器
func (controller *DeptController) SelectPageHandler(c echo.Context) error {
	req := new(request.DeptPageParam)
	if err := c.Bind(req); err != nil {
		response.NewRespCodeMsg(c, http.StatusInternalServerError, "部门分页参数解析失败")
		return err
	}

	list, total, page, err := controller.service.SelectPage(c, req)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	pageData := response.NewPageData(list, total, page)

	response.NewResponse(c, 200, "部门分页查询成功", pageData)
	return nil
}

// BatchDelete 批量删除
func (controller *DeptController) BatchDelete(c echo.Context) error {
	ids, err := GlobalDeleteHandler(c)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	count, err := controller.service.BatchDelete(c, ids)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}
	response.NewResponse(c, 200, "批量删除成功", count)
	return nil
}

func ConvertToDeptInsert(data [][]string) ([]*request.DeptInsert, error) {
	if len(data) == 0 {
		return nil, errors.New("没有可以转换的数据")
	}

	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		zap.L().Sugar().Errorf("LocalDateTime LoadLocation error %v", err)
	}
	var deptInserts []*request.DeptInsert

	for _, row := range data {
		if len(row) != 9 {
			return nil, fmt.Errorf("每行%d个字段数量必须是9", len(row))
		}
		deptInsert := &request.DeptInsert{
			ParentId:   cast.ToInt64(row[0]),
			Ancestors:  row[1],
			DeptName:   row[2],
			OrderNum:   cast.ToInt64(row[3]),
			DeptStatus: row[4],
			CreateBy:   row[5],
			CreateTime: time.Now().In(location),
			UpdateBy:   row[6],
			Remarks:    row[7],
			DelFlag:    row[8],
		}

		deptInserts = append(deptInserts, deptInsert)
	}

	return deptInserts, nil
}

func (controller *DeptController) TreeSelectHandler(c echo.Context) error {

	deptList, err := controller.service.SelectListByDept(c, &domain.Dept{
		DeptStatus: "0",
	})

	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	deptTreeList, err := controller.service.BuildDeptTree(deptList)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	treeSelect := controller.service.BuildDeptTreeSelect(deptTreeList)

	response.NewResponse(c, http.StatusOK, "查询成功", treeSelect)
	return nil
}

func (controller *DeptController) RoleDeptTreeSelectHandler(c echo.Context) error {
	roleId := c.Param("roleId")
	if roleId == "" {
		response.NewRespCodeMsg(c, 500, "角色ID不能为空")
		return errors.New("角色ID不能为空")
	}

	checkedKeys, err := controller.service.SelectListByRoleId(c, cast.ToInt64(roleId))
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	deptList, err := controller.service.SelectListByDept(c, &domain.Dept{
		DeptStatus: "0",
	})

	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	dataMap := map[string]any{
		"checkedKeys": checkedKeys,
		"depts":       controller.service.BuildDeptTreeSelect(deptList),
	}
	response.NewResponse(c, http.StatusOK, "查询成功", dataMap)
	return nil
}

// DownloadExcelBufferHandler 导出Excel，使用缓冲流
func (controller *DeptController) DownloadExcelBufferHandler(c echo.Context) error {
	data, err := controller.service.SelectAll(c)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}
	headers := []string{"ID", "父节点ID", "组级列表", "部门名称", "排序", "部门状态（0正常 1停用）", "创建人", "创建时间", "修改人", "修改时间", "备注信息", "负责人", "手机", "邮箱", "是否删除"}
	return DownloadExcelBuffer(c, "部门信息_"+time.Now().Format("20060102150405")+".xlsx", "部门信息", headers, data, true)
}