package controller

import "github.com/google/wire"

// ProviderSet wire注入使用
var ProviderSet = wire.NewSet(
	UserControllerInit,
	DeptControllerInit,
	SysNoticeControllerInit,
	SysMenuControllerInit,
	SysRoleControllerInit,
	SysDictDataControllerInit,
	FileControllerInit,
	BarcodeControllerInit,
	QrcodeControllerInit,
	SysUserRoleControllerInit,
	SysPostControllerInit,
	SysDictTypeControllerInit,
	LoginControllerInit,
	SysLoginRecordControllerInit,
	SysConfigControllerInit,
	SysOperationLogControllerInit,
	SysMonitorControllerInit,
	SysJobControllerInit,
	SysJobLogControllerInit,
)
