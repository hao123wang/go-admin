package controller

import (
	"go-admin-server/api/entity"
	"go-admin-server/common/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 创建部门
// @Description 创建部门
// @Tags 部门管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body entity.CreateDeptDto true "创建部门请求结构体"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/deptService/createDept [post]
func CreateDept(c *gin.Context) {
	var dto entity.CreateDeptDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.ValidationError(c, err)
		return
	}
	if err := SysDeptService.CreateDept(&dto); err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c)
}

// @Summary 查询部门列表
// @Description 查询部门列表
// @Tags 部门管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param deptName query string false "部门名称"
// @Param deptStatus query int false "部门状态"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/deptService/getDeptList [get]
func GetDeptList(c *gin.Context) {
	deptName := c.Query("deptName")
	deptStatus, _ := strconv.Atoi(c.Query("deptStatus"))
	sysDepts, err := SysDeptService.GetDeptList(deptName, deptStatus)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.SuccessWithData(c, sysDepts)
}

// @Summary 根据id查询部门
// @Description 根据id查询部门
// @Tags 部门管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body entity.GetDeptByIdDto true "根据id查询部门请求"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/deptService/getDeptById [post]
func GetDeptById(c *gin.Context) {
	var dto entity.GetDeptByIdDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.ValidationError(c, err)
		return
	}
	sysDept, err := SysDeptService.GetDeptById(dto.ID)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.SuccessWithData(c, sysDept)
}

// @Summary 修改部门信息
// @Description 修改部门信息
// @Tags 部门管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body entity.UpdateSysDeptDto true "修改部门请求结构体"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/deptService/updateDept [post]
func UpdateDept(c *gin.Context) {
	var dto entity.UpdateSysDeptDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.ValidationError(c, err)
		return
	}

	if err := SysDeptService.UpdateDept(&dto); err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c)
}

// @Summary 根据id删除部门
// @Description 根据id删除部门
// @Tags 部门管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body entity.DeleteDeptDto true "删除部门请求"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/deptService/deleteDept [post]
func DeleteDept(c *gin.Context) {
	var dto entity.DeleteDeptDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.ValidationError(c, err)
		return
	}
	if err := SysDeptService.DeleteDept(dto.ID); err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c)
}

// @Summary 部门下拉列表
// @Description 部门下拉列表
// @Tags 部门管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/deptService/getDeptDropdown [get]
func GetDeptDropdown(c *gin.Context) {
	dropdown, err := SysDeptService.GetDeptDropdown()
	if err != nil {
		response.Error(c, err)
		return
	}
	response.SuccessWithData(c, dropdown)
}
