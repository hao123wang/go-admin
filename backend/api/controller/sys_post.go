package controller

import (
	"go-admin-server/api/entity"
	"go-admin-server/common/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 创建岗位
// @Description 创建岗位
// @Tags 岗位管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body entity.CreateSysPostDto true "创建岗位的请求参数"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/postService/createPost [post]
func CreatePost(c *gin.Context) {
	var dto entity.CreateSysPostDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.ValidationError(c, err)
		return
	}
	if err := SysPostService.CreateSysPost(&dto); err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c)
}

// @Summary 获取岗位列表
// @Description 分页获取岗位列表
// @Tags 岗位管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param pageNum query int false "页码"
// @Param pageSize query int false "每页数量"
// @Param postStatus query int false "岗位状态：1->启用,2->禁用"
// @Param postName query string false "岗位名称"
// @Param beginTime query string false "开始时间"
// @Param endTime query string false "结束时间"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/postService/getPostList [get]
func GetPostList(c *gin.Context) {
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	postStatus, _ := strconv.Atoi(c.Query("postStatus"))

	postName := c.Query("postName")
	beginTime := c.Query("beginTime")
	endTime := c.Query("endTime")

	sysPostListVo, err := SysPostService.GetSysPostList(pageNum, pageSize, postStatus, postName, beginTime, endTime)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.SuccessWithData(c, sysPostListVo)
}

// @Summary 根据id查询岗位信息
// @Description 根据id查询岗位信息
// @Tags 岗位管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body entity.GetPostByIdDto true "获取岗位信息请求结构体"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/postService/getPostById [post]
func GetPostById(c *gin.Context) {
	var dto entity.GetPostByIdDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.Error(c, response.ErrInvalidParams)
		return
	}
	sysPost, err := SysPostService.GetSysPost(dto.ID)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.SuccessWithData(c, sysPost)
}

// @Summary 修改岗位
// @Description 修改岗位信息
// @Tags 岗位管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body entity.UpdateSysPostDto true "修改岗位请求结构体"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/postService/updatePost [post]
func UpdatePost(c *gin.Context) {
	// 绑定请求参数
	var dto entity.UpdateSysPostDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.ValidationError(c, err)
		return
	}

	// 更新岗位信息
	if err := SysPostService.UpdateSysPost(&dto); err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c)
}

// @Summary 删除单个岗位
// @Description 删除单个岗位
// @Tags 岗位管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body entity.DeletePostByIdDto true "删除单个岗位请求结构体"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/postService/deletePost [post]
func DeletePost(c *gin.Context) {
	var dto entity.DeletePostByIdDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.ValidationError(c, err)
		return
	}
	if err := SysPostService.DeleteSysPost(dto.ID); err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c)
}

// @Summary 批量删除岗位
// @Description 批量删除岗位
// @Tags 岗位管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body entity.BatchDeletePostsDto true "要删除的岗位id列表"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/postService/batchDeletePosts [post]
func BatchDeletePosts(c *gin.Context) {
	var dto entity.BatchDeletePostsDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.ValidationError(c, err)
		return
	}
	if len(dto.PostIds) == 0 {
		response.Error(c, response.ErrInvalidParams)
		return
	}
	rows, err := SysPostService.BatchDeletePosts(dto.PostIds)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.SuccessWithData(c, map[string]any{
		"rows": rows,
	})
}

// @Summary 修改岗位状态
// @Description 修改岗位状态
// @Tags 岗位管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body entity.UpdatePostStatusDto true "修改岗位状态请求"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/postService/updatePostStatus [post]
func UpdatePostStatus(c *gin.Context) {
	var dto entity.UpdatePostStatusDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.ValidationError(c, err)
		return
	}
	if err := SysPostService.UpdatePostStatus(&dto); err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c)
}

// @Summary 岗位下拉列表
// @Description 岗位下拉列表
// @Tags 岗位管理
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/postService/getPostDropdown [get]
func GetPostDropdown(c *gin.Context) {
	dropdownList, err := SysPostService.GetSysPostDropdown()
	if err != nil {
		response.Error(c, err)
		return
	}
	response.SuccessWithData(c, dropdownList)
}
