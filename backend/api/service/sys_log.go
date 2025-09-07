package service

import (
	"go-admin-server/api/entity"
	"go-admin-server/common/response"
)

type SysLogService struct{}

// 获取登录日志列表
func (s *SysLogService) GetLoginLogList(pageNum, pageSize int, username, beginTime, endTime string, loginStaus uint) (*entity.LoginLogListVo, error) {
	if pageNum < 1 {
		pageNum = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	loginLogList, total, err := SysLogDao.GetLoginLogList(pageNum, pageSize, username, beginTime, endTime, loginStaus)
	if err != nil {
		return nil, response.ErrServerError
	}
	logListVo := &entity.LoginLogListVo{
		Data: loginLogList,
		Pagination: response.PaginationMeta{
			Total:      total,
			PageNum:    pageNum,
			PageSize:   pageSize,
			TotalPages: (total - 1 + pageSize) / pageSize,
		},
	}
	return logListVo, nil
}

// 删除登录日志
func (s *SysLogService) DeleteLoginLog(logId uint) error {
	if err := SysLogDao.DeleteLoginLog(logId); err != nil {
		return response.ErrServerError
	}
	return nil
}

// 批量删除登录日志
func (s *SysLogService) BatchDeleteLoginLog(logIds []uint) error {
	if err := SysLogDao.BatchDeleteLoginLog(logIds); err != nil {
		return response.ErrServerError
	}
	return nil
}

// 获取操作日志列表
func (s *SysLogService) GetOperationLogList(pageNum, pageSize int, username, beginTime, endTime string) (*entity.OperationLogListVo, error) {
	if pageNum < 1 {
		pageNum = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	operationLogList, total, err := SysLogDao.GetOperationLogList(pageNum, pageSize, username, beginTime, endTime)
	if err != nil {
		return nil, response.ErrServerError
	}
	logListVo := &entity.OperationLogListVo{
		Data: operationLogList,
		Pagination: response.PaginationMeta{
			Total:      total,
			PageNum:    pageNum,
			PageSize:   pageSize,
			TotalPages: (total - 1 + pageSize) / pageSize,
		},
	}
	return logListVo, nil
}

// 删除操作日志
func (s *SysLogService) DeleteOpLog(logId uint) error {
	if err := SysLogDao.DeleteOpLog(logId); err != nil {
		return response.ErrServerError
	}
	return nil
}

// 批量删除操作日志
func (s *SysLogService) BatchDeleteOpLog(logIds []uint) error {
	if err := SysLogDao.BatchDeleteOpLog(logIds); err != nil {
		return response.ErrServerError
	}
	return nil
}
