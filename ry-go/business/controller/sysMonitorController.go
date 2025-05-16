package controller

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"ry-go/business/domain"
	"ry-go/business/service"
	"ry-go/business/service/serviceImpl"
	"ry-go/common/model"
	"ry-go/common/request"
	"ry-go/common/response"
	"ry-go/configuation"
	"ry-go/utils"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/spf13/cast"
)

type SysMonitorController struct {
	Monitor service.SysMonitorService
}

func SysMonitorControllerInit(ml *serviceImpl.SysMonitorServiceImpl) *SysMonitorController {
	return &SysMonitorController{
		Monitor: ml,
	}
}

// OnlineListHandler 在线用户分页查询
func (ct *SysMonitorController) OnlineListHandler(e echo.Context) error {
	req := new(request.OnlinePageParam)
	if err := e.Bind(req); err != nil {
		response.NewRespCodeMsg(e, http.StatusInternalServerError, "分页参数解析失败")
		return err
	}

	allLoginValues, err := ct.Monitor.OnlineList(e, configuation.CacheConfig.Cache.Login.Prefix)
	if err != nil {
		response.NewRespCodeErr(e, 500, err)
		return err
	}

	var filterList []*domain.LoginUser

	if len(allLoginValues) > 0 {
		filterList = filterUsers(allLoginValues, req.Username, req.Ip)
	}

	response.NewResponse(e, http.StatusOK, "查询成功", utils.BuildPageBySilence(filterList, req.Page, req.Size))
	return nil
}

// OnlineForceLogoutHandler 在线用户强行下线
func (ct *SysMonitorController) OnlineForceLogoutHandler(e echo.Context) error {
	uuid, err := GetSinglePathParam(e, "uuid")
	if err != nil {
		response.NewRespCodeErr(e, 500, err)
		return err
	}

	result, err := ct.Monitor.OnlineForceLogout(e, configuation.CacheConfig.Cache.Login.Prefix+uuid)
	if err != nil {
		response.NewRespCodeErr(e, 500, err)
		return err
	}

	if result {
		response.NewRespCodeMsg(e, http.StatusOK, "操作成功")
	}
	return nil
}

// ServerInfoHandler  获取服务器信息处理器
func (ct *SysMonitorController) ServerInfoHandler(e echo.Context) error {
	hostInfo, _, memoryStat, stats := utils.OsInfo()

	// 获取操作系统和架构
	goos := runtime.GOOS

	// 示例：构建跨平台工具路径
	tool := "go"
	if goos == "windows" {
		tool += ".exe"
	}

	// 处理路径分隔符
	path := fmt.Sprintf("%s/bin/%s", runtime.GOROOT(), tool)
	if strings.Contains(runtime.GOOS, "windows") {
		path = strings.ReplaceAll(path, "/", "\\")
	}

	physicalCores, _ := cpu.Counts(false)
	cpuPercent, _ := cpu.Percent(time.Second, true)

	absPath, _ := filepath.Abs(runtime.GOROOT())
	cmd := exec.Command("go", "version")
	output, _ := cmd.Output()

	// 两位小数的计算方式
	toGB := func(bytes uint64) float64 {
		gb := float64(bytes) / 1024 / 1024 / 1024
		return math.Round(gb*100) / 100
	}

	var sysFiles []*response.Disk
	if len(stats) > 0 {

		for _, f := range stats {
			usage, _ := disk.Usage(f.Mountpoint)
			sysFile := &response.Disk{
				DirName:     f.Device,
				Total:       toGB(usage.Total),
				Free:        cast.ToString(toGB(usage.Free)),
				Used:        toGB(usage.Used),
				Usage:       math.Round(usage.UsedPercent*100) / 100,
				SysTypeName: usage.Fstype,
				TypeName:    usage.Path,
			}
			sysFiles = append(sysFiles, sysFile)
		}
	}

	dataMap := map[string]any{
		"cpu": &response.Cpu{
			CpuNum: int64(physicalCores),
			Used:   math.Round(cpuPercent[0]*100) / 100,
		},

		"mem": &response.Memory{
			Total: toGB(memoryStat.Total),
			Used:  toGB(memoryStat.Used),
			Free:  toGB(memoryStat.Available),
			Usage: math.Round(memoryStat.UsedPercent*100) / 100,
		},
		"jvm": &response.Gvm{
			Home:      absPath,
			Version:   string(output),
			InputArgs: cast.ToString(os.Args),
			StartTime: time.Now().Format(time.DateTime),
			Runtime:   utils.TimeForHuman(configuation.GlobalStartTime),
		},
		"sys": &response.OsInfo{
			ComputerName: hostInfo.Hostname,
			Ip:           utils.GetClientIP(e),
			ProPath:      utils.GetCurrentAbPathByCaller(),
			OsName:       hostInfo.Platform + hostInfo.PlatformVersion,
			OsArch:       hostInfo.KernelArch,
		},
		"sysFiles": sysFiles,
	}
	response.NewResponse(e, http.StatusOK, "操作成功", dataMap)
	return nil
}

func (ct *SysMonitorController) CacheNameListHandler(e echo.Context) error {
	var list []*model.Cache
	list = append(list, &model.Cache{
		Name:   "user:login",
		Remark: "登录信息",
	})
	list = append(list, &model.Cache{
		Name:   "sys_config",
		Remark: "配置信息",
	})
	list = append(list, &model.Cache{
		Name:   "sys_dict",
		Remark: "字典信息",
	})
	response.NewResponse(e, http.StatusOK, "查询成功", list)
	return nil
}

func (ct *SysMonitorController) CacheListByKeyHandler(e echo.Context) error {
	cacheKey := e.Param("key")
	if cacheKey == "" {
		response.NewRespCodeMsg(e, http.StatusInternalServerError, "缓存的键不能为空")
		return errors.New("缓存的键不能为空")
	}
	keyList, err := ct.Monitor.CacheListByKey(e, cacheKey)
	if err != nil {
		response.NewRespCodeErr(e, http.StatusInternalServerError, err)
		return errors.New("缓存的键不能为空")
	}
	response.NewResponse(e, http.StatusOK, "查询成功", keyList)
	return nil
}

func (ct *SysMonitorController) CacheDetailHandler(e echo.Context) error {
	prefix := e.Param("prefix")
	suffix := e.Param("suffix")
	if prefix == "" && suffix == "" {
		response.NewRespCodeMsg(e, http.StatusInternalServerError, "缓存参数不能为空")
		return errors.New("缓存参数不能为空")
	}
	cache, err := ct.Monitor.CacheDetail(e, prefix, suffix)
	if err != nil {
		response.NewRespCodeErr(e, http.StatusInternalServerError, err)
		return err
	}
	response.NewResponse(e, http.StatusOK, "查询成功", cache)
	return nil
}

func (ct *SysMonitorController) CacheClearHandler(e echo.Context) error {
	prefix := e.Param("key")
	if prefix == "" {
		response.NewRespCodeMsg(e, http.StatusInternalServerError, "缓存的键不能为空")
		return errors.New("缓存的键不能为空")
	}

	if err := ct.Monitor.CacheClear(e, prefix); err != nil {
		response.NewRespCodeErr(e, http.StatusInternalServerError, err)
		return err
	}
	response.NewRespCodeMsg(e, http.StatusOK, "删除成功")
	return nil
}

func filterUsers(users []*domain.LoginUser, username, ip string) []*domain.LoginUser {
	if username == "" && ip == "" {
		return users
	}

	var filtered []*domain.LoginUser

	for _, user := range users {
		if user == nil {
			continue
		}

		// 统一处理匹配逻辑
		matchUsername := username == "" || strings.Contains(user.Username, username)
		matchIP := ip == "" || strings.Contains(user.IP, ip)

		location, err := utils.Ip2RegionSearch(user.IP)
		if err != nil {
			continue
		}
		user.LoginLocation = location

		if matchUsername && matchIP {
			filtered = append(filtered, user)
		}
	}
	return filtered
}
