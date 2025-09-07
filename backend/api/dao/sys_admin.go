package dao

import (
	"go-admin-server/api/entity"
	"go-admin-server/global"

	"gorm.io/gorm"
)

type SysAdminDao struct{}

// 检查用户名称是否已存在
func (d *SysAdminDao) ExistsByName(username string) (bool, error) {
	var count int64
	err := global.DB.Model(&entity.SysAdmin{}).Where("username = ?", username).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// 检查用户昵称是否已存在
func (d *SysAdminDao) ExistsNickname(nickname string) (bool, error) {
	var count int64
	err := global.DB.Model(&entity.SysAdmin{}).Where("nickname = ?", nickname).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// 根据id获取用户
func (d *SysAdminDao) GetAdminById(userId uint) (*entity.SysAdmin, error) {
	var sysAdmin entity.SysAdmin
	err := global.DB.Where("id = ?", userId).First(&sysAdmin).Error
	if err != nil {
		return nil, err
	}
	return &sysAdmin, nil
}

// 根据名称获取用户
func (d *SysAdminDao) GetAdminByName(username string) (*entity.SysAdmin, error) {
	var sysAdmin entity.SysAdmin
	if err := global.DB.Where("username = ?", username).First(&sysAdmin).Error; err != nil {
		return nil, err
	}
	return &sysAdmin, nil
}

// 创建用户，以及分配角色
func (d *SysAdminDao) CreateAdmin(roleID uint, user *entity.SysAdmin) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		sysAdminRole := &entity.SysAdminRole{
			AdminID: user.ID,
			RoleID:  roleID,
		}
		if err := tx.Create(sysAdminRole).Error; err != nil {
			return err
		}
		return nil
	})
}

// 联表查询用户信息列表(联表查询)
func (s *SysAdminDao) JointGetAdminList(pageNum, pageSize, status int, username, beginTime, endTime string) ([]entity.AdminList, int, error) {
	query := global.DB.Model(&entity.SysAdmin{}).
		Select("sys_admin.*,ar.role_id,d.dept_name,p.post_name,r.role_name").
		Joins("LEFT JOIN sys_dept d ON sys_admin.dept_id = d.id").
		Joins("LEFT JOIN sys_post p ON sys_admin.post_id = p.id").
		Joins("LEFT JOIN sys_admin_role ar ON sys_admin.id = ar.admin_id").
		Joins("LEFT JOIN sys_role r ON ar.role_id = r.id")

	query = query.Where("sys_admin.status = ?", status)
	if username != "" {
		query = query.Where("username = ?", username)
	}
	if beginTime != "" && endTime != "" {
		query = query.Where("created_at BETWEEN ? AND ?", beginTime, endTime)
	}

	var count int64
	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	var adminList []entity.AdminList
	err := query.Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("sys_admin.created_at DESC").Find(&adminList).Error
	if err != nil {
		return nil, 0, err
	}

	return adminList, int(count), nil
}

// 根据id查询单个用户信息(联表查询)
func (d *SysAdminDao) JointGetAdminById(userId uint) (*entity.GetAdminByIdVo, error) {
	var sysAdmin entity.GetAdminByIdVo
	err := global.DB.Model(&entity.SysAdmin{}).
		Select("sys_admin.*,sys_admin_role.role_id").
		Joins("LEFT JOIN sys_admin_role ON sys_admin.id = sys_admin_role.admin_id").
		Where("sys_admin.id = ?", userId).
		First(&sysAdmin).Error
	if err != nil {
		return nil, err
	}
	return &sysAdmin, nil
}

// 修改用户角色
func (d *SysAdminDao) UpdateAdminRole(userId, roleId uint) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		// 删除旧角色
		if err := tx.Model(&entity.SysAdminRole{}).Delete("admin_id = ?", userId).Error; err != nil {
			return err
		}
		// 分配新角色
		adminRole := &entity.SysAdminRole{
			AdminID: userId,
			RoleID:  roleId,
		}
		if err := tx.Create(adminRole).Error; err != nil {
			return err
		}
		return nil
	})
}

// 修改用户信息
func (d *SysAdminDao) UpdateAdmin(sysAdmin *entity.SysAdmin) error {
	return global.DB.Save(sysAdmin).Error
}

// 删除用户
func (d *SysAdminDao) DeleteAdmin(userId uint) error {
	// 要同时删除关联关系
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", userId).Delete(&entity.SysAdmin{}).Error; err != nil {
			return err
		}
		if err := tx.Where("admin_id = ?", userId).Delete(&entity.SysAdminRole{}).Error; err != nil {
			return err
		}
		return nil
	})
}

