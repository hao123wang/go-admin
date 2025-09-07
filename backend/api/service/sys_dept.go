package service

import (
	"errors"
	"go-admin-server/api/entity"
	"go-admin-server/common/response"
	"go-admin-server/common/utils"
	"time"

	"gorm.io/gorm"
)

type SysDeptService struct{}

// 用于在创建/更新部门时，设置父部门id
func (s *SysDeptService) setupParentId(parentId uint, sysDept *entity.SysDept) error {
	// 顶级类型，父id直接设置为nil
	if sysDept.DeptType == 1 {
		sysDept.ParentID = nil
		return nil
	}
	// 检查父部门
	parentDept, err := SysDeptDao.GetDeptById(parentId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ErrInvalidDeptParentID
		}
		return response.ErrServerError
	}
	if parentDept.DeptStatus == 2 {
		return response.ErrParentDeptDisabled
	}
	sysDept.ParentID = &parentId
	return nil
}

// 创建部门
func (s *SysDeptService) CreateDept(dto *entity.CreateDeptDto) error {
	// 检查部门名称是否已存在
	exists, err := SysDeptDao.ExistsByName(dto.DeptName)
	if err != nil {
		return response.ErrServerError
	}
	if exists {
		return response.ErrDeptNameExists
	}

	var sysDept entity.SysDept
	sysDept.DeptName = dto.DeptName // 部门名称
	sysDept.DeptType = dto.DeptType // 部门类型
	// 设置部门父id
	if err := s.setupParentId(dto.ParentID, &sysDept); err != nil {
		return err
	}
	sysDept.CreateAT = utils.HTime{Time: time.Now()}
	if err := SysDeptDao.CreateDept(&sysDept); err != nil {
		return response.ErrServerError
	}
	return nil
}

// 获取部门列表
func (s *SysDeptService) GetDeptList(deptName string, deptStatus int) ([]entity.SysDept, error) {
	if deptStatus < 1 {
		deptStatus = 1
	}
	sysDepts, err := SysDeptDao.GetDeptList(deptName, deptStatus)
	if err != nil {
		return nil, response.ErrServerError
	}
	return sysDepts, nil
}

// 根据id获取部门信息
func (s *SysDeptService) GetDeptById(deptId uint) (*entity.SysDept, error) {
	sysDept, err := SysDeptDao.GetDeptById(deptId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrNotFound
		}
		return nil, response.ErrServerError
	}
	return sysDept, nil
}

// 修改部门信息
func (s *SysDeptService) UpdateDept(dto *entity.UpdateSysDeptDto) error {
	sysDept, err := SysDeptDao.GetDeptById(dto.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ErrDeptNotExists
		}
		return response.ErrServerError
	}
	if dto.DeptName != nil && sysDept.DeptName != *dto.DeptName {
		// 检查名称是否被占用
		exists, err := SysDeptDao.ExistsByName(*dto.DeptName)
		if err != nil {
			return response.ErrServerError
		}
		if exists {
			return response.ErrDeptNameExists
		}
		sysDept.DeptName = *dto.DeptName
	}

	if dto.DeptType != nil && sysDept.DeptType != *dto.DeptType {
		sysDept.DeptType = *dto.DeptType
	}

	if dto.DeptStatus != nil && sysDept.DeptStatus != *dto.DeptStatus {
		sysDept.DeptStatus = *dto.DeptStatus
	}

	// 检查部门的父id是否需要更新
	if dto.ParentID != nil {
		// 不能将自身id设置为自己父部门id
		if *dto.ParentID == sysDept.ID {
			return response.ErrInvalidDeptParentID
		}
		// 当前部门的父id为空 || 当前部门的父id不等于新的父id
		if sysDept.ParentID == nil || *sysDept.ParentID != *dto.ParentID {
			if err := s.setupParentId(*dto.ParentID, sysDept); err != nil {
				return err
			}
		}
	}
	if err := SysDeptDao.UpdateDept(sysDept); err != nil {
		return response.ErrServerError
	}
	return nil
}

// 根据id删除部门
func (s *SysDeptService) DeleteDept(deptId uint) error {
	// 查询部门是有员工
	hasEmployees, err := SysDeptDao.HasEmployees(deptId)
	if err != nil {
		return response.ErrServerError
	}
	if hasEmployees {
		return response.ErrDeptHasEmployees
	}
	// 需要先判断是否有子部门
	hasChildDept, err := SysDeptDao.HasChildDept(deptId)
	if err != nil {
		return response.ErrServerError
	}
	if hasChildDept {
		return response.ErrDeptHasChildDepts
	}
	if err := SysDeptDao.DeleteDept(deptId); err != nil {
		return response.ErrServerError
	}
	return nil
}

// 获取部门下拉列表
func (s *SysDeptService) GetDeptDropdown() ([]entity.DeptDropdownVo, error) {
	dropdown, err := SysDeptDao.GetDeptDropdown()
	if err != nil {
		return nil, response.ErrServerError
	}
	return dropdown, nil
}
