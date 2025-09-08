# go-admin
gin+vue前后端分离的后台管理项目，基于用户-角色-权限的RBAC权限控制  
主要功能模块：菜单模块、角色模块、用户模块、部门模块、登录日志和操作日志  
后端技术栈：gin + gorm + mysql + redis  
**后端项目启动**  
<pre>
  # 克隆项目
  git clone https://github.com/hao123wang/go-admin.git  
  # 进入后端项目目录
  cd go-admin/backend  
  # 通过命令行迁移数据库表结构体
  go run main.go --sql  
  # 启动项目  
  go run main.go
</pre>
