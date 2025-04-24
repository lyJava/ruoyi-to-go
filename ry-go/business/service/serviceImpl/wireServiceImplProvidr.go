package serviceImpl

import (
	"github.com/google/wire"
)

// ProviderSet wire注入使用
var ProviderSet = wire.NewSet(
	NewSysUserServiceImpl,
	NewSysDeptServiceImpl,
	NewSysNoticeServiceImpl,
	NewSysMenuServiceImpl,
	NewSysRoleServiceImpl,
	NewSysDictDataServiceImpl,
	NewFileServiceImpl,
	NewCaptchaServiceImpl,
	NewSysUserRoleServiceImpl,
	NewSysPostServiceImpl,
	NewSysDictTypeServiceImpl,
	NewTokenServiceImpl,
	NewSysLoginRecordServiceImpl,
	NewSysConfigServiceImpl,
	NewPermissionServiceImpl,
	NewSysOperationLogServiceImpl,
	NewSysMonitorServiceImpl,
	NewSysJobServiceImpl,
	NewSysJobLogServiceImpl,
	NewCornSchedulerServiceImpl,
)
