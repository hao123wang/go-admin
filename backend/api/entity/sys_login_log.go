package entity

import (
	"go-admin-server/common/response"
	"go-admin-server/common/utils"
)

// 登录日志模型
type SysLoginLog struct {
	ID            uint        `json:"id" gorm:"column:id;primaryKey"`
	Username      string      `json:"username" gorm:"column:username;type:varchar(50)"`
	IpAddress     string      `json:"ipAddress" gorm:"column:ip_address;type:varchar(128)"`
	LoginLocation string      `json:"loginLocation" gorm:"column:login_location;type:varchar(255)"`
	Browser       string      `json:"browser" gorm:"column:browser;type:varchar(50);comment:'浏览器类型'"`
	Os            string      `json:"os" gorm:"column:os;type:varchar(50);comment:'操作系统'"`
	LoginStatus   uint        `json:"loginStatus" gorm:"column:login_status;comment:'登录状态: 1->成功,2->失败'"`
	Message       string      `json:"message" gorm:"column:message;type:varchar(255);comment:'提示信息'"`
	LoginAt       utils.HTime `json:"loginAt" gorm:"column:login_at;comment:'登录时间'"`
}

func (SysLoginLog) TableName() string {
	return "sys_login_log"
}

// 登录日志列表响应结构体
type LoginLogListVo response.PaginatedResult[SysLoginLog]

// 删除登录日志请求结构体
type DeleteLoginLogDto struct {
	ID uint `json:"id"`
}

// 批量删除登录日志请求结构体
type BatchDeleteLoginLogDto struct {
	Ids []uint `json:"ids"`
}
