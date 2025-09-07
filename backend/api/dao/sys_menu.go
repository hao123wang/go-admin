package dao

import (
	"go-admin-server/api/entity"
	"go-admin-server/global"
)

type SysMenuDao struct{}

func (d *SysMenuDao) ExistsByName(menuName string) (bool, error) {
	var count int64
	err := global.DB.Model(&entity.SysMenu{}).Where("menu_name = ?", menuName).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (d *SysMenuDao) GetMenuByID(menuID uint) (*entity.SysMenu, error) {
	var sysMenu entity.SysMenu
	if err := global.DB.Where("id = ?", menuID).First(&sysMenu).Error; err != nil {
		return nil, err
	}
	return &sysMenu, nil
}

func (d *SysMenuDao) CreateMenu(sysMenu *entity.SysMenu) error {
	return global.DB.Create(sysMenu).Error
}

func (d *SysMenuDao) GetMenuList(menuName string, menuStaus int) ([]entity.SysMenu, error) {
	var sysMenus []entity.SysMenu
	query := global.DB.Model(&entity.SysMenu{})
	if menuName != "" {
		query = query.Where("menu_name LIKE ?", "%"+menuName+"%")
	}
	if menuStaus != 0 {
		query = query.Where("menu_status = ?", menuStaus)
	}
	err := query.Find(&sysMenus).Error
	if err != nil {
		return nil, err
	}
	return sysMenus, nil
}

// 修改菜单
func (d *SysMenuDao) UpdateMenu(sysMenu *entity.SysMenu) error {
	return global.DB.Save(sysMenu).Error
}

// 删除单个菜单
func (d *SysMenuDao) DeleteMenu(menuID uint) error {
	return global.DB.Where("id = ?", menuID).Delete(&entity.SysMenu{}).Error
}

// 判断是否有子菜单
func (d *SysMenuDao) HasSubMenu(menuID uint) (bool, error) {
	var count int64
	err := global.DB.Model(&entity.SysMenu{}).Where("parent_id = ?", menuID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// 获取菜单下拉列表
func (d *SysMenuDao) GetMenuDropdown() ([]entity.MenuDropdownVo, error) {
	var dropdown []entity.MenuDropdownVo
	err := global.DB.Model(&entity.SysMenu{}).Select("id,menu_name,parent_id").Scan(&dropdown).Error
	if err != nil {
		return nil, err
	}
	return dropdown, nil
}

// 获取当前登录用户的左侧菜单列表（一级菜单和二级菜单展示）
func (d *SysMenuDao) LeftMenuList(adminId uint) ([]entity.FirstLevelMenuVo, error) {
	// 先查询用户有权限的所有菜单id
	var menuIds []uint
	err := global.DB.Model(&entity.SysAdmin{}).
		Joins("LEFT JOIN sys_admin_role ar ON sys_admin.id = ar.admin_id").
		Joins("LEFT JOIN sys_role r ON ar.role_id = r.id").
		Joins("LEFT JOIN sys_role_menu rm ON r.id = rm.menu_id").
		Where("sys_admin.id = ?", adminId).
		Where("r.role_status = ?", 1).
		Pluck("rm.menu_id", &menuIds).Error
	if err != nil {
		return nil, err
	}

	if len(menuIds) == 0 {
		return []entity.FirstLevelMenuVo{}, nil
	}

	// 查询一级菜单信息
	var firstMenuList []entity.FirstLevelMenuVo
	err = global.DB.Model(&entity.SysMenu{}).
		Select("id", "menu_name", "menu_icon", "url").
		Where("id IN (?)", menuIds).
		Where("menu_type = ?", 1).
		Scan(&firstMenuList).Error
	if err != nil {
		return nil, err
	}

	// 为每个一级菜单，查询对应的二级菜单列表
	for i := range firstMenuList {
		var children []entity.SecondLevelMenuVo
		err = global.DB.Model(&entity.SysMenu{}).
			Select("menu_name", "menu_icon", "url").
			Where("id IN (?)", menuIds).
			Where("menu_type = ?", 2).
			Where("parent_id = ?", firstMenuList[i].ID).
			Scan(&children).
			Scan(&children).Error
		if err != nil {
			return nil, err
		}
		firstMenuList[i].Children = children
	}
	return firstMenuList, nil
}

// 获取登录用户权限列表
func (s *SysMenuDao) GetPermissionList(adminId uint) ([]entity.PermissionListVo, error) {
	var permissionList []entity.PermissionListVo
	err := global.DB.Model(&entity.SysAdmin{}).Select("m.value").
		Joins("LEFT JOIN sys_admin_role ar ON sys_admin.id = ar.admin_id").
		Joins("LEFT JOIN sys_role r ON ar.role_id = r.id").
		Joins("LEFT JOIN sys_role_menu rm ON r.id = rm.role_id").
		Joins("LEFT JOIN sys_menu m ON rm.menu_id = m.id").
		Where("sys_admin.id = ?", adminId).
		Where("r.role_status = ?", 1).
		Where("m.menu_status = ?", 1).
		Where("m.menu_type = ?", 1).
		Scan(&permissionList).Error
	if err != nil {
		return nil, err
	}
	return permissionList, nil
}
