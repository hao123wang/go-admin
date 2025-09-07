package entity

import (
	"go-admin-server/common/response"
	"go-admin-server/common/utils"
)

// 角色模型
type SysRole struct {
	ID          uint        `gorm:"column:id;primaryKey" json:"id"`
	RoleName    string      `gorm:"column:role_name;type:varchar(64);unique;not null" json:"roleName"`
	RoleKey     string      `gorm:"column:role_key;type:varchar(64);comment:'权限字符串';unique;not null" json:"roleKey"`
	RoleStatus  uint        `gorm:"column:role_status;comment:'角色状态: 1->启用,2->禁用';not null;default:1" json:"roleStatus"`
	Description string      `gorm:"column:description;type:varchar(500)" json:"description"`
	CreatedAt   utils.HTime `gorm:"column:created_at" json:"createdAt"`
}

func (SysRole) TableName() string {
	return "sys_role"
}

// 创建角色请求结构体
type CreateRoleDto struct {
	RoleName    string `json:"roleName" binding:"required"`
	RoleKey     string `json:"roleKey" binding:"required"`
	RoleStatus  uint   `json:"roleStatus" binding:"omitempty,oneof=1 2"`
	Description string `json:"description" binding:"omitempty"`
}

// 查询角色列表响应结构体，将角色列表与分页信息封装到响应结构体中
type RoleListVo response.PaginatedResult[SysRole]

// 根据id查询角色请求结构体
type GetRoleByIdDto struct {
	ID uint `json:"id" binding:"required"`
}

// 修改角色请求结构体
type UpdateRoleDto struct {
	ID          uint    `json:"id" binding:"required"`
	RoleName    *string `json:"roleName"`
	RoleKey     *string `json:"roleKey"`
	RoleStatus  *uint   `json:"roleStatus" binding:"omitempty,oneof=1 2"`
	Description *string `json:"description"`
}

// 删除角色请求结构体
type DeleteRoleDto struct {
	ID uint `json:"id" binding:"required"`
}

// 修改角色状态请求结构体
type UpdateRoleStatusDto struct {
	ID        uint `json:"id" binding:"required"`
	NewStatus uint `json:"newStatus" binding:"omitempty,oneof=1 2"`
}

// 角色下拉列表
type RoleDropdownVo struct {
	ID       uint   `json:"id"`
	RoleName string `json:"role_name"`
}

// 查询角色权限请求结构体
type GetRoleMenusDto struct {
	ID uint `json:"id" binding:"required"`
}

// 角色分配的权限（菜单）列表
type AssignRoleMenusDto struct {
	ID      uint   `json:"id" binding:"required"` // 角色id
	MenuIDs []uint `json:"menuIds" binding:"required"`
}
