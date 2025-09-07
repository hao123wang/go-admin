package response

// 业务状态码
const (
	CodeSuccess = 200 // 成功

	// 1000~2000 HTTPStatus为BadRequest
	CodeInvalidParams = 1000 // 请求参数错误

	// 岗位模块
	CodePostCodeExists = 1101 // 岗位编号已存在
	CodePostNameExists = 1102 // 岗位名称已存在
	CodePostNotExists  = 1103 // 岗位不存在
	CodePostDisabled   = 1104 // 岗位已停用

	// 部门模块
	CodeDeptNameExists      = 1201 // 部门名称已存在
	CodeInvalidDeptParentID = 1202 // 无效的部门id
	CodeDeptDisabled        = 1203 // 当前部门已停用
	CodeParentDeptDisabled  = 1204 // 父部门已停用
	CodeDeptNotExists       = 1205 // 目标部门不存在
	CodeDeptHasEmployees    = 1206 // 部门中有员工
	CodeDeptHasChildDept    = 1207 // 存在子部门

	// 菜单模块
	CodeMenuNameExists      = 1301 // 菜单名称已存在
	CodeMenuNotExists       = 1302 // 目标菜单不存在
	CodeParentMenuDisabled  = 1303 // 父菜单已被禁用
	CodeInvalidMenuParentID = 1304 // 无效的菜单id
	CodeHasSubmenu          = 1305 // 存在子菜单

	// 角色模块
	CodeRoleNameExists = 1401 // 角色名称已存在
	CodeRoleKeyExists  = 1402 // 角色关键字已存在
	CodeRoleNotExists  = 1403 // 角色不存在
	CodeRoleDisabled   = 1404 // 角色已被禁用

	// 用户模块
	CodeAdmiNameExists       = 1501 // 用户名称已存在
	CodeAdminNicknameExists  = 1502 // 用户昵称已存在
	CodeAdminNotExists       = 1503 // 用户不存在
	CodeLoginError           = 1504 // 用户名或密码错误
	CodeCaptchaError         = 1505 // 验证码错误或已失效
	CodePasswordError        = 1506 // 旧密码错误
	CodePasswordInConsistent = 1507 // 两次密码不一致
	CodeAdminDisabled        = 1508 // 账号已停用

	CodeFileUploadFail = 1601 // 文件上传失败

	// 2000~3000 对应的HTTPStatus 为 Unauthorized
	CodeUnauthorized     = 2000 // 未认证
	CodeTokenFormatError = 2001 // token格式错误
	CodeTokenInvalid     = 2002 // 无效token

	// 3000~4000 对应的HTTPStatus 为 Forbidden

	CodeNotFound = 4000 // 请求资源不存在

	CodeServerError = 5000 // 服务器内部错误
)

// BusinessError 业务错误类型
type BusinessError struct {
	Code    int
	Message string
}

// 实现error接口
func (e *BusinessError) Error() string {
	return e.Message
}

// 创建业务错误
func NewBusinessError(code int, message string) *BusinessError {
	return &BusinessError{
		Code:    code,
		Message: message,
	}
}

// 统一错误注册
var (
	ErrServerError   = NewBusinessError(CodeServerError, "服务器内部错误")
	ErrNotFound      = NewBusinessError(CodeNotFound, "请求资源不存在")
	ErrInvalidParams = NewBusinessError(CodeInvalidParams, "请求参数错误")

	// 岗位模块
	ErrPostCodeExists = NewBusinessError(CodePostCodeExists, "岗位编号已存在")
	ErrPostNameExists = NewBusinessError(CodePostNameExists, "岗位名称已存在")
	ErrPostNotExists  = NewBusinessError(CodePostNotExists, "目标岗位不存在")
	ErrPostDisabled   = NewBusinessError(CodePostDisabled, "岗位已停用")

	// 部门模块
	ErrDeptNameExists      = NewBusinessError(CodeDeptNameExists, "部门名称已存在")
	ErrInvalidDeptParentID = NewBusinessError(CodeInvalidDeptParentID, "无效的父部门id")
	ErrDeptDisabled        = NewBusinessError(CodeDeptDisabled, "当前部门已停用")
	ErrParentDeptDisabled  = NewBusinessError(CodeParentDeptDisabled, "父部门已停用")
	ErrDeptNotExists       = NewBusinessError(CodeDeptNotExists, "目标部门不存在")
	ErrDeptHasEmployees    = NewBusinessError(CodeDeptHasEmployees, "部门中有员工")
	ErrDeptHasChildDepts   = NewBusinessError(CodeDeptHasChildDept, "存在子部门")

	// 菜单模块
	ErrMenuNameExists      = NewBusinessError(CodeMenuNameExists, "菜单名称已存在")
	ErrMenuNotExists       = NewBusinessError(CodeMenuNotExists, "目标菜单不存在")
	ErrParentMenuDisabled  = NewBusinessError(CodeParentMenuDisabled, "父菜单已被禁用")
	ErrInvalidMenuParentID = NewBusinessError(CodeInvalidMenuParentID, "无效的父菜单id")
	ErrHasSubmenu          = NewBusinessError(CodeHasSubmenu, "存在子菜单")

	// 角色模块
	ErrRoleNameExists = NewBusinessError(CodeRoleNameExists, "角色名称已存在")
	ErrRoleKeyExists  = NewBusinessError(CodeRoleKeyExists, "角色关键字已存在")
	ErrRoleNotExists  = NewBusinessError(CodeRoleNotExists, "角色不存在")
	ErrRoleDisabled   = NewBusinessError(CodeRoleDisabled, "角色已被禁用")

	// 用户模块
	ErrAdminNameExists      = NewBusinessError(CodeAdmiNameExists, "用户名称已存在")
	ErrAdminNicknameExists  = NewBusinessError(CodeAdminNicknameExists, "用户昵称已存在")
	ErrAdminNotExists       = NewBusinessError(CodeAdminNotExists, "用户不存在")
	ErrLoginError           = NewBusinessError(CodeLoginError, "用户名或密码错误")
	ErrCaptchaError         = NewBusinessError(CodeCaptchaError, "验证码错误或失效")
	ErrPasswordError        = NewBusinessError(CodePasswordError, "旧密码错误")
	ErrPasswordInConsistent = NewBusinessError(CodePasswordInConsistent, "两次新密码不一致")
	ErrAdminDisabled        = NewBusinessError(CodeAdminDisabled, "账号已停用")

	ErrAdminUnauthorized = NewBusinessError(CodeUnauthorized, "用户未认证")
	ErrTokenFormatError  = NewBusinessError(CodeTokenFormatError, "Token格式错误")
	ErrTokenInvalid      = NewBusinessError(CodeTokenInvalid, "无效的Token")

	ErrFileUploadFail = NewBusinessError(CodeFileUploadFail, "文件上传失败")
)
