package service

import "go-admin-server/api/dao"

// 注册dao层对象实例
var (
	SysPostDao  = &dao.SysPostDao{}
	SysDeptDao  = &dao.SysDeptDao{}
	SysMenuDao  = &dao.SysMenuDao{}
	SysRoleDao  = &dao.SysRoleDao{}
	SysAdminDao = &dao.SysAdminDao{}
	SysLogDao   = &dao.SysLogDao{}
)
