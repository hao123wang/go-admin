package entity

type SysAdminRole struct {
	AdminID uint `gorm:"admin_id"`
	RoleID  uint `gorm:"role_id"`
}

func (SysAdminRole) TableName() string {
	return "sys_admin_role"
}