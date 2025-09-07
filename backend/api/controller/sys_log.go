package controller

import (
	"go-admin-server/api/entity"
	"go-admin-server/common/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 查询登录日志列表
// @Description 查询登录日志列表
// @Tags 日志管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param pageNum query int false "页码"
// @Param pageSize query int false "页大小"
// @Param username query string false "用户名"
// @Param loginStatus query int false "登录状态"
// @Param beginTime query string false "开始时间"
// @Param endTime query string false "结束时间"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/logService/getLoginLogList [get]
func GetLoginLogList(c *gin.Context) {
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	username := c.Query("username")
	loginStaus, _ := strconv.ParseUint(c.Query("loginStatus"), 10, 64)
	beginTime := c.Query("beginTime")
	endTime := c.Query("endTime")

	logListVo, err := LogService.GetLoginLogList(pageNum, pageSize, username, beginTime, endTime, uint(loginStaus))
	if err != nil {
		response.Error(c, err)
		return
	}
	response.SuccessWithData(c, logListVo)
}

// @Summary 删除登录日志
// @Description 删除登录日志
// @Tags 日志管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body entity.DeleteLoginLogDto true "删除登录日志请求"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/logService/deleteLoginLog [post]
func DeleteLoginLog(c *gin.Context) {
	var dto entity.DeleteLoginLogDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.ValidationError(c, err)
		return
	}
	if err := LogService.DeleteLoginLog(dto.ID); err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c)
}

// @Summary 批量删除登录日志
// @Description 批量删除登录日志
// @Tags 日志管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body entity.BatchDeleteLoginLogDto true "批量删除登录日志请求"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/logService/batchDeleteLoginLog [post]
func BatchDeleteLoginLog(c *gin.Context) {
	var dto entity.BatchDeleteLoginLogDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.ValidationError(c, err)
		return
	}
	if err := LogService.BatchDeleteLoginLog(dto.Ids); err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c)
}

// @Summary 查询操作日志列表
// @Description 查询操作日志列表
// @Tags 日志管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param pageNum query int false "页码"
// @Param pageSize query int false "页大小"
// @Param username query string false "用户名"
// @Param beginTime query string false "开始时间"
// @Param endTime query string false "结束时间"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/logService/getOperationLogList [get]
func GetOpLogList(c *gin.Context) {
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	username := c.Query("username")
	beginTime := c.Query("beginTime")
	endTime := c.Query("endTime")

	logListVo, err := LogService.GetOperationLogList(pageNum, pageSize, username, beginTime, endTime)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.SuccessWithData(c, logListVo)
}

// @Summary 删除操作日志
// @Description 删除操作日志
// @Tags 日志管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body entity.DeleteOpLogDto true "删除操作日志请求"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/logService/deleteOpLog [post]
func DeleteOpLog(c *gin.Context) {
	var dto entity.DeleteOpLogDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.ValidationError(c, err)
		return
	}
	if err := LogService.DeleteOpLog(dto.ID); err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c)
}

// @Summary 批量删除操作日志
// @Description 批量操作登录日志
// @Tags 日志管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body entity.BatchDeleteOpLogDto true "批量删除操作日志请求"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/logService/batchDeleteOpLog [post]
func BatchDeleteOpLog(c *gin.Context) {
	var dto entity.BatchDeleteOpLogDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.ValidationError(c, err)
		return
	}
	if err := LogService.BatchDeleteOpLog(dto.Ids); err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c)
}
