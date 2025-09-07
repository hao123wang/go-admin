package entity

import (
	"go-admin-server/common/response"
	"go-admin-server/common/utils"
)

// 操作日志模型
type SysOperationLog struct {
	ID        uint        `json:"id" gorm:"column:id;primaryKey"`
	AdminID   uint        `json:"adminId" gorm:"column:admin_id;not null;"`
	Username  string      `json:"username" gorm:"column:user_name;type:varchar(64);not null"`
	Method    string      `json:"method" gorm:"column:method;type:varchar(64);not null"`
	Ip        string      `json:"ip" gorm:"column:ip;type:varchar(128)"`
	Url       string      `json:"url" gorm:"column:url;type:varchar(500)"`
	CreatedAt utils.HTime `json:"createdAt" gorm:"column:created_at"`
}

func (SysOperationLog) TableName() string {
	return "sys_operation_log"
}

// 操作日志列表响应结构体
type OperationLogListVo response.PaginatedResult[SysOperationLog]

// 删除操作日志请求结构体
type DeleteOpLogDto struct {
	ID uint `json:"id"`
}

// 批量删除操作日志请求结构体
type BatchDeleteOpLogDto struct {
	Ids []uint `json:"ids"`
}
