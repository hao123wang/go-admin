package entity

import "go-admin-server/common/utils"

// 菜单模型
type SysMenu struct {
	ID         uint        `gorm:"column:id;primaryKey" json:"id"`
	MenuName   string      `gorm:"column:menu_name;type:varchar(64);unique;not null" json:"menuName"`
	MenuIcon   string      `gorm:"column:menu_icon;type:varchar(64);not null" json:"menuIcon"`
	MenuType   uint        `gorm:"column:menu_type;comment:'菜单类型: 1->目录,2->菜单,3->按钮'" json:"menuType"`
	MenuStatus uint        `gorm:"column:menu_status;comment:'菜单状态: 1->启用,2->禁用';not null;default:1" json:"menuStatus"`
	Url        string      `gorm:"column:url;type:varchar(100);comment:'路由路径'" json:"url"`
	Value      string      `gorm:"column:value;type:varchar(64);comment:'权限值'" json:"value"`
	Sort       uint        `gorm:"column:sort" json:"sort"`
	ParentID   *uint       `gorm:"column:parent_id" json:"parentId"`
	Children   []SysMenu   `gorm:"foreignKey:ParentID" json:"children"`
	CreateAT   utils.HTime `gorm:"created_at" json:"createdAt"`
}

func (SysMenu) TableName() string {
	return "sys_menu"
}

// 创建菜单请求结构体
type CreateMenuDto struct {
	MenuName   string `json:"menuName" binding:"required"`
	MenuIcon   string `json:"menuIcon" binding:"required"`
	MenuType   uint   `json:"menuType" binding:"required,oneof=1 2 3"`
	MenuStatus uint   `json:"menyStatus" binding:"omitempty,oneof=1 2"`
	Url        string `json:"url,omitempty"`
	Value      string `json:"value,omitempty"`
	Sort       uint   `json:"sort"`
	ParentID   *uint  `json:"parentID"`
}

// 根据id查询菜单请求结构体
type GetMenuByIdDto struct {
	ID uint `json:"id" binding:"required"`
}

// 修改菜单请求结构体
type UpdateSysMenuDto struct {
	ID         uint    `json:"id" binding:"required"`
	MenuName   *string `json:"menuName,omitempty"`
	MenuIcon   *string `json:"menuIcon,omitempty"`
	MenuType   *uint   `json:"menuType,omitempty" binding:"omitempty,oneof=1 2 3"`
	MenuStatus *uint   `json:"menyStatus,omitempty" binding:"omitempty,oneof=1 2"`
	Url        *string `json:"url,omitempty"`
	Value      *string `json:"value,omitempty"`
	Sort       *uint   `json:"sort,omitempty"`
	ParentID   *uint   `json:"parentID,omitempty"`
}

// 删除菜单请求结构体
type DeleteMenuDto struct {
	ID uint `json:"id" binding:"required"`
}

// 菜单下拉列表响应结构体
type MenuDropdownVo struct {
	ID       uint   `json:"id"`
	MenuName string `json:"menuName"`
	ParentID *uint  `json:"parentID"`
}

// 二级菜单展示信息的响应结构体
type SecondLevelMenuVo struct {
	MenuName string `json:"menuName"`
	MenuIcon string `json:"menuIcon"`
	Url      string `json:"url"`
}

// 一级菜单展示信息的结构体
type FirstLevelMenuVo struct {
	ID       uint                `json:"id"`
	MenuName string              `json:"menuName"`
	MenuIcon string              `json:"menuIcon"`
	Url      string              `json:"url"`
	Children []SecondLevelMenuVo `json:"children,omitempty"`
}

// 展示用户所有的权限字符串列表
type PermissionListVo struct {
	Value string `json:"value"` // 权限
}
