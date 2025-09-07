package flag

import (
	"go-admin-server/api/entity"
	"go-admin-server/global"
)

// 通过命令行执行模型迁移
func SQL() error {
	return global.DB.Set("table_options", "ENGINE=InnoDB").AutoMigrate(
		&entity.SysPost{},         // 岗位表
		&entity.SysDept{},         // 部门表
		&entity.SysMenu{},         // 菜单表
		&entity.SysRole{},         // 角色表
		&entity.SysRoleMenu{},     // 角色-菜单关联表
		&entity.SysAdmin{},        // 用户表
		&entity.SysAdminRole{},    // 用户-角色关联表
		&entity.SysLoginLog{},     // 登录日志表
		&entity.SysOperationLog{}, // 操作日志表
	)
}
