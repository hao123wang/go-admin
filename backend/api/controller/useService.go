package controller

import "go-admin-server/api/service"

// 注册service层对象实例
var (
	SysPostService  = &service.SysPostService{}
	SysDeptService  = &service.SysDeptService{}
	SysMenuService  = &service.SysMenuService{}
	SysRoleService  = &service.SysRoleService{}
	SysAdminService = &service.SysAdminService{}
	UploadService   = &service.UploadService{}
	LogService      = &service.SysLogService{}
)
