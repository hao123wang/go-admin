package entity

import (
	"go-admin-server/common/response"
	"go-admin-server/common/utils"
)

// SysPost 岗位模型
type SysPost struct {
	ID          uint        `gorm:"column:id;primaryKey" json:"id"`
	PostName    string      `gorm:"column:post_name;type:varchar(64);comment:'岗位名称';not null" json:"postName"`
	PostCode    string      `gorm:"column:post_code;type:varchar(64);comment:'岗位编码';not null" json:"postCode"`
	Remark      string      `gorm:"column:remark;type:varchar(64);comment:'备注'" json:"remark"`
	PostStatus  uint        `gorm:"column:post_status;comment:'岗位状态:1->正常,2->停用';not null;default:1" json:"postStatus"`
	CreatedTime utils.HTime `gorm:"column:created_at" json:"createdAT"`
}

func (SysPost) TableName() string {
	return "sys_post"
}

// 新增岗位请求结构体
type CreateSysPostDto struct {
	PostName   string `json:"postName" binding:"required"`
	PostCode   string `json:"postCode" binding:"required"`
	Remark     string `json:"remark"`
	PostStauts uint   `json:"psotStatus" binding:"omitempty,oneof=1 2"`
}

// 获取岗位列表响应结构体，将数据与分页信息封装到一起返回
type SysPostListVo response.PaginatedResult[SysPost]

// 根据id查询岗位信息请求结构体
type GetPostByIdDto struct {
	ID uint `json:"id" binding:"required"`
}

// 修改岗位信息请求结构体
type UpdateSysPostDto struct {
	ID         uint    `json:"id" binding:"required"`
	PostName   *string `json:"postName,omitempty"`
	PostCode   *string `json:"postCode,omitempty"`
	Remark     *string `json:"remark,omitempty"`
	PostStauts *uint   `json:"postStatus,omitempty" binding:"omitempty,oneof=1 2"`
}

// 删除单个岗位请求结构体
type DeletePostByIdDto struct {
	ID uint `json:"id" binding:"required"`
}

// BatchDeletePostsDto 批量删除岗位请求结构体
type BatchDeletePostsDto struct {
	PostIds []uint `json:"postIds" binding:"required,min=1"`
}

// 修改岗位状态请求结构体
type UpdatePostStatusDto struct {
	ID        uint `json:"id" binding:"required"`
	NewStatus uint `json:"newStatus" binding:"omitempty,oneof=1 2"`
}

// SysPostDropDownVo 岗位下拉列表(创建用户时提供选择)
type SysPostDropdownVo struct {
	ID       uint   `json:"id"`
	PostName string `json:"postName"`
}
