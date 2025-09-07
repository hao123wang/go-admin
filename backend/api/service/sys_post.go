package service

import (
	"errors"
	"go-admin-server/api/entity"
	"go-admin-server/common/response"
	"go-admin-server/common/utils"
	"go-admin-server/global"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SysPostService struct{}

// 创建新岗位
func (s *SysPostService) CreateSysPost(dto *entity.CreateSysPostDto) error {
	// 判断岗位编码是否已存在
	codeExists, err := SysPostDao.ExistsByCode(dto.PostCode)
	if err != nil {
		global.Logger.Error("Failed to validate post code", zap.Error(err))
		return response.ErrServerError
	}
	if codeExists {
		return response.ErrPostCodeExists
	}

	// 判断岗位名称是否已存在
	nameExists, err := SysPostDao.ExistsByName(dto.PostName)
	if err != nil {
		return response.ErrServerError
	}
	if nameExists {
		return response.ErrPostNameExists
	}
	// 添加新岗位
	newSysPost := &entity.SysPost{
		PostName:    dto.PostName,
		PostCode:    dto.PostCode,
		Remark:      dto.Remark,
		CreatedTime: utils.HTime{Time: time.Now()},
	}
	if dto.PostStauts != 2 { // 默认岗位状态为1（启用）
		newSysPost.PostStatus = 1
	}
	if err := SysPostDao.CreateSysPost(newSysPost); err != nil {
		return response.ErrServerError
	}
	return nil
}

// 查询岗位列表
func (s *SysPostService) GetSysPostList(pageNum, pageSize, postStatus int, postName, beginTime, endTime string) (*entity.SysPostListVo, error) {
	if pageNum < 1 {
		pageNum = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if postStatus < 1 {
		postStatus = 1
	}
	postList, total, err := SysPostDao.GetSysPostList(pageNum, pageSize, postStatus, postName, beginTime, endTime)
	if err != nil {
		return nil, response.ErrServerError
	}
	paginatedReulst := &entity.SysPostListVo{
		Data: postList,
		Pagination: response.PaginationMeta{
			PageNum:    pageNum,
			PageSize:   pageSize,
			Total:      total,
			TotalPages: (total + pageSize - 1) / pageSize,
		},
	}
	return paginatedReulst, nil
}

// 根据id查询单个岗位
func (s *SysPostService) GetSysPost(postID uint) (*entity.SysPost, error) {
	sysPost, err := SysPostDao.GetSysPostById(postID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrPostNotExists
		}
		return nil, response.ErrServerError
	}
	return sysPost, nil
}

// 更新岗位信息
func (s *SysPostService) UpdateSysPost(dto *entity.UpdateSysPostDto) error {
	// 获取原来的岗位
	post, err := SysPostDao.GetSysPostById(dto.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ErrPostNotExists
		}
		return response.ErrServerError
	}

	// 如果传入了新的岗位名称
	if dto.PostName != nil && *dto.PostName != post.PostName {
		postExist, _ := SysPostDao.GetSysPostByName(*dto.PostName)
		if postExist != nil {
			return response.ErrPostNameExists
		}
	}
	// 如果传入了新的岗位编号
	if dto.PostCode != nil && *dto.PostCode != post.PostCode {
		postExist, _ := SysPostDao.GetSysPostByCode(*dto.PostCode)
		if postExist != nil {
			return response.ErrPostCodeExists
		}
	}
	if dto.PostName != nil {
		post.PostName = *dto.PostName
	}
	if dto.PostCode != nil {
		post.PostCode = *dto.PostCode
	}
	if dto.Remark != nil {
		post.Remark = *dto.Remark
	}
	if dto.PostStauts != nil {
		post.PostStatus = *dto.PostStauts
	}
	if err := SysPostDao.UpdatePost(post); err != nil {
		return response.ErrServerError
	}
	return nil
}

// 根据id删除单个岗位
func (s *SysPostService) DeleteSysPost(postId uint) error {
	// 判断岗位是否存在
	_, err := SysPostDao.GetSysPostById(postId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ErrPostNotExists
		}
		return response.ErrServerError
	}

	if err := SysPostDao.DeleteSysPost(postId); err != nil {
		return response.ErrServerError
	}
	return nil
}

// 根据id列表，批量删除岗位
func (s *SysPostService) BatchDeletePosts(postIds []uint) (int64, error) {
	rows, err := SysPostDao.BatchDeletePosts(postIds)
	if err != nil {
		return 0, response.ErrServerError
	}
	return rows, nil
}

// 修改岗位状态
func (s *SysPostService) UpdatePostStatus(dto *entity.UpdatePostStatusDto) error {
	// 根据id获取岗位
	post, err := s.GetSysPost(dto.ID)
	if err != nil {
		return err	
	}
	post.PostStatus = dto.NewStatus
	if err := SysPostDao.UpdatePost(post); err != nil {
		return response.ErrServerError
	}
	return nil
}

// 获取岗位下拉列表
func (s SysPostService) GetSysPostDropdown() ([]entity.SysPostDropdownVo, error) {
	dropdownList, err := SysPostDao.GetSysPostDropdown()
	if err != nil {
		return nil, response.ErrServerError
	}
	return dropdownList, nil
}
