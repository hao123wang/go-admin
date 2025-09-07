package dao

import (
	"errors"
	"go-admin-server/api/entity"
	"go-admin-server/global"

	"gorm.io/gorm"
)

type SysPostDao struct{}

// 判断岗位编码是否存在
func (d *SysPostDao) ExistsByCode(postCode string) (bool, error) {
	var count int64
	if err := global.DB.Model(&entity.SysPost{}).Where("post_code = ?", postCode).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// 判断岗位名称是否已经存在
func (d *SysPostDao) ExistsByName(postName string) (bool, error) {
	var count int64
	if err := global.DB.Model(&entity.SysPost{}).Where("post_name = ?", postName).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// 根据编码获取岗位
func (d *SysPostDao) GetSysPostByCode(postCode string) (*entity.SysPost, error) {
	var sysPost entity.SysPost
	err := global.DB.Where("post_code = ?", postCode).First(&sysPost).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &sysPost, nil
}

// 根据名称获取岗位
func (d *SysPostDao) GetSysPostByName(name string) (*entity.SysPost, error) {
	var sysPost entity.SysPost
	err := global.DB.Where("post_name = ?", name).First(&sysPost).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &sysPost, nil
}

// 创建新岗位
func (d *SysPostDao) CreateSysPost(newSysPost *entity.SysPost) error {
	return global.DB.Create(newSysPost).Error
}

// 分页获取岗位列表
func (d *SysPostDao) GetSysPostList(pageNum, pageSize, postStatus int, postName, beginTime, endTime string) ([]entity.SysPost, int, error) {
	var sysPosts []entity.SysPost
	query := global.DB.Table("sys_post")
	query = query.Where("post_status = ?", postStatus)

	if postName != "" {
		query = query.Where("post_name = ?", postName)
	}
	if beginTime != "" && endTime != "" {
		query = query.Where("created_at BETWEEN ? AND ?", beginTime, endTime)
	}
	var count int64
	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	err := query.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&sysPosts).Error
	if err != nil {
		return nil, 0, err
	}
	return sysPosts, int(count), nil

}

// 根据id获取岗位信息
func (d *SysPostDao) GetSysPostById(postID uint) (*entity.SysPost, error) {
	var sysPost entity.SysPost
	err := global.DB.Where("id = ?", postID).First(&sysPost).Error
	if err != nil {
		return nil, err
	}
	return &sysPost, nil
}

// 更新岗位信息
func (d *SysPostDao) UpdatePost(post *entity.SysPost) error {
	return global.DB.Save(post).Error
}

// 删除岗位
func (d *SysPostDao) DeleteSysPost(postId uint) error {
	return global.DB.Where("id = ?", postId).Delete(&entity.SysPost{}).Error
}

// 批量删除
func (d *SysPostDao) BatchDeletePosts(postIds []uint) (int64, error) {
	result := global.DB.Where("id IN (?)", postIds).Delete(&entity.SysPost{})
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

// 获取岗位下拉列表
func (d *SysPostDao) GetSysPostDropdown() ([]entity.SysPostDropdownVo, error) {
	var dropdown []entity.SysPostDropdownVo
	err := global.DB.Model(&entity.SysPost{}).
		Select("id,post_name").
		Where("post_status = ?", 1).
		Scan(&dropdown).Error
	if err != nil {
		return nil, err
	}
	return dropdown, nil
}
