package router

import (
	"go-admin-server/api/controller"
	"go-admin-server/global"
	"go-admin-server/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFile "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(global.Config.Server.Mode)
	router := gin.New()
	router.Use(middleware.Cors())
	router.Use(middleware.GinLogger(), middleware.GinRecovery(true))

	router.StaticFS("/uploads", http.Dir("./uploads"))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFile.Handler))

	router.GET("/api/captcha", controller.Captcha) // 生成验证码
	router.POST("/api/login", controller.Login)    // 用户登录

	// 私有路由（需要认证）
	private := router.Group("/api")
	private.Use(middleware.JWTAuth(), middleware.OperationLog())
	{
		private.POST("/upload", controller.Upload) // 单图片上传
		// 岗位管理
		postGroup := private.Group("/postService")
		{
			postGroup.POST("/createPost", controller.CreatePost)             // 创建岗位
			postGroup.GET("/getPostList", controller.GetPostList)            // 查询岗位列表
			postGroup.POST("/getPostById", controller.GetPostById)           // 根据id查询岗位信息
			postGroup.POST("/updatePost", controller.UpdatePost)             // 修改岗位信息
			postGroup.POST("/deletePost", controller.DeletePost)             // 根据id删除岗位
			postGroup.POST("/batchDeletePosts", controller.BatchDeletePosts) // 批量删除岗位
			postGroup.POST("/updatePostStatus", controller.UpdatePostStatus) // 修改岗位状态
			postGroup.GET("/getPostDropdown", controller.GetPostDropdown)    // 岗位下拉列表
		}

		// 部门管理
		deptGroup := private.Group("/deptService")
		{
			deptGroup.POST("/createDept", controller.CreateDept)          // 创建部门
			deptGroup.GET("/getDeptList", controller.GetDeptList)         // 查询部门列表
			deptGroup.POST("/getDeptById", controller.GetDeptById)        // 根据id查询部门
			deptGroup.POST("/updateDept", controller.UpdateDept)          // 修改部门信息
			deptGroup.POST("/deleteDept", controller.DeleteDept)          // 根据id删除单个部门
			deptGroup.GET("/getDeptDropdown", controller.GetDeptDropdown) // 部门下拉列表
		}

		// 菜单管理
		menuGroup := private.Group("/menuService")
		{
			menuGroup.POST("/createMenu", controller.CreateMenu)          // 创建菜单
			menuGroup.GET("/getMenuList", controller.GetMenuList)         // 查询菜单列表
			menuGroup.POST("/getMenuById", controller.GetMenuById)        // 根据id查询菜单
			menuGroup.POST("/updateMenu", controller.UpdateMenu)          // 修改菜单信息
			menuGroup.POST("/deleteMenu", controller.DeleteMenu)          // 根据id删除单个菜单
			menuGroup.GET("/getMenuDropdown", controller.GetMenuDropdown) // 菜单下拉列表
		}

		// 角色管理
		roleGroup := private.Group("/roleService")
		{
			roleGroup.POST("/createRole", controller.CreateRole)             // 创建角色
			roleGroup.GET("/getRoleList", controller.GetRoleList)            // 查询角色列表
			roleGroup.POST("/getRoleById", controller.GetRoleById)           // 根据id查询角色
			roleGroup.POST("/updateRole", controller.UpdateRole)             // 修改角色信息
			roleGroup.POST("/deleteRole", controller.DeleteRole)             // 删除角色
			roleGroup.POST("/updateRoleStatus", controller.UpdateRoleStatus) // 修改角色状态
			roleGroup.GET("/getRoleDropdown", controller.GetRoleDropdown)    // 角色下拉列表
			roleGroup.POST("/getRoleMenus", controller.GetRoleMenus)         // 查询角色的权限列表
			roleGroup.POST("/assignRoleMenus", controller.AssignRoleMenus)   // 分配角色权限
		}

		// 用户管理
		adminGroup := private.Group("/adminService")
		{
			adminGroup.POST("/createAdmin", controller.CreateAdmin)             // 创建用户
			adminGroup.GET("/getAdminList", controller.GetAdminList)            // 查询用户列表
			adminGroup.POST("/getAdminById", controller.GetAdminById)           // 根据id查询用户
			adminGroup.POST("/updateAdmin", controller.UpdateAdmin)             // 修改用户信息
			adminGroup.POST("/deleteAdmin", controller.DeleteAdmin)             // 删除用户
			adminGroup.POST("/updateAdminStatus", controller.UpdateAdminStatus) // 修改用户状态
			adminGroup.POST("/resetPassword", controller.ResetPassword)         // 重置密码
			adminGroup.POST("/updatePersonal", controller.UpdatePersonal)       // 修改个人资料
			adminGroup.POST("/updatePassword", controller.UpdatePassword)       // 修改个人密码
		}

		// 日志管理
		logGroup := private.Group("/logService")
		{
			logGroup.GET("/getLoginLogList", controller.GetLoginLogList)          // 查询登录日志列表
			logGroup.POST("/deleteLoginLog", controller.DeleteLoginLog)           // 删除登录日志
			logGroup.POST("/batchDeleteLoginLog", controller.BatchDeleteLoginLog) // 批量删除登录日志
			logGroup.GET("/getOpLogList", controller.GetOpLogList)                // 查询操作日志列表
			logGroup.POST("/deleteOpLog", controller.DeleteOpLog)                 // 删除操作日志
			logGroup.POST("/batchDeleteOpLog", controller.BatchDeleteOpLog)       // 批量删除操作日志
		}
	}
	return router
}
