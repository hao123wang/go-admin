package controller

import (
	"go-admin-server/api/entity"
	"go-admin-server/common/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 创建菜单
// @Description 创建菜单
// @Tags 菜单管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body entity.CreateMenuDto true "创建菜单请求结构体"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/menuService/createMenu [post]
func CreateMenu(c *gin.Context) {
	var dto entity.CreateMenuDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.ValidationError(c, err)
		return
	}
	if err := SysMenuService.CreateMenu(&dto); err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c)
}

// @Summary 查询菜单列表
// @Description 查询菜单列表
// @Tags 菜单管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param menuName query string false "菜单名称"
// @Param menuStatus query int false "菜单状态"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/menuService/getMenuList [get]
func GetMenuList(c *gin.Context) {
	menuName := c.Query("menuName")
	menuStatus, _ := strconv.Atoi(c.Query("menuStatus"))
	sysMenus, err := SysMenuService.GetMenuList(menuName, menuStatus)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.SuccessWithData(c, sysMenus)
}

// @Summary 根据id查询菜单
// @Description 根据id查询菜单
// @Tags 菜单管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body entity.GetMenuByIdDto true "根据id查询菜单请求"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/menuService/getMenuById [post]
func GetMenuById(c *gin.Context) {
	var dto entity.GetMenuByIdDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.Error(c, response.ErrInvalidParams)
		return
	}
	sysMenu, err := SysMenuService.GetMenuById(dto.ID)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.SuccessWithData(c, sysMenu)
}

// @Summary 修改菜单信息
// @Description 修改菜单信息
// @Tags 菜单管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body entity.UpdateSysMenuDto true "修改菜单请求结构体"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/menuService/updateMenu [post]
func UpdateMenu(c *gin.Context) {
	var dto entity.UpdateSysMenuDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.ValidationError(c, err)
		return
	}

	if err := SysMenuService.UpdateMenu(&dto); err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c)
}

// @Summary 根据id删除单个菜单
// @Description 根据id删除单个菜单
// @Tags 菜单管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body entity.DeleteMenuDto true "删除菜单请求"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/menuService/deleteMenu [post]
func DeleteMenu(c *gin.Context) {
	var dto entity.DeleteMenuDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.ValidationError(c, err)
		return
	}
	if err := SysMenuService.DeleteMenu(dto.ID); err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c)
}

// @Summary 菜单下拉列表
// @Description 菜单下拉列表
// @Tags 菜单管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/menuService/getMenuDropdown [get]
func GetMenuDropdown(c *gin.Context) {
	dropdown, err := SysMenuService.GetMenuDropdown()
	if err != nil {
		response.Error(c, err)
		return
	}
	response.SuccessWithData(c, dropdown)
}
