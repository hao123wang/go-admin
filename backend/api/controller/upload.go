package controller

import (
	"go-admin-server/common/response"

	"github.com/gin-gonic/gin"
)

// @Summary 单图片上传
// @Description 单图片上传
// @Tags 文件上传
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param image formData file true "图片"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/upload [post]
func Upload(c *gin.Context) {
	fileHeader, err := c.FormFile("image")
	if err != nil {
		response.Error(c, response.ErrFileUploadFail)
		return
	}

	fullPath, url, err := UploadService.UploadFile(fileHeader.Filename)
	if err != nil {
		response.Error(c, err)
		return
	}
	c.SaveUploadedFile(fileHeader, fullPath)
	response.SuccessWithData(c, map[string]any{
		"url": url,
	})
}
