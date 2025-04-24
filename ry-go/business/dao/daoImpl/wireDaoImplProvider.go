package daoImpl

import "github.com/google/wire"

// ProviderSet wire注入使用
var ProviderSet = wire.NewSet(
	NewSysUserDaoImpl,
	NewSysDeptDaoImpl,
	NewSysNoticeDaoImpl,
	NewSysMenuDaoImpl,
	NewSysRoleDaoImpl,
	NewSysDictDataDaoImpl,
	NewSysUserRoleDaoImpl,
	NewSysPostDaoImpl,
	NewSysDictTypeDaoImpl,
	NewSysLoginRecordDaoImpl,
	NewSysConfigDaoImpl,
	NewSysUserPostDaoImpl,
	NewSysRoleMenuDaoImpl,
	NewSysRoleDeptDaoImpl,
	NewSysOperationLogDaoImpl,
	NewSysJobDaoImpl,
	NewSysJobLogDaoImpl,
)
