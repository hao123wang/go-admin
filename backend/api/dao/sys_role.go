package dao

import (
	"go-admin-server/api/entity"
	"go-admin-server/global"

	"gorm.io/gorm"
)

type SysRoleDao struct{}

// 检查角色名称是否存在
func (d *SysRoleDao) ExistsByName(roleName string) (bool, error) {
	var count int64
	err := global.DB.Model(&entity.SysRole{}).Where("role_name = ?", roleName).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// 检查角色关键字是否存在
func (d *SysRoleDao) ExistsByKey(roleKey string) (bool, error) {
	var count int64
	err := global.DB.Model(&entity.SysRole{}).Where("role_key = ?", roleKey).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// 创建角色
func (d *SysRoleDao) CreateRole(sysRole *entity.SysRole) error {
	return global.DB.Create(sysRole).Error
}

// 获取角色列表
func (d *SysRoleDao) GetRoleList(pageNum, pageSize, roleStatus int, roleName, beginTime, endTime string) ([]entity.SysRole, int, error) {
	query := global.DB.Table("sys_role")
	query = query.Where("role_status = ?", roleStatus)
	if roleName != "" {
		query = query.Where("role_name LIKE ?", "%"+roleName+"%")
	}
	if beginTime != "" && endTime != "" {
		query = query.Where("created_at BETWEEN ? AND ?", beginTime, endTime)
	}
	var count int64
	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	var sysRoleList []entity.SysRole
	if err := query.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&sysRoleList).Error; err != nil {
		return nil, 0, err
	}
	return sysRoleList, int(count), nil
}

// 根据id获取角色
func (d *SysRoleDao) GetRoleByID(roleID uint) (*entity.SysRole, error) {
	var sysRole entity.SysRole
	if err := global.DB.Where("id = ?", roleID).First(&sysRole).Error; err != nil {
		return nil, err
	}
	return &sysRole, nil
}

// 更新角色信息
func (d *SysRoleDao) UpdateRole(sysRole *entity.SysRole) error {
	return global.DB.Save(sysRole).Error
}

// 根据id删除角色
func (d *SysRoleDao) DeleteRole(roleID uint) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		// 先删除角色表中的数据
		if err := tx.Where("id = ?", roleID).Delete(&entity.SysRole{}).Error; err != nil {
			return err
		}
		// 再删除角色-菜单关联表中的数据
		if err := tx.Where("role_id = ?", roleID).Delete(&entity.SysRoleMenu{}).Error; err != nil {
			return err
		}
		return nil
	})
}

// 获取角色下拉列表
func (d *SysRoleDao) GetRoleDropdown() ([]entity.RoleDropdownVo, error) {
	var dropdown []entity.RoleDropdownVo
	err := global.DB.Model(&entity.SysRole{}).Select("id,role_name").Scan(&dropdown).Error
	if err != nil {
		return nil, err
	}
	return dropdown, nil
}

// 判断角色id是否存在
func (d *SysRoleDao) ExistsByID(roleID uint) (bool, error) {
	var count int64
	err := global.DB.Model(&entity.SysRole{}).Where("id = ?", roleID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// 判断菜单列表中的菜单id是否都存在
func (d *SysRoleDao) ExistsMenuIds(menuIds []uint) (bool, error) {
	var count int64
	err := global.DB.Model(&entity.SysMenu{}).Where("id IN (?)", menuIds).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count == int64(len(menuIds)), nil
}

// 分配角色权限
func (d *SysRoleDao) AssignRoleMenus(roleID uint, menuIds []uint) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		// 删除旧权限
		if err := tx.Where("role_id = ?", roleID).Delete(&entity.SysRoleMenu{}).Error; err != nil {
			return err
		}
		// 新的权限列表为空，旧权限删除后即结束
		if len(menuIds) == 0 {
			return nil
		}

		// 准备批量插入的结构体实例
		var roleMenus []entity.SysRoleMenu
		for _, menuId := range menuIds {
			roleMenus = append(roleMenus, entity.SysRoleMenu{
				RoleID: roleID,
				MenuID: menuId,
			})
		}
		// 批量添加新权限
		if err := tx.Create(&roleMenus).Error; err != nil {
			return err
		}
		return nil
	})
}

// 根据角色id，获取对应的三级菜单列表
func (d *SysRoleDao) GetRoleMenus(roleID uint) ([]uint, error) {
	var roleMenus []uint
	err := global.DB.Model(&entity.SysRoleMenu{}).
		Select("sys_role_menu.menu_id").
		Joins("LEFT JOIN sys_menu ON sys_role_menu.menu_id = sys_menu.id").
		Where("sys_role_menu.role_id = ?", roleID).
		Where("sys_menu.menu_type = ?", 3).
		Scan(&roleMenus).Error
	if err != nil {
		return nil, err
	}
	return roleMenus, nil
}
