package service

import (
	"errors"
	"go-admin-server/api/entity"
	"go-admin-server/common/response"
	"go-admin-server/common/utils"
	"go-admin-server/pkg/encrypt"
	"go-admin-server/pkg/jwt"
	"time"

	"gorm.io/gorm"
)

type SysAdminService struct{}

// 用户登录
func (s *SysAdminService) Login(ip, browser, Os string, dto *entity.LoginDto) (*entity.SysAdmin, string, error) {
	// 先检查验证码
	if !captchaStore.Verify(dto.CaptchaID, dto.CaptchaImage, true) {
		SysLogDao.CreateLoginLog(dto.Username, ip, utils.GetRealAddressByIP(ip), browser, Os, "验证码错误或失效", 2)
		return nil, "", response.ErrCaptchaError
	}
	// 根据名称获取用户
	user, err := SysAdminDao.GetAdminByName(dto.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			SysLogDao.CreateLoginLog(dto.Username, ip, utils.GetRealAddressByIP(ip), browser, Os, "用户名或密码错误", 2)
			return nil, "", response.ErrLoginError
		}
		SysLogDao.CreateLoginLog(dto.Username, ip, utils.GetRealAddressByIP(ip), browser, Os, "服务器故障", 2)
		return nil, "", response.ErrServerError
	}
	// 检查密码
	if !encrypt.VerifyPassword(user.Password, dto.Password) {
		SysLogDao.CreateLoginLog(dto.Username, ip, utils.GetRealAddressByIP(ip), browser, Os, "用户名或密码错误", 2)
		return nil, "", response.ErrLoginError
	}

	// 检测账号状态
	if user.Status == 2 {
		SysLogDao.CreateLoginLog(dto.Username, ip, utils.GetRealAddressByIP(ip), browser, Os, "账号已停用", 2)
		return nil, "", response.ErrAdminDisabled
	}

	// 生成token
	tokenString, err := jwt.GenerateToken(user)
	if err != nil {
		SysLogDao.CreateLoginLog(dto.Username, ip, utils.GetRealAddressByIP(ip), browser, Os, "服务器故障", 2)
		return nil, "", response.ErrServerError
	}

	// 登录成功
	SysLogDao.CreateLoginLog(dto.Username, ip, utils.GetRealAddressByIP(ip), browser, Os, "登录成功", 1)
	return user, tokenString, nil
}

// 创建用户
func (s *SysAdminService) CreateAdmin(dto *entity.CreateAdminDto) error {
	// 检查名称是否已被占用
	nameExists, err := SysAdminDao.ExistsByName(dto.Username)
	if err != nil {
		return response.ErrServerError
	}
	if nameExists {
		return response.ErrAdminNameExists
	}

	// 检查昵称是否已被占用
	nicknameExists, err := SysAdminDao.ExistsNickname(dto.Nickname)
	if err != nil {
		return response.ErrServerError
	}
	if nicknameExists {
		return response.ErrAdminNicknameExists
	}

	// 检查部门
	sysDept, err := SysDeptDao.GetDeptById(dto.DeptID)
	if err != nil {
		// 如果部门不存在
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ErrDeptNotExists
		}
		return response.ErrServerError
	}
	// 部门已停用
	if sysDept.DeptStatus == 2 {
		return response.ErrDeptDisabled
	}

	// 检查岗位
	sysPost, err := SysPostDao.GetSysPostById(dto.PostID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ErrPostNotExists
		}
		return response.ErrServerError
	}
	if sysPost.PostStatus == 2 {
		return response.ErrPostDisabled
	}

	// 检查角色
	sysRole, err := SysRoleDao.GetRoleByID(dto.RoleID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ErrRoleNotExists
		}
		return response.ErrServerError
	}
	if sysRole.RoleStatus == 2 {
		return response.ErrRoleDisabled
	}

	// 密码加密
	hashPassword, _ := encrypt.EncryptPassword(dto.Password)

	// 创建用户的同时，分配角色
	sysAdmin := &entity.SysAdmin{
		Username:  dto.Username,
		Password:  hashPassword,
		Nickname:  dto.Nickname,
		Email:     dto.Email,
		Phone:     dto.Phone,
		Status:    dto.Status,
		Note:      dto.Note,
		DeptID:    dto.DeptID,
		PostID:    dto.PostID,
		CreatedAt: utils.HTime{Time: time.Now()},
	}
	if err := SysAdminDao.CreateAdmin(dto.RoleID, sysAdmin); err != nil {
		return response.ErrServerError
	}
	return nil
}

// 联表查询用户信息列表
func (s *SysAdminService) JointGetAdminList(pageNum, pageSize, status int, username, beginTime, endTime string) (*entity.AdminListVo, error) {
	if pageNum < 1 {
		pageNum = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	if status != 1 && status != 2 {
		status = 1
	}

	sysAdminList, total, err := SysAdminDao.JointGetAdminList(pageNum, pageSize, status, username, beginTime, endTime)
	if err != nil {
		return nil, response.ErrServerError
	}
	adminListVo := &entity.AdminListVo{
		Data: sysAdminList,
		Pagination: response.PaginationMeta{
			PageNum:    pageNum,
			PageSize:   pageSize,
			Total:      total,
			TotalPages: (total + pageSize - 1) / pageSize,
		},
	}
	return adminListVo, nil
}

// 联表查询单个用户信息
func (s *SysAdminService) JointGetAdminById(userId uint) (*entity.GetAdminByIdVo, error) {
	sysAdmin, err := SysAdminDao.JointGetAdminById(userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrAdminNotExists
		}
		return nil, response.ErrServerError
	}
	return sysAdmin, nil
}

// 修改用户信息
func (s *SysAdminService) UpdateSysAdmin(dto *entity.UpdateAdminDto) error {
	// 获取当前用户
	user, err := SysAdminDao.GetAdminById(dto.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ErrAdminNotExists
		}
		return response.ErrServerError
	}

	// 逐个字段检查
	if dto.Username != nil && *dto.Username != user.Username {
		// 检查新名字是否被占用
		exists, _ := SysAdminDao.ExistsByName(*dto.Username)
		if exists {
			return response.ErrAdminNameExists
		}
		user.Username = *dto.Username
	}
	if dto.Nickname != nil && *dto.Nickname != user.Nickname {
		// 检查新昵称是否被占用
		exists, _ := SysAdminDao.ExistsByName(*dto.Username)
		if exists {
			return response.ErrAdminNicknameExists
		}
		user.Nickname = *dto.Nickname
	}
	// 检查新的部门、岗位的存在性，以及状态
	if dto.DeptId != nil && *dto.DeptId != user.DeptID {
		dept, err := SysDeptDao.GetDeptById(*dto.DeptId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return response.ErrDeptNotExists
			}
			return response.ErrServerError
		}
		if dept.DeptStatus == 2 {
			return response.ErrDeptDisabled
		}
		user.DeptID = *dto.DeptId
	}
	if dto.PostId != nil && *dto.PostId != user.PostID {
		post, err := SysPostDao.GetSysPostById(*dto.PostId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return response.ErrPostNotExists
			}
			return response.ErrServerError
		}
		if post.PostStatus == 2 {
			return response.ErrPostDisabled
		}
		user.PostID = *dto.PostId
	}
	if dto.Phone != nil && *dto.Phone != user.Phone {
		user.Phone = *dto.Phone
	}
	if dto.Email != nil && *dto.Email != user.Email {
		user.Email = *dto.Email
	}
	if dto.Status != nil && *dto.Status != user.Status {
		user.Status = *dto.Status
	}
	if dto.Note != nil {
		user.Note = *dto.Note
	}
	// 修改用户信息
	if err := SysAdminDao.UpdateAdmin(user); err != nil {
		return response.ErrServerError
	}
	// 修改角色信息
	if err := SysAdminDao.UpdateAdminRole(dto.ID, *dto.RoleId); err != nil {
		return response.ErrServerError
	}
	return nil
}

// 删除用户
func (s *SysAdminService) DeleteAdmin(userId uint) error {
	// 先检查用户是否存在
	_, err := SysAdminDao.GetAdminById(userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ErrAdminNotExists
		}
		return response.ErrServerError
	}
	// 删除用户
	if err := SysAdminDao.DeleteAdmin(userId); err != nil {
		return response.ErrServerError
	}
	return nil
}

// 修改用户状态
func (s *SysAdminService) UpdateAdminStatus(dto *entity.UpdateAdminStatusDto) error {
	user, err := SysAdminDao.GetAdminById(dto.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ErrAdminNotExists
		}
		return response.ErrServerError
	}
	user.Status = dto.NewStatus
	if err := SysAdminDao.UpdateAdmin(user); err != nil {
		return response.ErrServerError
	}
	return nil
}

// 修改用户密码
func (s *SysAdminService) ResetPassword(dto *entity.ResetPasswordDto) error {
	user, err := SysAdminDao.GetAdminById(dto.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ErrAdminNotExists
		}
		return response.ErrServerError
	}
	newHashPassword, _ := encrypt.EncryptPassword(dto.NewPassword)
	user.Password = newHashPassword
	if err := SysAdminDao.UpdateAdmin(user); err != nil {
		return response.ErrServerError
	}
	return nil
}

// 修改个人资料
func (s *SysAdminService) UpdatePersonal(adminId uint, dto *entity.UpdatePersonalDto) error {
	// 获取用户
	admin, err := SysAdminDao.GetAdminById(adminId)
	if err != nil {
		return response.ErrServerError
	}
	// 修改用户信息
	if dto.Username != nil && *dto.Username != admin.Username {
		// 改名前先判断新名字是否被占用
		nameExists, _ := SysAdminDao.ExistsByName(*dto.Username)
		if nameExists {
			return response.ErrAdminNameExists
		}
		admin.Username = *dto.Username
	}
	if dto.Nickname != nil && *dto.Nickname != admin.Nickname {
		nicknameExists, _ := SysAdminDao.ExistsNickname(*dto.Nickname)
		if nicknameExists {
			return response.ErrAdminNicknameExists
		}
		admin.Nickname = *dto.Nickname
	}
	if dto.Icon != nil && *dto.Icon != admin.Icon {
		admin.Icon = *dto.Icon
	}
	if dto.Phone != nil && *dto.Phone != admin.Phone {
		admin.Phone = *dto.Phone
	}
	if dto.Email != nil && *dto.Email != admin.Email {
		admin.Email = *dto.Email
	}
	if dto.Note != nil {
		admin.Note = *dto.Note
	}
	if err := SysAdminDao.UpdateAdmin(admin); err != nil {
		return response.ErrServerError
	}
	return nil
}

// 修改个人密码
func (s *SysAdminService) UpdatePassword(adminId uint, dto *entity.UpdatePasswordDto) error {
	// 获取用户
	admin, err := SysAdminDao.GetAdminById(adminId)
	if err != nil {
		return response.ErrServerError
	}

	// 验证旧密码
	if !encrypt.VerifyPassword(admin.Password, dto.Password) {
		return response.ErrPasswordError
	}

	// 如果两次新密码不一致
	if dto.NewPassword != dto.RePassword {
		return response.ErrPasswordInConsistent
	}

	// 修改新密码
	hashNewPwd, _ := encrypt.EncryptPassword(dto.NewPassword)
	admin.Password = hashNewPwd
	if err := SysAdminDao.UpdateAdmin(admin); err != nil {
		return response.ErrServerError
	}
	return nil
}
