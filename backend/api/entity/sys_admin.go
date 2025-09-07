package entity

import (
	"go-admin-server/common/response"
	"go-admin-server/common/utils"
)

// 用户模型
type SysAdmin struct {
	ID        uint        `gorm:"column:id;primaryKey"`
	Username  string      `gorm:"column:username;type:varchar(64);unique;not null"`
	Password  string      `gorm:"column:password;type:varchar(64);not null"`
	Nickname  string      `gorm:"column:nickname;type:varchar(64);comment:'昵称';unique;not null"`
	Icon      string      `gorm:"column:icon;type:varchar(500);comment:'用户头像'"`
	Email     string      `gorm:"column:email;type:varchar(64)"`
	Phone     string      `gorm:"column:phone;type:char(11)"`
	Note      string      `gorm:"column:note;type:varchar(500);comment:'备注'"`
	Status    uint        `gorm:"column:status;comment:'账号状态:1->启用,2->禁用';not null;default:1"`
	DeptID    uint        `gorm:"column:dept_id;comment:'部门id'"`
	PostID    uint        `gorm:"column:post_id;comment:'岗位id'"`
	CreatedAt utils.HTime `gorm:"column:created_at"`
}

func (SysAdmin) TableName() string {
	return "sys_admin"
}

// 鉴权用户结构体
type JwtAdmin struct {
	ID       uint   `json:"id"`       // ID
	Username string `json:"username"` // 用户账号
	Nickname string `json:"nickname"` // 昵称
	Icon     string `json:"icon"`     // 头像
	Email    string `json:"email"`    // 邮箱
	Phone    string `json:"phone"`    // 电话
	Note     string `json:"note"`     // 备注
}

// 登录请求结构体
type LoginDto struct {
	Username     string `json:"username" binding:"required"`
	Password     string `json:"password" binding:"required"`
	CaptchaID    string `json:"captchaId" binding:"required"`
	CaptchaImage string `json:"captchaImage" binding:"required"`
}

// 创建用户请求结构体
type CreateAdminDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone" binding:"required"`
	Note     string `json:"note"`
	Status   uint   `json:"staus" binding:"required,oneof=1 2"`
	PostID   uint   `json:"postID" binding:"required"`
	DeptID   uint   `json:"deptID" binding:"required"`
	RoleID   uint   `json:"roleID" binding:"required"`
}

// 联表查询用户信息
type AdminList struct {
	ID       uint   `json:"id"`       // ID
	Username string `json:"username"` // 用户名
	Nickname string `json:"nickname"` // 昵称
	Status   uint   `json:"status"`   // 状态：1->启用,2->禁用
	PostId   uint   `json:"postId"`   // 岗位id
	DeptId   uint   `json:"deptId"`   // 部门id
	RoleId   uint   `json:"roleId" `  // 角色id
	PostName string `json:"postName"` // 岗位名称
	DeptName string `json:"deptName"` // 部门名称
	RoleName string `json:"roleName"` // 角色名称
	Icon     string `json:"icon"`     // 头像
	Email    string `json:"email"`    // 邮箱
	Phone    string `json:"phone"`    // 电话
	Note     string `json:"note"`     // 备注
}

// 将联表查询的用户信息列表与分页信息放到一起返回给前端
type AdminListVo response.PaginatedResult[AdminList]

// 联表查询单个用户的请求结构体
type GetAdminByIdDto struct {
	ID uint `json:"id" binding:"required"`
}

// 联表查询单个用户的响应结构体
type GetAdminByIdVo struct {
	ID       uint   `json:"id"`       // ID
	Username string `json:"username"` // 用户名
	Nickname string `json:"nickname"` // 昵称
	Status   uint   `json:"status"`   // 状态：1->启用,2->禁用
	PostId   uint   `json:"postId"`   // 岗位id
	DeptId   uint   `json:"deptId"`   // 部门id
	RoleId   uint   `json:"roleId" `  // 角色id
	Email    string `json:"email"`    // 邮箱
	Phone    string `json:"phone"`    // 手机号
	Note     string `json:"note"`     // 备注
}

// 修改用户请求结构体
type UpdateAdminDto struct {
	ID       uint    `json:"id" binding:"required"`
	PostId   *uint   `json:"postId"`
	DeptId   *uint   `json:"deptId"`
	RoleId   *uint   `json:"roleId"`
	Username *string `json:"username"`
	Nickname *string `json:"nickname"`
	Phone    *string `json:"phone"`
	Email    *string `json:"email" binding:"omitempty,email"`
	Note     *string `json:"note"`
	Status   *uint   `json:"status" binding:"omitempty,oneof=1 2"`
}

// 删除用户请求结构体
type DeleteAdminDto struct {
	ID uint `json:"id" binding:"required"`
}

// 修改用户状态请求结构体
type UpdateAdminStatusDto struct {
	ID        uint `json:"id" binding:"required"`
	NewStatus uint `json:"newStatus" binding:"omitempty,oneof=1 2"`
}

// 重置密码请求结构体
type ResetPasswordDto struct {
	ID          uint   `json:"id" binding:"required"`
	NewPassword string `json:"newPassword"`
}

// 修改个人资料请求结构体
type UpdatePersonalDto struct {
	Icon     *string `json:"icon"`
	Username *string `json:"username"`
	Nickname *string `json:"nichname"`
	Phone    *string `json:"phone"`
	Email    *string `json:"email"`
	Note     *string `json:"note"`
}

// 修改个人密码请求结构体
type UpdatePasswordDto struct {
	Password    string `json:"password" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
	RePassword  string `json:"rePassword" binding:"required"`
}
