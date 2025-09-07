package entity

import "go-admin-server/common/utils"

// SysDept 部门模型
type SysDept struct {
	ID         uint        `gorm:"column:id;primaryKey" json:"id"`
	DeptName   string      `gorm:"column:dept_name;unique;not null" json:"deptName"`
	DeptType   uint        `gorm:"column:dept_type;comment:'部门类型: 1->公司,2->中心,3->部门';not null" json:"deptType"`
	DeptStatus uint        `gorm:"column:dept_status;comment:'部门状态: 1->正常,2->停用';not null;default:1" json:"deptStatus"`
	ParentID   *uint       `gorm:"column:parent_id" json:"parentId"`
	Children   []SysDept   `gorm:"foreignKey:ParentID;references:ID" json:"children"`
	CreateAT   utils.HTime `gorm:"column:created_at" json:"createdAT"`
}

func (SysDept) TableName() string {
	return "sys_dept"
}

// 创建部门请求结构体
type CreateDeptDto struct {
	DeptName string `json:"deptName" binding:"required"`
	DeptType uint   `json:"deptType" binding:"required,oneof=1 2 3"`
	ParentID uint   `json:"parentID,omitempty" binding:"omitempty"`
}

// 根据id查询部门请求结构体
type GetDeptByIdDto struct {
	ID uint `json:"id" binding:"required"`
}

// UpdateSysDeptDto 修改部门信息请求结构体
type UpdateSysDeptDto struct {
	ID         uint    `json:"id" binding:"required"`
	DeptName   *string `json:"deptName,omitempty" binding:"omitempty"`
	DeptType   *uint   `json:"deptTyep,omitempty" binding:"omitempty,oneof=1 2 3"`
	DeptStatus *uint   `json:"deptStatus,omitempty" binding:"omitempty,oneof=1 2"`
	ParentID   *uint   `json:"parentID,omitempty" binding:"omitempty"`
}

// 删除部门请求结构体
type DeleteDeptDto struct {
	ID uint `json:"id" binding:"required"`
}

// 部门下拉列表响应结构体
type DeptDropdownVo struct {
	ID       uint   `json:"id"`
	DeptName string `json:"deptName"`
	ParentID *uint  `json:"parentID"`
}
