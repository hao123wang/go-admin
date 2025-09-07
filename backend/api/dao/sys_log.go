package dao

import (
	"go-admin-server/api/entity"
	"go-admin-server/common/utils"
	"go-admin-server/global"
	"time"
)

type SysLogDao struct{}

// 创建登录日志
func (d *SysLogDao) CreateLoginLog(username, ipAddr, loginLocation, browser, os, message string, loginStaus uint) {
	loginLog := &entity.SysLoginLog{
		Username:      username,
		IpAddress:     ipAddr,
		LoginLocation: loginLocation,
		Browser:       browser,
		Os:            os,
		Message:       message,
		LoginStatus:   loginStaus,
		LoginAt:       utils.HTime{Time: time.Now()},
	}
	global.DB.Create(loginLog)
}

// 获取登录日志列表
func (d *SysLogDao) GetLoginLogList(pageNum, pageSize int, username, beginTime, endTime string, loginStatus uint) ([]entity.SysLoginLog, int, error) {
	var loginLogList []entity.SysLoginLog
	query := global.DB.Model(&entity.SysLoginLog{})

	if loginStatus != 0 {
		query = query.Where("login_status = ?", loginStatus)
	}

	if username != "" {
		query = query.Where("username = ?", username)
	}

	if beginTime != "" && endTime != "" {
		query = query.Where("login_at BETWEEN ? AND ?", beginTime, endTime)
	}

	var count int64
	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	offset := (pageNum - 1) * pageSize
	if err := query.Limit(pageSize).Offset(offset).Order("login_at DESC").Find(&loginLogList).Error; err != nil {
		return nil, 0, err
	}
	return loginLogList, int(count), nil
}

// 根据id删除登录日志
func (d *SysLogDao) DeleteLoginLog(logId uint) error {
	return global.DB.Where("id = ?", logId).Delete(&entity.SysLoginLog{}).Error
}

// 批量删除登录日志
func (d *SysLogDao) BatchDeleteLoginLog(logIds []uint) error {
	return global.DB.Where("id IN (?)", logIds).Delete(&entity.SysLoginLog{}).Error
}

// 创建操作日志
func (d *SysLogDao) CreateOperationLog(operaLog *entity.SysOperationLog) error {
	return global.DB.Create(operaLog).Error
}

// 获取操作日志列表
func (d *SysLogDao) GetOperationLogList(pageNum, pageSize int, username, beginTime, endTime string) ([]entity.SysOperationLog, int, error) {
	var operationLogList []entity.SysOperationLog
	query := global.DB.Model(&entity.SysOperationLog{})

	if username != "" {
		query = query.Where("username = ?", username)
	}

	if beginTime != "" && endTime != "" {
		query = query.Where("login_at BETWEEN ? AND ?", beginTime, endTime)
	}

	var count int64
	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	offset := (pageNum - 1) * pageSize
	if err := query.Limit(pageSize).Offset(offset).Order("created_at DESC").Find(&operationLogList).Error; err != nil {
		return nil, 0, err
	}
	return operationLogList, int(count), nil
}

// 根据id删除操作日志
func (d *SysLogDao) DeleteOpLog(logId uint) error {
	return global.DB.Where("id = ?", logId).Delete(&entity.SysOperationLog{}).Error
}

// 批量删除操作日志
func (d *SysLogDao) BatchDeleteOpLog(logIds []uint) error {
	return global.DB.Where("id IN (?)", logIds).Delete(&entity.SysOperationLog{}).Error
}
