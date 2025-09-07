package service

import (
	"errors"
	"go-admin-server/api/entity"
	"go-admin-server/common/response"
	"go-admin-server/common/utils"
	"time"

	"gorm.io/gorm"
)

type SysMenuService struct{}

// 设置菜单的父id
func (s *SysMenuService) setupParentID(parentID uint, sysMenu *entity.SysMenu) error {
	if sysMenu.MenuType == 1 {
		// 顶级菜单，父id为nil
		sysMenu.ParentID = nil
	} else {
		// 根据id获取父菜单
		parentMenu, err := SysMenuDao.GetMenuByID(parentID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 父菜单不存在
				return response.ErrInvalidMenuParentID
			}
			return response.ErrServerError
		}
		if parentMenu.MenuStatus == 2 {
			// 父菜单已被禁用
			return response.ErrParentMenuDisabled
		}
		sysMenu.ParentID = &parentID
	}
	return nil
}

// 创建菜单
func (s *SysMenuService) CreateMenu(dto *entity.CreateMenuDto) error {
	// 检查名称存在性
	exists, err := SysMenuDao.ExistsByName(dto.MenuName)
	if err != nil {
		return response.ErrServerError
	}
	if exists {
		return response.ErrMenuNameExists
	}

	// 创建菜单实例
	sysMenu := &entity.SysMenu{
		MenuName:   dto.MenuName,
		MenuIcon:   dto.MenuIcon,
		MenuType:   dto.MenuType,
		MenuStatus: dto.MenuStatus,
		Url:        dto.Url,
		Sort:       dto.Sort,
		CreateAT:   utils.HTime{Time: time.Now()},
	}
	// 设置菜单的父id
	if err := s.setupParentID(*dto.ParentID, sysMenu); err != nil {
		return err
	}

	// 创建菜单
	if err := SysMenuDao.CreateMenu(sysMenu); err != nil {
		return response.ErrServerError
	}
	return nil
}

// 获取菜单列表
func (s *SysMenuService) GetMenuList(menuName string, menuStatus int) ([]entity.SysMenu, error) {
	sysMenus, err := SysMenuDao.GetMenuList(menuName, menuStatus)
	if err != nil {
		return nil, response.ErrServerError
	}
	return sysMenus, nil
}

// 根据id获取菜单
func (s *SysMenuService) GetMenuById(menuID uint) (*entity.SysMenu, error) {
	sysMenu, err := SysMenuDao.GetMenuByID(menuID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrMenuNotExists
		}
		return nil, response.ErrServerError
	}
	return sysMenu, nil
}

// 修改菜单信息
func (s *SysMenuService) UpdateMenu(dto *entity.UpdateSysMenuDto) error {
	// 根据id获取菜单
	sysMenu, err := SysMenuDao.GetMenuByID(dto.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ErrMenuNotExists
		}
		return response.ErrServerError
	}

	// 修改名称
	if dto.MenuName != nil && *dto.MenuName != sysMenu.MenuName {
		// 检查名称存在性
		exists, err := SysMenuDao.ExistsByName(*dto.MenuName)
		if err != nil {
			return response.ErrServerError
		}
		if exists {
			return response.ErrMenuNameExists
		}
		sysMenu.MenuName = *dto.MenuName
	}

	// 修改图标
	if dto.MenuIcon != nil && *dto.MenuIcon != sysMenu.MenuIcon {
		sysMenu.MenuIcon = *dto.MenuIcon
	}

	// 修改类型
	if dto.MenuType != nil && *dto.MenuType != sysMenu.MenuType {
		sysMenu.MenuType = *dto.MenuType
	}

	// 修改状态
	if dto.MenuStatus != nil && *dto.MenuStatus != sysMenu.MenuStatus {
		sysMenu.MenuStatus = *dto.MenuStatus
	}

	if dto.Value != nil && *dto.Value != sysMenu.Value {
		sysMenu.Value = *dto.Value
	}
	if dto.Url != nil && *dto.Url != sysMenu.Url {
		sysMenu.Url = *dto.Url
	}
	if dto.Sort != nil && *dto.Sort != sysMenu.Sort {
		sysMenu.Sort = *dto.Sort
	}

	// 修改父id，需进行对应判断
	if dto.ParentID != nil {
		// 不能将父菜单id设置为自身的id
		if sysMenu.ID == *dto.ParentID {
			return response.ErrInvalidMenuParentID
		}
		// 当前部门的父id为nil || 当前部门的父id不等于新的父id
		if sysMenu.ParentID == nil || *sysMenu.ParentID != *dto.ParentID {
			// 设置新的父id
			if err := s.setupParentID(*dto.ParentID, sysMenu); err != nil {
				return err
			}
		}
	}

	if err := SysMenuDao.UpdateMenu(sysMenu); err != nil {
		return response.ErrServerError
	}
	return nil
}

// 删除单个菜单
func (s *SysMenuService) DeleteMenu(menuID uint) error {
	hasSubmenu, err := SysMenuDao.HasSubMenu(menuID)
	if err != nil {
		return response.ErrServerError
	}
	if hasSubmenu {
		return response.ErrHasSubmenu
	}
	if err := SysMenuDao.DeleteMenu(menuID); err != nil {
		return response.ErrServerError
	}
	return nil
}

// 获取菜单下拉列表
func (s *SysMenuService) GetMenuDropdown() ([]entity.MenuDropdownVo, error) {
	dropdown, err := SysMenuDao.GetMenuDropdown()
	if err != nil {
		return nil, response.ErrServerError
	}
	return dropdown, nil
}

// 获取左侧菜单列表
func (s *SysMenuService) GetLeftMenuList(adminId uint) ([]entity.FirstLevelMenuVo, error) {
	firstMenuList, err := SysMenuDao.LeftMenuList(adminId)
	if err != nil {
		return nil, response.ErrServerError
	}
	return firstMenuList, nil
}

// 获取用户的权限列表
func (s *SysMenuService) GetPermissionList(adminId uint) ([]entity.PermissionListVo, error) {
	permissionList, err := SysMenuDao.GetPermissionList(adminId)
	if err != nil {
		return nil, response.ErrServerError
	}
	return permissionList, nil
}
