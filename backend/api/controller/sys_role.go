package controller

import (
	"go-admin-server/api/entity"
	"go-admin-server/common/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 创建角色
// @Description 创建角色
// @Tags 角色管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body entity.CreateRoleDto true "创建角色请求结构体"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/roleService/createRole [post]
func CreateRole(c *gin.Context) {
	var dto entity.CreateRoleDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.ValidationError(c, err)
		return
	}
	if err := SysRoleService.CreateRole(&dto); err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c)
}

// @Summary 查询角色列表
// @Description 查询角色列表
// @Tags 角色管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param pageNum query int false "页码"
// @Param pageSize query int false "页大小"
// @Param roleStatus query int false "角色状态"
// @Param roleName query string false "角色名称"
// @Param startTime query string false "开始时间"
// @Param endTiem query string false "结束时间"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/roleService/getRoleList [get]
func GetRoleList(c *gin.Context) {
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	roleStatus, _ := strconv.Atoi(c.Query("roleStatus"))
	roleName := c.Query("roleName")
	beginTime := c.Query("beginTime")
	endTime := c.Query("endTime")

	sysRoleList, err := SysRoleService.GetRoleList(pageNum, pageSize, roleStatus, roleName, beginTime, endTime)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.SuccessWithData(c, sysRoleList)
}

// @Summary 根据id查询角色
// @Description 根据id查询角色
// @Tags 角色管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body entity.GetRoleByIdDto true "根据id查询角色请求"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/roleService/getRoleById [post]
func GetRoleById(c *gin.Context) {
	var dto entity.GetRoleByIdDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.ValidationError(c, err)
		return
	}
	sysRole, err := SysRoleService.GetRoleByID(dto.ID)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.SuccessWithData(c, sysRole)
}

// @Summary 修改角色
// @Description 修改角色
// @Tags 角色管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body entity.UpdateRoleDto true "修改角色请求结构体"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/roleService/updateRole [post]
func UpdateRole(c *gin.Context) {
	var dto entity.UpdateRoleDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.ValidationError(c, err)
		return
	}

	if err := SysRoleService.UpdateRole(&dto); err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c)
}

// @Summary 删除角色
// @Description 删除角色
// @Tags 角色管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body entity.DeleteRoleDto true "删除角色请求"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/roleService/deleteRole [post]
func DeleteRole(c *gin.Context) {
	var dto entity.DeleteRoleDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.ValidationError(c, err)
		return
	}
	if err := SysRoleService.DeleteRole(dto.ID); err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c)
}

// @Summary 修改角色状态
// @Description 修改角色状态
// @Tags 角色管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body entity.UpdateRoleStatusDto true "修改角色状态请求结构体"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/roleService/updateRoleStatus [post]
func UpdateRoleStatus(c *gin.Context) {
	var dto entity.UpdateRoleStatusDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.ValidationError(c, err)
		return
	}

	if err := SysRoleService.UpdateRoleStatus(&dto); err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c)
}

// @Summary 角色下拉列表
// @Description 角色下拉列表
// @Tags 角色管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/roleService/getRoleDropdown [get]
func GetRoleDropdown(c *gin.Context) {
	dropdown, err := SysRoleService.GetRoleDropdown()
	if err != nil {
		response.Error(c, err)
		return
	}
	response.SuccessWithData(c, dropdown)
}

// @Summary 角色权限分配
// @Description 角色权限分配
// @Tags 角色管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body entity.AssignRoleMenusDto true "权限菜单列表"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/roleService/assignRoleMenus [post]
func AssignRoleMenus(c *gin.Context) {
	var dto entity.AssignRoleMenusDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.ValidationError(c, err)
		return
	}
	if err := SysRoleService.AssignRoleMenus(&dto); err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c)
}

// @Summary 查询角色权限列表
// @Description 查询角色权限列表
// @Tags 角色管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body entity.GetRoleMenusDto true "查询角色权限"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/roleService/getRoleMenus [post]
func GetRoleMenus(c *gin.Context) {
	var dto entity.GetRoleMenusDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.ValidationError(c, err)
		return
	}
	roleMenus, err := SysRoleService.GetRoleMenus(dto.ID)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.SuccessWithData(c, roleMenus)
}
