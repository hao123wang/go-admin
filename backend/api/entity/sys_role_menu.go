package entity

// 角色菜单模型
type SysRoleMenu struct {
	RoleID uint `gorm:"column:role_id;comment:'角色id';not null"`
	MenuID uint `gorm:"column:menu_id;comment:'菜单id';not null"`
}

func (SysRoleMenu) TableName() string {
	return "sys_role_menu"
}
