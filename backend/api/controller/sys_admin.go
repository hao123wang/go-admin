package controller

import (
	"go-admin-server/api/entity"
	"go-admin-server/common/response"
	"go-admin-server/common/utils"
	"go-admin-server/global"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 用户登录
// @Description 用户登录
// @Tags 无需认证接口
// @Accept json
// @Produce json
// @Param data body entity.LoginDto true "登录请求结构体"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/login [post]
func Login(c *gin.Context) {
	var dto entity.LoginDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.ValidationError(c, err)
		return
	}
	// 获取当前登录用户的ip、浏览器、操作系统
	ip := c.ClientIP()
	browser := utils.GetBrowser(c)
	Os := utils.GetOS(c)

	// 用户登录
	user, token, err := SysAdminService.Login(ip, browser, Os, &dto)
	if err != nil {
		response.Error(c, err)
		return
	}
	// 用于前端展示的左侧菜单列表
	leftMenuList, err := SysMenuService.GetLeftMenuList(user.ID)
	if err != nil {
		response.Error(c, err)
		return
	}
	// 当前登录用户的权限列表
	permissionList, err := SysMenuService.GetPermissionList(user.ID)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.SuccessWithData(c, map[string]any{
		"sysAdmin":       user,
		"token":          token,
		"leftMenuList":   leftMenuList,
		"permissionList": permissionList,
	})
}

// @Summary 创建用户
// @Description 创建用户
// @Tags 用户管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body entity.CreateAdminDto true "创建用户请求结构体"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/adminService/createAdmin [post]
func CreateAdmin(c *gin.Context) {
	var dto entity.CreateAdminDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.ValidationError(c, err)
		return
	}
	if err := SysAdminService.CreateAdmin(&dto); err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c)
}

// @Summary 查询用户列表
// @Description 查询用户列表
// @Tags 用户管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param pageNum query int false "页码"
// @Param pageSize query int false "页大小"
// @Param username query string false "用户名"
// @Param status query int false "状态"
// @Param beginTime query string false "开始时间"
// @Param endTime query string false "结束时间"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/adminService/getAdminList [get]
func GetAdminList(c *gin.Context) {
	// 接收查询参数
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	status, _ := strconv.Atoi(c.Query("status"))
	username := c.Query("username")
	beginTime := c.Query("beginTime")
	endTime := c.Query("endTime")

	adminListVo, err := SysAdminService.JointGetAdminList(pageNum, pageSize, status, username, beginTime, endTime)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.SuccessWithData(c, adminListVo)
}

// @Summary 根据id查询用户
// @Description 根据id查询用户
// @Tags 用户管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body entity.GetAdminByIdDto true "根据id查询用户请求"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/adminService/getAdminById [post]
func GetAdminById(c *gin.Context) {
	var dto entity.GetAdminByIdDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.ValidationError(c, err)
		return
	}
	sysAdmin, err := SysAdminService.JointGetAdminById(dto.ID)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.SuccessWithData(c, sysAdmin)
}

// @Summary 修改用户信息
// @Description 修改用户信息
// @Tags 用户管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body entity.UpdateAdminDto true "修改用户请求结构体"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/adminService/updateAdmin [post]
func UpdateAdmin(c *gin.Context) {
	var dto entity.UpdateAdminDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.ValidationError(c, err)
		return
	}
	if err := SysAdminService.UpdateSysAdmin(&dto); err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c)
}

// @Summary 删除用户
// @Description 删除用户
// @Tags 用户管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body entity.DeleteAdminDto true "删除用户请求"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/adminService/deleteAdmin [post]
func DeleteAdmin(c *gin.Context) {
	var dto entity.DeleteAdminDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.ValidationError(c, err)
		return
	}
	if err := SysAdminService.DeleteAdmin(dto.ID); err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c)
}

// @Summary 修改用户状态
// @Description 修改用户状态
// @Tags 用户管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body entity.UpdateAdminStatusDto true "修改用户状态请求"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/adminService/updateAdminStatus [post]
func UpdateAdminStatus(c *gin.Context) {
	var dto entity.UpdateAdminStatusDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.ValidationError(c, err)
		return
	}

	if err := SysAdminService.UpdateAdminStatus(&dto); err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c)
}

// @Summary 重置用户密码
// @Description 重置用户密码
// @Tags 用户管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body entity.ResetPasswordDto true "重置密码请求"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/adminService/resetPassword [post]
func ResetPassword(c *gin.Context) {
	var dto entity.ResetPasswordDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.ValidationError(c, err)
		return
	}

	if err := SysAdminService.ResetPassword(&dto); err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c)
}

// @Summary 修改个人资料
// @Description 修改个人资料
// @Tags 用户管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body entity.UpdatePersonalDto true "修改个人资料请求结构体"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/adminService/updatePersonal [post]
func UpdatePersonal(c *gin.Context) {
	// id要通过jwt获取(当前登录用户)
	id, _ := c.Get(global.LoggedUser)
	var dto entity.UpdatePersonalDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.ValidationError(c, err)
		return
	}
	if err := SysAdminService.UpdatePersonal(id.(uint), &dto); err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c)
}

// @Summary 修改个人密码
// @Description 修改个人密码
// @Tags 用户管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body entity.UpdatePasswordDto true "修改个人密码请求结构体"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/adminService/updatePassword [post]
func UpdatePassword(c *gin.Context) {
	// id要通过jwt获取(当前登录用户)
	id, _ := c.Get(global.LoggedUser)
	var dto entity.UpdatePasswordDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.ValidationError(c, err)
		return
	}
	if err := SysAdminService.UpdatePassword(id.(uint), &dto); err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c)
}
