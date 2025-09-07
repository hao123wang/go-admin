package dao

import (
	"errors"
	"go-admin-server/api/entity"
	"go-admin-server/global"

	"gorm.io/gorm"
)

type SysDeptDao struct{}

// 判断部门名称是否存在
func (d *SysDeptDao) ExistsByName(deptName string) (bool, error) {
	var count int64
	err := global.DB.Model(&entity.SysDept{}).Where("dept_name = ?", deptName).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// 根据名称获取部门信息
func (d *SysDeptDao) GetSysDeptByName(deptName string) (*entity.SysDept, error) {
	var sysDept entity.SysDept
	err := global.DB.Where("dept_name = ?", deptName).First(&sysDept).Error
	if err != nil {
		return nil, err
	}
	return &sysDept, nil
}

// 根据ID查询部门
func (d SysDeptDao) GetDeptById(deptId uint) (*entity.SysDept, error) {
	var sysDept entity.SysDept
	if err := global.DB.Where("id = ?", deptId).First(&sysDept).Error; err != nil {
		return nil, err
	}
	return &sysDept, nil
}

// 创建部门
func (d *SysDeptDao) CreateDept(sysDept *entity.SysDept) error {
	return global.DB.Create(sysDept).Error
}

// 获取部门列表
func (d *SysDeptDao) GetDeptList(deptName string, deptStaus int) ([]entity.SysDept, error) {
	var sysDepts []entity.SysDept
	query := global.DB.Model(entity.SysDept{})
	if deptName != "" {
		query = query.Where("dept_name LIKE ?", "%"+deptName+"%")
	}
	err := query.Where("dept_status = ?", deptStaus).Find(&sysDepts).Error
	if err != nil {
		return nil, err
	}
	return sysDepts, nil
}

// 更新部门信息
func (d *SysDeptDao) UpdateDept(sysDept *entity.SysDept) error {
	return global.DB.Save(sysDept).Error
}

// 查询部门中是否有员工
func (d *SysDeptDao) HasEmployees(deptID uint) (bool, error) {
	var count int64
	err := global.DB.Model(&entity.SysAdmin{}).Select("1").Where("dept_id = ?", deptID).First(&count).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return count > 1, nil
}

// 查询部门是否有子部门
func (d *SysDeptDao) HasChildDept(deptID uint) (bool, error) {
	var count int64
	err := global.DB.Model(&entity.SysDept{}).Where("parent_id = ?", deptID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// 根据id删除部门
func (d *SysDeptDao) DeleteDept(deptID uint) error {
	return global.DB.Where("id = ?", deptID).Delete(&entity.SysDept{}).Error
}

// 获取部门下拉列表
func (d *SysDeptDao) GetDeptDropdown() ([]entity.DeptDropdownVo, error) {
	var dropdown []entity.DeptDropdownVo
	err := global.DB.Model(&entity.SysDept{}).
		Select("id,dept_name,parent_id").
		Scan(&dropdown).Error
	if err != nil {
		return nil, err
	}
	return dropdown, nil
}
