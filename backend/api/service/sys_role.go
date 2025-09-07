package service

import (
	"errors"
	"go-admin-server/api/entity"
	"go-admin-server/common/response"
	"go-admin-server/common/utils"
	"time"

	"gorm.io/gorm"
)

type SysRoleService struct{}

// 创建角色
func (s *SysRoleService) CreateRole(dto *entity.CreateRoleDto) error {
	// 检查名称存在性
	nameExists, err := SysRoleDao.ExistsByName(dto.RoleName)
	if err != nil {
		return response.ErrServerError
	}
	if nameExists {
		return response.ErrRoleNameExists
	}
	// 检查关键字存在性
	keyExists, err := SysRoleDao.ExistsByKey(dto.RoleKey)
	if err != nil {
		return response.ErrServerError
	}
	if keyExists {
		return response.ErrRoleKeyExists
	}
	// 创建角色
	sysRole := &entity.SysRole{
		RoleName:    dto.RoleName,
		RoleKey:     dto.RoleKey,
		Description: dto.Description,
		CreatedAt:   utils.HTime{Time: time.Now()},
	}
	if dto.RoleStatus == 0 {
		sysRole.RoleStatus = 1
	} else {
		sysRole.RoleStatus = dto.RoleStatus
	}
	if err := SysRoleDao.CreateRole(sysRole); err != nil {
		return response.ErrServerError
	}
	return nil
}

// 获取角色列表
func (s *SysRoleService) GetRoleList(pageNum, pageSize, roleStatus int, roleName, beginTime, endTime string) (*entity.RoleListVo, error) {
	if pageNum < 1 {
		pageNum = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if roleStatus != 1 && roleStatus != 2 {
		roleStatus = 1
	}
	sysRoleList, total, err := SysRoleDao.GetRoleList(pageNum, pageSize, roleStatus, roleName, beginTime, endTime)
	if err != nil {
		return nil, response.ErrServerError
	}
	paginatedResult := &entity.RoleListVo{
		Data: sysRoleList,
		Pagination: response.PaginationMeta{
			PageNum:    pageNum,
			PageSize:   pageSize,
			Total:      total,
			TotalPages: (total + pageSize - 1) / pageSize,
		},
	}
	return paginatedResult, nil
}

// 根据id获取角色
func (s *SysRoleService) GetRoleByID(roleID uint) (*entity.SysRole, error) {
	sysRole, err := SysRoleDao.GetRoleByID(roleID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrRoleNotExists
		}
		return nil, response.ErrServerError
	}
	return sysRole, nil
}

// 修改角色
func (s *SysRoleService) UpdateRole(dto *entity.UpdateRoleDto) error {
	// 获取要修改的角色
	sysRole, err := s.GetRoleByID(dto.ID)
	if err != nil {
		return err
	}
	// 是否修改名称
	if dto.RoleName != nil && *dto.RoleName != sysRole.RoleName {
		// 检查名称存在性
		nameExists, _ := SysRoleDao.ExistsByName(*dto.RoleName)
		if nameExists {
			return response.ErrRoleNameExists
		}
		sysRole.RoleName = *dto.RoleName
	}
	// 是否修改关键字
	if dto.RoleKey != nil && *dto.RoleKey != sysRole.RoleKey {
		// 检查关键字存在性
		keyExists, _ := SysRoleDao.ExistsByKey(*dto.RoleKey)
		if keyExists {
			return response.ErrRoleKeyExists
		}
		sysRole.RoleKey = *dto.RoleKey
	}
	// 修改状态
	if dto.RoleStatus != nil && *dto.RoleStatus != sysRole.RoleStatus {
		sysRole.RoleStatus = *dto.RoleStatus
	}
	// 修改描述
	if dto.Description != nil {
		sysRole.Description = *dto.Description
	}
	// 更新数据库
	if err := SysRoleDao.UpdateRole(sysRole); err != nil {
		return response.ErrServerError
	}
	return nil
}

// 删除角色
func (s *SysRoleService) DeleteRole(roleID uint) error {
	// 先检查角色是否存在
	_, err := s.GetRoleByID(roleID)
	if err != nil {
		return err
	}
	// 删除角色
	if err := SysRoleDao.DeleteRole(roleID); err != nil {
		return response.ErrServerError
	}
	return nil
}

// 修改角色状态
func (s *SysRoleService) UpdateRoleStatus(dto *entity.UpdateRoleStatusDto) error {
	// 根据id获取角色
	role, err := SysRoleDao.GetRoleByID(dto.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ErrRoleNotExists
		}
		return response.ErrServerError
	}
	role.RoleStatus = dto.NewStatus
	if err := SysRoleDao.UpdateRole(role); err != nil {
		return response.ErrServerError
	}
	return nil
}

// 获取角色下拉列表
func (s *SysRoleService) GetRoleDropdown() ([]entity.RoleDropdownVo, error) {
	dropdown, err := SysRoleDao.GetRoleDropdown()
	if err != nil {
		return nil, response.ErrServerError
	}
	return dropdown, nil
}

// 角色权限分配
func (s *SysRoleService) AssignRoleMenus(dto *entity.AssignRoleMenusDto) error {
	// 判断角色id是否存在
	roleExists, err := SysRoleDao.ExistsByID(dto.ID)
	if err != nil {
		return response.ErrServerError
	}
	if !roleExists {
		return response.ErrRoleNotExists
	}

	// 判断列表中的菜单是否全部存在
	menuAllExists, err := SysRoleDao.ExistsMenuIds(dto.MenuIDs)
	if err != nil {
		return response.ErrServerError
	}
	if !menuAllExists {
		return response.ErrMenuNotExists
	}

	// 分配权限
	if err := SysRoleDao.AssignRoleMenus(dto.ID, dto.MenuIDs); err != nil {
		return response.ErrServerError
	}
	return nil
}

// 获取角色权限列表
func (s *SysRoleService) GetRoleMenus(roleID uint) ([]uint, error) {
	roleExists, err := SysRoleDao.ExistsByID(roleID)
	if err != nil {
		return nil, response.ErrServerError
	}
	if !roleExists {
		return nil, response.ErrRoleNotExists
	}
	roleMenus, err := SysRoleDao.GetRoleMenus(roleID)
	if err != nil {
		return nil, response.ErrServerError
	}
	return roleMenus, nil
}
